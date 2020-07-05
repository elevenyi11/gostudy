package main

import (
	"fmt"
	"log"

	"gopl.io/ch4/github"
)

func main() {
	key := []string{"a"}

	result, err := github.SearchIssues(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d	%9.9s	%.55s\n", item.Number, item.User.Login, item.Title)
	}
}
