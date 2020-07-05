package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var (
		mux    sync.Mutex
		state1 = map[string]int{
			"a": 65,
		}
		muxTotal uint64
		rw       sync.RWMutex
		state2   = map[string]int{
			"a": 65,
		}

		rwTotal uint64
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				_ = state1["a"]
				mux.Unlock()
				atomic.AddUint64(&muxTotal, 1)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				rw.RLock()
				_ = state2["a"]
				rw.RUnlock()
				atomic.AddUint64(&rwTotal, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("sync.Mutex redOps is ", muxTotal)

	fmt.Println("sync.RWMutex readOps is", rwTotal)
}
