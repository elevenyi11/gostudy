package main

import (
	"fmt"
	"sort"
)

func main() {
	sorts()
}

func sorts() {
	ages := map[string]int{
		"Fob": 10,
		"Bob": 15,
		"Aob": 18,
	}
	//var names []string
	names := make([]string, 0, len(ages)) //因为我们一开始就知道names的最终大小，因此给slice分配一个合适的大小将会更有效
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	if age, ok := ages["Bob"]; !ok {
		fmt.Println("not ok")
	} else {
		fmt.Println(age)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
