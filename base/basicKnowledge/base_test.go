package basicKnowledge

import (
	"fmt"
	"testing"
	"time"
)

func TestBase(t *testing.T){
	fmt.Print(1)
	fmt.Println(fmt.Sprintf("TestBase:%v ", time.Now().Format("2006-01-02 15:04:05")))
}
