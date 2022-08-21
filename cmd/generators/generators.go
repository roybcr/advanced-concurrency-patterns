package generators

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"time"
)

func Generators() {

	// ============================= Generator =============================
	// A simple generator implementation

	defer func() {
		ch1, ch2 := Generator()
		
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			value := <-ch1
			timestamp := <-ch2
			fmt.Printf("Generator value: %d | %s\n", value, timestamp)
		}
	}()

	// =============================== Repeat ===============================
	// Combining the Repeat and Take generators

	done := make(chan interface{})
	defer close(done)

	// Although the Repeat generator is infinite, the Take generator
	// will only take the first 5 values from the Repeat generator
	// and then both generators will exit.

	repeat := Repeat(done, 1, 2)
	for num := range Take(done, repeat, 5) {
		fmt.Printf("%v\t ", num) //=> 1 2 1 2 1
	}

	fmt.Println()

	// ============================ RepeatFn ============================
	// Here we use the RepeatFn generator combined with the take generator
	// to generate a stream of random numbers by passing a function to the 
	// RepeatFn generator then passing it to the Take generator.

	doneRepeatFn := make(chan interface{})
	
	defer close(doneRepeatFn)
	
	type ANY = interface{}
	
	rand := func() ANY {
		var reader io.Reader = rand.Reader 
		n, err := rand.Int(reader, big.NewInt(10000))
		if err != nil { panic(err) }
		return n
	}

	for num := range Take(done, RepeatFn(done, rand), 10) {
		fmt.Println(num)
	}

}
