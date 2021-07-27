package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)

	var n int
	title := 666

	_, _ = fmt.Fscanln(br, &n)

	for i := 0; n != 0; i++ {
		num := i

		for num >= 666 {
			if num%1000 == 666 {
				n--
				break
			}
			num /= 10
		}

		title = i
	}

	_, _ = fmt.Fprint(bw, title)
	_ = bw.Flush()
}
