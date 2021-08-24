package main

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func main() {
	Wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("in goroutine")

	}(&Wg)
	Wg.Wait()
	fmt.Println("vim-go")
}
