package crawler

import (
	"fmt"
	"goSpider/helpers"
	"goSpider/link"
	"goSpider/memorydb"
	"goSpider/tokenizer"
	"io"
	"net/http"
)

type spider struct {
	visited *memorydb.MemDB[string, bool]
	links   []link.Link
	depth   int
}

func NewSpider() *spider {
	return &spider{
		visited: memorydb.NewMemorydb[string, bool](),
		depth:   0,
	}
}

func (s *spider) Crawl(target string) {
	baseBody := s.callBase(target)
	defer baseBody.Close()
	s.markVisited(target)
	spiderLegs := tokenizer.NewTokenizer(baseBody, target)
	Links := spiderLegs.SplitAnchors()
	for _, Link := range *Links {
		if s.visited.Exist(Link.Url) != true {
			fmt.Println(Link.Url)
			s.Crawl(Link.Url)
		}
	}
}

func (s *spider) markVisited(t string) {
	s.visited.Store(t, true)
}

func (s *spider) callBase(target string) io.ReadCloser {
	response, err := http.Get(target)
	helpers.HnadleError(err)

	return response.Body
}
