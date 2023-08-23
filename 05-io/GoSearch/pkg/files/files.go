package main

import (
	"GOTR/05-io/GoSearch/pkg/search"
	"GOTR/GoSearch/pkg/crawler"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

// main TODO move to GoSearch/cmd/gosearch/gosearch.go
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
	var file []byte
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	data = search.ParseUrl(sts)
	if file, err = json.MarshalIndent(data, "", " "); err != nil {
		return nil, err
	}
	if err = os.WriteFile("test.json", file, 0644); err != nil {
		return nil, err
	}
	return data, err
}
