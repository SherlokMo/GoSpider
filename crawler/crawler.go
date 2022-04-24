package crawler

import (
	"goSpider/helpers"
	"goSpider/link"
	"goSpider/memorydb"
	"goSpider/tokenizer"
	"io"
	"log"
	"net/http"
	"sync"
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

func (s *spider) Crawl(target string, depth int) {
	var wg sync.WaitGroup
	baseBody := s.callBase(target)
	defer baseBody.Close()
	if depth < 0 {
		return
	}
	s.markVisited(target)
	spiderLegs := tokenizer.NewTokenizer(baseBody, target)
	Links := spiderLegs.SplitAnchors()
	for _, Link := range *Links {
		if s.visited.Exist(Link.Url) != true {
			log.Println(Link.Url)
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				s.Crawl(u, depth-1)
			}(Link.Url)
		}
	}
	wg.Wait()
}

func (s *spider) markVisited(t string) {
	s.visited.Store(t, true)
}

func (s *spider) callBase(target string) io.ReadCloser {
	response, err := http.Get(target)
	helpers.HnadleError(err)

	return response.Body
}
