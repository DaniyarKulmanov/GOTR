package main

import (
	search "GOTR/03-algorithms/1-search"
	"fmt"
)

func main() {
	sts := [2]string{
		"https://go.dev/",
		"https://www.programiz.com/golang/",
	}
	fmt.Println(search.ParseUrl(sts))
}
