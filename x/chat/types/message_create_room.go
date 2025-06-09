package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRoom{}

func NewMsgCreateRoom(creator string, title string, description string) *MsgCreateRoom {
	return &MsgCreateRoom{
		Creator:     creator,
		Title:       title,
		Description: description,
	}
}

func (msg *MsgCreateRoom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
