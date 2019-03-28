// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/customtx"
	"github.com/hyperledger/fabric/protos/common"
)

type Processor struct {
	GenerateSimulationResultsStub        func(*common.Envelope, ledger.TxSimulator, bool) error
	generateSimulationResultsMutex       sync.RWMutex
	generateSimulationResultsArgsForCall []struct {
		arg1 *common.Envelope
		arg2 ledger.TxSimulator
		arg3 bool
	}
	generateSimulationResultsReturns struct {
		result1 error
	}
	generateSimulationResultsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Processor) GenerateSimulationResults(arg1 *common.Envelope, arg2 ledger.TxSimulator, arg3 bool) error {
	fake.generateSimulationResultsMutex.Lock()
	ret, specificReturn := fake.generateSimulationResultsReturnsOnCall[len(fake.generateSimulationResultsArgsForCall)]
	fake.generateSimulationResultsArgsForCall = append(fake.generateSimulationResultsArgsForCall, struct {
		arg1 *common.Envelope
		arg2 ledger.TxSimulator
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("GenerateSimulationResults", []interface{}{arg1, arg2, arg3})
	fake.generateSimulationResultsMutex.Unlock()
	if fake.GenerateSimulationResultsStub != nil {
		return fake.GenerateSimulationResultsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.generateSimulationResultsReturns
	return fakeReturns.result1
}

func (fake *Processor) GenerateSimulationResultsCallCount() int {
	fake.generateSimulationResultsMutex.RLock()
	defer fake.generateSimulationResultsMutex.RUnlock()
	return len(fake.generateSimulationResultsArgsForCall)
}

func (fake *Processor) GenerateSimulationResultsCalls(stub func(*common.Envelope, ledger.TxSimulator, bool) error) {
	fake.generateSimulationResultsMutex.Lock()
	defer fake.generateSimulationResultsMutex.Unlock()
	fake.GenerateSimulationResultsStub = stub
}

func (fake *Processor) GenerateSimulationResultsArgsForCall(i int) (*common.Envelope, ledger.TxSimulator, bool) {
	fake.generateSimulationResultsMutex.RLock()
	defer fake.generateSimulationResultsMutex.RUnlock()
	argsForCall := fake.generateSimulationResultsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Processor) GenerateSimulationResultsReturns(result1 error) {
	fake.generateSimulationResultsMutex.Lock()
	defer fake.generateSimulationResultsMutex.Unlock()
	fake.GenerateSimulationResultsStub = nil
	fake.generateSimulationResultsReturns = struct {
		result1 error
	}{result1}
}

func (fake *Processor) GenerateSimulationResultsReturnsOnCall(i int, result1 error) {
	fake.generateSimulationResultsMutex.Lock()
	defer fake.generateSimulationResultsMutex.Unlock()
	fake.GenerateSimulationResultsStub = nil
	if fake.generateSimulationResultsReturnsOnCall == nil {
		fake.generateSimulationResultsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.generateSimulationResultsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Processor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateSimulationResultsMutex.RLock()
	defer fake.generateSimulationResultsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Processor) recordInvocation(key string, args []interface{}) {
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

var _ customtx.Processor = new(Processor)