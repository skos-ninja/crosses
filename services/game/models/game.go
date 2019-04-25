package models

import (
	"time"
)

type Game struct {
	ID        string `json:"id"`
	Player1ID string `json:"player1_id"`
	Player2ID string `json:"player2_id"`

	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`

	WinningPlayer string `json:"winning_player"`
}
