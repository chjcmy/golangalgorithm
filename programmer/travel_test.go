package programmer

import (
	"fmt"
	"sort"
	"testing"
)

func TestTravel(t *testing.T) {

	trip := [][]string{{"ICN", "JFK"}, {"HND", "IAD"}, {"JFK", "HND"}}

	fmt.Println(Travelsolution(trip))

}

func Travelsolution(tickets [][]string) []string {
	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i][1] < tickets[j][1]
	})

	maxVisitedCount := 0
	visited := make([]bool, len(tickets))
	answer, cities := []string{}, []string{}

	var dfs func(city string, count int)
	dfs = func(city string, count int) {
		cities = append(cities, city)

		if maxVisitedCount < count {
			maxVisitedCount = count
			answer = append(make([]string, 0, len(cities)), cities...)
		}

		for i, v := range tickets {
			if city == v[0] && !visited[i] {
				visited[i] = true
				dfs(v[1], count+1)
				visited[i] = false
			}
		}
		cities = cities[:len(cities)-1]
	}

	dfs("ICN", 0)
	return answer
}
