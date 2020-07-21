package reflects

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType1(t *testing.T){
	var x float64 = 3.14
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}


func TestType2(t *testing.T){
	var x float64 = 3.14
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	fmt.Println("settability of v:",v.CanSet())
	v.SetFloat(5.55)
	fmt.Println("value of x:",x)

}
