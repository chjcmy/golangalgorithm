package programmer

import (
	"strconv"
	"strings"
	"testing"
)

func TestNumWord(t *testing.T) {

	lines := "one4seveneight"

	NumSolution(lines)

}

func NumSolution(line string) (i int) {

	f := strings.NewReplacer(
		"zero", "0",
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	i, _ = strconv.Atoi(f.Replace(line))
	return i

}
