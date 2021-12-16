package models

import "time"

type Todo struct {
	Index     uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
