package systemMenuService

import (
	"ruoyi-gin/src/ruoyi-domain/sysMenu"
	"ruoyi-gin/src/ruoyi-domain/sysUser"
	"ruoyi-gin/src/ruoyi-mapper/sysMenuMapper"
)

func SelectMenuTreeByUserId(su *sysUser.SysUser, permissions []string) []sysMenu.SysMenu {
	var menus []sysMenu.SysMenu
	if len(permissions) == 1 && permissions[0] == "*:*:*" {
		// admin
		menus = sysMenuMapper.SelectMenuTreeAll()
	} else {
		menus = sysMenuMapper.SelectMenuTreeByUserId(su.UserId)
	}
	return fullMenu(menus, 0)
}

func fullMenu(list []sysMenu.SysMenu, p int) []sysMenu.SysMenu {
	r := make([]sysMenu.SysMenu, 0)

	for i, _ := range list {
		if list[i].ParentId == p {
			list[i].Children = fullMenu(list, list[i].MenuId)
			r = append(r, list[i])
		}
	}

	return r
}
