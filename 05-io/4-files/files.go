package main

import (
	"GOTR/GoSearch/pkg/crawler"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
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
