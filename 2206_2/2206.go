package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var xLay, yLay int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &xLay, &yLay)

	maps := make([][]string, xLay)

	for i := 0; i < len(maps); i++ {
		maps[i] = make([]string, yLay)
	}

	for i := 0; i < len(maps); i++ {
		var S string
		fmt.Fscanln(reader, &S)

		maps[i] = strings.Split(S, "")
	}
	fmt.Fprintln(writer, maps)
}
