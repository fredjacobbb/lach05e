package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// updatePayloadsCmd represents the updatePayloads command
var updatePayloadsCmd = &cobra.Command{
	Use: "update",
	Run: func(cmd *cobra.Command, args []string) {

		if _, errGit := exec.LookPath("git"); errGit != nil {
			fmt.Println("GIT is not installed or not reachable.")
			return
		}

		_, errFolder := os.Stat("SecLists")

		if errFolder != nil {
			err := exec.Command("git", "clone", updatePayloadsURL).Run()
			if err != nil {
				fmt.Println("- git clone - ERROR : ", err)
				return
			}
		}

		err := exec.Command("git", "-C", "SecLists", "pull").Run()
		if err != nil {
			fmt.Println("- git pull - ERROR : ", err)
			return
		}

		fmt.Println("Payloads update!")
	},
}

func init() {
	rootCmd.AddCommand(updatePayloadsCmd)
}
