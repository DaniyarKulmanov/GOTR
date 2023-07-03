package index

import (
	"GOTR/GoSearch/pkg/crawler"
	"sort"
	"strings"
)

var MapIndex = make(map[string][]int)
var Documents []crawler.Document

func Create(data crawler.Document) {
	words := strings.Fields(data.Title)
	Documents = append(Documents, data)

	for _, v := range words {
		v = strings.ToLower(v)
		MapIndex[v] = append(MapIndex[v], data.ID)
	}
	sort.Slice(Documents, func(i, j int) bool { return Documents[i].ID < Documents[j].ID })
}
