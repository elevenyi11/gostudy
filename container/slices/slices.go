package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

func main() {

	arr := [...]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("arr = ", arr)
	s1 = arr[2:6]
	s2 = arr[3:5]
	fmt.Printf("s1 =%v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 =%v, len(s2) =%d, cap(s2) = %d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	s6 := append(s4, 13)
	fmt.Println("s3,s4,s5 =", s3, s4, s5, s6)

	fmt.Println("arr = ", arr)
	ss := NewSlice()
	defer ss.Add(1).Add(2)
	ss.Add(3)
}
