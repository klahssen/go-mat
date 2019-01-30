package mat

import "fmt"

//mat "gonum.org/v1/gonum/mat"

/*
func useless() *mat.Dense {
	return mat.NewDense()
}
*/

//M64 represents a float64 matrix with r rows and c colomns
type M64 struct {
	r    int
	c    int
	data []float64
}

//Dims returns the number of rows and colomns
func (m *M64) Dims() (int, int) {
	if m == nil {
		return 0, 0
	}
	return m.r, m.c
}

//Size returns the size of the data array: rows*colomns
func (m *M64) Size() int {
	if m == nil {
		return 0
	}
	return m.r * m.c
}

//NewM64 returns a new M64 instance, initialized with data if len==r*c
func NewM64(r, c int, data []float64) *M64 {
	if r <= 0 {
		r = 1
	}
	if c <= 0 {
		c = 1
	}
	m := &M64{r: r, c: c}
	if len(data) == r*c {
		m.data = data
	} else {
		m.data = make([]float64, int(r*c))
	}
	return m
}

//Valid returns false if m is nil, and initiates with empty data of size=r*c if invalid size
func (m *M64) Valid() bool {
	if m == nil {
		return false
	}
	if m.r <= 0 {
		m.r = 1
	}
	if m.c <= 0 {
		m.c = 1
	}
	s := m.r * m.c
	if len(m.data) != s {
		m.data = make([]float64, s)
	}
	return true
}

func (m *M64) index(i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if i == 0 {
		return j
	}
	return m.c*i + j
}

//At returns the value at position row=i,col=j. panics if m is nil or index out of range
func (m *M64) At(i, j int) float64 {
	return m.data[m.index(i, j)]
}

//Set sets val at position row=i,col=j. panics if m is nil or index out of range
func (m *M64) Set(i, j int, val float64) {
	m.data[m.index(i, j)] = val
}

//Add adds n to m (element by element)
func (m *M64) Add(n *M64) error {
	return add(m, n, m)
}

//Sub substracts n to m (element by element)
func (m *M64) Sub(n *M64) error {
	return sub(m, n, m)
}

//Mul return mxn (matrix product)
func (m *M64) Mul(n *M64) error {
	return mul(m, n, m)
}

//MulElem return mxn (matrix product)
func (m *M64) MulElem(n *M64) error {
	return mulElem(m, n, m)
}

//MapElem applies fn to each element of the matrix
func (m *M64) MapElem(fn func(x float64) float64) error {
	return mapElemVal(m, m, fn)
}

//AtInd returns the value at index ind. panics if m is nil or index out of range
func (m *M64) AtInd(ind int) float64 {
	return m.data[ind]
}

//GetData returns inner matrix
func (m *M64) GetData() []float64 {
	if m == nil {
		return nil
	}
	return m.data
}

//SetData sets inner matrix
func (m *M64) SetData(data []float64) error {
	if m == nil {
		return fmt.Errorf("matrix is nil")
	}
	exp := m.Size()
	if len(data) != exp {
		return fmt.Errorf("data must have %d values", exp)
	}
	m.data = data
	return nil
}
