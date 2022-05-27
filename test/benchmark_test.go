package test

import (
	"fmt"
	"testing"
)

/*
第 1 行的-bench=.表示运行 benchmark_test.go 文件里的所有基准测试，和单元测试中的-run类似
基准测试框架对一个测试用例的默认测试时间是 1 秒。开始测试时，当以 Benchmark 开头的基准测试用例函数返回时还不到 1 秒，那么 testing.B 中的 N 值将按 1、2、5、10、20、50……递增，同时以递增后的值重新调用基准测试用例函数。
go test -v -bench=. benchmark_test.go
//通过-benchtime参数可以自定义测试时间
go test -v -bench=. -benchtime=5s benchmark_test.go

在命令行中添加-benchmem参数以显示内存分配情况
第 1 行的代码中-bench后添加了 Alloc，指定只测试 Benchmark_Alloc() 函数。
第 4 行代码的“16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。
go test -v -bench=Alloc -benchmem benchmark_test.go
*/
func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}

//有些测试需要一定的启动和初始化时间，如果从 Benchmark() 函数开始计时会很大程度上影响测试结果的精准性。
//testing.B 提供了一系列的方法可以方便地控制计时器，从而让计时器只在需要的区间进行测试。我们通过下面的代码来了解计时器的控制。
//基准测试中的计时器控制
func Benchmark_Add_TimerControl(b *testing.B) {
	// 重置计时器
	b.ResetTimer()
	// 停止计时器
	b.StopTimer()
	// 开始计时器
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}