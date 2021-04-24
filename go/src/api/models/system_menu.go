package models

import (
	"demo/conf"
	"demo/utils"
	"time"
)

// 菜单
type SystemMenu struct {
	Id       int       `json:"id"`
	Pid      int       `json:"pid"`
	Title    string    `json:"title"`
	Icon     string    `json:"icon"`
	Href     string    `json:"href"`
	Sort     string    `json:"sort"`
	Target   string    `json:"target"`
	Remark   string    `json:"remark"`
	State    int       `json:"state"`
	CreateAt time.Time `json:"create_at"`
}

// 菜单结构体
type MenuTreeList struct {
	Id     int             `json:"id"`
	Pid    int             `json:"pid"`
	Title  string          `json:"title"`
	Icon   string          `json:"icon"`
	Href   string          `json:"href"`
	Target string          `json:"target"`
	Remark string          `json:"remark"`
	Child  []*MenuTreeList `json:"child"`
}

// 获取菜单列表
func (m *SystemMenu) GetMenuList(menus []int) []*MenuTreeList {
	var menuList []SystemMenu
	_ = conf.Engine.Where("state=1").OrderBy("sort asc").Find(&menuList)

	return m.buildMenuChild(0, menuList, menus)
}

func (m *SystemMenu) GetAll() []SystemMenu {
	var menuList []SystemMenu
	_ = conf.Engine.OrderBy("sort asc").Find(&menuList)
	return menuList
}

//递归获取子菜单
func (m *SystemMenu) buildMenuChild(pid int, menuList []SystemMenu, menus []int) []*MenuTreeList {
	var treeList []*MenuTreeList
	for _, v := range menuList {
		if pid == v.Pid {
			node := &MenuTreeList{
				Id:     v.Id,
				Title:  v.Title,
				Icon:   v.Icon,
				Href:   v.Href,
				Target: v.Target,
				Pid:    v.Pid,
			}
			child := v.buildMenuChild(v.Id, menuList, menus)
			if len(child) != 0 {
				node.Child = child
			}
			// 权限判断
			if utils.InArray(v.Id, menus) == true {
				treeList = append(treeList, node)
			}
		}
	}
	return treeList
}
