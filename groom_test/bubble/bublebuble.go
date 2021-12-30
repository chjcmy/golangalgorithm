package main

import "fmt"

type student struct {
	name  string
	sex   string
	score map[string]int
}

func newStudent() *student {
	d := student{}
	d.score = map[string]int{}
	return &d
}

func main() {
	var a, b, num int
	var name, sex, sub string

	fmt.Scanln(&a, &b)

	s := make([]student, a)

	for i := 0; i < a; i++ {
		object := newStudent()
		fmt.Scanln(&name, &sex)

		object.name = name
		object.sex = sex

		for j := 0; j < b; j++ {
			fmt.Scanln(&sub, &num)
			object.score[sub] = num
		}
		s[i] = *object

	}

	for i := 0; i < len(s); i++ {

		fmt.Println("----------")
		fmt.Printf("%s %s\n", s[i].name, s[i].sex)
		for index, val := range s[i].score {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
