package search

import (
	"GOTR/GoSearch/pkg/crawler"
	"GOTR/GoSearch/pkg/crawler/spider"
	"GOTR/GoSearch/pkg/index"
	"flag"
	"fmt"
)

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
		BinaryDocs = binary(index.MapIndex[word], index.Documents)
	}

	if BinaryDocs != nil {
		docs = BinaryDocs
	} else {
		docs = index.Documents
	}

	return docs
}

func binary(words []int, allDocs []crawler.Document) []crawler.Document {
	var Docs []crawler.Document
	for _, word := range words {
		low := 0
		high := len(allDocs) - 1

		for low <= high {
			median := (low + high) / 2

			if allDocs[median].ID < word {
				low = median + 1
			} else {
				high = median - 1
			}
			if low == len(allDocs) || allDocs[low].ID != word {
				continue
			}
			if canAdd(allDocs[low], Docs) {
				Docs = append(Docs, allDocs[low])
			}
		}
	}

	return Docs
}

func canAdd(doc crawler.Document, docs []crawler.Document) bool {
	for _, v := range docs {
		if v == doc {
			return false
		}
	}
	return true
}
