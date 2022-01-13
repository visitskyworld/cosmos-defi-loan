package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sheldonlsides/loan/x/loan/types"
)

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//create a new loan with the following user input
	var loan = types.Loan{
		Amount: msg.Amount,
		Fee: msg.Fee,
		Collateral: msg.Collateral,
		Deadline: msg.Deadline,
		State: "requested",
		Borrower: msg.Creator,
	}

	//checks to see if collateral is more than the loan amount + fee
	

	//get the borrower address
	borrower, _ := sdk.AccAddressFromBech32(msg.Creator)

	//get the collateral as sdk.Coins
	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)

	if err != nil {
		panic(err)
	}

	//use the module account as escrow account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)

	if sdkError != nil {
		return nil, sdkError
	}

	k.AppendLoan(ctx, loan)

	return &types.MsgRequestLoanResponse{}, nil
}
