package index

import (
	"GO_thinknetica/03-algorithms/GoSearch/pkg/crawler"
	"sort"
	"strings"
)

var MapIndex = make(map[string][]int)
var Documents []crawler.Document

func Create(data crawler.Document) {
	words := strings.Fields(data.Title)
	data.ID = len(Documents)
	Documents = append(Documents, data)
	for _, v := range words {
		MapIndex[v] = append(MapIndex[v], data.ID)
	}
	sort.Slice(Documents, func(i, j int) bool { return Documents[i].ID < Documents[j].ID })
}
