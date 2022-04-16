package crawler

import (
	"gocrawler/helpers"
	"io"
	"net/http"
)

type crawler struct {
	baseurl string
	links   []Link
	depth   int
}

const MAX_DEPTH int = 5

func NewCrawler(baseurl string, depth int) *crawler {
	return &crawler{
		baseurl: baseurl,
		links:   make([]Link, 0),
		depth:   depth,
	}
}

func (c *crawler) Spider() {
	baseBody := c.callBase()
	defer baseBody.Close()
}

func (c *crawler) callBase() io.ReadCloser {
	response, err := http.Get(c.baseurl)
	helpers.HnadleError(err)

	return response.Body
}
