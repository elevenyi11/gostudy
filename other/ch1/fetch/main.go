package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := make(map[string]string)
	//urls["www.baidu.com"] = "www.baidu.com"
	urls["http://www.163.com"] = "http://www.163.com"
	//urls["http://www.google.com"] = "http://www.google.com"
	for _, url := range urls {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
