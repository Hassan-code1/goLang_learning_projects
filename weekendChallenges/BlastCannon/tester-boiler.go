package main

import (
	"net/http"
	"time"
)

// Result holds the outcome of a single HTTP request
type Result struct {
	StatusCode int

	// TODO 1:
	// Add a field named:
	// Latency
	// using the type:
	// time.Duration


	// TODO 2:
	// Add a field named:
	// Error
	// using the type:
	// error


}

// Metrics stores the aggregated benchmark statistics
type Metrics struct {
	TotalTime    time.Duration
	AvgLatency   time.Duration
	MinLatency   time.Duration
	MaxLatency   time.Duration
	SuccessCount int
	FailCount    int
}

// RunLoadTest coordinates the worker pool,
// dispatches jobs, and aggregates benchmark results.
func RunLoadTest(
	url string,
	totalRequests int,
	concurrency int,
	onResult func(completed int),
) Metrics {

	// TODO 3: Prevent zero-worker deadlocks
	// If concurrency is less than or equal to 0,
	// forcefully set it to 1.
	// (A pool with 0 workers will deadlock the application)


	testStartTime := time.Now()

	// Configure high-performance HTTP client (Provided for you)
	customClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: concurrency,
		},
		Timeout: 10 * time.Second,
	}

	// TODO 4: Initialize worker channels
	//
	// Create:
	// 1. A buffered 'jobs' channel of type int
	// 2. A buffered 'results' channel of type Result
	//
	// Both channels should use:
	// totalRequests
	// as their buffer capacity.



	// TODO 5: Launch the worker pool
	//
	// Create a loop that runs:
	// concurrency times.
	//
	// Inside the loop:
	// launch the worker() function as a Goroutine.
	//
	// Pass:
	// - customClient
	// - url
	// - jobs
	// - results



	// TODO 6: Dispatch benchmark jobs
	//
	// Create a loop that runs:
	// totalRequests times.
	//
	// Send the request number into the jobs channel.



	// TODO 7: Signal worker shutdown
	//
	// Close the jobs channel so workers know
	// no more requests are coming.



	// === Aggregation Variables (Do not modify) ===

	var totalLatency, minLatency, maxLatency time.Duration
	var successCount, failCount int

	// TODO 8: Aggregate worker results
	//
	// Create a loop that runs exactly:
	// totalRequests times.
	//
	// Inside the loop:
	//
	// 1. Receive a Result from the results channel.
	//
	// 2. Trigger the progress callback:
	//
	// if onResult != nil {
	//     onResult(i)
	// }
	//
	// 3. Detect failed requests:
	//    - res.Error != nil
	//    - OR non-200 status codes
	//
	//    Increment failCount and continue.
	//
	// 4. Successful requests should:
	//    - increment successCount
	//    - accumulate totalLatency
	//
	// 5. Track:
	//    - maximum latency
	//    - minimum latency



	// TODO 9: Calculate average latency
	//
	// Prevent divide-by-zero panics.
	//
	// Only calculate the average if:
	// successCount > 0

	var avgLatency time.Duration



	return Metrics{
		TotalTime:    time.Since(testStartTime),
		AvgLatency:   avgLatency,
		MinLatency:   minLatency,
		MaxLatency:   maxLatency,
		SuccessCount: successCount,
		FailCount:    failCount,
	}
}

// worker continuously pulls jobs from the queue
// and executes HTTP requests.
func worker(
	client *http.Client,
	url string,
	jobs chan int,
	results chan Result,
) {

	// TODO 10: Consume jobs continuously
	//
	// Use a for range
	// loop over the jobs channel.
	//
	// Workers should automatically stop once
	// the jobs channel closes.



		// TODO 11: Capture request start time
		//
		// Record the current timestamp before
		// firing the HTTP request.



		// TODO 12: Dispatch the HTTP request
		//
		// Use client.Get(url)
		//
		// Capture both:
		// - response
		// - error



		// TODO 13: Measure request latency
		//
		// Use time.Since(...)
		// to calculate how long the request took.



		// TODO 14: Handle request failures
		//
		// If a network error occurred:
		//
		// 1. Send a Result into the results channel
		//    containing the Error field.
		//
		// 2. Skip the remaining loop iteration
		//    using continue.



		// TODO 15: Extract status and release resources
		//
		// Requirements:
		// 1. Save resp.StatusCode
		// 2. Immediately close resp.Body
		//
		// Important:
		// NEVER use defer inside a loop,
		// or response bodies will accumulate
		// and leak memory.



		// TODO 16: Publish successful benchmark data
		//
		// Send a populated Result struct into
		// the results channel containing:
		//
		// - StatusCode
		// - Latency


		
}


