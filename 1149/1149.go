package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var (
	houses [][]int
	bens   []int
)

func main() {
	var count, sum int
	fmt.Fscanln(reader, &count)
	houses = make([][]int, count)
	for home := 0; home < count; home++ {
		var home1, home2, home3 int
		fmt.Fscanf(reader, "%d %d %d\n", &home1, &home2, &home3)
		houses[home] = append(houses[home], home1)
		houses[home] = append(houses[home], home2)
		houses[home] = append(houses[home], home3)
	}
	for i := range houses {
		var min int
		for first := range houses[i] {
			if first == 0 {
				bens[0] = first
				min = houses[i][first]
				continue
			}
			if i != 0 {
				if bens[i-1] == first {
					continue
				}
			}
			if houses[i][first-1] > houses[i][first] {
				min = first
			}
		}
		bens = append(bens, min)
	}
	for bus := range bens {
		sum += bens[bus]
	}
	fmt.Fprintln(writer, sum)
	writer.Flush()
}
