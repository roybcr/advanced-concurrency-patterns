package primitives

import (
	"fmt"
	"time"
)

func ForSelect() {

	done := make(chan interface{})
	go func() {
		defer close(done)
		time.Sleep(5 * time.Second)
		done <- 45
	}()
	workCounter := 0
loop:
	for {
		select {
		case x := <-done:
			fmt.Println("X: ", x)
			break loop
		default:
		}
		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}
