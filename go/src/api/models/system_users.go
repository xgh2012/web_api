package models

import "demo/conf"

// 角色表
type SystemUsers struct {
	Id            int    `json:"id"`
	UserName      string `json:"user_name"`
	PassWord      string `json:"pass_word"`
	Name          string `json:"name"`
	RoleId        int    `json:"role_id"`
	Avatar        string `json:"avatar"`
	RememberToken string `json:"remember_token"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (user *SystemUsers) One() *SystemUsers {
	conf.Engine.Get(user)
	return user
}
