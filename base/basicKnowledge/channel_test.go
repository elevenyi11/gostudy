package basicKnowledge

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelWithTimeout(t *testing.T){
	ch := make(chan string,5)
	go func() {
		ch <-"hello"
		time.Sleep(time.Second * 10)
		ch <- "world"
	}()
	for i := 0; i < 2; i++ {
		select{
			case msg := <- ch:
				fmt.Println(msg)
			case <- time.After(time.Second*5):
			fmt.Println("time out")
			default:
				fmt.Println("nothing received!")
				time.Sleep(time.Second)

		}
	}
}
