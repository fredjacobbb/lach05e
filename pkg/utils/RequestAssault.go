package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

type Request struct {
	Url          string
	Method       string
	Ua           string
	PayloadsPath string
	Cookie       []string
	Payload      string
	Header       []string
	Data         string
}

var body io.Reader

func RequestAssault(r Request) {

	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   5 * time.Second,
	}

	if r.Method == "POST" {
		r.Data = r.Payload
		body = strings.NewReader(r.Data)
		r.Method = "POST"
		r.Payload = ""
	} else {
		r.Method = "GET"
	}

	req, err := http.NewRequest(r.Method, r.Url+"/"+r.Payload, body)
	if err != nil {
		fmt.Printf("Domain is down or inexist !\n")
		os.Exit(1)
	}

	for _, c := range r.Cookie {
		req.Header.Add("Cookie", c)
	}

	// POST DATA

	fmt.Println(io.ReadAll(req.Body))

	if r.Ua == "" {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:39.0) Gecko/20100101 Firefox/39.0")
	}

	for _, h := range r.Header {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			pterm.Warning.Printfln("Format error ! %s", h)
			continue
		}
		req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Domain is down or inexist ! \n")
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		pterm.Printf("%s %s\n", pterm.NewStyle(pterm.BgRed, pterm.FgWhite).Sprint(resp.Status), pterm.NewStyle(pterm.FgWhite).Sprint(strings.TrimSpace(r.Url)+"/"+strings.TrimSpace(r.Payload)))
		return
	}

	pterm.FgGreen.Printf("%s %s\n", pterm.NewStyle(pterm.BgGreen, pterm.FgDarkGray).Sprint(resp.Status), pterm.NewStyle(pterm.FgWhite).Sprint(strings.TrimSpace(r.Url)+strings.TrimSpace(r.Payload)))

}
