package index

import (
	"GOTR/GoSearch/pkg/crawler"
	"fmt"
)

func CreateIndex(data []crawler.Document) {
	mapIndex := make(map[string][]int)
	mapIndex["For"] = []int{1}
	mapIndex["Go"] = []int{0, 1}

	fmt.Println(mapIndex)
	fmt.Println(data[0])
}
