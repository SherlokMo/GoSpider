package crawler

import (
	"goSpider/helpers"
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

func (s *spider) Crawl() {
	baseBody := s.callBase()
	defer baseBody.Close()
}

func (s *spider) callBase() io.ReadCloser {
	response, err := http.Get(s.baseurl)
	helpers.HnadleError(err)

	return response.Body
}
