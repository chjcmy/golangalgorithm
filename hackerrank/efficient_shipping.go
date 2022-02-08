package main

import "fmt"

func remove(slice []int64, s int) []int64 {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	boxes := []int64{3, 1, 6}
	unitsPerBox := []int64{2, 7, 4}
	truckSize := int64(6)
	fmt.Println(getMaxUnits(boxes, unitsPerBox, truckSize))

}

func getMaxUnits(boxes []int64, unitsPerBox []int64, truckSize int64) (value int64) {
	// Write your code here

	for truckSize != 0 {
		bs, bc := bigarr(unitsPerBox)
		if truckSize > boxes[bc] {
			truckSize -= boxes[bc]
			value += boxes[bc] * int64(bs)
			boxes = remove(boxes, bc)
			unitsPerBox = remove(unitsPerBox, bc)
		} else {
			value += truckSize * int64(bs)
			truckSize = 0
		}
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
