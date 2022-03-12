package programmer

import "testing"

func TestSquare(t *testing.T) {
	println(SquareSol(8, 12))
}

func SquareSol(w, h int) (result int64) {

	w1 := int64(w)
	h1 := int64(h)

	return (w1 * h1) - (w1 + h1 - gcd(w1, h1))
}

func gcd(w, h int64) int64 {
	big, small := bigsmall(w, h)

	for small != 0 {
		nmg := big % small
		big = small
		small = nmg
	}
	return big
}

func bigsmall(w, h int64) (big, small int64) {

	if w > h {
		big = w
		small = h
	} else {
		big = h
		small = w
	}

	return
}
