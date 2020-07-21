package basicKnowledge

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var i int = 10;
	var pInt *int = &i;
	fmt.Printf("整数i=[%d]，指针pInt=[%p]，指针指向*pInt=[%d]\n",
		i, pInt, *pInt)
	*pInt = 3
	fmt.Printf("整数i=[%d]，指针pInt=[%p]，指针指向*pInt=[%d]\n",
		i, pInt, *pInt)
	i = 5
	fmt.Printf("整数i=[%d]，指针pInt=[%p]，指针指向*pInt=[%d]\n",
		i, pInt, *pInt)
}

func TestMemAllocate(t *testing.T) {
	var pNil *[]int
	fmt.Println("Wild的数组指针：", pNil)
	fmt.Printf("Wild的数组指针==nil[%t]\n", pNil == nil)
	var p *[]int = new([]int)
	fmt.Println("New分配的数组指针：", p)
	fmt.Printf("New分配的数组指针[%p]，长度[%d]\n", p, len(*p))
	fmt.Printf("New分配的数组指针==nil[%t]\n", p == nil)
	//Error occurred
	//(*p)[3] = 23
	*p = make([]int, 10)
	fmt.Println("New分配的数组指针Make后：", p)
	(*p)[3] = 23
	fmt.Println("New分配的数组元素[3]：", (*p)[3])
	var v []int = make([]int, 10)
	fmt.Println("Make分配的数组引用：", v)
}