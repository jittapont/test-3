package main

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func Main() {
	Wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("in goroutine")

	}(&Wg)
	fmt.Println("vim-go")
}

func main() {
	Main()
}
