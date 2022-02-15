package programmer

import (
	"fmt"
	"testing"
)

var (
	word       []string
	result2    int
	targetting string
)

func TestWord(t *testing.T) {

	word = []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println(hitting("hit", "cog", word))
}

func hitting(begin string, target string, words []string) int {

	var inside bool

	word = words

	targetting = target

	for i := 0; i < len(words); i++ {
		if targetting == words[i] {
			inside = true
		}
	}

	if !inside {
		return 0
	}

	for i := 0; i < len(word[0]); i++ {
		var diff int
		if begin[i] != word[0][i] {
			diff++
			if diff > 1 {
				return 0
			}
		}
	}

	hitdfs(0, word[0])

	return result2
}

func hitdfs(value int, remind string) {

	var diff int

	for i := 0; i < len(word[value+1]); i++ {
		if remind[i] != word[value+1][i] {
			if diff > 1 {
				remind = word[value]
			}
			diff++
		}
	}
	if word[value] == targetting {
		return
	}
	result2++
	hitdfs(value+1, remind)
}
