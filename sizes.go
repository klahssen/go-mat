package mat

import "fmt"

func dotSize(m, n, dest *M64) error {
	if !m.Valid() {
		return fmt.Errorf("m is nil")
	}
	if !n.Valid() {
		return fmt.Errorf("n is nil")
	}
	if !dest.Valid() {
		return fmt.Errorf("dest is nil")
	}
	mr, mc := m.Dims()
	nr, nc := n.Dims()
	dr, dc := dest.Dims()
	if mc != nr {
		return fmt.Errorf("m colomns != n rows")
	}
	if dr != mr {
		return fmt.Errorf("m,dest rows not equal")
	}
	if dc != nc {
		return fmt.Errorf("n,dest colomns not equal")
	}
	return nil
}

func sameSize(m, n, dest *M64) error {
	if !m.Valid() {
		return fmt.Errorf("m is nil")
	}
	if !n.Valid() {
		return fmt.Errorf("n is nil")
	}
	if !dest.Valid() {
		return fmt.Errorf("dest is nil")
	}
	mr, mc := m.Dims()
	nr, nc := n.Dims()
	dr, dc := dest.Dims()
	if mr != nr {
		return fmt.Errorf("m,n rows not equal")
	}
	if mc != nc {
		return fmt.Errorf("m,n colomns not equal")
	}
	if mr != dr {
		return fmt.Errorf("m,dest rows not equal")
	}
	if mc != dc {
		return fmt.Errorf("m,dest colomns not equal")
	}
	return nil
}

func sameSize2(m, dest *M64) error {
	if !m.Valid() {
		return fmt.Errorf("m is nil")
	}
	if !dest.Valid() {
		return fmt.Errorf("dest is nil")
	}
	mr, mc := m.Dims()
	dr, dc := dest.Dims()
	if mr != dr {
		return fmt.Errorf("m,dest rows not equal")
	}
	if mc != dc {
		return fmt.Errorf("m,dest colomns not equal")
	}
	return nil
}

func inverseSize(m, dest *M64) error {
	if !m.Valid() {
		return fmt.Errorf("m is nil")
	}
	if !dest.Valid() {
		return fmt.Errorf("dest is nil")
	}
	mr, mc := m.Dims()
	dr, dc := dest.Dims()

	if mr != dc {
		return fmt.Errorf("m rows != dest colomns")
	}
	if mc != dr {
		return fmt.Errorf("m colomns != dest rows")
	}
	return nil
}
