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

	m := make(map[string][]string)

	result2 := make([]int, len(id_list))

	for i := 0; i < len(report); i++ {
		values := strings.Split(report[i], " ")

		if len(m[values[1]]) > 0 {
			for _, s := range m[values[1]] {
				if s != values[0] {
					m[values[1]] = append(m[values[1]], values[0])
				}
			}
		} else {
			m[values[1]] = append(m[values[1]], values[0])
		}

	}

	for i := 0; i < len(id_list); i++ {
		if len(m[id_list[i]]) >= k {
			for _, s := range m[id_list[i]] {
				for y, i3 := range id_list {
					if i3 == s {
						result2[y]++
						break
					}
				}
			}

		}
	}

	return result2
}
