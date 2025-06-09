package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "aaronetwork/testutil/keeper"
	"aaronetwork/x/chat/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ChatKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
