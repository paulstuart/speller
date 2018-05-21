package main

import (
	"fmt"
	"os"

	"github.com/paulstuart/speller"
)

func main() {
	speller.InitFile(speller.Sample)
	for _, word := range os.Args[1:] {
		fmt.Println(word, "--", speller.Correction(word))
	}
}
