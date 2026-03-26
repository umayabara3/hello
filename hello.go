package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// コンテキストをどうやって使うか？
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://httpbin.org/delay/3", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out")
		} else {
			fmt.Println("Error:", err)
		}
		return
	}
	defer resp.Body.Close()

	fmt.Println("Request succeeded")
}
