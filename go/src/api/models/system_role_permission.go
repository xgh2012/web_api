package models

import "demo/conf"

// 角色权限表
type SystemRolePermission struct {
	RoleId        int    `json:"role_id"`
	PermissionIds string `json:"permission_ids"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func (rolePermission *SystemRolePermission) One() *SystemRolePermission {
	conf.Engine.Get(rolePermission)
	return rolePermission
}
