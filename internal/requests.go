package internal

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type TestResult struct {
	Success        int64
	Failure        int64
	TotalTime      time.Duration
	FirstByteTime  time.Duration
	LastByteTime   time.Duration
	RequestsPerSec float64
	TotalTimeStats Stats
	FirstByteStats Stats
	LastByteStats  Stats
}

type Stats struct {
	Min   time.Duration
	Max   time.Duration
	Mean  time.Duration
	Count int
}

func PerformRequestTests(URL *string, numReq *int, concReq *int) TestResult {
	log.Println("***** Performing test *****")
	log.Println("URL:", *URL)
	log.Println("Number of Requests:", *numReq)
	log.Println("Number of Concurrent Requests:", *concReq)

	var wg sync.WaitGroup

	done := make(chan struct{})

	success := int64(0)
	failure := int64(0)
	totalTime := time.Duration(0)
	firstByteTime := time.Duration(0)
	lastByteTime := time.Duration(0)

	client := http.Client{}

	startTime := time.Now()

	for i := 0; i < *numReq; i++ {
		wg.Add(1)
		go performTest(*URL, *numReq / *concReq, &wg, done, &client, &success, &failure, &totalTime, &firstByteTime, &lastByteTime)
	}

	wg.Wait()

	close(done)

	totalElapsedTime := time.Since(startTime).Seconds()

	requestsPerSec := float64(*numReq) / totalElapsedTime

	return TestResult{
		Success:        success,
		Failure:        failure,
		TotalTime:      totalTime,
		FirstByteTime:  firstByteTime,
		LastByteTime:   lastByteTime,
		RequestsPerSec: requestsPerSec,
		TotalTimeStats: calculateStats(totalTime, *numReq),
		FirstByteStats: calculateStats(firstByteTime, *numReq),
		LastByteStats:  calculateStats(lastByteTime, *numReq),
	}
}

func calculateStats(duration time.Duration, count int) Stats {
	return Stats{
		Min:   time.Duration(0), // Placeholder for Min value
		Max:   duration,
		Mean:  duration / time.Duration(count),
		Count: count,
	}
}

func performTest(URL string, numReq int, wg *sync.WaitGroup, done chan struct{}, client *http.Client, success *int64, failure *int64, totalTime, firstByteTime, lastByteTime *time.Duration) {
	defer wg.Done()

	for i := 0; i < numReq; i++ {
		select {
		case <-done:
			return
		default:
			startTime := time.Now()
			res, errReq := client.Get(URL)
			if errReq != nil {
				log.Println(errReq)
				*failure++
				return
			}
			defer res.Body.Close()

			*totalTime += time.Since(startTime)

			// Read the first byte
			firstByteStart := time.Now()
			io.CopyN(io.Discard, res.Body, 1)
			*firstByteTime += time.Since(firstByteStart)

			// Read the rest of the response to get the last byte time
			_, err := io.Copy(io.Discard, res.Body)
			if err != nil {
				log.Println(err)
			}
			*lastByteTime += time.Since(startTime)

			if res.StatusCode == http.StatusOK {
				*success++
			} else {
				*failure++
			}
		}
	}
}
