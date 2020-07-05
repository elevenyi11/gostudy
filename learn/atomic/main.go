package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	var (
		total int64
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				atomic.AddInt64(&total, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("The total number is ", total)
}
