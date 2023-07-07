package main

import "fmt"

func sender(ch chan<- int, done chan<- string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Sended ", i)
		ch <- i
	}
	close(ch)
	done <- "Sender done"
}

func receiver(ch <-chan int, done chan<- string) {
	for v := range ch {
		fmt.Println("Received ", v)
	}
	done <- "Receiver done"
	close(done)
}

func main() {
	ch := make(chan int, 5)
	done := make(chan string)

	go sender(ch, done)
	go receiver(ch, done)

	for v := range done {
		fmt.Println(v)
	}
}
