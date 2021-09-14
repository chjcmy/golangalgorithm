package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var maps int

	fmt.Scan(&maps)

	sc := bufio.NewScanner(os.Stdin)

	var walk string

	var x, y = 1, 1

	var ny, nx = 1, 1

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	var moveType = [4]string{"L", "R", "U", "D"}

	sc.Scan()

	walk = sc.Text()

	plans := strings.Split(walk, " ")

	for plan := range plans {
		for i := range moveType {
			if plans[plan] == moveType[i] {
				nx = x + dx[i]
				ny = y + dy[i]
			}
		}
		if nx > maps || ny > maps || nx < 1 || ny < 1 {
			continue
		}
		x, y = nx, ny
	}
	fmt.Printf("x 는 %d, y는 %d", x, y)
}
