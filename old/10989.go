package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	var arr = make([]int, 10001)
	for i := 0; i < n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		arr[num]++
	}

	for i := 1; i < 10001; i++ {
		for j := 0; j < arr[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}
}
