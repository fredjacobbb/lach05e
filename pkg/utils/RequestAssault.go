package utils

import (
	"fmt"
	"net/http"
)

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
	}

	req, err := http.NewRequest(r.Method, "https://"+r.Url+r.Payload, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "How are You ?")
	req.Header.Set("Cookie", "Hey ok cookie")

	fmt.Println(req)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	fmt.Println(resp.Header)
}
