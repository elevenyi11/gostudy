package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func doSomething(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after doo job: (%d) \n", id)

}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go doSomething(1, &wg)
	go doSomething(2, &wg)
	go doSomething(3, &wg)
	wg.Wait()
	log.Printf("finish all jobs \n")

	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
