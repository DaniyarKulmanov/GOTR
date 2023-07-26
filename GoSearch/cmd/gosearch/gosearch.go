package main

import (
	"GOTR/03-algorithms/1-search"
	"GOTR/GoSearch/pkg/crawler"
	"GOTR/GoSearch/pkg/crawler/spider"
	"GOTR/GoSearch/pkg/index"
	"flag"
	"fmt"
)

// #TODO move search to 03-algorithms/1-search/search.go
func main() {
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	fmt.Println(ParseUrl(sts))
}

func ParseUrl(url [2]string) (docs []crawler.Document) {
	var word string
	var depth int
	var BinaryDocs []crawler.Document

	s := spider.New()
	flag.StringVar(&word, "s", "", "word links")
	flag.IntVar(&depth, "d", 1, "depth size")
	flag.Parse()

	for i := range url {
		result, err := s.Scan(url[i], depth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		index.Create(result[0])
	}
	if word != "" {
		BinaryDocs = __search.Binary(index.MapIndex[word], index.Documents)
	}

	if BinaryDocs != nil {
		docs = BinaryDocs
	} else {
		docs = index.Documents
	}

	return docs
}
