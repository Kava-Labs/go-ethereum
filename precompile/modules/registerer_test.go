// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package modules

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile/testutil"
	"github.com/stretchr/testify/require"
)

func TestInsertSortedByAddress(t *testing.T) {
	clearRegisteredModules()

	data := make([]Module, 0)
	// test that the module is registered in sorted order
	module1 := Module{
		Address: common.BigToAddress(big.NewInt(1)),
	}
	data = insertSortedByAddress(data, module1)
	require.Equal(t, []Module{module1}, data)

	module0 := Module{
		Address: common.BigToAddress(big.NewInt(0)),
	}
	data = insertSortedByAddress(data, module0)
	require.Equal(t, []Module{module0, module1}, data)

	module3 := Module{
		Address: common.BigToAddress(big.NewInt(3)),
	}
	data = insertSortedByAddress(data, module3)
	require.Equal(t, []Module{module0, module1, module3}, data)

	module2 := Module{
		Address: common.BigToAddress(big.NewInt(2)),
	}
	data = insertSortedByAddress(data, module2)
	require.Equal(t, []Module{module0, module1, module2, module3}, data)
}

func TestRegisterModule(t *testing.T) {
	clearRegisteredModules()

	const moduleNum = 4
	// create modules
	modules := make([]Module, moduleNum)
	for i := 0; i < moduleNum; i++ {
		modules[i] = Module{
			Address: common.BigToAddress(big.NewInt(int64(i))),
		}
	}

	// register modules
	for i := 0; i < moduleNum; i++ {
		err := RegisterModule(modules[i])
		require.NoError(t, err)
	}

	// get modules by address
	for i := 0; i < moduleNum; i++ {
		address := common.BigToAddress(big.NewInt(int64(i)))
		module, exists := GetPrecompileModuleByAddress(address)
		require.True(t, exists)
		require.Equal(t, module, Module{
			Address: address,
		})
	}

	// get unexisting module by address
	address := common.BigToAddress(big.NewInt(moduleNum))
	_, exists := GetPrecompileModuleByAddress(address)
	require.False(t, exists)

	// get all modules
	registeredModules := RegisteredModules()
	require.Equal(t, modules, registeredModules)
}

func TestRegisterModuleWithDuplicateAddress(t *testing.T) {
	clearRegisteredModules()

	modules := []Module{
		{
			Address: common.BigToAddress(big.NewInt(0)),
		},
	}

	err := RegisterModule(modules[0])
	require.NoError(t, err)

	err = RegisterModule(modules[0])
	require.Error(t, err)
	require.Contains(t, err.Error(), "address 0x0000000000000000000000000000000000000000 already used by a stateful precompile")

	// get all modules
	registeredModules := RegisteredModules()
	require.Equal(t, modules, registeredModules)
}

func TestRegisterIsEnabled(t *testing.T) {
	clearRegisteredModules()

	manager := testutil.NewManager()
	modules := []Module{
		{
			Address: common.BigToAddress(big.NewInt(0)),
		},
		{
			Address: common.BigToAddress(big.NewInt(1)),
		},
		{
			Address: common.BigToAddress(big.NewInt(3)),
		},
	}
	for _, module := range modules {
		err := RegisterModule(module)
		require.NoError(t, err, "expected no error when registering module for test")
	}

	// default is all modules enabled
	enabledModules := RegisteredModules()
	require.Equal(t, modules, enabledModules)
	for _, module := range modules {
		_, ok := GetPrecompileModuleByAddress(module.Address)
		require.True(t, ok, "expected precompile %s to be enabled", module.Address)
	}

	// set manager with no enabled modules
	SetPrecompileManager(manager)
	enabledModules = RegisteredModules()
	require.Equal(t, 0, len(enabledModules), "expected no registered modules ot be enabled")
	for _, module := range modules {
		_, ok := GetPrecompileModuleByAddress(module.Address)
		require.False(t, ok, "expected precompile %s to not enabled", module.Address)
	}

	// enable and test sorting is as expected
	manager.Enable(modules[0].Address)
	enabledModules = RegisteredModules()
	require.Equal(t, modules[:1], enabledModules, "expected one enabled precompile module")
	module, enabled := GetPrecompileModuleByAddress(modules[0].Address)
	require.Equal(t, modules[0], module, "expected module to be returned")
	require.True(t, enabled, "expected module to be enabled")

	// enable and test sorting is as expected
	manager.Enable(modules[2].Address)
	enabledModules = RegisteredModules()
	require.Equal(t, []Module{modules[0], modules[2]}, enabledModules, "expected two enabled precompile modules")
	module, enabled = GetPrecompileModuleByAddress(modules[2].Address)
	require.Equal(t, modules[2], module, "expected module to be returned")
	require.True(t, enabled, "expected module to be enabled")

	// test that disable works
	manager.Disable(modules[0].Address)
	enabledModules = RegisteredModules()
	require.Equal(t, modules[2:], enabledModules, "expected one enabled precompile modules")
	_, enabled = GetPrecompileModuleByAddress(modules[0].Address)
	require.False(t, enabled, "expected module to not be enabled")
}

func clearRegisteredModules() {
	registeredModules = make([]Module, 0)
}
