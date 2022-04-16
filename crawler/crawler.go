package crawler

import (
	"gocrawler/helpers"
	"io"
	"net/http"
)

type spider struct {
	baseurl string
	links   []Link
	depth   int
}

const MAX_DEPTH int = 5

func NewSpider(baseurl string, depth int) *spider {
	return &spider{
		baseurl: baseurl,
		links:   make([]Link, 0),
		depth:   depth,
	}
}

func (c *spider) Crawl() {
	baseBody := c.callBase()
	defer baseBody.Close()
}

func (c *spider) callBase() io.ReadCloser {
	response, err := http.Get(c.baseurl)
	helpers.HnadleError(err)

	return response.Body
}
