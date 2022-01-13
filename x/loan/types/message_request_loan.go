package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestLoan = "request_loan"

var _ sdk.Msg = &MsgRequestLoan{}

func NewMsgRequestLoan(creator string, amount string, fee string, collateral string, deadline string) *MsgRequestLoan {
	return &MsgRequestLoan{
		Creator:    creator,
		Amount:     amount,
		Fee:        fee,
		Collateral: collateral,
		Deadline:   deadline,
	}
}

func (msg *MsgRequestLoan) Route() string {
	return RouterKey
}

func (msg *MsgRequestLoan) Type() string {
	return TypeMsgRequestLoan
}

func (msg *MsgRequestLoan) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestLoan) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestLoan) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)

	amount, err  := sdk.ParseCoinsNormalized(msg.Amount)
	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(msg.Collateral)

	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%)", err)
	}

	if !amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}
	if amount.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	if !fee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
	}
	if !collateral.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is not a valid Coins object")
	}
	if collateral.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collateral is empty")
	}
	return nil
}
