package main

import (
	"fmt"
	"net/http"
	"sync"
)

func sendRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Send an HTTP GET request to the target URL
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Request sent successfully")
}

func main() {
	targetURL := "http://BigBoss-ALB-213346331.ap-south-1.elb.amazonaws.com/v1/bb/user/dinesh/votes" // Replace with your target URL
	numRequests := 10000                                                                             // Number of requests to send
	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go sendRequest(targetURL, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All requests completed")
}
