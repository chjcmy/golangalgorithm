package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var xLay, yLay int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &xLay, &yLay)

	maps := make([][]int, xLay)

	for i := 0; i < len(maps); i++ {
		maps[i] = make([]int, yLay)
	}

	for i := 0; i < len(maps); i++ {
		var S string
		fmt.Fscanln(reader, &S)
		values := strings.Split(S, "")
		maps[i] = scanconv(values)

	}
	fmt.Fprintln(writer, maps)
}

func scanconv(strs []string) (val []int) {
	for _, raw := range strs {
		n, _ := strconv.Atoi(raw)
		val = append(val, n)
	}
	return
}
