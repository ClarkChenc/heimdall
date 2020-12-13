package cli

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maticnetwork/bor/common"

	// "github.com/maticnetwork/heimdall/bridge/setu/util"
	"github.com/maticnetwork/heimdall/contracts/stakinginfo"
	"github.com/maticnetwork/heimdall/helper"
	hmTypes "github.com/maticnetwork/heimdall/types/common"
	"github.com/maticnetwork/heimdall/x/staking/types"
)

var logger = helper.Logger.With("module", "staking/client/cli")

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	stakingTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	stakingTxCmd.AddCommand(
		ValidatorJoinTxCmd(),
		StakeUpdateTxCmd(),
		SignerUpdateTxCmd(),
		ValidatorExitTxCmd(),
	)

	return stakingTxCmd
}

// ValidatorJoinTxCmd send validator join message
func ValidatorJoinTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-join",
		Short: "Join Heimdall as a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// get proposer
			proposer := hmTypes.HexToAccAddress(viper.GetString(FlagProposerAddress))
			if proposer.Empty() {
				proposer = helper.GetFromAddress(clientCtx)
			}

			// get txHash
			txhash := viper.GetString(FlagTxHash)
			if txhash == "" {
				return fmt.Errorf("transaction hash is required")
			}

			// get PubKey string
			pubkeyStr := viper.GetString(FlagSignerPubkey)
			if pubkeyStr == "" {
				return fmt.Errorf("pubkey is required")
			}

			// convert PubKey to bytes
			pubkeyBytes := common.FromHex(pubkeyStr)
			if len(pubkeyBytes) != 65 {
				return fmt.Errorf("Invalid public key length")
			}
			pubkey := hmTypes.NewPubKey(pubkeyBytes)

			// total stake amount
			amount, ok := sdk.NewIntFromString(viper.GetString(FlagAmount))
			if !ok {
				return errors.New("Invalid stake amount")
			}

			// Get contractCaller ref
			contractCallerObj, err := helper.NewContractCaller()
			if err != nil {
				return err
			}
			//ToDO
			// chainmanagerParams, err := util.GetChainmanagerParams(cliCtx)
			// if err != nil {
			// 	return err
			// }

			// get main tx receipt
			// NOTE: Use 'chainmanagerParams.MainchainTxConfirmations'. Now it is hard coded.
			receipt, err := contractCallerObj.GetConfirmedTxReceipt(hmTypes.HexToHeimdallHash(txhash).EthHash(), 6)
			if err != nil || receipt == nil {
				return errors.New("Transaction is not confirmed yet. Please wait for sometime and try again")
			}

			abiObject := &contractCallerObj.StakingInfoABI
			eventName := "Staked"
			event := new(stakinginfo.StakinginfoStaked)
			var logIndex uint64
			found := false
			for _, vLog := range receipt.Logs {
				topic := vLog.Topics[0].Bytes()
				selectedEvent := helper.EventByID(abiObject, topic)
				if selectedEvent != nil && selectedEvent.Name == eventName {
					if err := helper.UnpackLog(abiObject, event, eventName, vLog); err != nil {
						return err
					}

					logIndex = uint64(vLog.Index)
					found = true
					break
				}
			}

			if !found {
				return fmt.Errorf("Invalid tx for validator join")
			}

			if !bytes.Equal(event.SignerPubkey, pubkey.Bytes()[1:]) {
				return fmt.Errorf("Public key mismatch with event log")
			}

			// msg new ValidatorJion message
			msg := types.NewMsgValidatorJoin(
				proposer,
				event.ValidatorId.Uint64(),
				viper.GetUint64(FlagActivationEpoch),
				amount,
				pubkey,
				hmTypes.HexToHeimdallHash(txhash),
				logIndex,
				viper.GetUint64(FlagBlockNumber),
				event.Nonce.Uint64(),
			)

			// broadcast message
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().StringP(FlagProposerAddress, "p", "", "--proposer=<proposer-address>")
	cmd.Flags().String(FlagSignerPubkey, "", "--signer-pubkey=<signer pubkey here>")
	cmd.Flags().String(FlagTxHash, "", "--tx-hash=<transaction-hash>")
	cmd.Flags().Uint64(FlagBlockNumber, 0, "--block-number=<block-number>")
	cmd.Flags().String(FlagAmount, "0", "--amount=<amount>")
	cmd.Flags().Uint64(FlagActivationEpoch, 0, "--activation-epoch=<activation-epoch>")

	if err := cmd.MarkFlagRequired(FlagBlockNumber); err != nil {
		logger.Error("SendValidatorJoinTx | MarkFlagRequired | FlagBlockNumber", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagActivationEpoch); err != nil {
		logger.Error("SendValidatorJoinTx | MarkFlagRequired | FlagActivationEpoch", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagAmount); err != nil {
		logger.Error("SendValidatorJoinTx | MarkFlagRequired | FlagAmount", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagSignerPubkey); err != nil {
		logger.Error("SendValidatorJoinTx | MarkFlagRequired | FlagSignerPubkey", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagTxHash); err != nil {
		logger.Error("SendValidatorJoinTx | MarkFlagRequired | FlagTxHash", "Error", err)
	}
	return cmd
}

// SignerUpdateTxCmd send singer update transaction
func SignerUpdateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signer-update",
		Short: "Update signer for a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// get proposer
			proposer := hmTypes.HexToAccAddress(viper.GetString(FlagProposerAddress))
			if proposer.Empty() {
				proposer = helper.GetFromAddress(clientCtx)
			}

			// get validatorID from flags
			ValidatorID := viper.GetUint64(FlagValidatorID)
			if ValidatorID == 0 {
				return fmt.Errorf("validator ID cannot be 0")
			}

			// get PubKey string from flag
			pubkeyStr := viper.GetString(FlagNewSignerPubkey)
			if pubkeyStr == "" {
				return fmt.Errorf("Pubkey has to be supplied")
			}

			// convert PubKey string to bytes
			pubkeyBytes, err := hex.DecodeString(pubkeyStr)
			if err != nil {
				return err
			}
			pubkey := hmTypes.NewPubKey(pubkeyBytes)

			// get txHash from flag
			txhash := viper.GetString(FlagTxHash)
			if txhash == "" {
				return fmt.Errorf("transaction hash has to be supplied")
			}

			// draft new SingerUpdate message
			msg := types.NewMsgSignerUpdate(
				proposer,
				ValidatorID,
				pubkey,
				hmTypes.HexToHeimdallHash(txhash),
				viper.GetUint64(FlagLogIndex),
				viper.GetUint64(FlagBlockNumber),
				viper.GetUint64(FlagNonce),
			)

			// broadcast messages
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().StringP(FlagProposerAddress, "p", "", "--proposer=<proposer-address>")
	cmd.Flags().Uint64(FlagValidatorID, 0, "--id=<validator-id>")
	cmd.Flags().String(FlagNewSignerPubkey, "", "--new-pubkey=<new-signer-pubkey>")
	cmd.Flags().String(FlagTxHash, "", "--tx-hash=<transaction-hash>")
	cmd.Flags().Uint64(FlagLogIndex, 0, "--log-index=<log-index>")
	cmd.Flags().Uint64(FlagBlockNumber, 0, "--block-number=<block-number>")
	cmd.Flags().Int(FlagNonce, 0, "--nonce=<nonce>")

	if err := cmd.MarkFlagRequired(FlagValidatorID); err != nil {
		logger.Error("SendValidatorUpdateTx | MarkFlagRequired | FlagValidatorID", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagTxHash); err != nil {
		logger.Error("SendValidatorUpdateTx | MarkFlagRequired | FlagTxHash", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagNewSignerPubkey); err != nil {
		logger.Error("SendValidatorUpdateTx | MarkFlagRequired | FlagNewSignerPubkey", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagLogIndex); err != nil {
		logger.Error("SendValidatorUpdateTx | MarkFlagRequired | FlagLogIndex", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagBlockNumber); err != nil {
		logger.Error("SendValidatorUpdateTx | MarkFlagRequired | FlagBlockNumber", "Error", err)
	}
	if err := cmd.MarkFlagRequired(FlagNonce); err != nil {
		logger.Error("SendValidatorUpdateTx | MarkFlagRequired | FlagNonce", "Error", err)
	}

	return cmd
}

// StakeUpdateTxCmd send stake update transaction
func StakeUpdateTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-update",
		Short: "Update stake for a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// get proposer
			proposer := hmTypes.HexToAccAddress(viper.GetString(FlagProposerAddress))
			if proposer.Empty() {
				proposer = helper.GetFromAddress(clientCtx)
			}

			// get validatorID from flag
			validatorID := viper.GetUint64(FlagValidatorID)
			if validatorID == 0 {
				return fmt.Errorf("validator ID cannot be 0")
			}

			// get txHash from flag
			txhash := viper.GetString(FlagTxHash)
			if txhash == "" {
				return fmt.Errorf("transaction hash has to be supplied")
			}

			// total stake amount
			amount, ok := sdk.NewIntFromString(viper.GetString(FlagAmount))
			if !ok {
				return errors.New("Invalid new stake amount")
			}

			// draft new StakeUpdate message
			msg := types.NewMsgStakeUpdate(
				proposer,
				validatorID,
				amount,
				hmTypes.HexToHeimdallHash(txhash),
				viper.GetUint64(FlagLogIndex),
				viper.GetUint64(FlagBlockNumber),
				viper.GetUint64(FlagNonce),
			)

			// broadcast message
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().StringP(FlagProposerAddress, "p", "", "--proposer=<proposer-address>")
	cmd.Flags().Uint64(FlagValidatorID, 0, "--id=<validator-id>")
	cmd.Flags().String(FlagTxHash, "", "--tx-hash=<transaction-hash>")
	cmd.Flags().String(FlagAmount, "", "--amount=<amount>")
	cmd.Flags().Uint64(FlagLogIndex, 0, "--log-index=<log-index>")
	cmd.Flags().Uint64(FlagBlockNumber, 0, "--block-number=<block-number>")
	cmd.Flags().Int(FlagNonce, 0, "--nonce=<nonce>")

	cmd.MarkFlagRequired(FlagTxHash)
	cmd.MarkFlagRequired(FlagLogIndex)
	cmd.MarkFlagRequired(FlagValidatorID)
	cmd.MarkFlagRequired(FlagBlockNumber)
	cmd.MarkFlagRequired(FlagAmount)
	cmd.MarkFlagRequired(FlagNonce)

	return cmd
}

// ValidatorExitTxCmd sends validator exit transaction
func ValidatorExitTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-exit",
		Short: "Exit heimdall as a validator ",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			// get proposer
			proposer := hmTypes.HexToAccAddress(viper.GetString(FlagProposerAddress))
			if proposer.Empty() {
				proposer = helper.GetFromAddress(clientCtx)
			}

			// get validatorid from flag
			validatorID := viper.GetUint64(FlagValidatorID)
			if validatorID == 0 {
				return fmt.Errorf("validator ID cannot be 0")
			}

			// get txHash from flag
			txhash := viper.GetString(FlagTxHash)
			if txhash == "" {
				return fmt.Errorf("transaction hash has to be supplied")
			}

			// get nonce from flag
			nonce := viper.GetUint64(FlagNonce)

			// draf new ValidatorExit message
			msg := types.NewMsgValidatorExit(
				proposer,
				validatorID,
				viper.GetUint64(FlagDeactivationEpoch),
				hmTypes.HexToHeimdallHash(txhash),
				viper.GetUint64(FlagLogIndex),
				viper.GetUint64(FlagBlockNumber),
				nonce,
			)

			// broadcast message
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().StringP(FlagProposerAddress, "p", "", "--proposer=<proposer-address>")
	cmd.Flags().Uint64(FlagValidatorID, 0, "--id=<validator ID here>")
	cmd.Flags().String(FlagTxHash, "", "--tx-hash=<transaction-hash>")
	cmd.Flags().Uint64(FlagLogIndex, 0, "--log-index=<log-index>")
	cmd.Flags().Uint64(FlagDeactivationEpoch, 0, "--deactivation-epoch=<deactivation-epoch>")
	cmd.Flags().Uint64(FlagBlockNumber, 0, "--block-number=<block-number>")
	cmd.Flags().Int(FlagNonce, 0, "--nonce=<nonce>")

	cmd.MarkFlagRequired(FlagValidatorID)
	cmd.MarkFlagRequired(FlagTxHash)
	cmd.MarkFlagRequired(FlagLogIndex)
	cmd.MarkFlagRequired(FlagBlockNumber)
	cmd.MarkFlagRequired(FlagNonce)

	return cmd
}
