package _Stesting

import (
	"testing"
)

func TestTest1(t *testing.T) {
	println(test1Sol([]int{20, 8, 10, 1, 4, 15}))
}

var used []bool
var nums []int
var result int

func test1Sol(v []int) int {

	nums = v

	used = make([]bool, len(nums))

	retry(0, []int{})

	return result
}

func retry(index int, now []int) {
	if index == len(nums) {

		sum := 0

		for i := 0; i < len(nums)-1; i++ {

			x := now[i] - now[i+1]

			if x < 0 {

				x *= -1

			}

			sum += x

		}

		if result < sum {

			result = sum

		}
	}

	for i := 0; i < len(nums); i++ {

		if used[i] == false {

			used[i] = true

			now = append(now, nums[i])

			retry(index+1, now)

			used[i] = false

			now = now[:len(now)-1]

		}
	}
}
