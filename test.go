package main

import (
	"fmt"
	"net/http"
	"time"
)

func worker() {
	fmt.Println("worker started")
	time.Sleep(time.Second * 10)
	fmt.Println("worker completed")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	go worker()
	w.Write([]byte("Hello, World!"))
}

func cal() {
	fmt.Println("inside cal function")
}
