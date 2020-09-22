package cli

import (
	slashingcli "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	return slashingcli.GetQueryCmd()
}
