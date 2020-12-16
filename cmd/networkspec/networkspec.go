package networkspec

import (
	"github.com/spf13/cobra"
)

var wsID int64

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "network-spec",
		Short: "Manage network specs in the OnFinality platform",
	}
	cmd.PersistentFlags().Int64VarP(&wsID, "workspace", "w", 0, "Workspace ID")

	cmd.AddCommand(
		listCmd(),
		CreateCmd(),
		DeleteCmd(),
		ShowCmd(),
	)
	return cmd
}
