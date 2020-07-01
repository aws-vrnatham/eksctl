package cmdutils

import (
	"fmt"

	api "github.com/weaveworks/eksctl/pkg/apis/eksctl.io/v1alpha5"
)

// NewScaleNodeGroupLoader will load config or use flags for 'eksctl scale nodegroup'
func NewScaleNodeGroupLoader(cmd *Cmd, ng *api.NodeGroup) ClusterConfigLoader {
	l := newCommonClusterConfigLoader(cmd)
	l.nameArgumentAllowed = true

	l.flagsIncompatibleWithConfigFile.Insert(
		"nodes",
		"nodes-min",
		"nodes-max",
	)

	l.validateWithConfigFile = func() error {
		if err := validateNameArgument(cmd, ng); err != nil {
			return err
		}

		ngFilter := NewNodeGroupFilter()
		ngFilter.AppendIncludeNames(ng.Name)

		matchedNodeGroup := ngFilter.FilterMatching(l.ClusterConfig.NodeGroups)
		if len(matchedNodeGroup) == 0 {
			return fmt.Errorf("node group %s not found", ng.Name)
		}

		if err := validateNumberOfNodes(matchedNodeGroup[0]); err != nil {
			return err
		}
		ng = matchedNodeGroup[0]
		l.Plan = false
		return nil
	}

	l.validateWithoutConfigFile = func() error {
		if l.ClusterConfig.Metadata.Name == "" {
			return ErrMustBeSet(ClusterNameFlag(cmd))
		}

		if err := validateNameArgument(cmd, ng); err != nil {
			return err
		}

		if err := validateNumberOfNodes(ng); err != nil {
			return err
		}
		l.Plan = false
		return nil
	}

	return l
}

func validateNameArgument(cmd *Cmd, ng *api.NodeGroup) error {
	if ng.Name != "" && cmd.NameArg != "" {
		return ErrFlagAndArg("--name", ng.Name, cmd.NameArg)
	}

	if cmd.NameArg != "" {
		ng.Name = cmd.NameArg
	}

	if ng.Name == "" {
		return ErrMustBeSet("--name")
	}

	return nil
}

func validateNumberOfNodes(ng *api.NodeGroup) error {

	if ng.DesiredCapacity == nil || *ng.DesiredCapacity < 0 {
		return fmt.Errorf("number of nodes must be 0 or greater")
	}

	if ng.MaxSize != nil && *ng.MaxSize < 0 {
		return fmt.Errorf("maximum number of nodes must be 0 or greater")
	}

	if ng.MaxSize != nil && ng.MinSize != nil && (*ng.MinSize > *ng.DesiredCapacity || *ng.MaxSize < *ng.DesiredCapacity) {
		return fmt.Errorf("number of nodes must be within range of min nodes and max nodes")
	}

	if ng.MaxSize != nil && *ng.MaxSize < *ng.DesiredCapacity {
		return fmt.Errorf("maximum number of nodes must be greater than or equal to number of nodes")
	}

	if ng.MinSize != nil && *ng.MinSize < 0 {
		return fmt.Errorf("minimum number of nodes must be 0 or greater")
	}

	if ng.MinSize != nil && *ng.MinSize > *ng.DesiredCapacity {
		return fmt.Errorf("minimum number of nodes must be less than or equal to number of nodes")
	}

	return nil
}