package cli

import (
	upgradecli "github.com/cosmos/cosmos-sdk/x/upgrade/client/cli"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for upgrade module
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Querying commands for the upgrade module",
	}
	queryCmd.AddCommand(
		upgradecli.GetCurrentPlanCmd(),
		upgradecli.GetAppliedPlanCmd(),
	)
	return queryCmd
}
