package main

import "fmt"

func sender(ch chan int, done chan string) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	done <- "Sender done"
}

func receiver(ch chan int, done chan string) {
	for v := range ch {
		fmt.Println("Receiver: Received ", v)
	}
	done <- "Receiver done"
}

func main() {
	ch := make(chan int)
	done := make(chan string)

	go sender(ch, done)
	go receiver(ch, done)

	fmt.Printf("%#v\n", <-done)
	fmt.Printf("%#v\n", <-done)
}