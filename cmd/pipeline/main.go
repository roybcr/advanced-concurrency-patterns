package main

import (
	"fmt"
	"math"
)

// The Pipeline pattern is a series of stages connected by channels,
// where each stage is a group of goroutines running the same function.
// By using Pipeline we seperate the concerns of each stage.

func main() {

	done := make(chan interface{})
	defer close(done)

	in  := generateWork([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	out := half(square(filter(in, done), done), done)

	// 2,8,18,32,50,72,98,128,162,200
	for val := range out {
		fmt.Println(val)
	}

}

func filter(in <-chan int, done <-chan interface{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			select {
				case <-done: { return }
				default:
					if i % 2 == 0 { 
						out<- i 
					}
			}
		}
	}()

	return out
}

func square(in <-chan int, done <-chan interface{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			select {
				case <-done: { return }
				case out<- int(math.Pow(float64(i), 2)):
			}
			
		}
	}()

	return out
}

func half(in <-chan int, done <-chan interface{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			select {
				case <-done: { return }
				case out<- (i/2):
			}
		}
	}()
	return out
}

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}
