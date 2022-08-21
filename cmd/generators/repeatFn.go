package generators

// ReapeatFn is a generator that takes a function
// and repeats it until the done channel is closed.

func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	
	repeatStream := make(chan interface{})
	
	go func() {
		defer close(repeatStream)
		for {
			select {
			case <-done:{ return }
			case repeatStream <- fn():
			}
		}
	}()
	
	return repeatStream
}
