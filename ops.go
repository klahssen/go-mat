package mat

//Add returns a new matrix as m+n (element by element)
func Add(m, n *M64) (*M64, error) {
	r, c := m.Dims()
	res := NewM64(r, c, nil)
	if err := add(m, n, res); err != nil {
		return nil, err
	}
	return res, nil
}

//Sub returns a new matrix as m-n (element by element)
func Sub(m, n *M64) (*M64, error) {
	r, c := m.Dims()
	res := NewM64(r, c, nil)
	if err := sub(m, n, res); err != nil {
		return nil, err
	}
	return res, nil
}

//Mul returns a new matrix as the dot product of m and n
func Mul(m, n *M64) (*M64, error) {
	r, _ := m.Dims()
	_, c1 := n.Dims()
	res := NewM64(r, c1, nil)
	if err := mul(m, n, res); err != nil {
		return nil, err
	}
	return res, nil
}

//MulElem returns a new matrix as the dot product of m and n
func MulElem(m, n *M64) (*M64, error) {
	r, c := m.Dims()
	res := NewM64(r, c, nil)
	if err := mulElem(m, n, res); err != nil {
		return nil, err
	}
	return res, nil
}

//MapElem applies function fn to each elem of m
func MapElem(m *M64, fn func(x float64) float64) (*M64, error) {
	res := NewM64(m.r, m.c, nil)
	if err := mapElemVal(m, res, fn); err != nil {
		return nil, err
	}
	return res, nil
}

//Transpose returns a new matrix as the transpose of m
func Transpose(m *M64) (*M64, error) {
	r, c := m.Dims()
	res := NewM64(c, r, nil)
	if err := transpose(m, res); err != nil {
		return nil, err
	}
	return res, nil
}
