package generators

// A Generator function that takes only the first n items
// off of its incoming valueStream and them exits.

func Take(done <-chan interface{}, valueStream <-chan interface{}, n int) <-chan interface{} {
	
	takeStream := make(chan interface{})

	go func() {
		defer close(takeStream)
		for i := 0; i < n; i++ {
			select {
				case <-done: { return }
				case takeStream<- <-valueStream:
			}
		}
	}()

	return takeStream
}
