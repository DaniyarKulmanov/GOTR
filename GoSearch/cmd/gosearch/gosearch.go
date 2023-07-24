package main

import (
	"GOTR/03-algorithms/1-search"
	"GOTR/GoSearch/pkg/crawler"
	"GOTR/GoSearch/pkg/crawler/spider"
	"GOTR/GoSearch/pkg/index"
	"flag"
	"fmt"
)

func main() {
	var word string
	var depth int
	var BinaryDocs []crawler.Document
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
		BinaryDocs = __search.Binary(index.MapIndex[word], index.Documents)
	}
	fmt.Println(index.Documents)
	fmt.Println(BinaryDocs)
}
