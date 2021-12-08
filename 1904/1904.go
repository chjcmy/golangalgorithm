package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var tile int

	fmt.Fscanln(reader, &tile)
	target1, target2 := 1, 1
	for ; tile > 0; tile-- {
		target1, target2 = target2, (target1+target2)%15746
	}
	fmt.Fprintln(writer, target1)
	writer.Flush()
}
