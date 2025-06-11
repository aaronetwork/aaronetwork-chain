package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"time"

	"aaronetwork/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateRoom(goCtx context.Context, msg *types.MsgCreateRoom) (*types.MsgCreateRoomResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	ctx := sdk.WrapSDKContext(sdkCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if _, found := k.Keeper.GetRoomByHandle(sdkCtx, msg.Handle); found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("handle %s already exists", msg.Handle))
	}

	roomID := time.Now().UTC().Format("20060102150405") + "-" + uuid.NewString()

	room := types.Room{
		Id:          roomID,
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		Handle:      msg.Handle,
	}

	k.Keeper.SetRoom(ctx, room)

	return &types.MsgCreateRoomResponse{Id: roomID}, nil
}
