package main

import (
	"fmt"
	"os"

	"github.com/paulstuart/speller"
)

func main() {
	for _, word := range os.Args[1:] {
		fmt.Println(word, "--", speller.Correction(word))
	}
}
