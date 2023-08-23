package search

import (
	"GO_thinknetica/03-algorithms/GoSearch/pkg/crawler"
)

func Binary(words []int, allDocs []crawler.Document) []crawler.Document {
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
