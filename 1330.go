package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	strLine, _ := br.ReadString('\n')
	tokens := strings.Fields(strLine)
	A, _ := strconv.Atoi(tokens[0])
	B, _ := strconv.Atoi(tokens[1])

	if A > B {
		_, _ = os.Stdout.Write([]byte(">"))
	} else if A < B {
		_, _ = os.Stdout.Write([]byte("<"))
	} else {
		_, _ = os.Stdout.Write([]byte("=="))
	}
}
