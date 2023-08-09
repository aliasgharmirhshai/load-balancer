package main

import (
	"fmt"
	"net/http"
)

func sendReq() int {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:99.0) Gecko/20200101 Firefox/99.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer resp.Body.Close()

	return resp.StatusCode
}


func main() {
	sendReq()
}
