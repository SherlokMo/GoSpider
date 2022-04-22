package tokenizer

import (
	"goSpider/helpers"
	"goSpider/link"
	"io"

	"golang.org/x/net/html"
)

type tokenizationManager struct {
	baseurl   string
	tokenizer *html.Tokenizer
	currToken token
}

func NewTokenizer(stream io.Reader, baseurl string) *tokenizationManager {
	return &tokenizationManager{
		baseurl:   baseurl,
		tokenizer: html.NewTokenizer(stream),
	}
}

func (t *tokenizationManager) SplitAnchors() *[]link.Link {
	var anchors []link.Link
	for {
		t.updateToken()
		if err := t.currToken.IsSafeToken(); err == io.EOF {
			break
		} else if t.currToken.IsTypeOf(html.StartTagToken) {
			if "a" == t.currToken.Data() {
				hyperLink, err := t.getHyperLink()
				t.updateToken()
				if err == nil {
					anchors = append(anchors, *link.NewLink(hyperLink))
				}
			}
		}
	}

	return &anchors
}

func (t tokenizationManager) getHyperLink() (string, error) {
	attributes := t.currToken.Attr()
	for _, attr := range attributes {
		if attr.Key == "href" {
			HyperLink, err := helpers.ProvisionURL(attr.Val, t.baseurl)
			return HyperLink, err
		}
	}
	return "", nil
}

func (t *tokenizationManager) updateToken() {
	t.tokenizer.Next()
	t.currToken = token{
		htmlToken: t.tokenizer.Token(),
		err:       t.tokenizer.Err(),
	}
}
