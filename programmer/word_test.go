package programmer

import (
	"fmt"
	"testing"
)

func TestWord(t *testing.T) {

	word1 := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	word2 := []string{"hot", "dot", "dog", "lot", "log"}

	fmt.Println(hitting("hit", "cog", word1))
	fmt.Println(hitting("hit", "cog", word2))
}

func hitting(begin string, target string, words []string) int {

	found := false
	wordLength := make(map[string]int)
	stack := make([]string, 0)
	for _, word := range words {
		if word == target {
			found = true
		}
		wordLength[word] = -1
	}
	if !found {
		return 0
	}
	wordLength[begin] = 0

	stack = append(stack, begin)
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		length := wordLength[cur] + 1

		for _, word := range words {
			if wordLength[word] >= 0 {
				continue
			}
			cnt := 0
			for idx := range cur {
				if word[idx] != cur[idx] {
					cnt++
				}
			}
			if cnt == 1 {
				if word == target {
					return length
				} else {
					wordLength[word] = length
					stack = append(stack, word)
				}
			}
		}
	}
	return 0
}
