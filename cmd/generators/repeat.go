package generators

// A Repeat Generator that takes a value 
// and repeats it until the done channel is closed.

func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	
	repeatStream := make(chan interface{})
	
	go func() {
		defer close(repeatStream)
		for {
			for _, v := range values {
				select {
					case <-done: { return }
					case repeatStream<- v:
				}
			}
		}
	}()

	return repeatStream
}

