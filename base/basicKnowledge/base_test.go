package basicKnowledge

import (
	"fmt"
	"testing"
	"time"
)

const PtrSize = 4 << (^uintptr(0) >> 63)

func TestBase(t *testing.T) {
	fmt.Println(1 << 16)
	fmt.Println(PtrSize)
	fmt.Println(fmt.Sprintf("TestBase:%v ", time.Now().Format("2006-01-02 15:04:05")))

}

func TestChan(t *testing.T) {
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}

	go func() {
		for ch := range ch1 {
			fmt.Println("first func:", ch)
		}

	}()

	go func() {
		for ch := range ch1 {
			fmt.Println("second func:", ch)
		}

	}()

	time.Sleep(10 * time.Second)
}
