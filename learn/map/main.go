package main

import (
	"fmt"
	"sync"
)

func main() {
	cMap := make(map[string]int)
	fmt.Println(cMap == nil)
	cMap["bei jing"] = 1
	fmt.Println(cMap)
	cMap = make(map[string]int, 100)
	cMap["cd"] = 2

	fmt.Println(cMap)

	cMap1 := map[string]int{"sc": 3}

	fmt.Println(cMap1)
	code, ok := cMap1["my"]
	if ok {
		fmt.Println("key is exist.")
	} else {
		fmt.Println("key is not exist.")
	}
	fmt.Println(code)

	cMap = map[string]int{"beijing": 1, "shanghai": 2, "guangzhou": 3}
	for city, code := range cMap {
		fmt.Printf("%s:%d", city, code)
		fmt.Println()
	}

	cMap2 := make(map[string]int)
	var wg sync.WaitGroup
	var mux sync.Mutex
	wg.Add(2)

	go func() {
		mux.Lock()
		cMap2["beijing"] = 1
		mux.Unlock()
		wg.Done()
	}()

	go func() {
		mux.Lock()
		cMap2["shanghai"] = 2
		mux.Unlock()
		wg.Done()
	}()
	wg.Wait()
}
