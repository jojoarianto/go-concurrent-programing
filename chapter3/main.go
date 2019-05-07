package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var msg = make(chan string, 3)
var box = make([]string, 0)

func main() {

	fmt.Println("Chapter 3 - WaitGroup & Channel")

	// print the number of goroutines that currently exist.
	fmt.Println(runtime.NumGoroutine())

	go function1("goroutine2", 5) // run with go routine
	go function1("goroutine3", 6) // run with go routine
	go function1("goroutine4", 3) // run with go routine

	// print the number of goroutines that currently exist.
	fmt.Println(runtime.NumGoroutine())

	var message1 = <-msg
	fmt.Println(message1)

	var message2 = <-msg
	fmt.Println(message2)

	var message3 = <-msg
	fmt.Println(message3)

	// print the number of goroutines that currently exist.
	fmt.Println(runtime.NumGoroutine())

}

func function1(identifier string, n int) {

	for index := 0; index < n; index++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Goroutine ", identifier, " caunt on", index)

		if index == 7 {
			return // this unpredictable error example
		}
	}

	temp := "Goroutine " + identifier + " done ========="

	// fill channel
	msg <- temp

	return
}
