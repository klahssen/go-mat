package mat

import (
	"fmt"
	"testing"

	"github.com/twiggg/tester"
)

func TestDotProd(t *testing.T) {
	te := tester.New(t)
	tests := []struct {
		m   *M64
		n   *M64
		res *M64
		err error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), nil, fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, nil, fmt.Errorf("n is nil")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), nil, fmt.Errorf("m colomns != n rows")},
		{NewM64(3, 2, []float64{1, 1, 1, 1, 1, 1}), NewM64(2, 3, []float64{1, 1, 1, 1, 1, 1}), NewM64(3, 3, []float64{2, 2, 2, 2, 2, 2, 2, 2, 2}), nil},
		{NewM64(3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}), NewM64(3, 1, []float64{1, 2, 3}), NewM64(3, 1, []float64{6, 6, 6}), nil},
	}

	for ind, test := range tests {
		res, err := Mul(test.m, test.n)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, res)
		}
	}
}
