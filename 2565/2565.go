package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n        int
	min, max int
	at       [100]int
	bt       [100]int
	cn       [][]int
)

func main() {

	var acable, bcable, c int
	min = 1
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &c)
	for i := 1; i <= c; i++ {
		fmt.Fscanf(reader, "%d %d\n", &acable, &bcable)
		at[i] = acable
		bt[i] = bcable
		if min > acable && acable > bcable {
			min = acable
		} else if min > bcable && acable > bcable {
			min = bcable
		}
		if max < acable && acable > bcable {
			max = acable
		} else if max < bcable && acable < bcable {
			max = bcable
		} else if max < bcable && max < acable && acable == bcable {
			max = acable
		}
	}

	cn = make([][]int, max+1)

	for Atop := min; Atop <= max; Atop++ {
		comparison(Atop)
	}
	fmt.Fprintln(writer, n)
}

func comparison(a int) {
	atop := a
	btop := at[a]
	for i := 1; i <= max; i++ {
		if atop == i {
			continue
		}
		if atop < i && btop < at[i] || atop > i && btop > at[i] {
			cn[atop] = append(cn[atop], i)
		}
	}
}
