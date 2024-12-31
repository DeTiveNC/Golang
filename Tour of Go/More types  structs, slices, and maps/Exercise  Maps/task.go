package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	f := make(map[string]int)
	for _, word := range strings.Fields(s) {
		f[word]++
	}
	return f
}

func main() {
	wc.Test(WordCount)
}
