// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package contract

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

const (
	SelectorLen = 4
)

type callType int

const (
	call callType = iota
	callCode
	delegateCall
	staticCall
)

type RunStatefulPrecompileFunc func(
	accessibleState AccessibleState,
	caller common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error)

// StatefulPrecompileFunction defines a function implemented by a stateful precompile
type StatefulPrecompileFunction struct {
	// selector is the 4 byte function selector for this function
	selector []byte
	// execute is performed when this function is selected
	execute RunStatefulPrecompileFunc
}

// NewStatefulPrecompileFunction creates a stateful precompile function with the given arguments
func NewStatefulPrecompileFunction(selector []byte, execute RunStatefulPrecompileFunc) *StatefulPrecompileFunction {
	return &StatefulPrecompileFunction{
		selector: selector,
		execute:  execute,
	}
}

// statefulPrecompileWithFunctionSelectors implements StatefulPrecompiledContract by using 4 byte function selectors to pass
// off responsibilities to internal execution functions.
// Note: because we only ever read from [functions] there no lock is required to make it thread-safe.
type statefulPrecompileWithFunctionSelectors struct {
	functions map[string]*StatefulPrecompileFunction
}

// NewStatefulPrecompileContract generates new StatefulPrecompile using [functions] as the available functions and [fallback]
// as an optional fallback if there is no input data. Note: the selector of [fallback] will be ignored, so it is required to be left empty.
func NewStatefulPrecompileContract(functions []*StatefulPrecompileFunction) (StatefulPrecompiledContract, error) {
	// Construct the contract and populate [functions].
	contract := &statefulPrecompileWithFunctionSelectors{
		functions: make(map[string]*StatefulPrecompileFunction),
	}
	for _, function := range functions {
		if len(function.selector) != SelectorLen {
			return nil, fmt.Errorf("invalid length of function selector, want: %v, got: %v", SelectorLen, len(function.selector))
		}

		_, exists := contract.functions[string(function.selector)]
		if exists {
			return nil, fmt.Errorf("cannot create stateful precompile with duplicated function selector: %q", function.selector)
		}
		contract.functions[string(function.selector)] = function
	}

	return contract, nil
}

// Run selects the function using the 4 byte function selector at the start of the input and executes the underlying function on the
// given arguments.
func (s *statefulPrecompileWithFunctionSelectors) Run(
	accessibleState AccessibleState,
	caller common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	// Otherwise, an unexpected input size will result in an error.
	if len(input) < SelectorLen {
		return nil, suppliedGas, fmt.Errorf("missing function selector to precompile - input length (%d)", len(input))
	}

	// Use the function selector to grab the correct function
	selector := input[:SelectorLen]
	functionInput := input[SelectorLen:]
	function, ok := s.functions[string(selector)]
	if !ok {
		return nil, suppliedGas, fmt.Errorf("invalid function selector %#x", selector)
	}

	return function.execute(accessibleState, caller, addr, functionInput, suppliedGas, readOnly)
}
