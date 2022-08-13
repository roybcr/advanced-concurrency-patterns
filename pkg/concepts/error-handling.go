package concepts

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func ErrorHandling() {
	
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		
		results := make(chan Result)
		go func() {
			defer close(results)
		
			for _, url := range urls {
				
				var result Result
				
				resp, err := http.Get(url)
				result = Result{ Error: err, Response: resp }
				
				select {
					case <-done: { return }
					// Pass the result whether it's an error or response,
					// so that it can be handled out of the scope of this goroutine
					case results <- result:
				}
			}

		}()

		return results
	}

	done := make(chan interface{})
	urls := []string{"https://www.google.com", "https://badhost"}
	
	defer close(done)
	
	for result := range checkStatus(done, urls...) {

		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
			continue
		}

		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
