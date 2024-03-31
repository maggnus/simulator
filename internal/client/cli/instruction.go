package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"simulator/api"
)

var instructionCmd = &cobra.Command{
	Use:          "instruction",
	Short:        fmt.Sprintf("Robot movement instruction %v", commands.GetAll()),
	Args:         cobra.OnlyValidArgs,
	ValidArgs:    commands.GetAll(),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("instruction is not defined %v", commands.GetAll())
		}

		return Client.Command(api.Command(args[0]))
	},
}
