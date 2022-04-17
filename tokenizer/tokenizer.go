package tokenizer

import (
	"io"

	"golang.org/x/net/html"
)

type tokenizationManager struct {
	tokenizer *html.Tokenizer
	currToken token
}

func NewTokenizer(stream io.Reader) *tokenizationManager {
	return &tokenizationManager{
		tokenizer: html.NewTokenizer(stream),
	}
}

func (t *tokenizationManager) splitAnchors() {
	for {
		t.updateToken()
		if err := t.currToken.IsSafeToken(); err == io.EOF {
			break
		} else if t.currToken.IsTypeOf(html.StartTagToken) {
			if "a" == t.currToken.Data() {
				// TODO: Handle Link
				t.updateToken()
			}
		}
	}
}

func (t *tokenizationManager) updateToken() {
	t.tokenizer.Next()
	t.currToken = token{
		htmlToken: t.tokenizer.Token(),
		err:       t.tokenizer.Err(),
	}
}
