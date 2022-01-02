package loan_test

import (
	"testing"

	keepertest "github.com/sheldonlsides/loan/testutil/keeper"
	"github.com/sheldonlsides/loan/testutil/nullify"
	"github.com/sheldonlsides/loan/x/loan"
	"github.com/sheldonlsides/loan/x/loan/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LoanKeeper(t)
	loan.InitGenesis(ctx, *k, genesisState)
	got := loan.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
