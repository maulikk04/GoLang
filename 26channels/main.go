//channels are way where multiple go routines can talk to each other

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2) // buffer channel
	wg := &sync.WaitGroup{}
	// fmt.Println(<-myCh)
	// myCh <- 5 // 5 channels

	wg.Add(2)

	//recieving value , recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, ischanelopen := <-myCh

		fmt.Println(ischanelopen)
		fmt.Println(val)

		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	//putting value ,send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
