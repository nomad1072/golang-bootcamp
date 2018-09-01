package main

import (
	"strings"
	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	res := map[string]int{}

	words := strings.Fields(s)

	for _, word := range words {
		res[word] += 1
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
