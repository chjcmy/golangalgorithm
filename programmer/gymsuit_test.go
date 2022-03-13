package programmer

import "testing"

func TestGymSuit(t *testing.T) {

	println(GymSuitSol(5, []int{2, 4}, []int{1, 3, 5}))
	println(GymSuitSol(5, []int{2, 4}, []int{3}))
	println(GymSuitSol(3, []int{3}, []int{1}))
	println(GymSuitSol(30, []int{1, 2, 7, 9, 10, 11, 12, 14, 15, 16, 18, 20, 21, 24, 25, 27}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 20, 22, 23, 24, 25, 26, 27, 28}))
}

func GymSuitSol(n int, lost []int, reserve []int) int {

	SuitNotHave := make([]int, n+1)

	for i := 0; i < len(lost); i++ {
		SuitNotHave[lost[i]]--
	}
	for i := 0; i < len(reserve); i++ {
		SuitNotHave[reserve[i]]++
	}

	for i := 1; i <= n; i++ {
		if i == 1 {
			if SuitNotHave[i] == -1 && SuitNotHave[i+1] == 1 {
				SuitNotHave[i]++
				SuitNotHave[i+1]--
			}
			continue
		}

		if i == n {
			if SuitNotHave[i] == -1 && SuitNotHave[i-1] == 1 {
				SuitNotHave[i]++
				SuitNotHave[i-1]--
			}
			break
		}

		if SuitNotHave[i] == -1 {
			if SuitNotHave[i-1] == 1 {
				SuitNotHave[i]++
				SuitNotHave[i-1]--
				continue
			}
			if SuitNotHave[i+1] == 1 {
				SuitNotHave[i]++
				SuitNotHave[i+1]--
			}
		}
	}

	return HaveSuit(SuitNotHave)
}

func HaveSuit(Suit []int) (result int) {

	for i := 1; i < len(Suit); i++ {
		if Suit[i] >= 0 {
			result++
		}
	}

	return result
}
