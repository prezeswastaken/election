package models

import "time"

type Contestant struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" gorm:"unique"`
	Score     float32   `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
