package main

import (
	"strings"
)

type Queue struct {
	items chan string
}

func (q *Queue) Set(value string) { q.items <- value }

func (q *Queue) Get() string { return <-q.items }

func main() {

	str := "zzzazz"

	var nuns int32

	nuns = 2

	getLargestString(str, nuns)
}

func getLargestString(s string, k int32) (val string) {
	// Write your code here
	nstring := strings.Split(s, "")

	q := Queue{items: make(chan string, len(nstring))}
	q2 := Queue{items: make(chan string, len(nstring))}

	for i := 0; i < len(nstring); i++ {
		var nk int32 = 1

		for j := i + 1; j < len(nstring); j++ {

			if nk != k {
				if nstring[i] == nstring[j] {
					q.Set(nstring[i])
					i = j
					nk++
				}
			} else if nk == k {
				break
			}
		}
		if nk != k {
			q2.Set(nstring[i])
		} else if nk == k {
			q.Set(nstring[i])
			continue
		}
	}

	for len(q.items) != 0 {
		val += q.Get()
	}
	for len(q2.items) != 0 {
		val += q2.Get()
	}

	return
}
