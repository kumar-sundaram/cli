// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeRenameBuildpackActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	RenameBuildpackStub        func(oldName string, newName string, stackName string) (v2action.Warnings, error)
	renameBuildpackMutex       sync.RWMutex
	renameBuildpackArgsForCall []struct {
		oldName   string
		newName   string
		stackName string
	}
	renameBuildpackReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	renameBuildpackReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRenameBuildpackActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeRenameBuildpackActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeRenameBuildpackActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRenameBuildpackActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRenameBuildpackActor) RenameBuildpack(oldName string, newName string, stackName string) (v2action.Warnings, error) {
	fake.renameBuildpackMutex.Lock()
	ret, specificReturn := fake.renameBuildpackReturnsOnCall[len(fake.renameBuildpackArgsForCall)]
	fake.renameBuildpackArgsForCall = append(fake.renameBuildpackArgsForCall, struct {
		oldName   string
		newName   string
		stackName string
	}{oldName, newName, stackName})
	fake.recordInvocation("RenameBuildpack", []interface{}{oldName, newName, stackName})
	fake.renameBuildpackMutex.Unlock()
	if fake.RenameBuildpackStub != nil {
		return fake.RenameBuildpackStub(oldName, newName, stackName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.renameBuildpackReturns.result1, fake.renameBuildpackReturns.result2
}

func (fake *FakeRenameBuildpackActor) RenameBuildpackCallCount() int {
	fake.renameBuildpackMutex.RLock()
	defer fake.renameBuildpackMutex.RUnlock()
	return len(fake.renameBuildpackArgsForCall)
}

func (fake *FakeRenameBuildpackActor) RenameBuildpackArgsForCall(i int) (string, string, string) {
	fake.renameBuildpackMutex.RLock()
	defer fake.renameBuildpackMutex.RUnlock()
	return fake.renameBuildpackArgsForCall[i].oldName, fake.renameBuildpackArgsForCall[i].newName, fake.renameBuildpackArgsForCall[i].stackName
}

func (fake *FakeRenameBuildpackActor) RenameBuildpackReturns(result1 v2action.Warnings, result2 error) {
	fake.RenameBuildpackStub = nil
	fake.renameBuildpackReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeRenameBuildpackActor) RenameBuildpackReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.RenameBuildpackStub = nil
	if fake.renameBuildpackReturnsOnCall == nil {
		fake.renameBuildpackReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.renameBuildpackReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeRenameBuildpackActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.renameBuildpackMutex.RLock()
	defer fake.renameBuildpackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRenameBuildpackActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.RenameBuildpackActor = new(FakeRenameBuildpackActor)