package cmd

import (
	"lach05e/pkg/utils"

	"github.com/spf13/cobra"
)

// corsCmd represents the cors command
var corsCmd = &cobra.Command{
	Use: "cors",
	Run: func(cmd *cobra.Command, args []string) {
		req := &utils.Request{
			Url:    target,
			Method: method,
			Ua:     ua,
		}
		utils.RequestAssault(req)
	},
}

func init() {
	rootCmd.AddCommand(corsCmd)
}
