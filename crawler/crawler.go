package crawler

import (
	"context"
	"goSpider/helpers"
	"goSpider/link"
	"goSpider/memorydb"
	"goSpider/tokenizer"
	"io"
	"log"
	"net/http"
	"sync"
)

const depthKey DepthKey = "depth"

type spider struct {
	visited  *memorydb.MemDB[string, bool]
	links    []link.Link
	maxDepth int
}

func NewSpider(maxDepth int) *spider {
	return &spider{
		visited:  memorydb.NewMemorydb[string, bool](),
		maxDepth: maxDepth,
	}
}

func (s *spider) Crawl(ctx context.Context, target string) {
	var wg sync.WaitGroup
	currDepth := s.retrieveDepth(ctx)
	if currDepth > s.maxDepth {
		return
	}
	baseBody := s.callBase(target)
	defer baseBody.Close()
	s.markVisited(target)
	spiderLegs := tokenizer.NewTokenizer(baseBody, target)
	Links := spiderLegs.SplitAnchors()
	for _, Link := range *Links {
		if s.visited.Exist(Link.Url) != true {
			log.Println(Link.Url)
			wg.Add(1)
			go func(u string) {
				defer wg.Done()
				ctx := context.WithValue(ctx, DepthKey("depth"), currDepth+1)
				s.Crawl(ctx, u)
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

func (s *spider) retrieveDepth(ctx context.Context) int {
	return ctx.Value(depthKey).(int)
}
