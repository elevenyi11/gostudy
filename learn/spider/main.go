package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SpiderContent struct {
	Href  string
	Title string
	Body  string
}

/*
1.抓取博客园首页
2.抓取二级页面内容
3.多个goroutine去抓取
*/

var (
	url = "https://www.cnblogs.com"
)

func main() {
	_, err := findLinks(url)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 10)
}

func findLinks(url string) ([]string, error) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	links := make([]string, 10, 10)

	doc.Find(".post_item_body").Each(func(i int, s *goquery.Selection) {
		l, ok := s.Find(".titlelnk").Attr("href")
		if !ok {
			return
		}
		links = append(links, l)
		fmt.Println("%s %s", l, s.Find(".titlelnk").Text())
		getContent(l)
	})

	return links, nil
}

func getContent(link string) (SpiderContent, error) {
	var body SpiderContent

	doc, err := goquery.NewDocument(link)
	if err != nil {
		return body, fmt.Errorf("%v", err)
	}
	body.Href = link
	body.Title = doc.Find("#cb_post_title_url").First().Text()
	body.Body = doc.Find("#cnblogs_post_body").First().Text()
	fmt.Println("%s : %s", body.Title, body.Body)
	return body, nil
}
