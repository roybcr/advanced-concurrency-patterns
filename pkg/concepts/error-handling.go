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
					
					// Separating the concerns of error handling from our producer goroutine.
					// This is desirable because the goroutine that spawned the producer goroutine,
					// in this case our main goroutine, has more context about the running program, 
					// and can make more intelligent decisions about what to do with errors.
					case results <- result:
				}
			}

		}()

		return results
	}

	done := make(chan interface{})
	defer close(done)


	errCount := 0
	urls := []string{"https://www.google.com", "https://badhost", "a", "b", "https://github.com"}
	for result := range checkStatus(done, urls...) {

		if result.Error != nil {

			fmt.Printf("Error: %v\n", result.Error)
			errCount ++

			if errCount >= 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}

			continue
		}

		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
