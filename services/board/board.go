package board

import (
	"errors"
	"math"
)

// BoardSize is the size of the board that will be created
const BoardSize = 3

var boardSliceMedian = int(math.Ceil(BoardSize/2)) - 1

// Board is a 3x3 array of states for the position.
type Board struct {
	positions [BoardSize][BoardSize]State
}

// NewBoard initialises a brand new empty board.
func NewBoard() Board {
	return Board{}
}

// GetPosition will return the current state of the select position
func (board *Board) GetPosition(x int, y int) State {
	return board.positions[x][y]
}

// GetState will return the entire boards current state
func (board *Board) GetState() []*Position {
	positions := make([]*Position, 0)

	for x, arr := range board.positions {
		for y, state := range arr {
			position := &Position{
				X:     x,
				Y:     y,
				State: state,
			}

			positions = append(positions, position)
		}
	}

	return positions
}

// CanTakeTurn describes if it's possible for a turn to be taken for the position.
func (board *Board) CanTakeTurn(x int, y int) bool {
	return board.positions[x][y] == Empty
}

// CalculateCurrentPlayersTurn will determine who's turn it currently is on a board.
func (board *Board) CalculateCurrentPlayersTurn() State {
	player1 := 0
	player2 := 0

	for _, pos := range board.positions {
		for _, state := range pos {
			switch state {
			case X:
				player1++
				break
			case O:
				player2++
				break
			}
		}
	}

	// We use a square of the board size to allow for dynamic board sizes
	if player1+player2 == int(math.Pow(float64(BoardSize), 2)) {
		return Empty
	}

	if player1 <= player2 {
		return X
	}

	return O
}

// CalculateWinnerFromPosition will determine if the player for that position has won
// We take the position of what is intended to be the last move made to allow calculating the winner faster
// Even though the size of the board is dynamic we currently require all squares to match
func (board *Board) CalculateWinnerFromPosition(xPos int, yPos int) State {
	state := board.GetPosition(xPos, yPos)

	if board.isMatchingRow(xPos) == state {
		return state
	}

	if board.isMatchingColumn(yPos) == state {
		return state
	}

	// If in a corner or middle of board then check diagonally.
	// If this game was to support win sizes of a dynamic ration then this check would need to be rewritten.
	if (xPos%2 == 0 && yPos%2 == 0) || (xPos == boardSliceMedian && yPos == boardSliceMedian) {
		// Check right to left.
		if board.isMatchingDiagonally(BoardSize-1) == state {
			return state
		}

		// Check left to right.
		if board.isMatchingDiagonally(0) == state {
			return state
		}
	}

	return Empty
}

// CalculateWinner will determine if a game has a winner
func (board *Board) CalculateWinner() State {
	for i := 0; i <= BoardSize-1; i++ {
		if state := board.CalculateWinnerFromPosition(i, boardSliceMedian); state != Empty {
			return state
		}
	}

	return Empty
}

// isMatchingRow calculates if the entire row matches state and returns state.empty if not.
func (board *Board) isMatchingRow(xPos int) State {
	firstState := board.GetPosition(xPos, 0)

	for i := 1; i < BoardSize; i++ {
		if board.GetPosition(xPos, i) != firstState {
			return Empty
		}
	}

	return firstState
}

// isMatchingColumn calculates if the entire column matches state and returns state.empty if not.
func (board *Board) isMatchingColumn(yPos int) State {
	firstState := board.GetPosition(0, yPos)

	for i := 1; i < BoardSize; i++ {
		if board.GetPosition(i, yPos) != firstState {
			return Empty
		}
	}

	return firstState
}

// isMatchingDiagonally calculates if diagonally they match state and returns state.empty if not.
func (board *Board) isMatchingDiagonally(xPos int) State {
	firstState := board.GetPosition(xPos, 0)
	boardIndex := BoardSize - 1

	startingRight := (boardIndex == xPos)
	if startingRight {
		xPos--
	} else {
		xPos++
	}

	for yPos := 1; yPos < boardIndex; yPos++ {
		if pos := board.GetPosition(xPos, yPos); pos != firstState {
			return Empty
		}

		if startingRight {
			xPos--
		} else {
			xPos++
		}
	}

	return firstState
}

// SetState will set the state of every position on the board.
// This should only be run on initial setup and should use SetPosition if you want to update an individual position.
func (board *Board) SetState(positions []*Position) error {
	for _, pos := range positions {
		err := board.SetPosition(pos)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetPosition will set the state of the individual position passed on the board.
// This should be used when you are after updating an individual position only.
func (board *Board) SetPosition(position *Position) error {
	if position.X >= BoardSize || position.Y >= BoardSize {
		return errors.New("position_out_of_range")
	}

	if board.positions[position.X][position.Y] == Empty {
		board.positions[position.X][position.Y] = position.State
		return nil
	}

	return errors.New("position_already_set")
}
