package main

import (
	"GOTR/03-algorithms/1-search"
	"GOTR/GoSearch/pkg/crawler"
	"fmt"
	"io"
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
	f, err := os.Create("./file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	err = os.WriteFile(f.Name(), []byte("Текст"), 0666)
	if err != nil {
		log.Println(err)
		return
	}
	// save data to file
	//saveFile()
}

func saveFile(docs []crawler.Document, w io.Writer) error {
	_, err := w.Write([]byte("Текст"))
	fmt.Println(docs)
	return err
}
