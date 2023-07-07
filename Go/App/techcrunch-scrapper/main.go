package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	
	blogPosts, err := GetLatestBlogTitles("https://www.techcrunch.com/")
	if err != nil {
		log.Println(err)
		return
	}

	if blogPosts == "" {
		fmt.Println("No Blog Titles Found")
		return
	}

	fmt.Println("Blog Post Found: ")
	fmt.Println(blogPosts)
}

func GetLatestBlogTitles(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	blogPosts := ""

	doc.Find("div.river").Find("div.post-block").Each(func(i int, item *goquery.Selection) {
		title := strings.TrimSpace(item.Find("h2.post-block__title").Text())
		urls := strings.TrimSpace(item.Find("h2.post-block__title").Find("a").AttrOr("href", ""))
		author := strings.TrimSpace(item.Find("span.river-byline__authors").Find("a").Text())
		time := strings.TrimSpace(item.Find("time.river-byline__time").Text())

		blogPosts += "- " + title + " (by " + author + " at " + time + ")\n  " + urls + "\n\n"

	})

	return blogPosts, nil

}
