package programmer

import "testing"

func TestNoNum(t *testing.T) {
	nums1 := []int{1, 2, 3, 4, 6, 7, 8, 0}
	nums2 := []int{5, 8, 4, 0, 6, 7, 9}

	println(NoNumsol(nums1))
	println(NoNumsol(nums2))

}

func NoNumsol(nums []int) (result int) {

	result = 45

	for j := 0; j < len(nums); j++ {
		result -= nums[j]
	}
	return
}
