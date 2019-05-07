package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Chapter 2 - Wait Group")

	wg.Add(2)                      // indicate we are going to wait for one thing
	go function1("goroutine2", 5)  // run with go routine
	go function1("goroutine3", 10) // run with go routine
	wg.Wait()                      // wait for all things to be done

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

	fmt.Println("Goroutine ", identifier, " DONE")

	return
}
