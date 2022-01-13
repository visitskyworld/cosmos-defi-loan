package keeper

import (
	"context"

	"github.com/cosmonaut/loan/x/loan/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sheldonlsides/loan/x/loan/types"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id)

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "key %d doesn't exist", msg.Id)
	}

	if loan.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)

	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongLoanState, "Cannot parse coins in loan amount")
	}

	k.bankKeeper.SendCoins(ctx, lender, borrower, amount)

	//sets the loan leander and state of loan to approved
	loan.Lender = msg.Creator
	loan.State = "approved"

	//updates loan state
	k.SetLoan(ctx, loan)
	return &types.MsgApproveLoanResponse{}, nil
}
