package staking

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maticnetwork/heimdall/staking/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, data types.GenesisState) {
	// get current val set
	var vals []*hmTypes.Validator
	if len(data.CurrentValSet.Validators) == 0 {
		vals = data.Validators
	} else {
		vals = data.CurrentValSet.Validators
	}

	// result
	resultValSet := hmTypes.NewValidatorSet(vals)
	validatorRewards := make(map[hmTypes.ValidatorID]*big.Int)

	// add validators in store
	for _, validator := range resultValSet.Validators {
		// Add individual validator to state
		keeper.AddValidator(ctx, *validator)

	}

	// update validator set in store
	if err := keeper.UpdateValidatorSetInStore(ctx, *resultValSet); err != nil {
		panic(err)
	}

	// Add rewards for initial validators
	for _, validator := range data.Validators {
		if _, ok := data.ValidatorRewards[validator.ID.String()]; ok {
			validatorRewards[validator.ID] = data.ValidatorRewards[validator.ID.String()]
		} else {
			validatorRewards[validator.ID] = big.NewInt(0)
		}
	}

	// update validator rewards
	keeper.UpdateValidatorRewards(ctx, validatorRewards)

	// proposer bonus percent
	keeper.SetProposerBonusPercent(ctx, data.ProposerBonusPercent)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) types.GenesisState {
	// Validator rewards
	validatorRewards := make(map[string]*big.Int)

	// Add rewards for initial validators
	for validatorID, validator := range keeper.GetAllValidatorRewards(ctx) {
		validatorRewards[validatorID.String()] = validator
	}

	// return new genesis state
	return types.NewGenesisState(
		keeper.GetAllValidators(ctx),
		keeper.GetValidatorSet(ctx),
		validatorRewards,
		keeper.GetProposerBonusPercent(ctx),
	)
}
