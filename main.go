package main

import (
	"goSpider/crawler"
	"os"
)

const MAX_DEPTH int = 5

func main() {
	if len(os.Args) < 2 {
		panic("Insert a start page you want to crawl")
	}

	baseUrl := os.Args[1]
	spider := crawler.NewSpider(baseUrl, MAX_DEPTH)
	spider.Crawl()
}
