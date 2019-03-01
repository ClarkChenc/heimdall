package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"

	hmCommon "github.com/maticnetwork/heimdall/common"
	"github.com/maticnetwork/heimdall/types"
	hmTypes "github.com/maticnetwork/heimdall/types"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	// Get all delegations from a delegator
	r.HandleFunc(
		"/staking/signer/{address}",
		validatorByAddressHandlerFn(cdc, cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/staking/validator/{id}",
		validatorByIDHandlerFn(cdc, cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/staking/validator-set",
		validatorSetHandlerFn(cdc, cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/staking/proposer/{times}",
		proposerHandlerFn(cdc, cliCtx),
	).Methods("GET")
}

// Returns validator information by signer address
func validatorByAddressHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		signerAddress := common.HexToAddress(vars["address"])

		res, err := cliCtx.QueryStore(hmCommon.GetValidatorKey(signerAddress.Bytes()), "staker")
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// the query will return empty if there is no data
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("Validator not found"))
			return
		}

		var _validator types.Validator
		err = cdc.UnmarshalBinary(res, &_validator)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		result, err := json.Marshal(_validator)
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

	}
}

// Returns validator information by val ID
func validatorByIDHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		signerAddr, err := cliCtx.QueryStore(hmCommon.GetValidatorMapKey(hmTypes.NewValidatorID(id).Bytes()), "staker")
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, err := cliCtx.QueryStore(hmCommon.GetValidatorKey(signerAddr), "staker")
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// the query will return empty if there is no data
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("Validator not found"))
			return
		}

		var _validator types.Validator
		err = cdc.UnmarshalBinary(res, &_validator)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		result, err := json.Marshal(_validator)
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

	}
}

// get current validator set
func validatorSetHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, err := cliCtx.QueryStore(hmCommon.CurrentValidatorSetKey, "staker")
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// the query will return empty if there is no data
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		var _validatorSet hmTypes.ValidatorSet
		cdc.UnmarshalBinary(res, &_validatorSet)

		// todo format validator set to remove pubkey like we did for validator
		result, err := json.Marshal(&_validatorSet)
		if err != nil {
			RestLogger.Error("Error while marshalling resposne to Json", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)

	}
}

// get proposer for current validator set
func proposerHandlerFn(
	cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		times, err := strconv.Atoi(vars["times"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		RestLogger.Debug("Calculating proposers", "Count", times)

		res, err := cliCtx.QueryStore(hmCommon.CurrentValidatorSetKey, "staker")
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		// the query will return empty if there is no data
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var _validatorSet hmTypes.ValidatorSet
		err = cdc.UnmarshalBinary(res, &_validatorSet)
		if err != nil {
			RestLogger.Error("Error while marshalling validator set", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// proposers
		var proposers []hmTypes.Validator

		for index := 0; index < times; index++ {
			RestLogger.Info("Getting proposer for current validator set", "Index", index, "TotalProposers", times)
			proposers = append(proposers, *(_validatorSet.GetProposer()))
			_validatorSet.IncrementAccum(1)
		}

		// todo format validator set to remove pubkey like we did for validator
		result, err := json.Marshal(&proposers)
		if err != nil {
			RestLogger.Error("Error while marshalling response to Json", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}