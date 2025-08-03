package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

var errorStyle = pterm.NewStyle(pterm.FgWhite, pterm.BgRed)

type Request struct {
	Url          string
	Method       string
	Ua           string
	PayloadsPath string
	Cookie       string
	Payload      string
}

func RequestAssault(r Request) {
	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   5 * time.Second,
	}

	req, err := http.NewRequest(r.Method, "https://"+r.Url+r.Payload, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "How are You ?")
	req.Header.Set("Cookie", "Hey ok cookie")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		pterm.Printf("%s %s\n", pterm.NewStyle(pterm.BgRed, pterm.FgWhite).Sprint(resp.Status), pterm.NewStyle(pterm.FgWhite).Sprint("https://"+strings.TrimSpace(r.Url)+strings.TrimSpace(r.Payload)))
		return
	}

	pterm.FgGreen.Printf("%s %s\n", pterm.NewStyle(pterm.BgGreen, pterm.FgWhite).Sprint(resp.Status), pterm.NewStyle(pterm.FgWhite).Sprint("https://"+strings.TrimSpace(r.Url)+strings.TrimSpace(r.Payload)))

}
