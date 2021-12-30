package type_go

import (
	"fmt"
	"testing"
)

type student struct {
	name  string
	sex   string
	score map[string]int
}

var studentss = [5]student{
	{name: "kim", sex: "F", score: map[string]int{"kor": 57, "math": 90, "eng": 57, "science": 57}},
	{name: "lee", sex: "F", score: map[string]int{"kor": 78, "math": 45, "eng": 80, "science": 100}},
	{name: "yun", sex: "M", score: map[string]int{"kor": 67, "math": 40, "eng": 83, "science": 97}},
	{name: "nam", sex: "M", score: map[string]int{"kor": 88, "math": 91, "eng": 80, "science": 31}},
	{name: "kim", sex: "M", score: map[string]int{"kor": 42, "math": 29, "eng": 100, "science": 100}}}

func TestType(t *testing.T) {

	students := make([]student, 9)

	students[0].score["sale"] = 6

	for i := 0; i < len(students); i++ {
		fmt.Println("----------")
		fmt.Printf("%s %s\n", students[i].name, students[i].sex)
		for index, val := range students[i].score {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
