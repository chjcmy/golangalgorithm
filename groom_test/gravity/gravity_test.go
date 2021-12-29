package gravity

import (
	"testing"
)

const gravity float32 = 9.8

type cal func(float32, float32) float32

func calMechEnergy(f cal, a, b float32) float32 {
	result := f(a, b)
	return result
}

func Test(t *testing.T) {
	result1, result2, result3 := gravity1(10.5, 50.22, 2.5)
	if result1 != 13240.754 && result2 != 257.25 && result3 != 13498.004 {
		t.Error("Wrong result")
	}
}

func gravity1(m, v, h float32) (float32, float32, float32) {

	kinEnergy := func(m float32, v float32) float32 {
		result := (m * v * v) / 2
		return result
	}
	potEnergy := func(m float32, h float32) float32 {
		result := m * gravity * h
		return result
	}

	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)

	return ke, pe, ke + pe
}
