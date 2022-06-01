package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Println(err)
		}
	}
}
