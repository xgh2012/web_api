package adapts

import (
	"demo/api/models"
	"demo/utils"
	"github.com/spf13/cast"
	"strings"
)

type Permissions struct {
	Uid  int
	User *models.SystemUsers
	Menu []int
}

//用户权限 三种授权方式
//用户->角色->权限->菜单
//用户->权限->菜单
//用户->菜单
func (permission *Permissions) Get() []int {
	//用户->角色
	user := &models.SystemUsers{
		UserName: "admin",
	}
	user.One()
	if user.Id == 0 {
		return permission.Menu
	}
	permission.User = user
	permission.One()
	permission.Two()
	permission.Three()

	return permission.Menu
}

//用户->角色->权限->菜单
func (permission *Permissions) One() {
	//角色->权限
	rolePermission := &models.SystemRolePermission{
		RoleId: permission.User.RoleId,
	}
	rolePermission.One()
	if rolePermission.PermissionIds == "" {
		return
	}

	//权限->菜单
	permissionIds := strings.Split(rolePermission.PermissionIds, ",")
	for _, permissionId := range permissionIds {
		menu := &models.SystemPermissionMenu{
			PermissionId: cast.ToInt(permissionId),
		}
		menu.One()
		if menu.MenuIds == "" {
			continue
		}
		menus := strings.Split(menu.MenuIds, ",")
		for _, menuId := range menus {
			if utils.InArray(cast.ToInt(menuId), permission.Menu) == false {
				permission.Menu = append(permission.Menu, cast.ToInt(menuId))
			}
		}
	}
}

//用户->权限->菜单
func (permission *Permissions) Two() {
	//用户->权限
	userPermission := &models.SystemUserPermission{
		UserId: permission.User.Id,
	}
	userPermission.One()
	if userPermission.PermissionIds == "" {
		return
	}

	//权限->菜单
	permissionIds := strings.Split(userPermission.PermissionIds, ",")
	for _, permissionId := range permissionIds {
		menu := &models.SystemPermissionMenu{
			PermissionId: cast.ToInt(permissionId),
		}
		menu.One()
		if menu.MenuIds == "" {
			continue
		}
		menus := strings.Split(menu.MenuIds, ",")
		for _, menuId := range menus {
			if utils.InArray(cast.ToInt(menuId), permission.Menu) == false {
				permission.Menu = append(permission.Menu, cast.ToInt(menuId))
			}
		}
	}
}

//用户->菜单
func (permission *Permissions) Three() {
	//用户->菜单
	userMenu := &models.SystemUserMenu{
		UserId: permission.User.Id,
	}
	userMenu.One()
	if userMenu.MenuIds == "" {
		return
	}
	for _, menuId := range userMenu.MenuIds {
		if utils.InArray(cast.ToInt(menuId), permission.Menu) == false {
			permission.Menu = append(permission.Menu, cast.ToInt(menuId))
		}
	}
}
