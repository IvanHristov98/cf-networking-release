// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/handlers"
	"code.cloudfoundry.org/policy-server/store"
)

type EgressDestinationStoreDeleter struct {
	DeleteStub        func(...string) ([]store.EgressDestination, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []string
	}
	deleteReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressDestinationStoreDeleter) Delete(arg1 ...string) ([]store.EgressDestination, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressDestinationStoreDeleter) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *EgressDestinationStoreDeleter) DeleteCalls(stub func(...string) ([]store.EgressDestination, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *EgressDestinationStoreDeleter) DeleteArgsForCall(i int) []string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressDestinationStoreDeleter) DeleteReturns(result1 []store.EgressDestination, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreDeleter) DeleteReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreDeleter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressDestinationStoreDeleter) recordInvocation(key string, args []interface{}) {
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

var _ handlers.EgressDestinationStoreDeleter = new(EgressDestinationStoreDeleter)
