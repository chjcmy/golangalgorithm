package programmer

import (
	"fmt"
	"strings"
	"testing"
)

type Sol struct {
	id_list []string
	report  []string
	k       int
}

func TestDeclaration(t *testing.T) {

	Sol1 := Sol{id_list: []string{"muzi", "frodo", "apeach", "neo"},
		report: []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"},
		k:      2,
	}

	Sol2 := Sol{id_list: []string{"con", "ryan"},
		report: []string{"ryan con", "ryan con", "ryan con", "ryan con"},
		k:      3,
	}

	fmt.Println(Declarationsolution(Sol1.id_list, Sol1.report, Sol1.k))
	fmt.Println(Declarationsolution(Sol2.id_list, Sol2.report, Sol2.k))
}

func Declarationsolution(id_list []string, report []string, k int) []int {

	a := map[string]map[string]bool{}
	var r = make([]int, len(id_list))

	for _, v := range report {
		c := strings.Split(v, " ")
		if _, ok := a[c[1]]; !ok {
			a[c[1]] = make(map[string]bool)
		}
		a[c[1]][c[0]] = true
	}

	for _, v := range a {
		if len(v) > k-1 {
			for k, _ := range v {
				for i, n := range id_list {
					if n == k {
						r[i]++
					}
				}
			}
		}
	}

	return r
}
