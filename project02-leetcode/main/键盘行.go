package main

import (
	"fmt"
	"unicode"
)
//转自他人题解
func findWords(words []string) (ans []string) {
	const rowIdx = "12210111011122000010020202"
next:
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		for _, ch := range word[1:] {
			if rowIdx[unicode.ToLower(ch)-'a'] != idx {
				continue next
			}
		}
		ans = append(ans, word)
	}
	return
}

func main() {
	var words =[] string{"Hello","Alaska","Dad","Peace"}

	fmt.Println(findWords(words))
}

