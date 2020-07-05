package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	data, err := ioutil.ReadFile("D:\\temp\\test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
