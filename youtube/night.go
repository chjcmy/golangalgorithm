package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var count int
	var a string
	var stay [2]int
	var move = [][]int{{-2, -1}, {-1, -1}, {-2, 1}, {-1, 2}, {1, -2}, {2, -1}, {1, 2}, {2, 1}}
	m := [2]int{8, 8}

	fmt.Scan(&a)

	slicemove := strings.Split(a, "")

	runes := []rune(slicemove[0])

	stay[0] = int(runes[0]) - int('a') + 1
	stay[1], _ = strconv.Atoi(slicemove[1])

	for i := range move {
		x := stay[0] + move[i][0]
		y := stay[1] + move[i][1]

		if x < m[0] && y < m[1] && x > 0 && y > 0 {
			count += 1
		}
	}

	fmt.Print(count)

}
