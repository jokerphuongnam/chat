package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	url         = "http://localhost:3031/v1/rooms/1866c185-771b-4b14-8f24-82abaf05ad5a/messages" // Update with your server URL
	numRequests = 10000                                                                          // Total number of requests to send
	concurrency = 1000                                                                           // Number of concurrent requests at a time
)

func sendRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA4MDE1NjAsInN1YiI6IjhhNmUxMjk0LWVkZWMtNDVhNi04YTYzLTE5YTM1YzBlOGVkZiJ9.bZa8cFEeauY-N54aNpkIiahLz9Qe5ueBBMx5wf9a7w8")

	client := &http.Client{}
	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)

	// Check for errors
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Log the response status and duration
	fmt.Printf("Response status: %s, Duration: %v\n", resp.Status, duration)
}

// func RunLoadTest() {
// 	var wg sync.WaitGroup
// 	startTime := time.Now()

// 	for i := 0; i < numRequests; i++ {
// 		wg.Add(1)
// 		go sendRequest(&wg)

// 		if (i+1)%concurrency == 0 {
// 			wg.Wait()
// 		}
// 	}

// 	wg.Wait()
// 	fmt.Printf("Load test completed in %v seconds\n", time.Since(startTime).Seconds())
// }

// func main() {
//     RunLoadTest()
// }
