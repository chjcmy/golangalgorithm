package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var N, M int
	var card []int

	if sc.Scan() {
		tmp := strings.Fields(sc.Text())
		N, _ = strconv.Atoi(tmp[0])
		M, _ = strconv.Atoi(tmp[1])
	}
	card = make([]int, N)
	if sc.Scan() {
		tmp := strings.Fields(sc.Text())
		for i, v := range tmp {
			card[i], _ = strconv.Atoi(v)
		}

	}

	var max int
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for K := j - 1; K < N; K++ {
				tmp := card[i] + card[j] + card[K]
				if tmp > max && tmp <= M {
					max = tmp
				}
			}
		}

	}
	fmt.Fprint(wr, max)
}
