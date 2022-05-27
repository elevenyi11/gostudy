package test

import (
	"fmt"
	"testing"
	"time"
)
/*
单元测试框架提供的日志方法
方  法	备  注
Log	打印日志，同时结束测试
Logf	格式化打印日志，同时结束测试
Error	打印错误日志，同时结束测试
Errorf	格式化打印错误日志，同时结束测试
Fatal	打印致命日志，同时结束测试
Fatalf	格式化打印致命日志，同时结束测试

go test -v select_test.go
go test -v -run TestA$ select_test.go  单个
go test -v -run TestA select_test.go

*/
func TestHelloWorld(t *testing.T) {
	t.Log("hello world")
}

func TestA(t *testing.T) {
	t.Log("A")
}
func TestAK(t *testing.T) {
	t.Log("AK")
}
func TestB(t *testing.T) {
	t.Log("B")
}

/*
for和select一同使用，有坑
break只能跳出select，无法跳出for
*/
/*
func TestBreak(t *testing.T) {
	tick := time.Tick(time.Second)
	for {
		select {
		case t := <-tick:
			fmt.Println(t)
			break
		}
	}
	fmt.Println("end")
}*/

/*
break无法跳出select的解决方案
1、标签
2、goto
*/

func TestBreakLable(t *testing.T) {
	tick := time.Tick(time.Second)
	//FOR是标签
FOR:
	for {
		select {
		case t := <-tick:
			fmt.Println(t)
			//break出FOR标签标识的代码
			break FOR
		}
	}
	fmt.Println("end")
}

func TestBreakGoto(t *testing.T) {
	tick := time.Tick(time.Second)
	for {
		select {
		case t := <-tick:
			fmt.Println(t)
			//跳到指定位置
			goto END
		}
	}
END:
	fmt.Println("end")
}

/*
单独在select中是不能使用continue，会编译错误，只能用在for-select中。
continue的语义就类似for中的语义，select后的代码不会被执行到
*/
/*
func TestBreakContinue(t *testing.T) {
	tick := time.Tick(time.Second)
	for {
		select {
		case t := <-tick:
			fmt.Println(t)
			continue
			fmt.Println("test")
		}
	}
	fmt.Println("end")
}*/

/*
// 终止当前测试用例时，可以使用 FailNow
func TestFailNow(t *testing.T) {
	t.FailNow()
}

// 只标记错误不终止测试的方法
func TestFail(t *testing.T) {
	fmt.Println("before fail")
	t.Fail()
	fmt.Println("after fail")
}*/

func TestC(t *testing.T) {
	t.Log("C")
}

/*
func TestBreak(t *testing.T) {
	tick := time.Tick(time.Second)
	for {
		select {
		case t := <-tick:
			fmt.Println(t)
			continue
			fmt.Println("test")
		}
	}
	fmt.Println("end")
}*/

func TestD(t *testing.T) {
	t.Log("D")
}