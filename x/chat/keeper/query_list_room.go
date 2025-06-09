package keeper

import (
	"context"

	"aaronetwork/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListRoom(goCtx context.Context, req *types.QueryListRoomRequest) (*types.QueryListRoomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryListRoomResponse{}, nil
}
