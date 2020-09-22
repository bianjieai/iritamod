package cli

import (
	"github.com/spf13/cobra"

	slashingcli "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	return slashingcli.GetQueryCmd()
}
