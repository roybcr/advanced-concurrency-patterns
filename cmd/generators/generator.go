package generators

import (
	"fmt"
	"time"
)

func Generator() (<-chan int, <-chan string) {

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


