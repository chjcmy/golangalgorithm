package Calc

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var ac = []int{rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20)}
var bc = []int{rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20)}
var cc = []int{rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20)}
var dc = []int{rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20)}
var ec = []int{rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20)}
var ic = []int{rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20),
	rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20), rand.Intn(20)}

func TestAdd(t *testing.T) {
	cal(10, ac)
	cal(10, bc)
	cal(10, cc)
	cal(10, dc)
	cal(10, ec)
	cal(10, ic)
}

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
