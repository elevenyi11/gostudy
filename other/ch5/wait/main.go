package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	err := WaitforServer("https://golang.org")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	if err := WaitforServer("https://golang.org"); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
	//if err := Ping(); err != nil {
	//log.Printf("Ping failed: %v; networking disabled", err)
	//}
}
func WaitforServer(url string) error {
	const timeout = 20 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("Server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("Server %s failed to respond after %s", url, timeout)
}
