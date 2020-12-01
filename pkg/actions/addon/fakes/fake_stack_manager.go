// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/weaveworks/eksctl/pkg/actions/addon"
	"github.com/weaveworks/eksctl/pkg/cfn/builder"
	"github.com/weaveworks/eksctl/pkg/cfn/manager"
)

type FakeStackManager struct {
	CreateStackStub        func(string, builder.ResourceSet, map[string]string, map[string]string, chan error) error
	createStackMutex       sync.RWMutex
	createStackArgsForCall []struct {
		arg1 string
		arg2 builder.ResourceSet
		arg3 map[string]string
		arg4 map[string]string
		arg5 chan error
	}
	createStackReturns struct {
		result1 error
	}
	createStackReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStackByNameStub        func(string) (*cloudformation.Stack, error)
	deleteStackByNameMutex       sync.RWMutex
	deleteStackByNameArgsForCall []struct {
		arg1 string
	}
	deleteStackByNameReturns struct {
		result1 *cloudformation.Stack
		result2 error
	}
	deleteStackByNameReturnsOnCall map[int]struct {
		result1 *cloudformation.Stack
		result2 error
	}
	ListStacksMatchingStub        func(string, ...string) ([]*cloudformation.Stack, error)
	listStacksMatchingMutex       sync.RWMutex
	listStacksMatchingArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	listStacksMatchingReturns struct {
		result1 []*cloudformation.Stack
		result2 error
	}
	listStacksMatchingReturnsOnCall map[int]struct {
		result1 []*cloudformation.Stack
		result2 error
	}
	UpdateStackStub        func(string, string, string, manager.TemplateData, map[string]string) error
	updateStackMutex       sync.RWMutex
	updateStackArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 manager.TemplateData
		arg5 map[string]string
	}
	updateStackReturns struct {
		result1 error
	}
	updateStackReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStackManager) CreateStack(arg1 string, arg2 builder.ResourceSet, arg3 map[string]string, arg4 map[string]string, arg5 chan error) error {
	fake.createStackMutex.Lock()
	ret, specificReturn := fake.createStackReturnsOnCall[len(fake.createStackArgsForCall)]
	fake.createStackArgsForCall = append(fake.createStackArgsForCall, struct {
		arg1 string
		arg2 builder.ResourceSet
		arg3 map[string]string
		arg4 map[string]string
		arg5 chan error
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("CreateStack", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createStackMutex.Unlock()
	if fake.CreateStackStub != nil {
		return fake.CreateStackStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createStackReturns
	return fakeReturns.result1
}

func (fake *FakeStackManager) CreateStackCallCount() int {
	fake.createStackMutex.RLock()
	defer fake.createStackMutex.RUnlock()
	return len(fake.createStackArgsForCall)
}

func (fake *FakeStackManager) CreateStackCalls(stub func(string, builder.ResourceSet, map[string]string, map[string]string, chan error) error) {
	fake.createStackMutex.Lock()
	defer fake.createStackMutex.Unlock()
	fake.CreateStackStub = stub
}

func (fake *FakeStackManager) CreateStackArgsForCall(i int) (string, builder.ResourceSet, map[string]string, map[string]string, chan error) {
	fake.createStackMutex.RLock()
	defer fake.createStackMutex.RUnlock()
	argsForCall := fake.createStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeStackManager) CreateStackReturns(result1 error) {
	fake.createStackMutex.Lock()
	defer fake.createStackMutex.Unlock()
	fake.CreateStackStub = nil
	fake.createStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) CreateStackReturnsOnCall(i int, result1 error) {
	fake.createStackMutex.Lock()
	defer fake.createStackMutex.Unlock()
	fake.CreateStackStub = nil
	if fake.createStackReturnsOnCall == nil {
		fake.createStackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createStackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) DeleteStackByName(arg1 string) (*cloudformation.Stack, error) {
	fake.deleteStackByNameMutex.Lock()
	ret, specificReturn := fake.deleteStackByNameReturnsOnCall[len(fake.deleteStackByNameArgsForCall)]
	fake.deleteStackByNameArgsForCall = append(fake.deleteStackByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteStackByName", []interface{}{arg1})
	fake.deleteStackByNameMutex.Unlock()
	if fake.DeleteStackByNameStub != nil {
		return fake.DeleteStackByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteStackByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackManager) DeleteStackByNameCallCount() int {
	fake.deleteStackByNameMutex.RLock()
	defer fake.deleteStackByNameMutex.RUnlock()
	return len(fake.deleteStackByNameArgsForCall)
}

func (fake *FakeStackManager) DeleteStackByNameCalls(stub func(string) (*cloudformation.Stack, error)) {
	fake.deleteStackByNameMutex.Lock()
	defer fake.deleteStackByNameMutex.Unlock()
	fake.DeleteStackByNameStub = stub
}

func (fake *FakeStackManager) DeleteStackByNameArgsForCall(i int) string {
	fake.deleteStackByNameMutex.RLock()
	defer fake.deleteStackByNameMutex.RUnlock()
	argsForCall := fake.deleteStackByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStackManager) DeleteStackByNameReturns(result1 *cloudformation.Stack, result2 error) {
	fake.deleteStackByNameMutex.Lock()
	defer fake.deleteStackByNameMutex.Unlock()
	fake.DeleteStackByNameStub = nil
	fake.deleteStackByNameReturns = struct {
		result1 *cloudformation.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) DeleteStackByNameReturnsOnCall(i int, result1 *cloudformation.Stack, result2 error) {
	fake.deleteStackByNameMutex.Lock()
	defer fake.deleteStackByNameMutex.Unlock()
	fake.DeleteStackByNameStub = nil
	if fake.deleteStackByNameReturnsOnCall == nil {
		fake.deleteStackByNameReturnsOnCall = make(map[int]struct {
			result1 *cloudformation.Stack
			result2 error
		})
	}
	fake.deleteStackByNameReturnsOnCall[i] = struct {
		result1 *cloudformation.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) ListStacksMatching(arg1 string, arg2 ...string) ([]*cloudformation.Stack, error) {
	fake.listStacksMatchingMutex.Lock()
	ret, specificReturn := fake.listStacksMatchingReturnsOnCall[len(fake.listStacksMatchingArgsForCall)]
	fake.listStacksMatchingArgsForCall = append(fake.listStacksMatchingArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	fake.recordInvocation("ListStacksMatching", []interface{}{arg1, arg2})
	fake.listStacksMatchingMutex.Unlock()
	if fake.ListStacksMatchingStub != nil {
		return fake.ListStacksMatchingStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStacksMatchingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStackManager) ListStacksMatchingCallCount() int {
	fake.listStacksMatchingMutex.RLock()
	defer fake.listStacksMatchingMutex.RUnlock()
	return len(fake.listStacksMatchingArgsForCall)
}

func (fake *FakeStackManager) ListStacksMatchingCalls(stub func(string, ...string) ([]*cloudformation.Stack, error)) {
	fake.listStacksMatchingMutex.Lock()
	defer fake.listStacksMatchingMutex.Unlock()
	fake.ListStacksMatchingStub = stub
}

func (fake *FakeStackManager) ListStacksMatchingArgsForCall(i int) (string, []string) {
	fake.listStacksMatchingMutex.RLock()
	defer fake.listStacksMatchingMutex.RUnlock()
	argsForCall := fake.listStacksMatchingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStackManager) ListStacksMatchingReturns(result1 []*cloudformation.Stack, result2 error) {
	fake.listStacksMatchingMutex.Lock()
	defer fake.listStacksMatchingMutex.Unlock()
	fake.ListStacksMatchingStub = nil
	fake.listStacksMatchingReturns = struct {
		result1 []*cloudformation.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) ListStacksMatchingReturnsOnCall(i int, result1 []*cloudformation.Stack, result2 error) {
	fake.listStacksMatchingMutex.Lock()
	defer fake.listStacksMatchingMutex.Unlock()
	fake.ListStacksMatchingStub = nil
	if fake.listStacksMatchingReturnsOnCall == nil {
		fake.listStacksMatchingReturnsOnCall = make(map[int]struct {
			result1 []*cloudformation.Stack
			result2 error
		})
	}
	fake.listStacksMatchingReturnsOnCall[i] = struct {
		result1 []*cloudformation.Stack
		result2 error
	}{result1, result2}
}

func (fake *FakeStackManager) UpdateStack(arg1 string, arg2 string, arg3 string, arg4 manager.TemplateData, arg5 map[string]string) error {
	fake.updateStackMutex.Lock()
	ret, specificReturn := fake.updateStackReturnsOnCall[len(fake.updateStackArgsForCall)]
	fake.updateStackArgsForCall = append(fake.updateStackArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 manager.TemplateData
		arg5 map[string]string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("UpdateStack", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.updateStackMutex.Unlock()
	if fake.UpdateStackStub != nil {
		return fake.UpdateStackStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateStackReturns
	return fakeReturns.result1
}

func (fake *FakeStackManager) UpdateStackCallCount() int {
	fake.updateStackMutex.RLock()
	defer fake.updateStackMutex.RUnlock()
	return len(fake.updateStackArgsForCall)
}

func (fake *FakeStackManager) UpdateStackCalls(stub func(string, string, string, manager.TemplateData, map[string]string) error) {
	fake.updateStackMutex.Lock()
	defer fake.updateStackMutex.Unlock()
	fake.UpdateStackStub = stub
}

func (fake *FakeStackManager) UpdateStackArgsForCall(i int) (string, string, string, manager.TemplateData, map[string]string) {
	fake.updateStackMutex.RLock()
	defer fake.updateStackMutex.RUnlock()
	argsForCall := fake.updateStackArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeStackManager) UpdateStackReturns(result1 error) {
	fake.updateStackMutex.Lock()
	defer fake.updateStackMutex.Unlock()
	fake.UpdateStackStub = nil
	fake.updateStackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) UpdateStackReturnsOnCall(i int, result1 error) {
	fake.updateStackMutex.Lock()
	defer fake.updateStackMutex.Unlock()
	fake.UpdateStackStub = nil
	if fake.updateStackReturnsOnCall == nil {
		fake.updateStackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStackManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createStackMutex.RLock()
	defer fake.createStackMutex.RUnlock()
	fake.deleteStackByNameMutex.RLock()
	defer fake.deleteStackByNameMutex.RUnlock()
	fake.listStacksMatchingMutex.RLock()
	defer fake.listStacksMatchingMutex.RUnlock()
	fake.updateStackMutex.RLock()
	defer fake.updateStackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStackManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ addon.StackManager = new(FakeStackManager)
