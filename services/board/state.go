package board

// State describes what state a given position on the board is in.
type State int

const (
	// Empty state means no turn has been taken for that position.
	Empty State = 0
	// X state means player 1 has taken a turn for that position.
	X State = 1
	// O state means player 2 has taken a turn for that position.
	O State = 2
)
