package board

import (
	"testing"
)

func TestEnsurePlayer1Turn(t *testing.T) {
	brd := NewBoard()

	turn := brd.CalculateCurrentPlayersTurn()
	if turn != X {
		t.Errorf("expected player turn to be 1; got %d", turn)
	}
}

func TestEnsurePlayer2Turn(t *testing.T) {
	brd := NewBoard()

	pos1 := &Position{
		X:     1,
		Y:     1,
		State: X,
	}
	err := brd.SetPosition(pos1)
	if err != nil {
		t.Error(err)
	}

	turn := brd.CalculateCurrentPlayersTurn()
	if turn != O {
		t.Errorf("expected player turn to be 2; got %d", turn)
	}
}

func TestEnsureCompleteBoardNoTurn(t *testing.T) {
	brd := NewBoard()
	index := 0

	// Fill every position on the board
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

	turn := brd.CalculateCurrentPlayersTurn()
	if turn != Empty {
		t.Errorf("expected 0; got %d", turn)
	}
}
