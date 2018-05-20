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

type split struct {
	L, R string
}

type splits []split

func cleaves(word string) (results splits) {
	for i := range word {
		results = append(results, split{word[:i], word[i:]})
	}
	results = append(results, split{word, ""})
	return results
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

func deletes(list splits) (results []string) {
	for _, s := range list {
		if s.R != "" {
			results = append(results, s.L+s.R[1:])
		}
	}
	return results
}

func transposes(list splits) (results []string) {
	for _, s := range list {
		if len(s.R) > 1 {
			results = append(results, s.L+s.R[1:2]+s.R[0:1]+s.R[2:])
		}
	}
	return results
}

// replaces replaces the "split" character with alphabetic permutations
func replaces(list splits) (results []string) {
	for _, s := range list {
		if len(s.R) > 0 {
			for _, b := range letters {
				c := string(b)
				results = append(results, s.L+c+s.R[1:])
			}
		}
	}
	return results
}

func inserts(list splits) (results []string) {
	for _, s := range list {
		for _, c := range letters {
			results = append(results, s.L+string(c)+s.R)
		}
	}
	return results
}

func edits1(word string) []string {
	list := cleaves(word)
	return sets(deletes(list), transposes(list), replaces(list), inserts(list))
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
