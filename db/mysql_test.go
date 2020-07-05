package main

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	result, err := query()
	if err != nil {
		fmt.Println(err)
	}
	if result != nil {
		for _, v := range result {
			fmt.Println(v)
		}
	}
}

func TestOrmQuery(t *testing.T) {
	result := selectall()
	if result != nil {
		for _, v := range result {
			fmt.Println(v)
		}
	}
}
