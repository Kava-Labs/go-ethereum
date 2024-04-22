package testutil_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	addr1 = common.HexToAddress("0x9000000000000000000000000000000000000001")
	addr2 = common.HexToAddress("0x9000000000000000000000000000000000000002")
)

func TestPrecompileManager(t *testing.T) {
	require.False(t, addr1 == addr2, "test requires two unqiue addresses")

	m := testutil.NewManager()

	// defaults to no contracts enabled
	assert.False(t, m.IsEnabled(addr1))
	assert.False(t, m.IsEnabled(addr2))

	m.Enable(addr1)

	// enable 1 contract, keep other disabled
	assert.True(t, m.IsEnabled(addr1))
	assert.False(t, m.IsEnabled(addr2))

	m.Enable(addr2)

	// enable both contracts
	assert.True(t, m.IsEnabled(addr1))
	assert.True(t, m.IsEnabled(addr2))

	m.Disable(addr1)

	// disable one contract
	assert.False(t, m.IsEnabled(addr1))
	assert.True(t, m.IsEnabled(addr2))

	m.Disable(addr2)

	// disable both contracts
	assert.False(t, m.IsEnabled(addr1))
	assert.False(t, m.IsEnabled(addr2))

	// ensure multiple calls don't panic
	assert.NotPanics(t, func() {
		m.Enable(addr1)
		m.Enable(addr1)
		m.Enable(addr1)

		m.Disable(addr2)
		m.Disable(addr2)
		m.Disable(addr2)
	})
}
