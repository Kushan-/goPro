package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	/*
		//block
		<-ch
		fmt.Printf("Here")
	*/
	go func() {
		// Send number of the channel
		ch <- 353

	}()

	//Receive from the channel
	val := <-ch
	fmt.Printf("got %d\n", val)

	//send multiple
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Print("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 5; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Print("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}
