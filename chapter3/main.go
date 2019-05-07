package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var msg = make(chan string, 3)

func main() {

	fmt.Println("Chapter 3 - WaitGroup & Channel")

	// print the number of goroutines that currently exist.
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)                     // indicate we are going to wait for one thing
	go function1("goroutine2", 5) // run with go routine
	go function1("goroutine3", 8) // run with go routine
	wg.Add(1)                     // indicate we are going to wait for one thing
	go function1("goroutine4", 3) // run with go routine

	// print the number of goroutines that currently exist.
	fmt.Println(runtime.NumGoroutine())

	wg.Wait() // wait for all things to be done

	// close the channel so that we not longer expect writes to it
	close(msg)

	// read remaining values in the channel
	for item := range msg {
		// read and print chan
		fmt.Println(item)
	}

	// print the number of goroutines that currently exist.
	fmt.Println(runtime.NumGoroutine())

}

func function1(identifier string, n int) {

	defer wg.Done()

	for index := 0; index < n; index++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Goroutine ", identifier, " caunt on", index)

		if index == 7 {
			return // this unpredictable error example
		}
	}

	temp := "Goroutine " + identifier + " done"

	// fill channel
	msg <- temp

	return
}
