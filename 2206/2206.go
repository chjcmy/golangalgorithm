package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	maps  [][]string
	graph [][]int
	X, Y  int
	hint  = 1
	stack []s
)

type s struct {
	X int
	Y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var S string

	var sum int

	fmt.Fscanf(reader, "%d %d\n", &Y, &X)

	maps = make([][]string, Y)

	for i := 0; i < Y; i++ {
		maps[i] = make([]string, X)
	}

	graph = make([][]int, Y)

	for i := 0; i < Y; i++ {
		graph[i] = make([]int, X)
	}

	for i := 0; i < Y; i++ {
		fmt.Fscanln(reader, &S)
		maps[i] = strings.Split(S, "")
	}

	stack = append(stack, s{0, 0})

	graph[0][0] = 1

	for stack[len(stack)-1].Y != Y-1 && stack[len(stack)-1].X != X-1 || len(stack) != 0 {

		n := len(stack) - 1

		beforeY := stack[n].Y
		beforeX := stack[n].X
		if beforeX+1 < X && maps[beforeY][beforeX+1] == "0" && hint >= 0 && graph[beforeY][beforeX+1] == 0 {
			sum++
			graph[beforeY][beforeX+1] = 1
			stack = append(stack, s{beforeX + 1, beforeY})
			continue
		}
		if beforeY+1 < Y && maps[beforeY+1][beforeX] == "0" && hint >= 0 && graph[beforeY+1][beforeX] == 0 {
			sum++
			graph[beforeY+1][beforeX] = 1
			stack = append(stack, s{beforeX, beforeY + 1})
			continue
		}
		if beforeX+1 < X && maps[beforeY][beforeX+1] == "1" && hint == 1 && graph[beforeY][beforeX+1] == 0 {
			hint--
			sum++
			graph[beforeY][beforeX+1] = 1
			stack = append(stack, s{beforeX + 1, beforeY})
			continue
		}
		if beforeY+1 < Y && maps[beforeY+1][beforeX] == "1" && hint == 1 && graph[beforeY+1][beforeX] == 0 {
			hint--
			sum++
			graph[beforeY+1][beforeX] = 1
			stack = append(stack, s{beforeX, beforeY + 1})
			continue
		}
		if beforeY-1 >= 0 && maps[beforeY-1][beforeX] == "0" && hint >= 0 && graph[beforeY-1][beforeX] == 0 {
			sum++
			graph[beforeY-1][beforeX] = 1
			stack = append(stack, s{beforeX, beforeY - 1})
			continue
		}
		if beforeX-1 >= 0 && maps[beforeY][beforeX-1] == "0" && hint >= 0 && graph[beforeY][beforeX-1] == 0 {
			sum++
			graph[beforeY][beforeX-1] = 1
			stack = append(stack, s{beforeX - 1, beforeY})
			continue
		}
		if beforeY-1 < Y && maps[beforeY-1][beforeX] == "1" && hint == 1 && graph[beforeY-1][beforeX] == 0 {
			hint--
			sum++
			graph[beforeY-1][beforeX] = 1
			stack = append(stack, s{beforeX, beforeY - 1})
			continue
		}
		if beforeX-1 > 0 && maps[beforeY][beforeX-1] == "1" && hint == 1 && graph[beforeY][beforeX-1] == 0 {
			hint--
			sum++
			graph[beforeY][beforeX-1] = 1
			stack = append(stack, s{beforeX - 1, beforeY})
			continue
		}
		if beforeX < X || beforeX-1 >= 0 || beforeY < Y || beforeY >= 0 ||
			maps[beforeY][beforeX+1] == "1" || maps[beforeY][beforeX-1] == "1" ||
			maps[beforeY+1][beforeX] == "1" || maps[beforeY-1][beforeX] == "1" {
			sum--
			stack = remove(stack, n)
			if hint == 0 && maps[beforeY][beforeX] == "1" {
				hint++
			}
			if hint == 0 && maps[beforeY][beforeX] == "0" {
				sum = -1
				break
			}
			continue
		}
	}
	fmt.Fprintln(writer, sum)

	writer.Flush()
}

func remove(s []s, i int) []s {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

//func normalMove(y, x, sum int) int {
//
//	if graph[X-1][Y-1] > 0 {
//		stack = stack
//		return sum
//	}
//	if y-1 > 0 && x-1 > 0 && maps[y-1][x-1] == "0" && hint >= 0 {
//		normalMove(y-1, x-1, sum+1)
//	}
//	if y+1 < Y && x+1 < X && maps[y+1][x+1] == "0" && hint >= 0 {
//		normalMove(y+1, x+1, sum+1)
//	}
//	if y-1 > 0 && x+1 < X && maps[y-1][x+1] == "0" && hint >= 0 {
//		normalMove(y-1, x+1, sum+1)
//	}
//	if y+1 < Y && x-1 > 0 && maps[y+1][x-1] == "0" && hint >= 0 {
//		normalMove(y+1, x-1, sum+1)
//	}
//	if y+1 < Y && x+1 < X && maps[y+1][x+1] == "1" && hint >= 1 {
//		hintMove(y+1, x+1, sum+1)
//	}
//	if y-1 >= 0 && x-1 >= 0 && maps[y-1][x-1] == "1" && hint >= 1 {
//		hintMove(y-1, x-1, sum+1)
//	}
//	if y-1 >= 0 && x+1 < X && maps[y-1][x+1] == "1" && hint >= 1 {
//		hintMove(y-1, x+1, sum+1)
//	}
//	if y+1 < Y && x-1 >= 0 && maps[y+1][x-1] == "1" && hint >= 1 {
//		hintMove(y+1, x-1, sum+1)
//	}
//	return -1
//}
//
//func hintMove(y, x, sum int) {
//	hint--
//
//	if y-1 > 0 && x-1 > 0 && maps[y-1][x-1] == "0" && hint >= 0 {
//		normalMove(y-1, x-1, sum+1)
//	}
//	if y-1 > 0 && x+1 < X && maps[y-1][x+1] == "0" && hint >= 0 {
//		normalMove(y-1, x+1, sum+1)
//	}
//	if y+1 < Y && x+1 < X && maps[y-1][x+1] == "0" && hint >= 0 {
//		normalMove(y+1, x+1, sum+1)
//	}
//	if y+1 < Y && x-1 > 0 && maps[y-1][x+1] == "0" && hint >= 0 {
//		normalMove(y+1, x-1, sum+1)
//	}
//}
