package main

import (
	"GOTR/GoSearch/pkg/crawler"
	"GOTR/GoSearch/pkg/crawler/spider"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var data []crawler.Document
	var search string
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	s := spider.New()
	flag.StringVar(&search, "s", "", "search links")
	flag.Parse()

	for i := range sts {
		result, err := s.Scan(sts[i], 1)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if search != "" {
			if strings.Contains(sts[i], search) {
				data = append(data, result...)
			}
		} else {
			data = append(data, result...)
		}
	}
	fmt.Println(data)
}
