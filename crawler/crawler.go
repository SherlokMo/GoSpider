package crawler

import (
	"fmt"
	"goSpider/helpers"
	"goSpider/link"
	"goSpider/tokenizer"
	"io"
	"net/http"
)

type spider struct {
	baseurl string
	links   []link.Link
	depth   int
}

func NewSpider(baseurl string, depth int) *spider {
	return &spider{
		baseurl: baseurl,
		links:   make([]link.Link, 0),
		depth:   depth,
	}
}

func (s *spider) Crawl() {
	baseBody := s.callBase()
	defer baseBody.Close()
	spiderLegs := tokenizer.NewTokenizer(baseBody, s.baseurl)
	Links := spiderLegs.SplitAnchors()
	for _, Link := range *Links {
		fmt.Println(Link.Url)
	}
}

func (s *spider) callBase() io.ReadCloser {
	response, err := http.Get(s.baseurl)
	helpers.HnadleError(err)

	return response.Body
}
