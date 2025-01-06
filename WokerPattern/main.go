package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Add a timeout for the overall process to be 500 milliseconds
// sample output
// Error from API: API 2 canceled
// Final Result:  API1: Data from API 1
// API3: Data from API 3
// Took: 451.432ms

func getData(ctx context.Context) {
	ApiCount := 10
	var wg sync.WaitGroup
	wg.Add(ApiCount)
	resultChan := make(chan string, ApiCount)
	errorChan := make(chan error, ApiCount)
	for i := 0; i < ApiCount; i++ {
		go func(id int) {
			defer wg.Done()
			val, err := callThirdPartyAPI(i, ctx)
			if err != nil {
				errorChan <- err
			}
			resultChan <- val
		}(i)
	}
	wg.Wait()
	close(resultChan)
	close(errorChan)

	for result := range resultChan {
		fmt.Println(result)
	}
	for err := range errorChan {
		fmt.Println(err)
	}
}

// callAPI simulates an API call with a random delay and respects the timeout
// Exits early if the timeout limit exceeded
// return the data to the getData function
func callThirdPartyAPI(apiID int, ctx context.Context) (string, error) {
	delay := time.Duration(rand.Intn(800)+100) * time.Millisecond // Random delay between 100ms and 900ms
	fmt.Println("TIME DELAY: ", delay)
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("API %d canceled", apiID)
	case <-time.After(delay):
		return "Data returned", nil
	}
}

func main() {
	start := time.Now()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	getData(ctx)
	fmt.Println("Took: ", time.Since(start))

}
