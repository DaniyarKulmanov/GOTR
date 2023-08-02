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
	if !errors.Is(err, fs.ErrNotExist) && err != nil {
		log.Println(err)
		return
	}

	if content == nil {
		docs, err = scanUrl()
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		err = json.Unmarshal(content, &docs)
		if err != nil {
			log.Println(err)
			return
		}
	}
	fmt.Println(docs)
}

func scanUrl() (data []crawler.Document, err error) {
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	data = search.ParseUrl(sts)
	file, _ := json.MarshalIndent(data, "", " ")
	err = os.WriteFile("test.json", file, 0644)
	return data, err
}
