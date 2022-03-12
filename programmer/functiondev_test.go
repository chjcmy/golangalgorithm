package programmer

import "testing"

func TestFunc(t *testing.T) {
	funcSolution([]int{93, 30, 55}, []int{1, 30, 5})
}

type Queue []int

func (q Queue) Add(data int) Queue {
	return append(q, data)
}

func (q Queue) Remove() (Queue, int) {
	return q[1:], q[0]
}

func (q Queue) Peek() int {
	return q[0]
}

func funcSolution(progresses Queue, speeds Queue) (result []int) {

	for len(progresses) != 0 {
		progress(progresses, speeds)

		var completed int
		for _, p := range progresses {
			if p >= 100 {
				completed++
				progresses, _ = progresses.Remove()
				speeds, _ = speeds.Remove()
			} else {
				break
			}
		}
		if completed > 0 {
			result = append(result, completed)
		}
	}

	return
}

func progress(progresses Queue, speeds Queue) {
	for i, _ := range progresses {
		progresses[i] += speeds[i]
	}
}
