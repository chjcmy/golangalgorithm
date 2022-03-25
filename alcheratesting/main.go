package main

import (
	"fmt"
	"sync"
)

// 완성하지 못함
// 1. 두개의 쓰레드 프로그램을 만든 후 첫번째에서 들어온 값을 받는다.
// 2. 두번째에서 첫번째 에서 받은 채널 값을 받는다.
// 3. 두번째 쓰레드 함수에 bool 배열을 만든다.
// 4. truefalse 라는 함수를 만들어서 출력 값들이 나오게끔 계산 후 출력한다
// 5. 마지막 값이 배열의 마지막 값이 0 이 아닐시 main 함수를 다시 실행한다
// 6. 첫번째 쓰레드 함수에 0 이 들어 올시 if 문으로 메인 함수를 끝낸다

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	nameChannel := make(chan int)

	go func() {

		N := 1

		fmt.Scanf("%d\n", &N)

		nameChannel <- N

		close(nameChannel)
	}()

	go func() {

		var Nums []int

		for name := range nameChannel {
			Nums = append(Nums, name)
		}

		for i := 0; i < len(Nums); i++ {
			truefalse()
			fmt.Printf("%d", Nums[i])
		}

		if Nums[len(Nums)-1] != 0 {
		}

		wg.Done()
	}()

	wg.Wait()

}

func truefalse() {

}
