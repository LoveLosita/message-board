package model

import "time"

type User struct {
	ID        int       `json:"ID"`
	NickName  string    `json:"NickName"`
	UserName  string    `json:"UserName"`
	PassWord  string    `json:"PassWord"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Role      string    `json:"Role"`
}
