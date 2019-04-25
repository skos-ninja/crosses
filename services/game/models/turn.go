package models

import (
	"time"
)

type Turn struct {
	ID       string `json:"id"`
	GameID   string `json:"game_id"`
	PlayerID string `json:"player_id"`

	XAxis int `json:"x_axis"`
	YAxis int `json:"y_axis"`

	CreatedAt *time.Time `json:"created_at"`
}
