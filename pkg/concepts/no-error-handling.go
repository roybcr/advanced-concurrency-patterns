package concepts

import (
	"fmt"
	"net/http"
)

func NoErrorHandling() {
	checkStatus := func( done <-chan interface{}, urls ...string, ) <-chan *http.Response {

		responses := make(chan *http.Response)

		go func() {
			defer close(responses)
			
			for _, url := range urls {
				resp, err := http.Get(url)
				// The goroutine is in a position where it can only
				// prints out the error and hopes something is paying attention
				if err != nil {
					fmt.Println(err)
					continue
				}
				
				select {
					case <-done: { return }
					case responses <- resp:
				}
			}

		}()

		return responses
	}

	done := make(chan interface{})
	urls := []string{"https://www.google.com", "https://badhost"}
	defer close(done)

	for response := range checkStatus(done, urls...) { fmt.Printf("Response: %v\n", response.Status) }
}
