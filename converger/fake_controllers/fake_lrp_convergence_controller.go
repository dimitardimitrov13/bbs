// Code generated by counterfeiter. DO NOT EDIT.
package fake_controllers

import (
	"sync"

	"code.cloudfoundry.org/bbs/converger"
	"code.cloudfoundry.org/lager"
)

type FakeLrpConvergenceController struct {
	ConvergeLRPsStub        func(logger lager.Logger)
	convergeLRPsMutex       sync.RWMutex
	convergeLRPsArgsForCall []struct {
		logger lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLrpConvergenceController) ConvergeLRPs(logger lager.Logger) {
	fake.convergeLRPsMutex.Lock()
	fake.convergeLRPsArgsForCall = append(fake.convergeLRPsArgsForCall, struct {
		logger lager.Logger
	}{logger})
	fake.recordInvocation("ConvergeLRPs", []interface{}{logger})
	fake.convergeLRPsMutex.Unlock()
	if fake.ConvergeLRPsStub != nil {
		fake.ConvergeLRPsStub(logger)
	}
}

func (fake *FakeLrpConvergenceController) ConvergeLRPsCallCount() int {
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	return len(fake.convergeLRPsArgsForCall)
}

func (fake *FakeLrpConvergenceController) ConvergeLRPsArgsForCall(i int) lager.Logger {
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	return fake.convergeLRPsArgsForCall[i].logger
}

func (fake *FakeLrpConvergenceController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLrpConvergenceController) recordInvocation(key string, args []interface{}) {
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

var _ converger.LrpConvergenceController = new(FakeLrpConvergenceController)
