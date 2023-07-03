package index

import (
	"GOTR/GoSearch/pkg/crawler"
	"strings"
)

var MapIndex = make(map[string][]int)
var Documents []crawler.Document

func Create(data crawler.Document) {
	words := strings.Fields(data.Title)
	Documents = append(Documents, data)

	for _, v := range words {
		//dont collect if ID already exist
		MapIndex[v] = append(MapIndex[v], data.ID)
	}
}
