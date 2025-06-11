package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"regexp"
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
	if len(msg.Title) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "title cannot be empty")
	}
	if len(msg.Description) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description cannot be empty")
	}
	if len(msg.Handle) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "handle cannot be empty")
	}
	if len(msg.Handle) > 32 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "handle too long (max 32 chars)")
	}
	if !isValidHandle(msg.Handle) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "handle must contain only lowercase letters, numbers, underscores")
	}

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func isValidHandle(handle string) bool {
	return regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(handle)
}
