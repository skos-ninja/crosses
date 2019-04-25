package board

import (
	"testing"
)

func TestEnsureCanSetPosition(t *testing.T) {
	brd := NewBoard()
	x := 0
	y := 0

	pos := &Position{
		X:     x,
		Y:     y,
		State: X,
	}

	err := brd.SetPosition(pos)
	if err != nil {
		t.Error(err)
	}

	if brd.GetPosition(x, y) != X {
		t.Errorf("position failed to set (%d,%d) expected 1; got %d", x, y, brd.GetPosition(x, y))
	}
}

func TestEnsureCannotSetPosition(t *testing.T) {
	brd := NewBoard()
	x := 0
	y := 0

	pos := &Position{
		X:     x,
		Y:     y,
		State: X,
	}

	// This is unsafe and as such isn't exported outside of this board service
	brd.positions[x][y] = X

	err := brd.SetPosition(pos)
	if err != nil {
		if err.Error() != "position_already_set" {
			t.Error(err)
		} else {
			return
		}
	}

	t.Errorf("board position set successfully when already set")
}

func TestEnsureCannotSetOutOfBoundsPosition(t *testing.T) {
	brd := NewBoard()

	pos := &Position{
		X:     BoardSize,
		Y:     BoardSize,
		State: X,
	}

	err := brd.SetPosition(pos)
	if err != nil {
		if err.Error() != "position_out_of_range" {
			t.Error(err)
		} else {
			return
		}
	}

	t.Errorf("board position set successfully when out of bounds")
}

func TestEnsureCanSetState(t *testing.T) {
	brd := NewBoard()

	pos1 := &Position{
		X:     0,
		Y:     0,
		State: X,
	}

	pos2 := &Position{
		X:     1,
		Y:     1,
		State: O,
	}

	err := brd.SetState([]*Position{pos1, pos2})
	if err != nil {
		t.Error(err)
	}

	if brd.positions[pos1.X][pos1.Y] != pos1.State {
		t.Errorf("failed to set position for %d,%d; expected %d; got %d", pos1.X, pos1.Y, pos1.State, brd.positions[pos1.X][pos1.Y])
	}

	if brd.positions[pos2.X][pos2.Y] != pos2.State {
		t.Errorf("failed to set position for %d,%d; expected %d; got %d", pos2.X, pos2.Y, pos2.State, brd.positions[pos2.X][pos2.Y])
	}
}

func TestEnsureCannotSetState(t *testing.T) {
	brd := NewBoard()

	pos1 := &Position{
		X:     0,
		Y:     0,
		State: X,
	}

	pos2 := &Position{
		X:     1,
		Y:     1,
		State: O,
	}

	// Unsafe set the state of the position
	brd.positions[pos2.X][pos2.Y] = pos2.State

	err := brd.SetState([]*Position{pos1, pos2})
	if err == nil {
		t.Errorf("board state set successfully")
	} else if err.Error() != "position_already_set" {
		t.Error(err)
	}
}

func TestEnsureCanGetPosition(t *testing.T) {
	brd := NewBoard()
	x := 0
	y := 0

	// Unsafe set board position
	brd.positions[x][y] = X

	state := brd.GetPosition(x, y)
	if state != X {
		t.Errorf("expected 1; got %d", state)
	}
}

func TestEnsureCanTakeTurn(t *testing.T) {
	brd := NewBoard()

	canTake := brd.CanTakeTurn(0, 0)
	if !canTake {
		t.Errorf("expected true; got false")
	}
}

func TestEnsureCannotTakeTurn(t *testing.T) {
	brd := NewBoard()
	x := 0
	y := 0

	// Unsafe set board position
	brd.positions[x][y] = X

	canTake := brd.CanTakeTurn(x, y)
	if canTake {
		t.Errorf("expected false; got true")
	}
}
