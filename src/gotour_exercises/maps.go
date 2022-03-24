package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		v, exist := m[word]
		if exist == true {
			m[word] = v + 1
		} else {
			m[word] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
