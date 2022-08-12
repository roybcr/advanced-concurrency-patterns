package main

import (
con	"concurrencyPatterns/pkg/concepts"
pri	"concurrencyPatterns/pkg/primitives"
	"sync"
)

type runner []func()

type Runner struct {
	runner
}

// This function executes all Runner's operations concurrently.
func (r *Runner) Run() {

	var wg sync.WaitGroup
	wg.Add(len((*r).runner))

	func(wg *sync.WaitGroup) {
		for _, fn := range (*r).runner {
			go func(wg **sync.WaitGroup, fn func()) {
				defer (**wg).Done()
				fn()
			}(&wg, fn)
		}
	}(&wg)

	wg.Wait()
}

var ConceptsRunner 	 = Runner{[]func(){con.LexicalConfinment}}
var PrimitivesRunner = Runner{
	[]func(){
		pri.Cond,
		pri.RWMutex,
		pri.Broadcast,
		pri.KnightsTour,
	},
}

func runAll() {
	ConceptsRunner.Run()
	PrimitivesRunner.Run()
}

func main() { runAll() }
