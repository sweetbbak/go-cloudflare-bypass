package main

import (
	"fmt"
	"go-cloudflare-bypass/reqwest"
	"net/http"
	"os"
)

func main() {
	u := os.Args[1]
	cl := tls_request()
	req, err := http.NewRequest("GET", u, nil)

	// user agent must be set
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(req.Body)
}
