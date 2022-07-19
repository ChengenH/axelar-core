package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/axelarnetwork/axelar-core/x/evm/types"
)

const (
	flagAddress = "address"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	evmTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		TraverseChildren:           true,
		RunE:                       client.ValidateCmd,
	}

	evmTxCmd.AddCommand(
		GetCmdSetGateway(),
		GetCmdLink(),
		GetCmdConfirmERC20TokenDeployment(),
		GetCmdConfirmERC20Deposit(),
		GetCmdConfirmTransferOperatorship(),
		GetCmdCreateConfirmGatewayTx(),
		GetCmdCreatePendingTransfers(),
		GetCmdCreateDeployToken(),
		GetCmdCreateBurnTokens(),
		GetCmdCreateTransferOwnership(),
		GetCmdCreateTransferOperatorship(),
		GetCmdSignCommands(),
		GetCmdAddChain(),
		GetCmdRetryFailedEvent(),
	)

	return evmTxCmd
}

// GetCmdSetGateway sets the gateway address for the given evm chain
func GetCmdSetGateway() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-gateway [chain] [address]",
		Short: "Set the gateway address for the given evm chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			addressHex := args[1]
			if !common.IsHexAddress(addressHex) {
				return fmt.Errorf("invalid address %s", addressHex)
			}

			msg := types.NewSetGatewayRequest(cliCtx.GetFromAddress(), chain, types.Address(common.HexToAddress(addressHex)))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdLink links a cross chain address to an EVM chain address created by Axelar
func GetCmdLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link [chain] [recipient chain] [recipient address] [asset name]",
		Short: "Link a cross chain address to an EVM chain address created by Axelar",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewLinkRequest(cliCtx.GetFromAddress(), args[0], args[1], args[2], args[3])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdConfirmERC20TokenDeployment returns the cli command to confirm a ERC20 token deployment
func GetCmdConfirmERC20TokenDeployment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-erc20-token [chain] [origin chain] [origin asset] [txID]",
		Short: "Confirm an ERC20 token deployment in an EVM chain transaction for a given asset of some origin chain and gateway address",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			originChain := args[1]
			originAsset := args[2]
			asset := types.NewAsset(originChain, originAsset)
			txID := common.HexToHash(args[3])
			msg := types.NewConfirmTokenRequest(cliCtx.GetFromAddress(), chain, asset, txID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdConfirmERC20Deposit returns the cli command to confirm an ERC20 deposit
func GetCmdConfirmERC20Deposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-erc20-deposit [chain] [txID] [burnerAddr]",
		Short: "Confirm ERC20 deposits in an EVM chain transaction to a burner address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			txID := common.HexToHash(args[1])
			burnerAddr := common.HexToAddress(args[2])

			msg := types.NewConfirmDepositRequest(cliCtx.GetFromAddress(), chain, txID, burnerAddr)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdConfirmTransferOperatorship returns the cli command to confirm a transfer operatorship for the gateway contract
func GetCmdConfirmTransferOperatorship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-transfer-operatorship [chain] [txID]",
		Short: "Confirm a transfer operatorship in an EVM chain transaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			txID := common.HexToHash(args[1])
			msg := types.NewConfirmTransferKeyRequest(cliCtx.GetFromAddress(), chain, txID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateConfirmGatewayTx returns the cli command to confirm a gateway transaction
func GetCmdCreateConfirmGatewayTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-gateway-tx [chain] [txID]",
		Short: "Confirm a gateway transaction in an EVM chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			txID := types.Hash(common.HexToHash(args[1]))

			msg := types.NewConfirmGatewayTxRequest(cliCtx.GetFromAddress(), chain, txID)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreatePendingTransfers returns the cli command to create commands for handling all pending token transfers to an EVM chain
func GetCmdCreatePendingTransfers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pending-transfers [chain]",
		Short: "Create commands for handling all pending transfers to an EVM chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewCreatePendingTransfersRequest(cliCtx.GetFromAddress(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateDeployToken returns the cli command to create deploy-token command for an EVM chain
func GetCmdCreateDeployToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deploy-token [evm chain] [origin chain] [origin asset] [token name] [symbol] [decimals] [capacity] [dailyMintLimit]",
		Short: "Create a deploy token command with the AxelarGateway contract",
		Args:  cobra.ExactArgs(8),
	}
	address := cmd.Flags().String(flagAddress, types.ZeroAddress.Hex(), "existing ERC20 token's address")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		cliCtx, err := client.GetClientTxContext(cmd)
		if err != nil {
			return err
		}

		chain := args[0]
		originChain := args[1]
		originAsset := args[2]
		tokenName := args[3]
		symbol := args[4]
		decs, err := strconv.ParseUint(args[5], 10, 8)
		if err != nil {
			return fmt.Errorf("could not parse decimals")
		}
		capacity, ok := sdk.NewIntFromString(args[6])
		if !ok {
			return fmt.Errorf("could not parse capacity")
		}

		if !common.IsHexAddress(*address) {
			return fmt.Errorf("could not parse address")
		}

		dailyMintLimit := args[7]

		asset := types.NewAsset(originChain, originAsset)
		tokenDetails := types.NewTokenDetails(tokenName, symbol, uint8(decs), capacity)
		msg := types.NewCreateDeployTokenRequest(cliCtx.GetFromAddress(), chain, asset, tokenDetails, types.Address(common.HexToAddress(*address)), dailyMintLimit)
		if err = msg.ValidateBasic(); err != nil {
			return err
		}

		return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateBurnTokens returns the cli command to create burn commands for all confirmed token deposits in an EVM chain
func GetCmdCreateBurnTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-burn-tokens [chain]",
		Short: "Create burn commands for all confirmed token deposits in an EVM chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewCreateBurnTokensRequest(cliCtx.GetFromAddress(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateTransferOwnership returns the cli command to create transfer-ownership command for an EVM chain contract
func GetCmdCreateTransferOwnership() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-ownership [chain] [keyID]",
		Short: "Create transfer ownership command for an EVM chain contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewCreateTransferOwnershipRequest(cliCtx.GetFromAddress(), args[0], args[1])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateTransferOperatorship returns the cli command to create transfer-operatorship command for an EVM chain contract
func GetCmdCreateTransferOperatorship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-operatorship [chain] [keyID]",
		Short: "Create transfer operatorship command for an EVM chain contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewCreateTransferOperatorshipRequest(cliCtx.GetFromAddress(), args[0], args[1])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdSignCommands returns the cli command to sign pending commands for an EVM chain contract
func GetCmdSignCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-commands [chain]",
		Short: "Sign pending commands for an EVM chain contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewSignCommandsRequest(cliCtx.GetFromAddress(), args[0])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdAddChain returns the cli command to add a new evm chain command
func GetCmdAddChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-chain [name] [key type] [chain config]",
		Short: "Add a new EVM chain",
		Long:  "Add a new EVM chain. The chain config parameter should be the path to a json file containing the key requirements and the evm module parameters",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			name := args[0]
			jsonFile := args[2]

			byteValue, err := ioutil.ReadFile(jsonFile)
			if err != nil {
				return err
			}
			var chainConf struct {
				Params types.Params `json:"params"`
			}
			err = json.Unmarshal([]byte(byteValue), &chainConf)
			if err != nil {
				return err
			}

			msg := types.NewAddChainRequest(cliCtx.GetFromAddress(), name, chainConf.Params)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdRetryFailedEvent returns the cli command to retry a failed event
func GetCmdRetryFailedEvent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retry-event [chain] [event ID]",
		Short: "Retry a failed event",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewRetryFailedEventRequest(cliCtx.GetFromAddress(), args[0], args[1])
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
