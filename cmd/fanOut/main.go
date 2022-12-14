package main

import "fmt"

type T interface {
	int32 | string
}

func fanOut[k T](inputChannel <-chan k) <-chan k {
	outputChannel := make(chan k)
	go func() {
		defer close(outputChannel)
		for data := range inputChannel { outputChannel <- data }
	}()
	return outputChannel
}

func main() {
	input 	:= []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	rch   	:= makeRChan(input)
	
	outputA := fanOut(rch)
	outputB := fanOut(rch)
	outputC := fanOut(rch)
	outputD := fanOut(rch)
	
	for i, v := range []<-chan int32 {outputA, outputB, outputC, outputD} {
		go func(ch <-chan int32, idx int) {
			val := <-ch 
			fmt.Printf("%d) value: %d", idx, val)
		}(v, i)
	}
	fmt.Printf("%T, %T, %T, %T\n", outputA, outputB, outputC, outputD)

	for range rch {
		select {
			case val := <-outputA:
			    fmt.Println("Output A got:", val)
			case val := <-outputB:
			    fmt.Println("Output B got:", val)
			case val := <-outputC:
			    fmt.Println("Output C got:", val)
			case val := <-outputD:
			    fmt.Println("Output D got:", val)
		}
	}
}

func makeRChan[k T](work []k) <-chan k {
	inputChannel := make(chan k)
	go func() {
		defer close(inputChannel)
		for _, v := range work { inputChannel <- v }	
	}()
	return inputChannel
} 
