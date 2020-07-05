package main

import "sync"
import "runtime"
import "fmt"

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func main() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)

}
