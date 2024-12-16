package model

import "time"

type Message struct { //留言的结构体
	ID        int       `json:"id"`
	UserID    int       `json:"userid"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated-at"`
	IsDeleted int       `json:"is_deleted"`
	ParentID  *int      `json:"parent_id"`
	Likes     int       `json:"likes"`
	Replies   []Message `json:"replies"` //数据库中没有这个字段，但是在前端展示的时候需要，所以在这里定义
}

type SearchParams struct { //定义一个结构体，用于接收前端传来搜索留言的参数
	CommentID int    `json:"comment_id"`
	Content   string `json:"content"`
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
}

type Reply struct { //定义一个结构体，用于接收前端传来的回复留言的参数
	ParentID int    `json:"parent_id"`
	Message  string `json:"message"`
}

type ReplyMessage struct { //回复留言的结构体
	ID        int       `json:"id"`
	UserID    int       `json:"userid"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated-at"`
	IsDeleted int       `json:"is_deleted"`
	ParentID  int       `json:"parent_id"`
	Likes     int       `json:"likes"`
}
