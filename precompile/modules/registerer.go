// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package modules

import (
	"bytes"
	"fmt"
	"slices"
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// registeredModules is a list of Module to preserve order
	// for deterministic iteration
	registeredModules = make([]Module, 0)
)

// RegisterModule registers a stateful precompile module
func RegisterModule(stm Module) error {
	address := stm.Address

	for _, registeredModule := range registeredModules {
		if registeredModule.Address == address {
			return fmt.Errorf("address %s already used by a stateful precompile", address)
		}
	}
	// sort by address to ensure deterministic iteration
	registeredModules = insertSortedByAddress(registeredModules, stm)
	return nil
}

// GetPrecompileModuleByAddress returns a precompile module by address and true
// if found. Otherwise, it returns false. Uses binary search to find the module,
// as the list is sorted by address.
func SearchPrecompileModuleByAddress(address common.Address) (Module, bool) {
	idx, found := slices.BinarySearchFunc(registeredModules, Module{
		Address: address,
	}, func(a, b Module) int {
		return bytes.Compare(a.Address.Bytes(), b.Address.Bytes())
	})

	if !found {
		return Module{}, false
	}

	return registeredModules[idx], true
}

func GetPrecompileModuleByAddress(address common.Address) (Module, bool) {
	for _, stm := range registeredModules {
		if stm.Address == address {
			return stm, true
		}
	}
	return Module{}, false
}

func RegisteredModules() []Module {
	return registeredModules
}

func insertSortedByAddress(data []Module, stm Module) []Module {
	data = append(data, stm)
	sort.Sort(moduleArray(data))
	return data
}

func ClearRegisteredModules() {
	registeredModules = make([]Module, 0)
}
