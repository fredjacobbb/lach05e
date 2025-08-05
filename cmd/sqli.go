package cmd

import (
	"bufio"
	"fmt"
	"lach05e/pkg/utils"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var payload string

// sqliCmd represents the sqli command
var sqliCmd = &cobra.Command{
	Use: "sqli",
	Run: func(cmd *cobra.Command, args []string) {
		start_injection()
	},
}

func start_injection() {
	file, err := os.Open(payloadsPath)
	if err != nil {
		fmt.Println("Erreur : ", err)
	}
	if payloadsPath == "" {
		fmt.Println("Oops you've forgotten the Payloads Path !")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payload = strings.TrimSpace(scanner.Text())
		req := utils.Request{
			Url:          target,
			Method:       method,
			Ua:           ua,
			PayloadsPath: payloadsPath,
			Cookie:       cookie,
			Payload:      payload,
			Header:       header,
		}
		utils.RequestAssault(req)
	}

}

func init() {
	rootCmd.AddCommand(sqliCmd)
}
