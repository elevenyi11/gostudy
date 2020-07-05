package main

import (
	"GoStudy/functional/fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			//panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC, 0666)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s,%s,%s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	goPath := os.Getenv("GOPATH")

	currentPath := goPath + "\\src\\GoStudy\\errhandling\\defer\\fib.txt"
	writeFile(currentPath)
}
