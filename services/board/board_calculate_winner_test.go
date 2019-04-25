package board

import (
	"math"
	"testing"
)

func TestEnsureNoWinner(t *testing.T) {
	brd := NewBoard()
	index := 0

	// Fill every position on the board
	// The fill is done in an alternating fill format and as such no winner could be found
	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			pos := &Position{
				X:     x,
				Y:     y,
				State: O,
			}

			if index%2 == 0 {
				pos.State = X
			}

			brd.SetPosition(pos)
			index++
		}
	}

	middle := int(math.Ceil(BoardSize / 2))
	state := brd.CalculateWinnerFromPosition(middle, middle)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}
}

func TestEnsureRowWinner(t *testing.T) {
	brd := NewBoard()
	xPos := 0

	for yPos := 0; yPos < BoardSize; yPos++ {
		pos := &Position{
			X:     xPos,
			Y:     yPos,
			State: X,
		}

		brd.SetPosition(pos)
	}

	state := brd.isMatchingRow(xPos)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinnerFromPosition(xPos, 0)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinner()
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}
}

func TestEnsureNoRowWinner(t *testing.T) {
	brd := NewBoard()
	xPos := 0

	for yPos := 0; yPos < BoardSize; yPos++ {
		pos := &Position{
			X:     xPos,
			Y:     yPos,
			State: O,
		}

		if yPos%2 == 0 {
			pos.State = X
		}

		brd.SetPosition(pos)
	}

	state := brd.isMatchingRow(xPos)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinnerFromPosition(xPos, 0)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinner()
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}
}

func TestEnsureColumnWinner(t *testing.T) {
	brd := NewBoard()
	yPos := 0

	for xPos := 0; xPos < BoardSize; xPos++ {
		pos := &Position{
			X:     xPos,
			Y:     yPos,
			State: X,
		}

		brd.SetPosition(pos)
	}

	state := brd.isMatchingColumn(yPos)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinnerFromPosition(0, yPos)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinner()
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}
}

func TestEnsureNoColumnWinner(t *testing.T) {
	brd := NewBoard()
	yPos := 0

	for xPos := 0; xPos < BoardSize; xPos++ {
		pos := &Position{
			X:     xPos,
			Y:     yPos,
			State: O,
		}

		if xPos%2 == 0 {
			pos.State = X
		}

		brd.SetPosition(pos)
	}

	state := brd.isMatchingColumn(yPos)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinnerFromPosition(0, yPos)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinner()
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}
}

func TestEnsureLeftDiagonalWinner(t *testing.T) {
	brd := NewBoard()

	for axisPos := 0; axisPos < BoardSize; axisPos++ {
		pos := &Position{
			X:     axisPos,
			Y:     axisPos,
			State: X,
		}

		brd.SetPosition(pos)
	}

	state := brd.isMatchingDiagonally(0)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinnerFromPosition(0, 0)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinner()
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}
}

func TestEnsureNoLeftDiagonalWinner(t *testing.T) {
	brd := NewBoard()

	for axisPos := 0; axisPos < BoardSize; axisPos++ {
		pos := &Position{
			X:     axisPos,
			Y:     axisPos,
			State: O,
		}

		if axisPos%2 == 0 {
			pos.State = X
		}

		brd.SetPosition(pos)
	}

	state := brd.isMatchingDiagonally(0)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinnerFromPosition(0, 0)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinner()
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}
}

func TestEnsureRightDiagonalWinner(t *testing.T) {
	brd := NewBoard()
	startPos := BoardSize - 1
	xPos := 0

	for yPos := startPos; yPos >= 0; yPos-- {
		pos := &Position{
			X:     xPos,
			Y:     yPos,
			State: X,
		}

		brd.SetPosition(pos)
		xPos++
	}

	state := brd.isMatchingDiagonally(startPos)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinnerFromPosition(0, startPos)
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}

	state = brd.CalculateWinner()
	if state != X {
		t.Errorf("expected %d; got %d", X, state)
	}
}

func TestEnsureNoRightDiagonalWinner(t *testing.T) {
	brd := NewBoard()
	startPos := BoardSize - 1
	xPos := 0

	for yPos := startPos; yPos >= 0; yPos-- {
		pos := &Position{
			X:     xPos,
			Y:     yPos,
			State: O,
		}

		if yPos%2 == 0 {
			pos.State = X
		}

		brd.SetPosition(pos)
		xPos++
	}

	state := brd.isMatchingDiagonally(startPos)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinnerFromPosition(0, startPos)
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}

	state = brd.CalculateWinner()
	if state != Empty {
		t.Errorf("expected %d; got %d", Empty, state)
	}
}
