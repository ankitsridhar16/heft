package main

import (
	"flag"
	"github.com/ankitsridhar16/heft/internal"
	"log"
)

const (
	defaultNumRequests   = 1
	defaultConcurrentRqs = 1
)

func main() {
	// Define command line flags
	URL := flag.String("u", "", "URL to test")
	numReq := flag.Int("n", defaultNumRequests, "Number of requests to perform at a time")
	concReq := flag.Int("c", defaultConcurrentRqs, "Number of concurrent requests to perform at a time")
	// Read a URL value in go
	flag.Parse()

	if *URL == "" {
		log.Fatal("Error: -u flag is required")
		return
	}

	// Parse URL
	parsedURLErr := internal.ParseURL(*URL)
	if parsedURLErr != nil {
		log.Fatal(parsedURLErr)
	}

	internal.PerformRequestTests(URL, numReq, concReq)

}
