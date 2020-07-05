package main

import (
	"GoStudy/util/file"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var (
	url          = "http://books.studygolang.com/gobyexample"
	fileBasePath = "D:\\temp\\goexample"
)

func main() {
	/* _, err := findLinks("http://books.studygolang.com/gobyexample/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2:	%v\n", err)
	} */
	/* for _, link := range links {
		fmt.Println(link)
	} */
	findLinks(url)
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	respBody := resp.Body
	doc, err := html.Parse(respBody)
	body, err := ioutil.ReadAll(respBody)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	saveToFile("index.html", string(body))
	links := visit(nil, doc)
	for _, link := range links {
		fmt.Println(link)
	}
	return nil, nil
}

func saveToFile(name, centent string) {
	fileName := fileBasePath + "\\" + name
	if file.FileExists(fileName) {
		fmt.Println("exist")
		os.Remove(fileName)

	} else {
		fmt.Println("not exist")

	}
	f, _ := os.Create(fileName)
	f.WriteString(centent)
	f.Close()
	fmt.Println("end")
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
