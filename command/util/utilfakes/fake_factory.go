// Code generated by counterfeiter. DO NOT EDIT.
package utilfakes

import (
	sync "sync"

	opkv "github.com/cappyzawa/op-kv"
	util "github.com/cappyzawa/op-kv/command/util"
)

type FakeFactory struct {
	CommandRunnerStub        func() opkv.Runner
	commandRunnerMutex       sync.RWMutex
	commandRunnerArgsForCall []struct {
	}
	commandRunnerReturns struct {
		result1 opkv.Runner
	}
	commandRunnerReturnsOnCall map[int]struct {
		result1 opkv.Runner
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFactory) CommandRunner() opkv.Runner {
	fake.commandRunnerMutex.Lock()
	ret, specificReturn := fake.commandRunnerReturnsOnCall[len(fake.commandRunnerArgsForCall)]
	fake.commandRunnerArgsForCall = append(fake.commandRunnerArgsForCall, struct {
	}{})
	fake.recordInvocation("CommandRunner", []interface{}{})
	fake.commandRunnerMutex.Unlock()
	if fake.CommandRunnerStub != nil {
		return fake.CommandRunnerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.commandRunnerReturns
	return fakeReturns.result1
}

func (fake *FakeFactory) CommandRunnerCallCount() int {
	fake.commandRunnerMutex.RLock()
	defer fake.commandRunnerMutex.RUnlock()
	return len(fake.commandRunnerArgsForCall)
}

func (fake *FakeFactory) CommandRunnerCalls(stub func() opkv.Runner) {
	fake.commandRunnerMutex.Lock()
	defer fake.commandRunnerMutex.Unlock()
	fake.CommandRunnerStub = stub
}

func (fake *FakeFactory) CommandRunnerReturns(result1 opkv.Runner) {
	fake.commandRunnerMutex.Lock()
	defer fake.commandRunnerMutex.Unlock()
	fake.CommandRunnerStub = nil
	fake.commandRunnerReturns = struct {
		result1 opkv.Runner
	}{result1}
}

func (fake *FakeFactory) CommandRunnerReturnsOnCall(i int, result1 opkv.Runner) {
	fake.commandRunnerMutex.Lock()
	defer fake.commandRunnerMutex.Unlock()
	fake.CommandRunnerStub = nil
	if fake.commandRunnerReturnsOnCall == nil {
		fake.commandRunnerReturnsOnCall = make(map[int]struct {
			result1 opkv.Runner
		})
	}
	fake.commandRunnerReturnsOnCall[i] = struct {
		result1 opkv.Runner
	}{result1}
}

func (fake *FakeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commandRunnerMutex.RLock()
	defer fake.commandRunnerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFactory) recordInvocation(key string, args []interface{}) {
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

var _ util.Factory = new(FakeFactory)