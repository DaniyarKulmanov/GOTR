package main

import (
	"GOTR/GoSearch/pkg/crawler/spider"
	"GOTR/GoSearch/pkg/index"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var search string
	var depth int
	sts := [4]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
		"https://en.wikipedia.org/wiki/Merge_sort",
		"https://golangify.com/array",
	}
	s := spider.New()
	flag.StringVar(&search, "s", "", "search links")
	flag.IntVar(&depth, "d", 1, "depth size")
	flag.Parse()

	for i := range sts {
		result, err := s.Scan(sts[i], depth)
		result[0].ID = len(index.Documents) + 1

		if err != nil {
			fmt.Println(err)
			continue
		}

		if search != "" {
			if strings.Contains(sts[i], search) {
				index.Create(result[0])
			}
		} else {
			index.Create(result[0])

		}
	}
	fmt.Println(index.MapIndex)
	fmt.Println(index.Documents)
}
