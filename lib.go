package speller

// These functions are the "library" functions the python version uses

import (
	"compress/gzip"
	"io/ioutil"
	"os"
	"strings"
)

func set(words ...string) []string {
	m := make(map[string]struct{})
	for _, word := range words {
		m[word] = struct{}{}
	}
	list := make([]string, 0, len(m))
	for word := range m {
		list = append(list, word)
	}
	return list
}

func sets(lists ...[]string) []string {
	m := make(map[string]struct{})
	for _, list := range lists {
		for _, word := range list {
			m[word] = struct{}{}
		}
	}
	list := make([]string, 0, len(m))
	for word := range m {
		list = append(list, word)
	}
	return list
}

func max(words []string, key func(string) float64) string {
	var top float64
	var word string
	for _, w := range words {
		if n := key(w); n > top {
			top = n
			word = w
		}
	}
	return word
}

func sum(list ...int) int {
	var total int
	for _, item := range list {
		total += item
	}
	return total
}

func fileReader(filename string) ([]byte, error) {
	if strings.HasSuffix(filename, ".gz") {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		r, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
		defer r.Close()
		return ioutil.ReadAll(r)
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data, nil
}

func readFile(filename string) string {
	file, err := fileReader(filename)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func counter(words []string) map[string]int {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}
