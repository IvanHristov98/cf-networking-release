// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/policy-server/handlers"
	"code.cloudfoundry.org/policy-server/store"
)

type EgressDestinationStoreLister struct {
	AllStub        func() ([]store.EgressDestination, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct {
	}
	allReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	GetByGUIDStub        func(...string) ([]store.EgressDestination, error)
	getByGUIDMutex       sync.RWMutex
	getByGUIDArgsForCall []struct {
		arg1 []string
	}
	getByGUIDReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	getByGUIDReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	GetByNameStub        func(...string) ([]store.EgressDestination, error)
	getByNameMutex       sync.RWMutex
	getByNameArgsForCall []struct {
		arg1 []string
	}
	getByNameReturns struct {
		result1 []store.EgressDestination
		result2 error
	}
	getByNameReturnsOnCall map[int]struct {
		result1 []store.EgressDestination
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressDestinationStoreLister) All() ([]store.EgressDestination, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct {
	}{})
	stub := fake.AllStub
	fakeReturns := fake.allReturns
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressDestinationStoreLister) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *EgressDestinationStoreLister) AllCalls(stub func() ([]store.EgressDestination, error)) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = stub
}

func (fake *EgressDestinationStoreLister) AllReturns(result1 []store.EgressDestination, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreLister) AllReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.allMutex.Lock()
	defer fake.allMutex.Unlock()
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreLister) GetByGUID(arg1 ...string) ([]store.EgressDestination, error) {
	fake.getByGUIDMutex.Lock()
	ret, specificReturn := fake.getByGUIDReturnsOnCall[len(fake.getByGUIDArgsForCall)]
	fake.getByGUIDArgsForCall = append(fake.getByGUIDArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.GetByGUIDStub
	fakeReturns := fake.getByGUIDReturns
	fake.recordInvocation("GetByGUID", []interface{}{arg1})
	fake.getByGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressDestinationStoreLister) GetByGUIDCallCount() int {
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	return len(fake.getByGUIDArgsForCall)
}

func (fake *EgressDestinationStoreLister) GetByGUIDCalls(stub func(...string) ([]store.EgressDestination, error)) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = stub
}

func (fake *EgressDestinationStoreLister) GetByGUIDArgsForCall(i int) []string {
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	argsForCall := fake.getByGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressDestinationStoreLister) GetByGUIDReturns(result1 []store.EgressDestination, result2 error) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = nil
	fake.getByGUIDReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreLister) GetByGUIDReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = nil
	if fake.getByGUIDReturnsOnCall == nil {
		fake.getByGUIDReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.getByGUIDReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreLister) GetByName(arg1 ...string) ([]store.EgressDestination, error) {
	fake.getByNameMutex.Lock()
	ret, specificReturn := fake.getByNameReturnsOnCall[len(fake.getByNameArgsForCall)]
	fake.getByNameArgsForCall = append(fake.getByNameArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.GetByNameStub
	fakeReturns := fake.getByNameReturns
	fake.recordInvocation("GetByName", []interface{}{arg1})
	fake.getByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressDestinationStoreLister) GetByNameCallCount() int {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	return len(fake.getByNameArgsForCall)
}

func (fake *EgressDestinationStoreLister) GetByNameCalls(stub func(...string) ([]store.EgressDestination, error)) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = stub
}

func (fake *EgressDestinationStoreLister) GetByNameArgsForCall(i int) []string {
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	argsForCall := fake.getByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressDestinationStoreLister) GetByNameReturns(result1 []store.EgressDestination, result2 error) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = nil
	fake.getByNameReturns = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreLister) GetByNameReturnsOnCall(i int, result1 []store.EgressDestination, result2 error) {
	fake.getByNameMutex.Lock()
	defer fake.getByNameMutex.Unlock()
	fake.GetByNameStub = nil
	if fake.getByNameReturnsOnCall == nil {
		fake.getByNameReturnsOnCall = make(map[int]struct {
			result1 []store.EgressDestination
			result2 error
		})
	}
	fake.getByNameReturnsOnCall[i] = struct {
		result1 []store.EgressDestination
		result2 error
	}{result1, result2}
}

func (fake *EgressDestinationStoreLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	fake.getByNameMutex.RLock()
	defer fake.getByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressDestinationStoreLister) recordInvocation(key string, args []interface{}) {
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

var _ handlers.EgressDestinationStoreLister = new(EgressDestinationStoreLister)
