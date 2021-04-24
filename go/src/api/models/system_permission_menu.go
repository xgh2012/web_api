package models

import "demo/conf"

// 角色权限表
type SystemPermissionMenu struct {
	PermissionId int    `json:"permission_id"`
	MenuIds      string `json:"menu_ids"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (permissionMenu *SystemPermissionMenu) One() *SystemPermissionMenu {
	conf.Engine.Get(permissionMenu)
	return permissionMenu
}
