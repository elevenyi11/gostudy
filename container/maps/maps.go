package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouce",
		"course":  "golang",
		"site":    "goland",
		"quality": "notbad",
	}
	m2 := make(map[string]int)

	var m3 map[string]int
	fmt.Println("m, m1, m3:")
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map m")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(`m["course"]`, courseName)
	if caseName, ok := m["cause"]; ok {
		fmt.Println(caseName)
	} else {
		fmt.Println("key 'cause' does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Printf("m[%q] before delete: %q, %v\n", "name", name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Printf("m[%q] after delete : %q, %v\n", "name", name, ok)
}
