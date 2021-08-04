package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	return n
}

func nextString() string {
	in.Scan()
	n := in.Text()
	return n
}

func main() {
	nx, ny := -1, -1
	x, y := 1, 1
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}
	moveTypes := []string{"L", "R", "U", "D"}
	in.Split(bufio.ScanWords)
	n := nextInt()
	var plans []string

	for i := range plans {
		plans[i] = nextString()
	}

	for plan := 0; plan < len(plans); plan++ {
		for i := 0; i < 4; i++ {
			if plans[plan] == moveTypes[i] {
				nx = x + dx[i]
				ny = y + dy[i]
			}
		}
		if nx < 1 || ny < 1 || nx > n || ny > n {
			continue
		}
		x, y = nx, ny
	}

	fmt.Printf(
		"%d %d",
		x,
		y,
	)

}
