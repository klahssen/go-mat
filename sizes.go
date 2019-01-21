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
	if m.c != n.r {
		return fmt.Errorf("m colomns != n rows")
	}
	if dest.r != m.r {
		return fmt.Errorf("m,dest rows not equal")
	}
	if dest.c != n.c {
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
	if m.r != n.r {
		return fmt.Errorf("m,n rows not equal")
	}
	if m.c != n.c {
		return fmt.Errorf("m,n colomns not equal")
	}
	if m.r != dest.r {
		return fmt.Errorf("m,dest rows not equal")
	}
	if m.c != dest.c {
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

	if m.r != dest.r {
		return fmt.Errorf("m,dest rows not equal")
	}
	if m.c != dest.c {
		return fmt.Errorf("m,dest colomns not equal")
	}
	return nil
}
