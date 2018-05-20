package speller

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

func (list splits) comp(fn func(split, *[]string)) []string {
	results := []string{}
	for _, s := range list {
		fn(s, &results)
	}
	return results
}
