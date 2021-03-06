package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	global_start := time.Now()
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Println(err)
		}
		// Prevent resource drain
		resp.Body.Close()
		secs := time.Since(start).Seconds()
		fmt.Println("Page status: ", resp.Status)
		fmt.Println("Processing time: ", secs)
	}
	global_secs := time.Since(global_start).Seconds()
	fmt.Println("Total processing time: ", global_secs)
}
