package model

import "time"

type Message struct {
	ID        int       `json:"ID"`
	UserID    int       `json:"UserID"`
	Content   string    `json:"Content"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	IsDeleted int       `json:"IsDeleted"`
	ParentID  int       `json:"ParentID"`
}
