package keeper

import (
	"aaronetwork/x/chat/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListRoom(goCtx context.Context, req *types.QueryListRoomRequest) (*types.QueryListRoomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	roomsValue := k.GetAllRoom(sdk.WrapSDKContext(ctx)) // []types.Room

	roomsPtr := make([]*types.Room, len(roomsValue))
	for i := range roomsValue {
		roomsPtr[i] = &roomsValue[i]
	}

	return &types.QueryListRoomResponse{
		Room: roomsPtr,
	}, nil
}
