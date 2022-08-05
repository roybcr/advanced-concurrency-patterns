package main

import (
	"fmt"
	"time"
)

func generator() (<-chan int, <-chan string) {

	ch1 := make(chan int   )
	ch2 := make(chan string)

	// the generator generates the next value from within an 
	// infinite loop whenever a ready receiver asks for it

	go func() { for i := 0; ; i++ { ch1 <- i } }()
	go func() {

		for {

			ch2 <- fmt.Sprintf(

				"%s %d %d, %s, %d:%d:%d",
				
				time.Now().Month   (),
				time.Now().Day     (),
				time.Now().Year    (),
				time.Now().Weekday (),
				time.Now().Hour    (),
				time.Now().Minute  (),
				time.Now().Second  (),
			)
			
		}

	}()

	return ch1, ch2;
}

func main() {

	ch1, ch2 := generator()

	for i := 0; i < 10; i++ {
		
		time.Sleep(time.Second * 2)

		value 	  := <-ch1
		timestamp := <-ch2

		fmt.Printf("Generator value: %d | %s\n",
		value, timestamp)
		
	}
}
