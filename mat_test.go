package mat

import (
	"testing"

	"github.com/twiggg/tester"
)

func TestIndex(t *testing.T) {
	te := tester.New(t)
	m := NewM64(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2 := NewM64(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2.Transpose()
	tests := []struct {
		m   *M64
		i   int
		j   int
		ind int
	}{
		{m, 1, 1, 4},
		{m2, 1, 1, 4},
		{m, 1, 2, 5},
		{m2, 1, 2, 7},
	}
	for ind, test := range tests {
		te.DeepEqual(ind, "ind", test.ind, test.m.index(test.i, test.j))
	}
}

//git commit -m "added no-touch transposer"

func TestAt(t *testing.T) {
	te := tester.New(t)
	m := NewM64(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2 := NewM64(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	m2.Transpose()
	tests := []struct {
		m   *M64
		i   int
		j   int
		val float64
	}{
		{m, 1, 1, 5},
		{m2, 1, 1, 5},
		{m, 1, 2, 6},
		{m2, 1, 2, 8},
	}
	for ind, test := range tests {
		te.DeepEqual(ind, "ind", test.val, test.m.At(test.i, test.j))
	}
}
