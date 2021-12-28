package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	nums := inputNums()
	total, higher, min := calExam(nums...)
	printResult(total, higher, min)
}

func inputNums() []int {
	nums := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d\n", &nums[i])
	}

	return nums
}

func calExam(arr ...int) (total, higher, min int) {

	for i := 0; i < len(arr); i++ {
		total += arr[i]

		if arr[i] >= 90 {
			higher++
		}
		if arr[i] < 70 {
			min++
		}
	}
	return
}

func printResult(total, higher, min int) {
	var result bool = true

	if total < 400 {
		fmt.Printf("총점이 400점 미만입니다.\n")
		result = false
	}
	if higher < 2 {
		fmt.Printf("90점 이상 과목 수가 2개 미만입니다.\n")
		result = false
	}
	if min > 0 {
		fmt.Printf("70점 미만 과목이 있습니다.\n")
		result = false
	}

	if !result {
		fmt.Printf("아이패드를 살 수 없습니다.\n")
	} else {
		fmt.Printf("아이패드를 살 수 있습니다.\n")
	}
}
