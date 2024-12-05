package model

import "time"

type Message struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userid"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated-at"`
	IsDeleted int       `json:"is_deleted"`
	ParentID  *int      `json:"parent_id"`
}
