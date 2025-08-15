package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var target string
var payloadsPath string
var method string
var ua string
var cookie []string
var header []string
var data string
var bufferSize int

var updatePayloadsURL = "https://github.com/danielmiessler/SecLists"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "lach05e",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func banner() {
	banner, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("LA", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("CH05E", pterm.FgLightMagenta.ToStyle()),
	).Srender()
	fmt.Printf("\n")
	fmt.Println(banner)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	banner()
	rootCmd.PersistentFlags().StringVarP(&target, "target", "u", "", "Target")
	rootCmd.PersistentFlags().StringVarP(&payloadsPath, "payloadsPath", "p", "", "PayloadsPath")
	rootCmd.PersistentFlags().StringVarP(&method, "method", "X", "", "Method")
	rootCmd.PersistentFlags().StringSliceVarP(&header, "header", "H", []string{}, "Header")
	rootCmd.PersistentFlags().StringSliceVarP(&cookie, "cookie", "C", []string{}, "Cookie")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "Data")
	rootCmd.PersistentFlags().IntVarP(&bufferSize, "threads", "T", 1, "Threads")
}
