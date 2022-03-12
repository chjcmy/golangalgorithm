package programmer

import (
	"math"
	"testing"
)

func TestSosu(t *testing.T) {
	println(sosu([]int{1, 2, 3, 4}))
}

func sosu(nums []int) (result int) {

	count := 2

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {

				come := nums[i] + nums[j] + nums[k]

				if isPrime(come) {
					result++
				}
				count++
			}
		}
	}

	return
}

func isPrime(num int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(num)); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
