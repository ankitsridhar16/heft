package internal

import (
	"io"
	"log"
	"net/http"
	"sync"
)

func PerformRequestTests(URL *string, numReq *int, concReq *int) {
	var wg sync.WaitGroup

	done := make(chan struct{})

	for i := 0; i < *numReq; i++ {
		wg.Add(1)
		go performTest(*URL, *numReq / *concReq, &wg, done)
	}

	wg.Wait()

	close(done)
}

func performTest(URL string, numReq int, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()

	client := http.Client{}
	for i := 0; i < numReq; i++ {
		select {
		case <-done:
			return
		default:
			res, errReq := client.Get(URL)
			if errReq != nil {
				log.Fatal(errReq)
				return
			}
			defer func(Body io.ReadCloser) {
				errBody := Body.Close()
				if errBody != nil {
					log.Fatal(errBody)
				}
			}(res.Body)
			log.Printf("Request %d - Status Code: %d\n\n", i+1, res.StatusCode)
		}
	}
}
