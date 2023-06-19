package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/bianjieai/iritamod/modules/params/types"
)

func NewTxCmd() *cobra.Command {
	paramsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Params transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	paramsTxCmd.AddCommand(
		NewCmdUpdateParams(),
		//NewDraftUpdateParams(),
	)

	return paramsTxCmd
}

// NewCmdUpdateParams implements the update params command handler.
func NewCmdUpdateParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [params.json]",
		Short: "update module params",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msgs, err := parseUpdateParamsFile(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			msg, err := types.NewMsgUpdateParams(msgs, clientCtx.GetFromAddress().String())

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

//func NewDraftUpdateParams() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "draft",
//		Short: "Generate a draft update params file",
//		RunE: func(cmd *cobra.Command, _ []string) error {
//			clientCtx, err := client.GetClientTxContext(cmd)
//			if err != nil {
//				return err
//			}
//
//			var updateParam updateParamsType
//			// select available update params message
//			msgPrompt := promptui.Select{
//				Label: "Select update params message type",
//				Items: func() []string {
//					umsgs := make([]string, 0)
//					msgs := clientCtx.InterfaceRegistry.ListImplementations(sdk.MsgInterfaceProtoName)
//					// filter only update params msg types out
//					for _, msg := range msgs {
//						if strings.HasSuffix(msg, types.ParamsMsgSuffix) && msg != types.ParamsMsgTypeURL {
//							umsgs = append(umsgs, msg)
//						}
//					}
//					sort.Strings(umsgs)
//					return umsgs
//				},
//			}
//
//			_, result, err := msgPrompt.Run()
//			if err != nil {
//				return fmt.Errorf("failed to prompt update params types: %w", err)
//			}
//
//			updateParam.MsgType = result
//
//			if updateParam.MsgType != "" {
//				updateParam.Msg, err = sdk.GetMsgFromTypeURL(clientCtx.Codec, updateParam.MsgType)
//				if err != nil {
//					// should never happen
//					panic(err)
//				}
//			}
//
//			res, err := updateParam.Prompt(clientCtx.Codec)
//			if err != nil {
//				return err
//			}
//
//			if err := writeFile(draftUpdateParamsFileName, res); err != nil {
//				return err
//			}
//
//			fmt.Printf("The draft update params file has successfully been generated.")
//
//			return nil
//
//			return nil
//		},
//	}
//
//	return cmd
//}
