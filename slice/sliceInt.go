package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type Student struct {
	id   string
	name string
}
type Comparable interface {
	IsEqual(obj interface{}) bool
}

func (this Student) IsEqual(obj interface{}) bool {
	if student, ok := obj.(Student); ok {
		return this.GetId() == student.GetId()
	}
	panic("unexpected type")
}

func (this Student) GetId() string {
	return this.id
}

func isEqual(a, b interface{}) bool {
	if comparable, ok := a.(Comparable); ok {
		return comparable.IsEqual(b)
	} else {
		return a == b
	}
}

type Slice []interface{}

func NewSlice() Slice {
	return make(Slice, 0)
}

func (this *Slice) Add(elem interface{}) error {
	for _, v := range *this {
		if isEqual(v, elem) {
			fmt.Printf("Slice:Add elem: %v already exist\n", elem)
			return errors.New("Elem exist")
		}
	}

	*this = append(*this, elem)
	fmt.Printf("Slice:Add elem: %v succ\n", elem)
	return nil
}
func (this *Slice) Remove(elem interface{}) error {
	found := false
	for i, v := range *this {
		if isEqual(v, elem) {
			if i == len(*this)-1 {
				*this = (*this)[:i]
			} else {
				*this = append((*this)[:i], (*this)[i+1:]...)
			}
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Slice:Remove elem: %v not exist\n", elem)
		return errors.New("Elem Not Exist")
	}
	fmt.Printf("Slice:Remove Elem: %v succ\n", elem)
	return nil
}

func IntSliceExec() {
	fmt.Println("int slice start")
	slice1 := NewSlice()
	slice1.Add(1)
	fmt.Println("current int slice:", slice1)
	slice1.Add(2)
	fmt.Println("current int slice:", slice1)
	slice1.Add(2)
	fmt.Println("current int slice:", slice1)
	slice1.Add(3)
	fmt.Println("current int slice:", slice1)
	slice1.Remove(2)
	fmt.Println("current int slice:", slice1)
	slice1.Remove(2)
	fmt.Println("current int slice:", slice1)
	slice1.Remove(3)
	fmt.Println("current int slice:", slice1)
	fmt.Println("int slice end")
}
func StringSliceExec() {
	fmt.Println("string slice start")
	slice := NewSlice()
	slice.Add("hello")
	fmt.Println("current string slice:", slice)
	slice.Add("golang")
	fmt.Println("current string slice:", slice)
	slice.Add("golang")
	fmt.Println("current string slice:", slice)
	slice.Add("generic")
	fmt.Println("current string slice:", slice)
	slice.Remove("golang")
	fmt.Println("current string slice:", slice)
	slice.Remove("golang")
	fmt.Println("current string slice:", slice)
	slice.Remove("generic")
	fmt.Println("current string slice:", slice)
	fmt.Println("string slice end")
}
func structSliceExec() {
	fmt.Println("struct slice start")
	xiaoMing := Student{"1001", "xiao ming"}
	xiaoLei := Student{"1002", "xiao lei"}
	xiaoFang := Student{"1003", "xiao fang"}
	slice := NewSlice()
	slice.Add(xiaoMing)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Add(xiaoFang)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoLei)
	fmt.Println("current struct slice:", slice)
	slice.Remove(xiaoFang)
	fmt.Println("current struct slice:", slice)
	fmt.Println("struct slice end")
}

func main() {
	IntSliceExec()

	StringSliceExec()

	structSliceExec()

}
