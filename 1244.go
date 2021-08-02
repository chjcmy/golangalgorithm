package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)
var out2 = bufio.NewWriter(os.Stdout)

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func main() {
	in.Split(bufio.ScanWords)
	n := nextInt()

	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	a[0], a[n+1] = -1, -2
	for m := nextInt(); m > 0; m-- {
		p, q := nextInt(), nextInt()
		if p == 1 {
			for i := q; i <= n; i += q {
				a[i] = 1 - a[i]
			}
		} else {
			a[q] = 1 - a[q]
			for i, j := q-1, q+1; a[i] == a[j]; i, j = i-1, j+1 {
				a[i] = 1 - a[i]
				a[j] = 1 - a[j]
			}
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Fprint(out, a[i])
		if i%20 == 0 {
			out.WriteByte('\n')
		} else {
			out.WriteByte(' ')
		}
	}
	out2.Flush()
}
