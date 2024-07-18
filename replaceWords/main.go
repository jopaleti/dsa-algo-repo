package main

import (
	"fmt"
	"strings"
)

// Defining Replace word function
func replaceWords(dictionary []string, sentence string) string {
	// Creating a Set for Dictionary Roots
	rootSet := make(map[string]struct{})

	for _, root := range dictionary {
		rootSet[root] = struct{}{}
	}

	// Splitting the Sentence into Words
	words := strings.Split(sentence, " ")

	// Replacing Words with Roots after iteration
	for i, word := range words {
		for j:=1; j<= len(word); j++ {
			if _, found := rootSet[word[:j]]; found {
				words[i] = word[:j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	result := replaceWords(dictionary, sentence)
	fmt.Println(result) // Output: "the cat was rat by the bat"
}