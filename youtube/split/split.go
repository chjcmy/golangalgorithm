package split

import (
	"bufio"
	"os"
	"strconv"
)

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	return n
}

func nextString() string {
	in.Scan()
	n := in.Text()
	return n
}
