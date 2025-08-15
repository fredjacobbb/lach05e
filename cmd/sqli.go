package cmd

import (
	"bufio"
	"fmt"
	"io"
	"lach05e/pkg/utils"
	"os"
	"strings"
	"sync"

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
		fmt.Println("No payload path given ! We sending Request like that !")
		return
	}

	defer file.Close()

	scannerThread := bufio.NewScanner(file)

	for scannerThread.Scan() {
		utils.PayloadsLines++
	}

	_, err = file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)

	sem := make(chan struct{}, bufferSize)
	var wg sync.WaitGroup

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
			Data:         data,
		}

		sem <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			utils.RequestAssault(req)
			utils.CurrentPayloadLine++
		}()
	}

	wg.Wait()

}

func init() {
	rootCmd.AddCommand(sqliCmd)
}
