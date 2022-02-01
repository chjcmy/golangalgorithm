package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	sc      *bufio.Scanner
	wr      *bufio.Writer
	N, M    int
	visited [1000][1000][2]bool
	arr     [1000][1000]int
	dx, dy  [4]int
	q       []status
)

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
	dx = [4]int{0, 1, 0, -1}
	dy = [4]int{1, 0, -1, 0}
}

type status struct {
	index
	cnt int
	brk int
}

type index struct {
	y, x int
}

func main() {
	defer wr.Flush()
	N = scanInt()
	M = scanInt()
	for i := 0; i < N; i++ {
		sc.Scan()
		for j, x := range sc.Bytes() {
			arr[i][j] = int(x - '0')
		}
	}
	wr.WriteString(strconv.Itoa(bfs()))
}

func bfs() int {
	q = append(q, status{index{0, 0}, 1, 0})
	for len(q) > 0 {
		tmp := q[0]
		q = q[1:]
		y := tmp.index.y
		x := tmp.index.x
		if y == N-1 && x == M-1 {
			return tmp.cnt
		}
		for i := 0; i < 4; i++ {
			newY := y + dy[i]
			newX := x + dx[i]
			if !check(newY, newX) {
				continue
			}
			if arr[newY][newX] == 1 && tmp.brk == 0 && !visited[newY][newX][1] {
				visited[newY][newX][1] = true
				q = append(q, status{index{newY, newX}, tmp.cnt + 1, 1})
			}
			if arr[newY][newX] == 0 && !visited[newY][newX][tmp.brk] {
				visited[newY][newX][tmp.brk] = true
				q = append(q, status{index{newY, newX}, tmp.cnt + 1, tmp.brk})
			}
		}
	}
	return -1
}

func check(i, j int) bool {
	if i < 0 || i >= N || j < 0 || j >= M {
		return false
	} else {
		return true
	}
}

func scanInt() (n int) {
	sc.Scan()
	n = 0
	for _, i := range sc.Bytes() {
		n *= 10
		n += int(i - '0')
	}
	return
}
