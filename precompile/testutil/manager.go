package testutil

import "github.com/ethereum/go-ethereum/common"

// Manager implements precompile IsEnabled interface
// and allows explicit control over which addresses
// are enabled and which are disabled
type Manager struct {
	enabledPrecompiles map[string]struct{}
}

// NewManager returns a new *Manager that has no precompiled enabled by default
func NewManager() *Manager {
	return &Manager{
		enabledPrecompiles: make(map[string]struct{}),
	}
}

// IsEnabled returns true if the precompile address is enabled
func (m *Manager) IsEnabled(addr common.Address) (enabled bool) {
	_, enabled = m.enabledPrecompiles[addr.String()]
	return enabled
}

// Enable enables a precompile address
func (m *Manager) Enable(addr common.Address) {
	m.enabledPrecompiles[addr.String()] = struct{}{}
}

// Disable disables a precompile address
func (m *Manager) Disable(addr common.Address) {
	delete(m.enabledPrecompiles, addr.String())
}
