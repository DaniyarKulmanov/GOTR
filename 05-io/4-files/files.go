package main

import (
	"GOTR/03-algorithms/1-search"
	"GOTR/GoSearch/pkg/crawler"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// scan data
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	docs := search.ParseUrl(sts)
	fmt.Println(docs)
	// create file
	f, err := os.Create("./scan.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	//
	//err = os.WriteFile(f.Name(), []byte("Текст"), 0666)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	// convert to json, save to file with json
	file, _ := json.MarshalIndent(docs, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
	//readFile json()
	content, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File contents: %s", content)
	// unmarshal json to struct
	var m []crawler.Document
	err = json.Unmarshal(content, &m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
}

func saveFile(docs []crawler.Document, w io.Writer) error {
	_, err := w.Write([]byte("Текст"))
	fmt.Println(docs)
	return err
}
