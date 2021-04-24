package models

import "demo/conf"

// 角色权限表
type SystemUserPermission struct {
	UserId        int    `json:"user_id"`
	PermissionIds string `json:"permission_ids"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (userPermission *SystemUserPermission) One() *SystemUserPermission {
	conf.Engine.Get(userPermission)
	return userPermission
}
