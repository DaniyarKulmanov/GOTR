package main

import (
	search "GOTR/03-algorithms/1-search"
	"GOTR/GoSearch/pkg/crawler"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	var docs []crawler.Document
	content, err := os.ReadFile("test.json")
	if errors.Is(err, fs.ErrNotExist) {
		docs = scanUrl()
	} else {
		if err != nil {
			log.Println(err)
			return
		}
	}

	if content != nil {
		err = json.Unmarshal(content, &docs)
		if err != nil {
			log.Println(err)
			return
		}
	}
	fmt.Println(docs)
}

func scanUrl() (data []crawler.Document) {
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	data = search.ParseUrl(sts)
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("test.json", file, 0644)
	return data
}
