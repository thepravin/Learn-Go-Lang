// Write a program that prints numbers from 1 to 10 using two goroutines in sequence
//  (odd-even problem).

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	oddChan := make(chan bool)
	evenChan := make(chan bool)

	wg.Add(2)

	// ODD
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-oddChan // waiting for the signals from main or even goroutine
			fmt.Println("ODD : ", i)
			evenChan <- true
		}
	}()

	// EVEN
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-evenChan // waiting for the signals from odd gorouting
			fmt.Println("EVEN : ", i)
			if i < 10 {
				oddChan <- true
			}

		}
	}()

	oddChan <- true // kickstart the process signaling the odd channel

	wg.Wait()
	fmt.Println("DONE...")

}
