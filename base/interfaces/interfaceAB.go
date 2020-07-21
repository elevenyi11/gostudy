package interfaces

import "reflect"

type A struct {
	Parent interface{}
}

func (this A) Run(){
	c := reflect.ValueOf(this.Parent)
	method := c.MethodByName("Test")
	println(method.IsValid())

}

type B struct{
	A
}

func (this B) Test(s string){
	println("b")
}

func (this B) Run(){
	this.A.Run()
}