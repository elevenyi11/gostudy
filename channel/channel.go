package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()

	str := "1,"
	arr := strings.Trim(str, ",")
	fmt.Println(len(arr))

	t := time.NewTicker(1 * time.Minute)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println(time.Now())
		}
	}
}
func GetSeasonId() string {
	t1 := time.Now()
	var month string
	if int(t1.Month()) < 10 {
		month = "0" + strconv.Itoa(int(t1.Month()))
	} else {
		month = strconv.Itoa(int(t1.Month()))
	}

	num := strconv.Itoa(t1.Year()) + month
	return num
}
