package networkspec

import (
	"fmt"

	"github.com/OnFinality-io/onf-cli/cmd/helpers"
	"github.com/OnFinality-io/onf-cli/pkg/printer"
	"github.com/OnFinality-io/onf-cli/pkg/service"
	"github.com/spf13/cobra"
)

func BootstrapCmd() *cobra.Command {
	var filePath, networkID string
	c := &cobra.Command{
		Use:   "bootstrap",
		Short: "bootstrap chain spec",
		Run: func(cmd *cobra.Command, args []string) {
			wsID, err := helpers.GetWorkspaceID(cmd)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			payload := &service.BootstrapChainSpecPayload{}
			err = helpers.ApplyDefinitionFile(filePath, payload)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			specs, err := service.BootstrapChainSpec(wsID, networkID, payload)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			printer.New().Print(specs)
		},
	}
	c.Flags().StringVarP(&filePath, "file", "f", "", "definition file for create network, yaml or json")
	c.Flags().StringVarP(&networkID, "network", "n", "", "Network id")
	_ = c.MarkFlagRequired("network")
	return c
}
