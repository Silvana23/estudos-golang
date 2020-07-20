package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int = 0
var wg sync.WaitGroup

func executar(s int) {
	for i := 0; i < s; i++ {

		go func(x int) {
			temp := count
			runtime.Gosched()
			temp++
			count = temp
		}(i)
	}
}

func main() {

	n := 1000
	wg.Add(n)
	executar(n)
	wg.Wait()
	fmt.Println("Numero de goroutines: ", n)
	fmt.Println("Valor do contador: ", count)
}