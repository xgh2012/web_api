package adapts

import ()

// 初始化结构体
type SystemInit struct {
	HomeInfo struct {
		Title string `json:"title"`
		Href  string `json:"href"`
	} `json:"homeInfo"`
	LogoInfo struct {
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"logoInfo"`
	MenuInfo []*models.MenuTreeList `json:"menuInfo"`
}

// 获取初始化数据
func (m *SystemInit) GetSystemInit() SystemInit {
	var systemInit SystemInit
	// 首页
	systemInit.HomeInfo.Title = "首页"
	systemInit.HomeInfo.Href = "page/welcome-1.html?t=1"

	// logo
	systemInit.LogoInfo.Title = "LAYUI MINI"
	systemInit.LogoInfo.Image = "images/logo.png"

	menus := &Permissions{
		Uid: 1,
	}
	// 菜单
	systemInit.MenuInfo = new(models.SystemMenu).GetMenuList(menus.Get())

	return systemInit
}
