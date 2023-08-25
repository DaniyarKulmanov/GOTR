package main

import (
	"GOTR/05-io/GoSearch/pkg/crawler"
	"GOTR/05-io/GoSearch/pkg/crawler/spider"
	"GOTR/05-io/GoSearch/pkg/index"
	"GOTR/05-io/GoSearch/pkg/search"
	"flag"
	"fmt"
)

func main() {
	var word string
	var depth int
	var docs []crawler.Document

	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	s := spider.New()
	flag.StringVar(&word, "s", "", "word links")
	flag.IntVar(&depth, "d", 1, "depth size")
	flag.Parse()

	for i := range sts {
		result, err := s.Scan(sts[i], depth)
		if err != nil {
			fmt.Println(err)
			continue
		}
		index.Create(result[0])
	}
	if word != "" {
		docs = search.Binary(index.MapIndex[word], index.Documents)
	}
	fmt.Println(index.Documents)
	fmt.Println(docs)
}
