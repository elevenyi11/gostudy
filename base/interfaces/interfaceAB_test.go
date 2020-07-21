package interfaces

import "testing"

func TestAB(t *testing.T){
	b := new(B)
	b.A.Parent = b
	b.Run()

}
