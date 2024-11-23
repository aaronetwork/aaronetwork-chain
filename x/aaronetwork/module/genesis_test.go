package aaronetwork_test

import (
	"testing"

	keepertest "aaronetwork/testutil/keeper"
	"aaronetwork/testutil/nullify"
	aaronetwork "aaronetwork/x/aaronetwork/module"
	"aaronetwork/x/aaronetwork/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AaronetworkKeeper(t)
	aaronetwork.InitGenesis(ctx, k, genesisState)
	got := aaronetwork.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
