package main

import "fmt"

func main(){
	cMap := make(map[string] int)

	cMap["北京"] = 1
	fmt.Println(cMap)

	code := cMap["北京"]
	fmt.Println(code)

	code, ok := cMap["shanghai"]
	if ok{
		fmt.Println(code)
	}else{
		fmt.Println("shagnhai not exist")
	}

	delete (cMap, "北京")

	fmt.Println(cMap);

	 
}