package main

import (
	"fmt"

	"github.com/detivenc/webscraping-golang/pkg/search"
)

func main() {
	s := search.GoogleSearcher{}
	sm := search.NewSearchManager(s, false)

	sm.Search("What is 2+2")
	fmt.Println(sm.GetResults())
}
