package basicKnowledge

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T){
	// CAS toolkit
	var cnt int32 = 0
	for i := 0; i <	10 ; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				atomic.AddInt32(&cnt,1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("cnt", atomic.LoadInt32(&cnt))
}
