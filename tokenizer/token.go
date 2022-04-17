package tokenizer

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

type token struct {
	htmlToken html.Token
	err       error
}

func (t token) Attr() []html.Attribute {
	return t.htmlToken.Attr
}

func (t token) Data() string {
	return t.htmlToken.Data
}

func (t token) IsTypeOf(tokenType html.TokenType) bool {
	return t.htmlToken.Type == tokenType
}

func (t token) IsSafeToken() error {

	if t.err != nil && t.err != io.EOF {
		log.Fatalf("Error tokenizing HTML %v", t.err)
	}

	return t.err
}
