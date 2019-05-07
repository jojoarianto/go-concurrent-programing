package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Chapter 1 - Basic")

	go function1("goroutine2", 5) // run with go routine
	go function1("goroutine3", 5) // run with go routine

	function1("main1", 15)

}

func function1(identifier string, n int) {

	for index := 0; index < n; index++ {
		time.Sleep(time.Second * 1)
		fmt.Println("Goroutine ", identifier, " caunt on", index)
	}

	return
}
