package math

import "testing"

type testpair struct {
	values []float64
	result float64
}

var averageTests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"Expected", pair.result,
				"got", v,
			)
		}
	}
}

var minTests = []testpair{
	{[]float64{1, 2}, 1},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, -1},
	{[]float64{-5, -8}, -8},
	{[]float64{}, 0},
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"Expected", pair.result,
				"got", v,
			)
		}
	}
}

var maxTests = []testpair{
	{[]float64{1, 2}, 2},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 1},
	{[]float64{-5, -8}, -5},
	{[]float64{}, 0},
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"Expected", pair.result,
				"got", v,
			)
		}
	}
}
