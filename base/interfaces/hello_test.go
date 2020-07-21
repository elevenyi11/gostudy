package interfaces

import "testing"

func TestSayHello(t *testing.T){
	c := Chinese{"小王"}
	e := English{"Ruby"}
	m := map[int] Human{}
	m[0] = &c
	m[1] = &e
	for i := 0; i< len(m); i++ {
		m[i].sayHello()
	}
}
