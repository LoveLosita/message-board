package model

import "time"

type User struct {
	ID        int       `json:"id"`
	NickName  string    `json:"nickname"`
	UserName  string    `json:"username"`
	PassWord  string    `json:"password"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
	Role      string    `json:"role"`
}

type JsonInquiry struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type NewUser struct {
	TargetID int    `json:"target_id"`
	NickName string `json:"nickname"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Role     string `json:"role"`
}

type DisplayUser struct { //去除密码的用户信息
	ID        int       `json:"id"`
	NickName  string    `json:"nickname"`
	UserName  string    `json:"username"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
	Role      string    `json:"role"`
}

type UpdatePwdUser struct {
	UserName    string `json:"username"`
	OldPassWord string `json:"old_pwd"`
	NewPassWord string `json:"new_pwd"`
}
