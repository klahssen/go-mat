package mat

import (
	"fmt"
	"testing"

	"github.com/twiggg/tester"
)

func TestSmallAdd(t *testing.T) {
	te := tester.New(t)
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		res  *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), nil, nil, fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, nil, nil, fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m,n rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(2, 3, nil), NewM64(3, 3, nil), fmt.Errorf("m,n colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m,dest colomns not equal")},
		{NewM64(3, 2, []float64{1, 1, 1, 1, 1, 1}), NewM64(3, 2, []float64{2, 2, 2, 1, 1, 1}), NewM64(3, 2, []float64{3, 3, 3, 2, 2, 2}), NewM64(3, 2, []float64{3, 3, 3, 2, 2, 2}), nil},
	}

	for ind, test := range tests {
		err := add(test.m, test.n, test.dest)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, test.dest)
		}
	}
}

func TestSmallTranspose(t *testing.T) {
	te := tester.New(t)
	tests := []struct {
		m    *M64
		dest *M64
		res  *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), nil, fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(2, 3, nil), fmt.Errorf("m rows != dest colomns")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m colomns != dest rows")},
		{NewM64(3, 2, []float64{1, 2, 3, 4, 5, 6}), NewM64(2, 3, nil), NewM64(2, 3, []float64{1, 3, 5, 2, 4, 6}), nil},
	}

	for ind, test := range tests {
		err := transpose(test.m, test.dest)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, test.dest)
		}
	}
}

func TestSmallSub(t *testing.T) {
	te := tester.New(t)
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		res  *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), nil, nil, fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, nil, nil, fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m,n rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(2, 3, nil), NewM64(3, 3, nil), fmt.Errorf("m,n colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m,dest colomns not equal")},
		{NewM64(3, 2, []float64{1, 1, 1, 1, 1, 1}), NewM64(3, 2, []float64{2, 2, 2, 1, 1, 1}), NewM64(3, 2, []float64{-1, -1, -1, 0, 0, 0}), NewM64(3, 2, []float64{-1, -1, -1, 0, 0, 0}), nil},
	}

	for ind, test := range tests {
		err := sub(test.m, test.n, test.dest)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, test.dest)
		}
	}
}

func TestSmallMul(t *testing.T) {
	te := tester.New(t)
	m := NewM64(1, 3, []float64{1, 2, 3})
	m.Transpose()
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		res  *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), nil, nil, fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, nil, nil, fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), NewM64(2, 3, nil), fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), fmt.Errorf("n,dest colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), fmt.Errorf("m colomns != n rows")},
		{NewM64(3, 2, []float64{1, 1, 1, 1, 1, 1}), NewM64(2, 3, []float64{1, 1, 1, 1, 1, 1}), NewM64(3, 3, []float64{2, 2, 2, 2, 2, 2, 2, 2, 2}), NewM64(3, 3, []float64{2, 2, 2, 2, 2, 2, 2, 2, 2}), nil},
		{NewM64(3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}), NewM64(3, 1, []float64{1, 2, 3}), NewM64(3, 1, nil), NewM64(3, 1, []float64{6, 6, 6}), nil},
		{NewM64(3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}), m, NewM64(3, 1, nil), NewM64(3, 1, []float64{6, 6, 6}), nil},
	}

	for ind, test := range tests {
		err := mul(test.m, test.n, test.dest)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, test.dest)
		}
	}

	/*
		fmt.Printf("check with mat64.Dense\n")
		w := mat64.NewDense(3, 3, []float64{1, 1, 1, 1, 1, 1, 1, 1, 1})
		x := mat64.NewDense(3, 1, []float64{1, 2, 3})
		res := mat64.NewDense(3, 1, []float64{6, 6, 6})
		res.Mul(w, x)
		fmt.Printf("w:\n%+v\n\nx:\n%+v\n\nres:\n%+v\n", w, x, res)
	*/
}

func BenchmarkSmallMul500x500x1(b *testing.B) {
	r, c, c2 := 500, 500, 1
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = float64(i)*100 + 1
	}
	data2 := make([]float64, c*c2)
	for i := 0; i < c*c2; i++ {
		data2[i] = float64(i)*100 + 4
	}
	m := NewM64(r, c, data)
	n := NewM64(c, c2, data2)
	dest := NewM64(r, c2, nil)
	// run the function b.N times
	//t0 := time.Now()
	for k := 0; k < b.N; k++ {
		//t0 = time.Now()
		mul(m, n, dest)
		//fmt.Printf("%s\n", time.Since(t0))
	}
}

func TestSmallMapElem(t *testing.T) {
	f0 := func(valm, valn float64) float64 {
		return 0.0
	}
	f1 := func(valm, valn float64) float64 {
		return valm * valn
	}
	te := tester.New(t)
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		fn   func(valm, valn float64) float64
		res  *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), f0, NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), NewM64(3, 3, nil), f0, nil, fmt.Errorf("m is nil")},
		{NewM64(3, 3, nil), nil, NewM64(3, 3, nil), f0, nil, fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, f0, nil, fmt.Errorf("dest is nil")},
		{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(3, 2, nil), f0, nil, fmt.Errorf("m,n rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(3, 2, nil), f0, nil, fmt.Errorf("m,n colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(2, 3, nil), f0, nil, fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), f0, nil, fmt.Errorf("m,dest colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 2, nil), f0, NewM64(3, 2, []float64{}), nil},
		{NewM64(3, 2, []float64{1, 1, 1, 1, 1, 1}), NewM64(3, 2, []float64{1, 2, 3, 4, 5, 6}), NewM64(3, 2, nil), f1, NewM64(3, 2, []float64{1, 2, 3, 4, 5, 6}), nil},
	}

	for ind, test := range tests {
		err := mapElem(test.m, test.n, test.dest, test.fn)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, test.dest)
		}
	}
}

func TestSmallMapElemIJ(t *testing.T) {
	f0 := func(i, j int) float64 {
		return 0.0
	}
	f1 := func(i, j int) float64 {
		return float64(i * j)
	}
	te := tester.New(t)
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		fn   func(i, j int) float64
		res  *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), f0, NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), NewM64(3, 3, nil), f0, nil, fmt.Errorf("m is nil")},
		//{NewM64(3, 3, nil), nil, NewM64(3, 3, nil), f0, nil, fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, f0, nil, fmt.Errorf("dest is nil")},
		//{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(3, 2, nil), f0, nil, fmt.Errorf("m,n rows not equal")},
		//{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(3, 2, nil), f0, nil, fmt.Errorf("m,n colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(2, 3, nil), f0, nil, fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), f0, nil, fmt.Errorf("m,dest colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 2, nil), f0, NewM64(3, 2, []float64{0, 0, 0, 0, 0, 0}), nil},
		{NewM64(3, 2, []float64{1, 1, 1, 1, 1, 1}), NewM64(3, 2, []float64{1, 2, 3, 4, 5, 6}), NewM64(3, 2, nil), f1, NewM64(3, 2, []float64{0, 0, 0, 1, 0, 2}), nil},
		{NewM64(2, 3, []float64{1, 1, 1, 1, 1, 1}), NewM64(2, 3, []float64{1, 2, 3, 4, 5, 6}), NewM64(2, 3, nil), f1, NewM64(2, 3, []float64{0, 0, 0, 0, 1, 2}), nil},
	}

	for ind, test := range tests {
		err := mapElemIJ(test.m, test.dest, test.fn)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, test.dest)
		}
	}
}

func TestSmallMapElemVal(t *testing.T) {
	f0 := func(val float64) float64 {
		return 0.0
	}
	f1 := func(val float64) float64 {
		return val * 2.0
	}
	te := tester.New(t)
	tests := []struct {
		m    *M64
		n    *M64
		dest *M64
		fn   func(val float64) float64
		res  *M64
		err  error
	}{
		{NewM64(3, 3, nil), NewM64(3, 3, nil), NewM64(3, 3, nil), f0, NewM64(3, 3, nil), nil},
		{nil, NewM64(3, 3, nil), NewM64(3, 3, nil), f0, nil, fmt.Errorf("m is nil")},
		//{NewM64(3, 3, nil), nil, NewM64(3, 3, nil), f0, nil, fmt.Errorf("n is nil")},
		{NewM64(3, 3, nil), NewM64(3, 3, nil), nil, f0, nil, fmt.Errorf("dest is nil")},
		//{NewM64(3, 2, nil), NewM64(2, 3, nil), NewM64(3, 2, nil), f0, nil, fmt.Errorf("m,n rows not equal")},
		//{NewM64(3, 2, nil), NewM64(3, 3, nil), NewM64(3, 2, nil), f0, nil, fmt.Errorf("m,n colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(2, 3, nil), f0, nil, fmt.Errorf("m,dest rows not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 3, nil), f0, nil, fmt.Errorf("m,dest colomns not equal")},
		{NewM64(3, 2, nil), NewM64(3, 2, nil), NewM64(3, 2, nil), f0, NewM64(3, 2, []float64{0, 0, 0, 0, 0, 0}), nil},
		{NewM64(3, 2, []float64{1, 1, 1, 2, 2, 2}), NewM64(3, 2, []float64{1, 2, 3, 4, 5, 6}), NewM64(3, 2, nil), f1, NewM64(3, 2, []float64{2, 2, 2, 4, 4, 4}), nil},
	}

	for ind, test := range tests {
		err := mapElemVal(test.m, test.dest, test.fn)
		te.CompareError(ind, test.err, err)
		if err == nil {
			te.DeepEqual(ind, "res", test.res, test.dest)
		}
	}
}
