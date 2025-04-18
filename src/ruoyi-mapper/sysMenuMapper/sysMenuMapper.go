package sysMenuMapper

import (
	"ruoyi-gin/src/ruoyi-common/core/db"
	"ruoyi-gin/src/ruoyi-domain/sysMenu"
)

func SelectMenuTreeAll() []sysMenu.SysMenu {

	sql := "select distinct m.menu_id as MenuId ," +
		" m.parent_id as ParentId, " +
		"m.menu_name as MenuName, " +
		"m.path as Path," +
		" m.component as Component," +
		" m.`" + "query` as Query," +
		" m.route_name as RouteName, " +
		" m.visible as Visible," +
		" m.status as Status, " +
		"ifnull(m.perms,'') as Perms, " +
		"m.is_frame as IsFrame, " +
		"m.is_cache as IsCache," +
		" m.menu_type as MenuType," +
		" m.icon as Icon," +
		" m.order_num as OrderNum, " +
		"m.create_time as CreateTime " +
		"from sys_menu m where m.menu_type in ('M', 'C') and m.status = 0 order by m.parent_id, m.order_num"

	var list []sysMenu.SysMenu
	db.SqlDB.Raw(sql).Scan(&list)
	return list
}

func SelectMenuTreeByUserId(userId int) []sysMenu.SysMenu {
	sql := "select distinct m.menu_id as MenuId ," +
		" m.parent_id as ParentId, " +
		"m.menu_name as MenuName, " +
		"m.path as Path," +
		" m.component as Component," +
		" m.`" + "query` as Query," +
		" m.route_name as RouteName, " +
		" m.visible as Visible," +
		" m.status as Status, " +
		"ifnull(m.perms,'') as Perms, " +
		"m.is_frame as IsFrame, " +
		"m.is_cache as IsCache," +
		" m.menu_type as MenuType," +
		" m.icon as Icon," +
		" m.order_num as OrderNum, " +
		"m.create_time as CreateTime " +
		`
		from sys_menu m
			 left join sys_role_menu rm on m.menu_id = rm.menu_id
			 left join sys_user_role ur on rm.role_id = ur.role_id
			 left join sys_role ro on ur.role_id = ro.role_id
			 left join sys_user u on ur.user_id = u.user_id
		where u.user_id = ? and m.menu_type in ('M', 'C') and m.status = 0  AND ro.status = 0
		order by m.parent_id, m.order_num
        `
	var list []sysMenu.SysMenu
	db.SqlDB.Raw(sql, userId).Scan(&list)
	return list
}
