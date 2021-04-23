package models

import "demo/conf"

// 用户菜单表
type SystemUserMenu struct {
	UserId    int    `json:"user_id"`
	MenuIds   string `json:"menu_ids"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (userMenu *SystemUserMenu) One() *SystemUserMenu {
	conf.Engine.Get(userMenu)
	return userMenu
}
