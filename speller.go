package speller

const (
	letters = "abcdefghijklmnopqrstuvwxyz"

	// Sample provides a corpus used to train the dictionary
	Sample = "big.txt.gz"
)

var (
	wordFreq  map[string]int
	wordTotal int
)

// probability returns the percent of times `word` is found in the corpus
func probability(word string) float64 {
	return float64(wordFreq[word]) / float64(wordTotal)
}

// Correction returns the most probable spelling correction for word
func Correction(word string) string {
	return max(Candidates(word), probability)
}

// deleted returns splits with non-empty R side
func deleted(s split, save func(string)) {
	if s.R != "" {
		save(s.L + s.R[1:])
	}
}

// transpose transposes the first 2 characters
func transpose(s split, save func(string)) {
	if len(s.R) > 1 {
		save(s.L + s.R[1:2] + s.R[0:1] + s.R[2:])
	}
}

// replace replaces the "split" character with alphabetic permutations
func replace(s split, save func(string)) {
	if len(s.R) > 0 {
		for _, b := range letters {
			c := string(b)
			save(s.L + c + s.R[1:])
		}
	}
}

// insert inserts the alphabet mid-split
func insert(s split, save func(string)) {
	for _, c := range letters {
		save(s.L + string(c) + s.R)
	}
}

// Candidates generate possible spelling corrections for word
func Candidates(word string) []string {
	self := []string{word}
	if list := set(known(self)...); len(list) > 0 {
		return list
	}
	if list := set(edits1(word)...); len(list) > 0 {
		return list
	}
	if list := set(edits2(word)...); len(list) > 0 {
		return list
	}
	return self
}

// known returns the subset of `words` that appear in the dictionary of wordFreq
func known(words []string) []string {
	m := make(map[string]struct{})
	for _, word := range words {
		if _, ok := wordFreq[word]; ok {
			m[word] = struct{}{}
		}
	}
	return slice(m)
}

// edits1 returns all edits that are one edit away from `word`
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

// InitFile initializes the dictionary using the given filename
func InitFile(filename string) {
	wordFreq = counter(words(readFile(filename)))
	for _, occurs := range wordFreq {
		wordTotal += occurs
	}
}
