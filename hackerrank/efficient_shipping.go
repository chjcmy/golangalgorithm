package main

import "fmt"

func main() {
	boxes := []int64{3, 1, 6}
	unitsPerBox := []int64{2, 7, 4}
	truckSize := int64(6)
	truefalse := make([]bool, len(boxes))
	fmt.Println(getMaxUnits(boxes, unitsPerBox, truckSize, truefalse))

}

func getMaxUnits(boxes []int64, unitsPerBox []int64, truckSize int64, bol []bool) (value int64) {
	// Write your code here

	for i := 0; i < len(boxes); i++ {

	}

	return
}

func bigarr(arr []int64) (val, wh int) {
	for i, i2 := range arr {
		for j := i + 1; j < len(arr); j++ {
			if i2 > arr[j] {
				val = int(i2)
				wh = i
			} else {
				val = int(arr[i+1])
				wh = i + 1
			}
		}
	}
	return
}
