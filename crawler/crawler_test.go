package crawler

import (
	"bytes"
	"context"
	"testing"
)

func TestVisited(t *testing.T) {
	testSite := "https://example.com"
	spider := NewSpider(1)
	spider.addToVisited(testSite)
	if !spider.isVisited(testSite) {
		t.Error("Expected visited recieved false")
	}
}

func TestWebMethod(t *testing.T) {
	testSite := "https://google.com"
	stream := []byte(`
			<a href="https://example.com">E1</a>
			<a href="https://facebook.com">E2</a> 
			<a href="https://gmail.com">E3</a>
		`)
	reader := bytes.NewReader(stream)
	spider := NewSpider(1)
	Links := spider.web(reader, testSite)
	if len(*Links) != 3 {
		t.Errorf("Expected 3 sites recieved %v", len(*Links))
	}
}

func TestRetrievDepth(t *testing.T) {
	ctx := context.WithValue(context.Background(), DepthKey("depth"), 10)
	spider := NewSpider(1)

	if d := spider.retrieveDepth(ctx); d != 10 {
		t.Errorf("Expected 10 recieved %v", d)
	}
}
