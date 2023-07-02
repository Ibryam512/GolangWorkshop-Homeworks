package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func censor() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := scanner.Text()
	scanner.Scan()
	missingLetters := scanner.Text()

	for i := 0; i < len(missingLetters); i++ {
		missingLetter := string(missingLetters[i])
		word = strings.Replace(word, "*", missingLetter, 1)
	}

	fmt.Println(word)
}
