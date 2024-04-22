// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package modules

import (
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

type Manager interface {
	IsEnabled(common.Address) bool
}

type defaultManager struct {
}

func newDefaultManager() Manager {
	return defaultManager{}
}

func (_ defaultManager) IsEnabled(_ common.Address) bool {
	return true
}

var (
	// registeredModules is a list of Module to preserve order
	// for deterministic iteration
	registeredModules = make([]Module, 0)

	// manager implements Manager interface
	isEnabledManager = newDefaultManager()
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

func GetPrecompileModuleByAddress(address common.Address) (Module, bool) {
	for _, stm := range registeredModules {
		if stm.Address == address && isEnabledManager.IsEnabled(stm.Address) {
			return stm, true
		}
	}
	return Module{}, false
}

func RegisteredModules() (enabledModules []Module) {
	for _, module := range registeredModules {
		if isEnabledManager.IsEnabled(module.Address) {
			enabledModules = append(enabledModules, module)
		}
	}
	return enabledModules
}

func SetPrecompileManager(manager Manager) {
	isEnabledManager = manager
}

func insertSortedByAddress(data []Module, stm Module) []Module {
	data = append(data, stm)
	sort.Sort(moduleArray(data))
	return data
}
