package mat

import (
	"testing"

	"github.com/twiggg/tester"
)

func TestIndex(t *testing.T) {
	te := tester.New(t)
	tests := []struct {
		m   *M64
		i   int
		j   int
		ind int
	}{
		{NewM64(3, 3, nil), 1, 1, 4},
	}
	for ind, test := range tests {
		te.DeepEqual(ind, "ind", test.ind, test.m.index(test.i, test.j))
	}
}
