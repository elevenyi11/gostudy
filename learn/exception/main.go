package main

import (
	"fmt"
	"os"
)

func main() {
	outerFunc()

	content, err := closeFile("main.go")
	if err == nil {
		fmt.Println(content)
	}

	printNumber()
}

func outerFunc() {
	defer fmt.Print(" World \n")
	fmt.Print("Hello")
}

func closeFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	defer f.Close()
	buf := make([]byte, 16)

	f.Read(buf)
	return string(buf), err
}

func printNumber() {
	for i := 0; i < 5; i++ {
		defer func(v int) {
			fmt.Println(v)
		}(i * 2)
	}
}
