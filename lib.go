package speller

// These functions are the "library" functions the python version uses
// as well as Go helpers to reduce verbosity

import (
	"compress/gzip"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// words breaks the given string into a slice of 'words'
func words(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z]+`) // original used `\w+` but we're only dealing with alpha characters so....
	return re.FindAllString(strings.ToLower(text), -1)
}

func slice(m map[string]struct{}) []string {
	list := make([]string, 0, len(m))
	for key := range m {
		list = append(list, key)
	}
	return list
}

func _set(m *map[string]struct{}, words ...string) {
	for _, word := range words {
		(*m)[word] = struct{}{}
	}
}

func set(words ...string) []string {
	m := make(map[string]struct{})
	_set(&m, words...)
	return slice(m)
}

func sets(lists ...[]string) []string {
	m := make(map[string]struct{})
	for _, list := range lists {
		_set(&m, list...)
	}
	return slice(m)
}

func max(words []string, key func(string) float64) (word string) {
	var top float64
	for _, w := range words {
		if n := key(w); n > top {
			top = n
			word = w
		}
	}
	return
}

func sum(list ...int) int {
	var total int
	for _, item := range list {
		total += item
	}
	return total
}

func fileReader(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if strings.HasSuffix(filename, ".gz") {
		r, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
		defer r.Close()
		return ioutil.ReadAll(r)
	}
	return ioutil.ReadAll(f)
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
