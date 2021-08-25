package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var Wg sync.WaitGroup

func Main() {
	Wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("in goroutine")

	}(&Wg)
	fmt.Println("vim-go")
	cal()
}

func Test2() {
	go func() {
		time.Sleep(30 * time.Minute)

	}()
}

func main() {
	Main()
	http.HandleFunc("/home", HomeHandler)
	http.ListenAndServe(":8081", nil)
}
