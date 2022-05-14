package main

import (
	"fmt"
	"sync"
)

func num_server(increment, decrement, read_i <-chan bool, wg *sync.WaitGroup) {

	var i = 0

	for {
		select {
		case <-increment:
			for j := 0; j < 1000000; j++ {
				i++
			}
			wg.Done()
		case <-decrement:
			for j := 0; j < 1000000; j++ {
				i--
			}
			wg.Done()
		case <-read_i:
			fmt.Println("The magic number is: ", i)
			wg.Done()
		}
	}

}

func increment(increment chan<- bool) {
	increment <- true
}

func decrement(decrement chan<- bool) {
	decrement <- true
}

func main() {

	read_i := make(chan bool)
	decr := make(chan bool)
	incr := make(chan bool)

	wg := new(sync.WaitGroup)

	/* do the increment and decrement */
	go num_server(incr, decr, read_i, wg)
	go increment(incr)
	go decrement(decr)
	wg.Add(2)
	wg.Wait()

	/* read i after increment and decrement is done */
	read_i <- true
	wg.Add(1)
	wg.Wait()
}
