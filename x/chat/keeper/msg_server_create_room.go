package keeper

import (
	"context"
	"github.com/google/uuid"
	"time"

	"aaronetwork/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRoom(goCtx context.Context, msg *types.MsgCreateRoom) (*types.MsgCreateRoomResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	ctx := sdk.WrapSDKContext(sdkCtx)

	roomID := time.Now().UTC().Format("20060102150405") + "-" + uuid.NewString()

	room := types.Room{
		Id:          roomID,
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
	}

	k.Keeper.SetRoom(ctx, room)

	return &types.MsgCreateRoomResponse{Id: roomID}, nil
}
