package main

import (
	alg "concurrencyPatterns/pkg/algorithms"
	con "concurrencyPatterns/pkg/concepts"
	pri "concurrencyPatterns/pkg/primitives"
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

	for _, fn := range (*r).runner {
		go func(wg *sync.WaitGroup, fn func()) {
			defer (*wg).Done()
			fn()
		}(&wg, fn)
	}

	wg.Wait()
}

var AlgorithmsRunner = Runner{[]func(){alg.BinarySearch}}
var ConceptsRunner = Runner{
	[]func(){
		con.OrChannel,
		con.LexicalConfinment,
		con.NoErrorHandling,
		con.ErrorHandling,
	},
}
var PrimitivesRunner = Runner{
	[]func(){
		pri.Cond,
		pri.RWMutex,
		pri.Broadcast,
		pri.KnightsTour,
	},
}

func RunAll() {
	ConceptsRunner.Run()
	PrimitivesRunner.Run()
	AlgorithmsRunner.Run()
}

func main() { RunAll() }
