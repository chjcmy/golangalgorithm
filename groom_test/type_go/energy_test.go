package type_go

import (
	"fmt"
	"testing"
)

const g float32 = 9.8

type msh struct {
	m  float32
	v  float32
	h  float32
	ke float32
	pe float32
	me float32
}

func (e *msh) kenergy() float32 {
	return e.m * e.v * e.v / 2
}

func (e *msh) Penergy() float32 {
	return e.m * g * e.h
}

func TestEnergy(t *testing.T) {

	total := make([]msh, 5)

	total[0].m = 30
	total[0].v = 2
	total[0].h = 10
	total[1].m = 10
	total[1].v = 2
	total[1].h = 40
	total[2].m = 4
	total[2].v = 5
	total[2].h = 1
	total[3].m = 30
	total[3].v = 2
	total[3].h = 33
	total[4].m = 19
	total[4].v = 21
	total[4].h = 9

	for i := 0; i < len(total); i++ {
		total[i].ke = total[i].kenergy()
		total[i].pe = total[i].Penergy()
		total[i].me = total[i].ke + total[i].pe
		fmt.Println(total[i].ke, total[i].pe, total[i].me)
	}
}
