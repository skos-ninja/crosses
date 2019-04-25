package board

import (
	"testing"
)

func TestEnsureEmptyCreate(t *testing.T) {
	brd := NewBoard()
	if len(brd.positions) != BoardSize {
		t.Errorf("board x size incorrect; want %d; got %d", BoardSize, len(brd.positions))
	}

	for x, v := range brd.positions {
		if len(v) != BoardSize {
			t.Errorf("board y size incorrect; want %d; got %d", BoardSize, len(v))
		}

		for y, v := range v {
			if v != Empty {
				t.Errorf("board position state not empty; got %d at %d,%d", v, x, y)
			}
		}
	}
}
