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
	Likes     int       `json:"likes"`
}

type SearchParams struct { //定义一个结构体，用于接收前端传来的参数
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
}
