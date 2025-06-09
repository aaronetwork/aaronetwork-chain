package keeper

import (
	"context"

	"aaronetwork/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRoom(goCtx context.Context, msg *types.MsgCreateRoom) (*types.MsgCreateRoomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateRoomResponse{}, nil
}
