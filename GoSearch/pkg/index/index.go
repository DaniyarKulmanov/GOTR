package index

import (
	"GOTR/GoSearch/pkg/crawler"
	"strings"
)

var MapIndex = make(map[string][]int)

func Create(data crawler.Document) {
	words := strings.Fields(data.Title)

	for _, v := range words {
		MapIndex[v] = append(MapIndex[v], data.ID)
	}
}
