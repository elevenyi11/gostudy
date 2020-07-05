package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

type Calculate interface {
	Do(numbers ...int) int
}

type Sum struct {
}

func (s Sum) Do(numbers ...int) int {
	count := 0
	for n := range numbers {
		count += n
	}
	return count
}

type Div struct {
}

func (d Div) Do(numbers ...int) int {
	if len(numbers) <= 1 {
		panic("参数错误")
	}
	reslut := numbers[0]
	for k, v := range numbers {
		if k == 0 {
			continue
		}
		reslut = reslut / v
	}
	return reslut
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int2 int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Error handling")
	if result, err := eval(3, 4, "*"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3,4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a,b after swap is:", a, b)
}
