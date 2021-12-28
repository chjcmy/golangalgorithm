package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	for num1 := 1; num1 < 10; num1++ {
		if num1 == 8 {
			continue
		}
		fmt.Fprintf(writer, "2 * %d = %d  3 * %d = %d 4 * %d = %d 5 * %d = %d \n",
			num1, 2*num1, num1, 3*num1, num1, 4*num1, num1, 5*num1)
	}
	fmt.Fprintln(writer)
	for num1 := 1; num1 < 10; num1++ {
		if num1 == 8 {
			continue
		}
		fmt.Fprintf(writer, "6 * %d = %d  7 * %d = %d 8 * %d = %d 9 * %d = %d \n",
			num1, 6*num1, num1, 7*num1, num1, 8*num1, num1, 9*num1)
	}

	writer.Flush()
}

func swap(p, q int) {
	oldP := p
	p = q
	q = oldP
}
