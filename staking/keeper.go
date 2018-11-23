package staking

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/maticnetwork/heimdall/checkpoint"
	"github.com/maticnetwork/heimdall/types"
	"github.com/tendermint/tendermint/crypto"
)

type Keeper struct {
	storeKey         sdk.StoreKey
	cdc              *codec.Codec
	checkpointKeeper checkpoint.Keeper

	// codespace
	codespace sdk.CodespaceType
}

var (
	ValidatorsKey = []byte{0x01} // prefix for each key to a validator
)

// getValidatorKey drafts the validator key for addresses
func getValidatorKey(address []byte) []byte {
	return append(ValidatorsKey, address...)
}

// NewKeeper creates new keeper for staking
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, codespace sdk.CodespaceType) Keeper {
	keeper := Keeper{
		storeKey:  key,
		cdc:       cdc,
		codespace: codespace,
	}
	return keeper
}

// Add validator indexed with address
func (k Keeper) AddValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)

	// marshall validator
	bz := k.cdc.MustMarshalBinary(validator)

	// store validator with address prefixed with validator key as index
	store.Set(getValidatorKey(validator.Pubkey.Address().Bytes()), bz)
}

// GetAllValidators returns all validators added for a specific validator key
func (k Keeper) GetAllCurrentValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	// get ACK count
	ACKs := k.checkpointKeeper.GetACKCount(ctx)

	// create iterator to iterate with Validator Key prefix
	iterator := sdk.KVStorePrefixIterator(store, ValidatorsKey)
	defer iterator.Close()

	// loop through validators to get valid validators
	for i := 0; ; i++ {
		if !iterator.Valid() {
			break
		}

		// unmarshall validator
		var validator types.Validator
		k.cdc.MustUnmarshalBinary(iterator.Value(), &validator)

		// check if validator is valid for current epoch
		if validator.IsCurrentValidator(ACKs) {
			// append if validator is current valdiator
			validators = append(validators, validator)
		}

		// increment iterator
		iterator.Next()
	}
	return
}

// GetValidatorInfo returns validator info for given the address
func (k Keeper) GetValidatorInfo(ctx sdk.Context, valAddr common.Address) (validator types.Validator, error error) {
	store := ctx.KVStore(k.storeKey)

	// get validator and unmarshall
	validatorBytes := store.Get(getValidatorKey(valAddr.Bytes()))
	if validatorBytes == nil {
		error = fmt.Errorf("Validator Not Found")
		return
	}

	// unmarshall validator (TODO: we might want to shift to mustUnmarshallBinary)
	error = k.cdc.UnmarshalBinary(validatorBytes, &validator)
	if error != nil {
		StakingLogger.Error("Error unmarshalling validator while fetching validator from store", "Error", error, "ValidatorAddress", valAddr)
		return
	}

	return validator, nil
}

func (k Keeper) RemoveValidator(ctx sdk.Context, valAddr common.Address) error {
	store := ctx.KVStore(k.storeKey)

	var validator types.Validator

	// get validator and unmarshall
	validatorBytes := store.Get(getValidatorKey(valAddr.Bytes()))
	if validatorBytes == nil {
		err := fmt.Errorf("Validator Not Found")
		return err
	}

	err := k.cdc.UnmarshalBinary(validatorBytes, &validator)
	if err != nil {
		StakingLogger.Error("Error unmarshalling validator while fetching validator from store", "Error", err, "ValidatorAddress", valAddr)
		return err
	}

	// generate empty validator
	validator = types.CreateValidatorWithAddr(validator.Address)

	// add updated validator to store with same key
	k.AddValidator(ctx, validator)

	return nil
}

// update validator with signer and pubkey
func (k Keeper) UpdateSigner(ctx sdk.Context, signer common.Address, pubkey crypto.PubKey, valAddr common.Address) error {
	store := ctx.KVStore(k.storeKey)

	var validator types.Validator

	// get validator and unmarshall
	validatorBytes := store.Get(getValidatorKey(valAddr.Bytes()))
	if validatorBytes == nil {
		err := fmt.Errorf("Validator Not Found")
		return err
	}

	err := k.cdc.UnmarshalBinary(validatorBytes, &validator)
	if err != nil {
		StakingLogger.Error("Error unmarshalling validator while fetching validator from store", "Error", err, "ValidatorAddress", valAddr)
		return err
	}
	//update validator
	validator.Signer = signer
	validator.Pubkey = pubkey

	// add updated validator to store with same key
	k.AddValidator(ctx, validator)

	return nil
}
