package mat

import (
	"fmt"
	"testing"

	"github.com/twiggg/tester"
)

func TestSameSize(t *testing.T) {
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), NewM64(3, 3, nil), fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, NewM64(3, 3, nil), fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(3, 2, nil), fmt.Errorf("m,n rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(3, 2, nil), fmt.Errorf("m,n colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(2, 3, nil), fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), fmt.Errorf("m,dest colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 2, nil), nil},
	}
	te := tester.New(t)
	for ind, test := range tests {
		err := sameSize(test.m, test.n, test.dest)
		te.CompareError(ind, test.err, err)
	}

}

func TestSameSize2(t *testing.T) {
	tests := []struct {
		m    *M64
		dest *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), fmt.Errorf("m,dest colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), nil},
	}

	te := tester.New(t)
	for ind, test := range tests {
		err := sameSize2(test.m, test.dest)
		te.CompareError(ind, test.err, err)
	}

}

func TestDotSize(t *testing.T) {
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), NewM64(3, 3, nil), fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, NewM64(3, 3, nil), fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(3, 2, nil), fmt.Errorf("n,dest colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(3, 2, nil), fmt.Errorf("m colomns != n rows")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(2, 2, nil), fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(3, 3, nil), nil},
	}

	te := tester.New(t)
	for ind, test := range tests {
		err := dotSize(test.m, test.n, test.dest)
		te.CompareError(ind, test.err, err)
	}

}

func testError(t *testing.T, testInd int, exp, err error) {
	if (exp != nil && err == nil) || (exp == nil && err != nil) || (err != nil && err.Error() != exp.Error()) {
		t.Errorf("test %d: expected %v received %v", testInd, exp, err)
	}
}
