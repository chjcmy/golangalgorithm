package main

import "fmt"

func inputSubNum() (num int, err error) {

	fmt.Scanln(&num)

	if num <= 0 {
		return 0, fmt.Errorf("잘못된 과목 수입니다.")
	}
	return
}

func average(num int) (avg float64, err error) {
	var score, total int

	for i := 0; i < num; i++ {
		fmt.Scanln(&score)

		if score < 0 || score > 100 {
			return 0, fmt.Errorf("잘못된 점수입니다.")
		}

		total += score

	}

	avg = float64(total) / float64(num)

	return
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	num, err := inputSubNum()

	if err != nil {
		panic(err)
		return
	}

	result, err := average(num)

	if err != nil {
		panic(err)
		return
	}

	fmt.Println(result)

}
