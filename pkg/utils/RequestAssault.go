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
var PayloadsLines uint
var CurrentPayloadLine uint = 1

func RequestAssault(r Request) {

	client := &http.Client{
		Transport: &http.Transport{},
		Timeout:   5 * time.Second,
	}

	if r.Method == "POST" {
		r.Data = strings.Replace(r.Data, "FUZZ", r.Payload, 1)
		r.Payload = ""
		body = strings.NewReader(r.Data)
	}

	req, err := http.NewRequest(r.Method, r.Url+"/"+r.Payload, body)
	if err != nil {
		fmt.Printf("Domain is down or inexist !\n")
		os.Exit(1)
	}

	for _, c := range r.Cookie {
		req.Header.Add("Cookie", c)
	}

	req.Body = io.NopCloser(body)

	// dump, err := httputil.DumpRequest(req, true)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

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

	fmt.Printf("\n\nRequest : %d / %d \n", CurrentPayloadLine, PayloadsLines)
	// fmt.Println(string(dump))

	pterm.FgGreen.Printf("%s %s\n", pterm.NewStyle(pterm.BgGreen, pterm.FgDarkGray).Sprint(resp.Status), pterm.NewStyle(pterm.FgWhite).Sprint(strings.TrimSpace(r.Url)+strings.TrimSpace(r.Payload)))

}
