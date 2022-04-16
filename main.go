package main

import (
	"fmt"
	"goSpider/helpers"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		panic("Insert a start page you want to crawl")
	}

	startPage := os.Args[1]
	response, err := http.Get(startPage)
	helpers.HnadleError(err)
	defer response.Body.Close()
	tokenizer := html.NewTokenizer(response.Body)
	for {
		token := tokenizer.Next()
		if token == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}

			log.Fatalf("Error tokenizing HTML %v", tokenizer.Err())
		} else if token == html.StartTagToken {
			tokeen := tokenizer.Token()
			if "a" == tokeen.Data {
				handleLink(tokeen.Attr, startPage)
				token = tokenizer.Next()
				if token == html.TextToken {
					fmt.Println(tokenizer.Token().Data)
				}
			}
		}
	}
}

func handleLink(attributes []html.Attribute, baseURL string) {

	for i := range attributes {
		if attributes[i].Key == "href" {
			if attributes[i].Val[0] == '/' || attributes[i].Val[0] == '.' {
				fmt.Println(baseURL + "/" + attributes[i].Val[1:])
			} else {
				fmt.Println(attributes[i].Val)
			}
		}
	}

}
