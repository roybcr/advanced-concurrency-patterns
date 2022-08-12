package primitives

import (
	"fmt"
	"sync"
	"time"
)

func SelectDemo() {

	var wg sync.WaitGroup
	defer wg.Done()

	// c1 := make(chan int)
	// c2 := make(chan int)

	var c1 chan int
	var c2 chan int
	var c3 chan int

	go func(c1, c2, c3 chan<- int) {

		defer close(c1)
		defer close(c2)
		defer close(c3)

		for i := 100; i > 0; i-- {
			wg.Add(1)
			if i%2 == 0 {
				c1 <- i
			} else {
				c2 <- i
			}
		}
	}(c1, c2, c3)

	var sum1, sum2, sum3, c1Count, c2Count, c3Count int
	for i := 10; i >= 0; i-- {
		select {
		case x := <-c1:
			fmt.Println("New x1: ", x)
			sum1 = sum1 + x
			c1Count++
		case x := <-c2:
			fmt.Println("New x2: ", x)
			sum2 += x
			c2Count++
		case x := <-c3:
			fmt.Println("New x3: ", x)
			sum3 += x
			c3Count++
		case <-time.After(time.Second * 2):
			fmt.Println("Timed out...")
		}
	}

	fmt.Printf("Sum1: %d\nSum2: %d\nc1Count: %d\nc2Count: %d\n", sum1, sum2, c1Count, c2Count)

}
