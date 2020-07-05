package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	receive(ch1, "pass message")
	send(ch1, ch2)
	fmt.Println(<-ch2)

	ch3 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch3 <- i
	}
	close(ch3)
	res := 0
	for v := range ch3 {
		res += v
	}
	fmt.Println(res)

	chStr := make(chan string)
	chInt := make(chan int)

	go strWorker(chStr)
	go intWorker(chInt)

	for i := 0; i < 2; i++ {
		select {
		case <-chStr:
			fmt.Println("get value form strWorker")
		case <-chInt:
			fmt.Println("get value from intWorker")
		}
	}

	done := make(chan bool)

	go worker(done)

	<-done
}

func worker(done chan bool) {
	fmt.Println("start working...")
	done <- true
	fmt.Println("end working...")
}

func receive(receiver chan<- string, msg string) {
	receiver <- msg
}

func send(sender <-chan string, receiver chan<- string) {
	msg := <-sender
	receiver <- msg
}

func strWorker(ch chan string) {
	time.Sleep(time.Second)
	fmt.Println("do something with strworker")
	ch <- "str"
}

func intWorker(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("do something with intworker")
	ch <- 1
}
