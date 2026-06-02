package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	statusURL = "https://status.claude.com/"
	timeout   = 10 * time.Second
)

func main() {
	fmt.Printf("Checking Claude server status at: %s\n", statusURL)

	client := &http.Client{Timeout: timeout}
	start := time.Now()

	resp, err := client.Get(statusURL)
	elapsed := time.Since(start).Milliseconds()

	if err != nil {
		fmt.Printf("[DOWN] Connection failed: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Printf("[UP] Status: %d | Response time: %dms\n", resp.StatusCode, elapsed)
	} else {
		fmt.Printf("[DOWN] HTTP error: %d %s | Response time: %dms\n", resp.StatusCode, resp.Status, elapsed)
		os.Exit(1)
	}
}
