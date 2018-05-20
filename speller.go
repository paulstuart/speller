package speller

import (
	"regexp"
	"strings"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyz"
	sample  = "big.txt.gz"
)

var (
	WORDS map[string]int
	TOTAL int
)

func words(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z]+`) // original used `\w+` but we're only dealing with alpha characters so....
	return re.FindAllString(strings.ToLower(text), -1)
}

func init() {
	WORDS = counter(words(readFile(sample)))
	for _, occurs := range WORDS {
		TOTAL += occurs
	}
}

// Correction returns the most probable spelling correction for word
func Correction(word string) string {
	return max(candidates(word), probability)
}

// transpose transposes the first 2 characters
func transpose(s split, list *[]string) {
	if len(s.R) > 1 {
		*list = append(*list, s.L+s.R[1:2]+s.R[0:1]+s.R[2:])
	}
}

// replace replaces the "split" character with alphabetic permutations
func replace(s split, list *[]string) {
	if len(s.R) > 0 {
		for _, b := range letters {
			c := string(b)
			*list = append(*list, s.L+c+s.R[1:])
		}
	}
}

// deleted returns splits with non-empty R side
func deleted(s split, list *[]string) {
	if s.R != "" {
		*list = append(*list, s.L+s.R[1:])
	}
}

// insert inserts the alphabet mid-split
func insert(s split, list *[]string) {
	for _, c := range letters {
		*list = append(*list, s.L+string(c)+s.R)
	}
}

func edits1(word string) []string {
	list := cleaves(word)
	return sets(list.comp(deleted), list.comp(transpose), list.comp(replace), list.comp(insert))
}

// edits2 returns all edits that are two edits away from `word`
func edits2(word string) (results []string) {
	for _, e1 := range edits1(word) {
		for _, e2 := range edits1(e1) {
			results = append(results, e2)
		}
	}
	return results
}

// includes returns a subset of words that exist in corpus
func includes(words []string) (results []string) {
	for _, word := range words {
		if _, ok := WORDS[word]; ok {
			results = append(results, word)
		}
	}
	return
}

// known returns a unique subset of words that exist in the corpus
func known(words []string) []string {
	return set(includes(words)...)
}

// candidates generate possible spelling corrections for word
func candidates(word string) []string {
	self := []string{word}
	if list := sets(known(self)); len(list) > 0 {
		return list
	}
	if list := sets(edits1(word)); len(list) > 0 {
		return list
	}
	if list := sets(edits2(word)); len(list) > 0 {
		return list
	}
	return self
}

// probability returns the percent of times `word` is found in the corpus
func probability(word string) float64 {
	return float64(WORDS[word]) / float64(TOTAL)
}
