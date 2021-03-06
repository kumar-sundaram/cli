// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeTasksActor struct {
	GetApplicationByNameAndSpaceStub        func(string, string) (v7action.Application, v7action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	GetApplicationTasksStub        func(string, v7action.SortOrder) ([]v7action.Task, v7action.Warnings, error)
	getApplicationTasksMutex       sync.RWMutex
	getApplicationTasksArgsForCall []struct {
		arg1 string
		arg2 v7action.SortOrder
	}
	getApplicationTasksReturns struct {
		result1 []v7action.Task
		result2 v7action.Warnings
		result3 error
	}
	getApplicationTasksReturnsOnCall map[int]struct {
		result1 []v7action.Task
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v7action.Application, v7action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) GetApplicationTasks(arg1 string, arg2 v7action.SortOrder) ([]v7action.Task, v7action.Warnings, error) {
	fake.getApplicationTasksMutex.Lock()
	ret, specificReturn := fake.getApplicationTasksReturnsOnCall[len(fake.getApplicationTasksArgsForCall)]
	fake.getApplicationTasksArgsForCall = append(fake.getApplicationTasksArgsForCall, struct {
		arg1 string
		arg2 v7action.SortOrder
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationTasks", []interface{}{arg1, arg2})
	fake.getApplicationTasksMutex.Unlock()
	if fake.GetApplicationTasksStub != nil {
		return fake.GetApplicationTasksStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationTasksReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeTasksActor) GetApplicationTasksCallCount() int {
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	return len(fake.getApplicationTasksArgsForCall)
}

func (fake *FakeTasksActor) GetApplicationTasksCalls(stub func(string, v7action.SortOrder) ([]v7action.Task, v7action.Warnings, error)) {
	fake.getApplicationTasksMutex.Lock()
	defer fake.getApplicationTasksMutex.Unlock()
	fake.GetApplicationTasksStub = stub
}

func (fake *FakeTasksActor) GetApplicationTasksArgsForCall(i int) (string, v7action.SortOrder) {
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	argsForCall := fake.getApplicationTasksArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTasksActor) GetApplicationTasksReturns(result1 []v7action.Task, result2 v7action.Warnings, result3 error) {
	fake.getApplicationTasksMutex.Lock()
	defer fake.getApplicationTasksMutex.Unlock()
	fake.GetApplicationTasksStub = nil
	fake.getApplicationTasksReturns = struct {
		result1 []v7action.Task
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) GetApplicationTasksReturnsOnCall(i int, result1 []v7action.Task, result2 v7action.Warnings, result3 error) {
	fake.getApplicationTasksMutex.Lock()
	defer fake.getApplicationTasksMutex.Unlock()
	fake.GetApplicationTasksStub = nil
	if fake.getApplicationTasksReturnsOnCall == nil {
		fake.getApplicationTasksReturnsOnCall = make(map[int]struct {
			result1 []v7action.Task
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationTasksReturnsOnCall[i] = struct {
		result1 []v7action.Task
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTasksActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationTasksMutex.RLock()
	defer fake.getApplicationTasksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTasksActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.TasksActor = new(FakeTasksActor)
