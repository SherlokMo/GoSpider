package main

import (
	"context"
	"goSpider/crawler"
	"os"
)

const MAX_DEPTH int = 1

func main() {
	if len(os.Args) < 2 {
		panic("Insert a start page you want to crawl")
	}

	baseUrl := os.Args[1]
	ctx := context.WithValue(context.Background(), crawler.DepthKey("depth"), 0)
	spider := crawler.NewSpider(MAX_DEPTH)
	spider.Crawl(ctx, baseUrl)
}
