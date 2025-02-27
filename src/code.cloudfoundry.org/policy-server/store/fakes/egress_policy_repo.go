// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/cf-networking-helpers/db"
	"code.cloudfoundry.org/policy-server/store"
)

type EgressPolicyRepo struct {
	CreateAppStub        func(db.Transaction, string, string) error
	createAppMutex       sync.RWMutex
	createAppArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
		arg3 string
	}
	createAppReturns struct {
		result1 error
	}
	createAppReturnsOnCall map[int]struct {
		result1 error
	}
	CreateDefaultStub        func(db.Transaction, string) error
	createDefaultMutex       sync.RWMutex
	createDefaultArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	createDefaultReturns struct {
		result1 error
	}
	createDefaultReturnsOnCall map[int]struct {
		result1 error
	}
	CreateEgressPolicyStub        func(db.Transaction, string, string, string) (string, error)
	createEgressPolicyMutex       sync.RWMutex
	createEgressPolicyArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
		arg3 string
		arg4 string
	}
	createEgressPolicyReturns struct {
		result1 string
		result2 error
	}
	createEgressPolicyReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	CreateSpaceStub        func(db.Transaction, string, string) error
	createSpaceMutex       sync.RWMutex
	createSpaceArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
		arg3 string
	}
	createSpaceReturns struct {
		result1 error
	}
	createSpaceReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteAppStub        func(db.Transaction, string) error
	deleteAppMutex       sync.RWMutex
	deleteAppArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	deleteAppReturns struct {
		result1 error
	}
	deleteAppReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteDefaultStub        func(db.Transaction, string) error
	deleteDefaultMutex       sync.RWMutex
	deleteDefaultArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	deleteDefaultReturns struct {
		result1 error
	}
	deleteDefaultReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteEgressPolicyStub        func(db.Transaction, string) error
	deleteEgressPolicyMutex       sync.RWMutex
	deleteEgressPolicyArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	deleteEgressPolicyReturns struct {
		result1 error
	}
	deleteEgressPolicyReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteSpaceStub        func(db.Transaction, string) error
	deleteSpaceMutex       sync.RWMutex
	deleteSpaceArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	deleteSpaceReturns struct {
		result1 error
	}
	deleteSpaceReturnsOnCall map[int]struct {
		result1 error
	}
	GetAllPoliciesStub        func() ([]store.EgressPolicy, error)
	getAllPoliciesMutex       sync.RWMutex
	getAllPoliciesArgsForCall []struct {
	}
	getAllPoliciesReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getAllPoliciesReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetByFilterStub        func([]string, []string, []string, []string, []string) ([]store.EgressPolicy, error)
	getByFilterMutex       sync.RWMutex
	getByFilterArgsForCall []struct {
		arg1 []string
		arg2 []string
		arg3 []string
		arg4 []string
		arg5 []string
	}
	getByFilterReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getByFilterReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetByGUIDStub        func(db.Transaction, ...string) ([]store.EgressPolicy, error)
	getByGUIDMutex       sync.RWMutex
	getByGUIDArgsForCall []struct {
		arg1 db.Transaction
		arg2 []string
	}
	getByGUIDReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getByGUIDReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetBySourceGuidsAndDefaultsStub        func([]string) ([]store.EgressPolicy, error)
	getBySourceGuidsAndDefaultsMutex       sync.RWMutex
	getBySourceGuidsAndDefaultsArgsForCall []struct {
		arg1 []string
	}
	getBySourceGuidsAndDefaultsReturns struct {
		result1 []store.EgressPolicy
		result2 error
	}
	getBySourceGuidsAndDefaultsReturnsOnCall map[int]struct {
		result1 []store.EgressPolicy
		result2 error
	}
	GetTerminalByAppGUIDStub        func(db.Transaction, string) (string, error)
	getTerminalByAppGUIDMutex       sync.RWMutex
	getTerminalByAppGUIDArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	getTerminalByAppGUIDReturns struct {
		result1 string
		result2 error
	}
	getTerminalByAppGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetTerminalBySpaceGUIDStub        func(db.Transaction, string) (string, error)
	getTerminalBySpaceGUIDMutex       sync.RWMutex
	getTerminalBySpaceGUIDArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	getTerminalBySpaceGUIDReturns struct {
		result1 string
		result2 error
	}
	getTerminalBySpaceGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	IsTerminalInUseStub        func(db.Transaction, string) (bool, error)
	isTerminalInUseMutex       sync.RWMutex
	isTerminalInUseArgsForCall []struct {
		arg1 db.Transaction
		arg2 string
	}
	isTerminalInUseReturns struct {
		result1 bool
		result2 error
	}
	isTerminalInUseReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EgressPolicyRepo) CreateApp(arg1 db.Transaction, arg2 string, arg3 string) error {
	fake.createAppMutex.Lock()
	ret, specificReturn := fake.createAppReturnsOnCall[len(fake.createAppArgsForCall)]
	fake.createAppArgsForCall = append(fake.createAppArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateAppStub
	fakeReturns := fake.createAppReturns
	fake.recordInvocation("CreateApp", []interface{}{arg1, arg2, arg3})
	fake.createAppMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressPolicyRepo) CreateAppCallCount() int {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return len(fake.createAppArgsForCall)
}

func (fake *EgressPolicyRepo) CreateAppCalls(stub func(db.Transaction, string, string) error) {
	fake.createAppMutex.Lock()
	defer fake.createAppMutex.Unlock()
	fake.CreateAppStub = stub
}

func (fake *EgressPolicyRepo) CreateAppArgsForCall(i int) (db.Transaction, string, string) {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	argsForCall := fake.createAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *EgressPolicyRepo) CreateAppReturns(result1 error) {
	fake.createAppMutex.Lock()
	defer fake.createAppMutex.Unlock()
	fake.CreateAppStub = nil
	fake.createAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) CreateAppReturnsOnCall(i int, result1 error) {
	fake.createAppMutex.Lock()
	defer fake.createAppMutex.Unlock()
	fake.CreateAppStub = nil
	if fake.createAppReturnsOnCall == nil {
		fake.createAppReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createAppReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) CreateDefault(arg1 db.Transaction, arg2 string) error {
	fake.createDefaultMutex.Lock()
	ret, specificReturn := fake.createDefaultReturnsOnCall[len(fake.createDefaultArgsForCall)]
	fake.createDefaultArgsForCall = append(fake.createDefaultArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.CreateDefaultStub
	fakeReturns := fake.createDefaultReturns
	fake.recordInvocation("CreateDefault", []interface{}{arg1, arg2})
	fake.createDefaultMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressPolicyRepo) CreateDefaultCallCount() int {
	fake.createDefaultMutex.RLock()
	defer fake.createDefaultMutex.RUnlock()
	return len(fake.createDefaultArgsForCall)
}

func (fake *EgressPolicyRepo) CreateDefaultCalls(stub func(db.Transaction, string) error) {
	fake.createDefaultMutex.Lock()
	defer fake.createDefaultMutex.Unlock()
	fake.CreateDefaultStub = stub
}

func (fake *EgressPolicyRepo) CreateDefaultArgsForCall(i int) (db.Transaction, string) {
	fake.createDefaultMutex.RLock()
	defer fake.createDefaultMutex.RUnlock()
	argsForCall := fake.createDefaultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) CreateDefaultReturns(result1 error) {
	fake.createDefaultMutex.Lock()
	defer fake.createDefaultMutex.Unlock()
	fake.CreateDefaultStub = nil
	fake.createDefaultReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) CreateDefaultReturnsOnCall(i int, result1 error) {
	fake.createDefaultMutex.Lock()
	defer fake.createDefaultMutex.Unlock()
	fake.CreateDefaultStub = nil
	if fake.createDefaultReturnsOnCall == nil {
		fake.createDefaultReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createDefaultReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) CreateEgressPolicy(arg1 db.Transaction, arg2 string, arg3 string, arg4 string) (string, error) {
	fake.createEgressPolicyMutex.Lock()
	ret, specificReturn := fake.createEgressPolicyReturnsOnCall[len(fake.createEgressPolicyArgsForCall)]
	fake.createEgressPolicyArgsForCall = append(fake.createEgressPolicyArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.CreateEgressPolicyStub
	fakeReturns := fake.createEgressPolicyReturns
	fake.recordInvocation("CreateEgressPolicy", []interface{}{arg1, arg2, arg3, arg4})
	fake.createEgressPolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) CreateEgressPolicyCallCount() int {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	return len(fake.createEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) CreateEgressPolicyCalls(stub func(db.Transaction, string, string, string) (string, error)) {
	fake.createEgressPolicyMutex.Lock()
	defer fake.createEgressPolicyMutex.Unlock()
	fake.CreateEgressPolicyStub = stub
}

func (fake *EgressPolicyRepo) CreateEgressPolicyArgsForCall(i int) (db.Transaction, string, string, string) {
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	argsForCall := fake.createEgressPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturns(result1 string, result2 error) {
	fake.createEgressPolicyMutex.Lock()
	defer fake.createEgressPolicyMutex.Unlock()
	fake.CreateEgressPolicyStub = nil
	fake.createEgressPolicyReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateEgressPolicyReturnsOnCall(i int, result1 string, result2 error) {
	fake.createEgressPolicyMutex.Lock()
	defer fake.createEgressPolicyMutex.Unlock()
	fake.CreateEgressPolicyStub = nil
	if fake.createEgressPolicyReturnsOnCall == nil {
		fake.createEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.createEgressPolicyReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) CreateSpace(arg1 db.Transaction, arg2 string, arg3 string) error {
	fake.createSpaceMutex.Lock()
	ret, specificReturn := fake.createSpaceReturnsOnCall[len(fake.createSpaceArgsForCall)]
	fake.createSpaceArgsForCall = append(fake.createSpaceArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateSpaceStub
	fakeReturns := fake.createSpaceReturns
	fake.recordInvocation("CreateSpace", []interface{}{arg1, arg2, arg3})
	fake.createSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressPolicyRepo) CreateSpaceCallCount() int {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	return len(fake.createSpaceArgsForCall)
}

func (fake *EgressPolicyRepo) CreateSpaceCalls(stub func(db.Transaction, string, string) error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = stub
}

func (fake *EgressPolicyRepo) CreateSpaceArgsForCall(i int) (db.Transaction, string, string) {
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	argsForCall := fake.createSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *EgressPolicyRepo) CreateSpaceReturns(result1 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	fake.createSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) CreateSpaceReturnsOnCall(i int, result1 error) {
	fake.createSpaceMutex.Lock()
	defer fake.createSpaceMutex.Unlock()
	fake.CreateSpaceStub = nil
	if fake.createSpaceReturnsOnCall == nil {
		fake.createSpaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createSpaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteApp(arg1 db.Transaction, arg2 string) error {
	fake.deleteAppMutex.Lock()
	ret, specificReturn := fake.deleteAppReturnsOnCall[len(fake.deleteAppArgsForCall)]
	fake.deleteAppArgsForCall = append(fake.deleteAppArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteAppStub
	fakeReturns := fake.deleteAppReturns
	fake.recordInvocation("DeleteApp", []interface{}{arg1, arg2})
	fake.deleteAppMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressPolicyRepo) DeleteAppCallCount() int {
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	return len(fake.deleteAppArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteAppCalls(stub func(db.Transaction, string) error) {
	fake.deleteAppMutex.Lock()
	defer fake.deleteAppMutex.Unlock()
	fake.DeleteAppStub = stub
}

func (fake *EgressPolicyRepo) DeleteAppArgsForCall(i int) (db.Transaction, string) {
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	argsForCall := fake.deleteAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) DeleteAppReturns(result1 error) {
	fake.deleteAppMutex.Lock()
	defer fake.deleteAppMutex.Unlock()
	fake.DeleteAppStub = nil
	fake.deleteAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteAppReturnsOnCall(i int, result1 error) {
	fake.deleteAppMutex.Lock()
	defer fake.deleteAppMutex.Unlock()
	fake.DeleteAppStub = nil
	if fake.deleteAppReturnsOnCall == nil {
		fake.deleteAppReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAppReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteDefault(arg1 db.Transaction, arg2 string) error {
	fake.deleteDefaultMutex.Lock()
	ret, specificReturn := fake.deleteDefaultReturnsOnCall[len(fake.deleteDefaultArgsForCall)]
	fake.deleteDefaultArgsForCall = append(fake.deleteDefaultArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteDefaultStub
	fakeReturns := fake.deleteDefaultReturns
	fake.recordInvocation("DeleteDefault", []interface{}{arg1, arg2})
	fake.deleteDefaultMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressPolicyRepo) DeleteDefaultCallCount() int {
	fake.deleteDefaultMutex.RLock()
	defer fake.deleteDefaultMutex.RUnlock()
	return len(fake.deleteDefaultArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteDefaultCalls(stub func(db.Transaction, string) error) {
	fake.deleteDefaultMutex.Lock()
	defer fake.deleteDefaultMutex.Unlock()
	fake.DeleteDefaultStub = stub
}

func (fake *EgressPolicyRepo) DeleteDefaultArgsForCall(i int) (db.Transaction, string) {
	fake.deleteDefaultMutex.RLock()
	defer fake.deleteDefaultMutex.RUnlock()
	argsForCall := fake.deleteDefaultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) DeleteDefaultReturns(result1 error) {
	fake.deleteDefaultMutex.Lock()
	defer fake.deleteDefaultMutex.Unlock()
	fake.DeleteDefaultStub = nil
	fake.deleteDefaultReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteDefaultReturnsOnCall(i int, result1 error) {
	fake.deleteDefaultMutex.Lock()
	defer fake.deleteDefaultMutex.Unlock()
	fake.DeleteDefaultStub = nil
	if fake.deleteDefaultReturnsOnCall == nil {
		fake.deleteDefaultReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteDefaultReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteEgressPolicy(arg1 db.Transaction, arg2 string) error {
	fake.deleteEgressPolicyMutex.Lock()
	ret, specificReturn := fake.deleteEgressPolicyReturnsOnCall[len(fake.deleteEgressPolicyArgsForCall)]
	fake.deleteEgressPolicyArgsForCall = append(fake.deleteEgressPolicyArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteEgressPolicyStub
	fakeReturns := fake.deleteEgressPolicyReturns
	fake.recordInvocation("DeleteEgressPolicy", []interface{}{arg1, arg2})
	fake.deleteEgressPolicyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyCallCount() int {
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	return len(fake.deleteEgressPolicyArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyCalls(stub func(db.Transaction, string) error) {
	fake.deleteEgressPolicyMutex.Lock()
	defer fake.deleteEgressPolicyMutex.Unlock()
	fake.DeleteEgressPolicyStub = stub
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyArgsForCall(i int) (db.Transaction, string) {
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	argsForCall := fake.deleteEgressPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyReturns(result1 error) {
	fake.deleteEgressPolicyMutex.Lock()
	defer fake.deleteEgressPolicyMutex.Unlock()
	fake.DeleteEgressPolicyStub = nil
	fake.deleteEgressPolicyReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteEgressPolicyReturnsOnCall(i int, result1 error) {
	fake.deleteEgressPolicyMutex.Lock()
	defer fake.deleteEgressPolicyMutex.Unlock()
	fake.DeleteEgressPolicyStub = nil
	if fake.deleteEgressPolicyReturnsOnCall == nil {
		fake.deleteEgressPolicyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteEgressPolicyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteSpace(arg1 db.Transaction, arg2 string) error {
	fake.deleteSpaceMutex.Lock()
	ret, specificReturn := fake.deleteSpaceReturnsOnCall[len(fake.deleteSpaceArgsForCall)]
	fake.deleteSpaceArgsForCall = append(fake.deleteSpaceArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteSpaceStub
	fakeReturns := fake.deleteSpaceReturns
	fake.recordInvocation("DeleteSpace", []interface{}{arg1, arg2})
	fake.deleteSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *EgressPolicyRepo) DeleteSpaceCallCount() int {
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	return len(fake.deleteSpaceArgsForCall)
}

func (fake *EgressPolicyRepo) DeleteSpaceCalls(stub func(db.Transaction, string) error) {
	fake.deleteSpaceMutex.Lock()
	defer fake.deleteSpaceMutex.Unlock()
	fake.DeleteSpaceStub = stub
}

func (fake *EgressPolicyRepo) DeleteSpaceArgsForCall(i int) (db.Transaction, string) {
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	argsForCall := fake.deleteSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) DeleteSpaceReturns(result1 error) {
	fake.deleteSpaceMutex.Lock()
	defer fake.deleteSpaceMutex.Unlock()
	fake.DeleteSpaceStub = nil
	fake.deleteSpaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) DeleteSpaceReturnsOnCall(i int, result1 error) {
	fake.deleteSpaceMutex.Lock()
	defer fake.deleteSpaceMutex.Unlock()
	fake.DeleteSpaceStub = nil
	if fake.deleteSpaceReturnsOnCall == nil {
		fake.deleteSpaceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteSpaceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *EgressPolicyRepo) GetAllPolicies() ([]store.EgressPolicy, error) {
	fake.getAllPoliciesMutex.Lock()
	ret, specificReturn := fake.getAllPoliciesReturnsOnCall[len(fake.getAllPoliciesArgsForCall)]
	fake.getAllPoliciesArgsForCall = append(fake.getAllPoliciesArgsForCall, struct {
	}{})
	stub := fake.GetAllPoliciesStub
	fakeReturns := fake.getAllPoliciesReturns
	fake.recordInvocation("GetAllPolicies", []interface{}{})
	fake.getAllPoliciesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) GetAllPoliciesCallCount() int {
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	return len(fake.getAllPoliciesArgsForCall)
}

func (fake *EgressPolicyRepo) GetAllPoliciesCalls(stub func() ([]store.EgressPolicy, error)) {
	fake.getAllPoliciesMutex.Lock()
	defer fake.getAllPoliciesMutex.Unlock()
	fake.GetAllPoliciesStub = stub
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturns(result1 []store.EgressPolicy, result2 error) {
	fake.getAllPoliciesMutex.Lock()
	defer fake.getAllPoliciesMutex.Unlock()
	fake.GetAllPoliciesStub = nil
	fake.getAllPoliciesReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetAllPoliciesReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.getAllPoliciesMutex.Lock()
	defer fake.getAllPoliciesMutex.Unlock()
	fake.GetAllPoliciesStub = nil
	if fake.getAllPoliciesReturnsOnCall == nil {
		fake.getAllPoliciesReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getAllPoliciesReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByFilter(arg1 []string, arg2 []string, arg3 []string, arg4 []string, arg5 []string) ([]store.EgressPolicy, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	var arg4Copy []string
	if arg4 != nil {
		arg4Copy = make([]string, len(arg4))
		copy(arg4Copy, arg4)
	}
	var arg5Copy []string
	if arg5 != nil {
		arg5Copy = make([]string, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.getByFilterMutex.Lock()
	ret, specificReturn := fake.getByFilterReturnsOnCall[len(fake.getByFilterArgsForCall)]
	fake.getByFilterArgsForCall = append(fake.getByFilterArgsForCall, struct {
		arg1 []string
		arg2 []string
		arg3 []string
		arg4 []string
		arg5 []string
	}{arg1Copy, arg2Copy, arg3Copy, arg4Copy, arg5Copy})
	stub := fake.GetByFilterStub
	fakeReturns := fake.getByFilterReturns
	fake.recordInvocation("GetByFilter", []interface{}{arg1Copy, arg2Copy, arg3Copy, arg4Copy, arg5Copy})
	fake.getByFilterMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) GetByFilterCallCount() int {
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	return len(fake.getByFilterArgsForCall)
}

func (fake *EgressPolicyRepo) GetByFilterCalls(stub func([]string, []string, []string, []string, []string) ([]store.EgressPolicy, error)) {
	fake.getByFilterMutex.Lock()
	defer fake.getByFilterMutex.Unlock()
	fake.GetByFilterStub = stub
}

func (fake *EgressPolicyRepo) GetByFilterArgsForCall(i int) ([]string, []string, []string, []string, []string) {
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	argsForCall := fake.getByFilterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *EgressPolicyRepo) GetByFilterReturns(result1 []store.EgressPolicy, result2 error) {
	fake.getByFilterMutex.Lock()
	defer fake.getByFilterMutex.Unlock()
	fake.GetByFilterStub = nil
	fake.getByFilterReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByFilterReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.getByFilterMutex.Lock()
	defer fake.getByFilterMutex.Unlock()
	fake.GetByFilterStub = nil
	if fake.getByFilterReturnsOnCall == nil {
		fake.getByFilterReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getByFilterReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGUID(arg1 db.Transaction, arg2 ...string) ([]store.EgressPolicy, error) {
	fake.getByGUIDMutex.Lock()
	ret, specificReturn := fake.getByGUIDReturnsOnCall[len(fake.getByGUIDArgsForCall)]
	fake.getByGUIDArgsForCall = append(fake.getByGUIDArgsForCall, struct {
		arg1 db.Transaction
		arg2 []string
	}{arg1, arg2})
	stub := fake.GetByGUIDStub
	fakeReturns := fake.getByGUIDReturns
	fake.recordInvocation("GetByGUID", []interface{}{arg1, arg2})
	fake.getByGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) GetByGUIDCallCount() int {
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	return len(fake.getByGUIDArgsForCall)
}

func (fake *EgressPolicyRepo) GetByGUIDCalls(stub func(db.Transaction, ...string) ([]store.EgressPolicy, error)) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = stub
}

func (fake *EgressPolicyRepo) GetByGUIDArgsForCall(i int) (db.Transaction, []string) {
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	argsForCall := fake.getByGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) GetByGUIDReturns(result1 []store.EgressPolicy, result2 error) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = nil
	fake.getByGUIDReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetByGUIDReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.getByGUIDMutex.Lock()
	defer fake.getByGUIDMutex.Unlock()
	fake.GetByGUIDStub = nil
	if fake.getByGUIDReturnsOnCall == nil {
		fake.getByGUIDReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getByGUIDReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetBySourceGuidsAndDefaults(arg1 []string) ([]store.EgressPolicy, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	ret, specificReturn := fake.getBySourceGuidsAndDefaultsReturnsOnCall[len(fake.getBySourceGuidsAndDefaultsArgsForCall)]
	fake.getBySourceGuidsAndDefaultsArgsForCall = append(fake.getBySourceGuidsAndDefaultsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.GetBySourceGuidsAndDefaultsStub
	fakeReturns := fake.getBySourceGuidsAndDefaultsReturns
	fake.recordInvocation("GetBySourceGuidsAndDefaults", []interface{}{arg1Copy})
	fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) GetBySourceGuidsAndDefaultsCallCount() int {
	fake.getBySourceGuidsAndDefaultsMutex.RLock()
	defer fake.getBySourceGuidsAndDefaultsMutex.RUnlock()
	return len(fake.getBySourceGuidsAndDefaultsArgsForCall)
}

func (fake *EgressPolicyRepo) GetBySourceGuidsAndDefaultsCalls(stub func([]string) ([]store.EgressPolicy, error)) {
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	defer fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	fake.GetBySourceGuidsAndDefaultsStub = stub
}

func (fake *EgressPolicyRepo) GetBySourceGuidsAndDefaultsArgsForCall(i int) []string {
	fake.getBySourceGuidsAndDefaultsMutex.RLock()
	defer fake.getBySourceGuidsAndDefaultsMutex.RUnlock()
	argsForCall := fake.getBySourceGuidsAndDefaultsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *EgressPolicyRepo) GetBySourceGuidsAndDefaultsReturns(result1 []store.EgressPolicy, result2 error) {
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	defer fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	fake.GetBySourceGuidsAndDefaultsStub = nil
	fake.getBySourceGuidsAndDefaultsReturns = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetBySourceGuidsAndDefaultsReturnsOnCall(i int, result1 []store.EgressPolicy, result2 error) {
	fake.getBySourceGuidsAndDefaultsMutex.Lock()
	defer fake.getBySourceGuidsAndDefaultsMutex.Unlock()
	fake.GetBySourceGuidsAndDefaultsStub = nil
	if fake.getBySourceGuidsAndDefaultsReturnsOnCall == nil {
		fake.getBySourceGuidsAndDefaultsReturnsOnCall = make(map[int]struct {
			result1 []store.EgressPolicy
			result2 error
		})
	}
	fake.getBySourceGuidsAndDefaultsReturnsOnCall[i] = struct {
		result1 []store.EgressPolicy
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUID(arg1 db.Transaction, arg2 string) (string, error) {
	fake.getTerminalByAppGUIDMutex.Lock()
	ret, specificReturn := fake.getTerminalByAppGUIDReturnsOnCall[len(fake.getTerminalByAppGUIDArgsForCall)]
	fake.getTerminalByAppGUIDArgsForCall = append(fake.getTerminalByAppGUIDArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.GetTerminalByAppGUIDStub
	fakeReturns := fake.getTerminalByAppGUIDReturns
	fake.recordInvocation("GetTerminalByAppGUID", []interface{}{arg1, arg2})
	fake.getTerminalByAppGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDCallCount() int {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	return len(fake.getTerminalByAppGUIDArgsForCall)
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDCalls(stub func(db.Transaction, string) (string, error)) {
	fake.getTerminalByAppGUIDMutex.Lock()
	defer fake.getTerminalByAppGUIDMutex.Unlock()
	fake.GetTerminalByAppGUIDStub = stub
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDArgsForCall(i int) (db.Transaction, string) {
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	argsForCall := fake.getTerminalByAppGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturns(result1 string, result2 error) {
	fake.getTerminalByAppGUIDMutex.Lock()
	defer fake.getTerminalByAppGUIDMutex.Unlock()
	fake.GetTerminalByAppGUIDStub = nil
	fake.getTerminalByAppGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalByAppGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getTerminalByAppGUIDMutex.Lock()
	defer fake.getTerminalByAppGUIDMutex.Unlock()
	fake.GetTerminalByAppGUIDStub = nil
	if fake.getTerminalByAppGUIDReturnsOnCall == nil {
		fake.getTerminalByAppGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getTerminalByAppGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUID(arg1 db.Transaction, arg2 string) (string, error) {
	fake.getTerminalBySpaceGUIDMutex.Lock()
	ret, specificReturn := fake.getTerminalBySpaceGUIDReturnsOnCall[len(fake.getTerminalBySpaceGUIDArgsForCall)]
	fake.getTerminalBySpaceGUIDArgsForCall = append(fake.getTerminalBySpaceGUIDArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.GetTerminalBySpaceGUIDStub
	fakeReturns := fake.getTerminalBySpaceGUIDReturns
	fake.recordInvocation("GetTerminalBySpaceGUID", []interface{}{arg1, arg2})
	fake.getTerminalBySpaceGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDCallCount() int {
	fake.getTerminalBySpaceGUIDMutex.RLock()
	defer fake.getTerminalBySpaceGUIDMutex.RUnlock()
	return len(fake.getTerminalBySpaceGUIDArgsForCall)
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDCalls(stub func(db.Transaction, string) (string, error)) {
	fake.getTerminalBySpaceGUIDMutex.Lock()
	defer fake.getTerminalBySpaceGUIDMutex.Unlock()
	fake.GetTerminalBySpaceGUIDStub = stub
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDArgsForCall(i int) (db.Transaction, string) {
	fake.getTerminalBySpaceGUIDMutex.RLock()
	defer fake.getTerminalBySpaceGUIDMutex.RUnlock()
	argsForCall := fake.getTerminalBySpaceGUIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDReturns(result1 string, result2 error) {
	fake.getTerminalBySpaceGUIDMutex.Lock()
	defer fake.getTerminalBySpaceGUIDMutex.Unlock()
	fake.GetTerminalBySpaceGUIDStub = nil
	fake.getTerminalBySpaceGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) GetTerminalBySpaceGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getTerminalBySpaceGUIDMutex.Lock()
	defer fake.getTerminalBySpaceGUIDMutex.Unlock()
	fake.GetTerminalBySpaceGUIDStub = nil
	if fake.getTerminalBySpaceGUIDReturnsOnCall == nil {
		fake.getTerminalBySpaceGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getTerminalBySpaceGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) IsTerminalInUse(arg1 db.Transaction, arg2 string) (bool, error) {
	fake.isTerminalInUseMutex.Lock()
	ret, specificReturn := fake.isTerminalInUseReturnsOnCall[len(fake.isTerminalInUseArgsForCall)]
	fake.isTerminalInUseArgsForCall = append(fake.isTerminalInUseArgsForCall, struct {
		arg1 db.Transaction
		arg2 string
	}{arg1, arg2})
	stub := fake.IsTerminalInUseStub
	fakeReturns := fake.isTerminalInUseReturns
	fake.recordInvocation("IsTerminalInUse", []interface{}{arg1, arg2})
	fake.isTerminalInUseMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EgressPolicyRepo) IsTerminalInUseCallCount() int {
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	return len(fake.isTerminalInUseArgsForCall)
}

func (fake *EgressPolicyRepo) IsTerminalInUseCalls(stub func(db.Transaction, string) (bool, error)) {
	fake.isTerminalInUseMutex.Lock()
	defer fake.isTerminalInUseMutex.Unlock()
	fake.IsTerminalInUseStub = stub
}

func (fake *EgressPolicyRepo) IsTerminalInUseArgsForCall(i int) (db.Transaction, string) {
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	argsForCall := fake.isTerminalInUseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EgressPolicyRepo) IsTerminalInUseReturns(result1 bool, result2 error) {
	fake.isTerminalInUseMutex.Lock()
	defer fake.isTerminalInUseMutex.Unlock()
	fake.IsTerminalInUseStub = nil
	fake.isTerminalInUseReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) IsTerminalInUseReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isTerminalInUseMutex.Lock()
	defer fake.isTerminalInUseMutex.Unlock()
	fake.IsTerminalInUseStub = nil
	if fake.isTerminalInUseReturnsOnCall == nil {
		fake.isTerminalInUseReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isTerminalInUseReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *EgressPolicyRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	fake.createDefaultMutex.RLock()
	defer fake.createDefaultMutex.RUnlock()
	fake.createEgressPolicyMutex.RLock()
	defer fake.createEgressPolicyMutex.RUnlock()
	fake.createSpaceMutex.RLock()
	defer fake.createSpaceMutex.RUnlock()
	fake.deleteAppMutex.RLock()
	defer fake.deleteAppMutex.RUnlock()
	fake.deleteDefaultMutex.RLock()
	defer fake.deleteDefaultMutex.RUnlock()
	fake.deleteEgressPolicyMutex.RLock()
	defer fake.deleteEgressPolicyMutex.RUnlock()
	fake.deleteSpaceMutex.RLock()
	defer fake.deleteSpaceMutex.RUnlock()
	fake.getAllPoliciesMutex.RLock()
	defer fake.getAllPoliciesMutex.RUnlock()
	fake.getByFilterMutex.RLock()
	defer fake.getByFilterMutex.RUnlock()
	fake.getByGUIDMutex.RLock()
	defer fake.getByGUIDMutex.RUnlock()
	fake.getBySourceGuidsAndDefaultsMutex.RLock()
	defer fake.getBySourceGuidsAndDefaultsMutex.RUnlock()
	fake.getTerminalByAppGUIDMutex.RLock()
	defer fake.getTerminalByAppGUIDMutex.RUnlock()
	fake.getTerminalBySpaceGUIDMutex.RLock()
	defer fake.getTerminalBySpaceGUIDMutex.RUnlock()
	fake.isTerminalInUseMutex.RLock()
	defer fake.isTerminalInUseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EgressPolicyRepo) recordInvocation(key string, args []interface{}) {
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
