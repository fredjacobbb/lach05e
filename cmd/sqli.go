package cmd

import (
	"bufio"
	"fmt"
	"lach05e/pkg/utils"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// sqliCmd represents the sqli command
var sqliCmd = &cobra.Command{
	Use: "sqli",
	Run: func(cmd *cobra.Command, args []string) {
		req := utils.Request{
			Url:          target,
			Method:       method,
			Ua:           ua,
			PayloadsPath: payloadsPath,
			Cookie:       cookie,
			Payload:      payload,
		}
		if req.PayloadsPath == "" {
			fmt.Println("Oops you've forgotten the Payloads Path !")
			return
		}
		start_injection()
	},
}

func start_injection() {
	file, err := os.Open(payloadsPath)
	if err != nil {
		fmt.Println("Erreur : ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payload := strings.TrimSpace(scanner.Text())
		time.Sleep(1 * time.Second)
		fmt.Println(payload)
	}

}

func init() {
	rootCmd.AddCommand(sqliCmd)
}
