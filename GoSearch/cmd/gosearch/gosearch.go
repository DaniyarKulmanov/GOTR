package main

import (
	"GOTR/GoSearch/pkg/crawler"
	"GOTR/GoSearch/pkg/crawler/spider"
	"GOTR/GoSearch/pkg/index"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var data []crawler.Document
	var search string
	sts := [4]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
		"https://en.wikipedia.org/wiki/Merge_sort",
		"https://golangify.com/array",
	}
	s := spider.New()
	flag.StringVar(&search, "s", "", "search links")
	flag.Parse()

	for i := range sts {
		result, err := s.Scan(sts[i], 1)
		result[0].ID = len(data) + 1
		if err != nil {
			fmt.Println(err)
			continue
		}

		// append data and collect index
		if search != "" {
			if strings.Contains(sts[i], search) {
				data = append(data, result...)
			}
		} else {
			data = append(data, result...)
		}
		index.Create(result[0])
	}
	//fmt.Println(data)
	fmt.Println(index.MapIndex)
	fmt.Println(index.Documents)
}
