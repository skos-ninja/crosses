package board

// Position is the data format that the board accepts commands for.
type Position struct {
	X     int   `json:"x"`
	Y     int   `json:"y"`
	State State `json:"state"`
}
