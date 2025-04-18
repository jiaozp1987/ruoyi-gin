package sysMenu

import "ruoyi-gin/src/ruoyi-common/core/domain/entity/baseEntity"

type SysMenu struct {
	baseEntity.BaseEntity
	//菜单ID
	MenuId int `json:"menuid"`
	//菜单名称
	MenuName string `json:"menuname"`
	//父菜单名称
	ParentName string `json:"parentname"`
	//父菜单ID
	ParentId int `json:"parentid"`
	//显示顺序
	OrderNum int `json:"ordernum"`
	//路由地址
	Path string `json:"path"`
	//组件路径
	Component string `json:"component"`
	//路由参数
	Query string `json:"query"`
	//路由名称，默认和路由地址相同的驼峰格式（注意：因为vue3版本的router会删除名称相同路由，为避免名字的冲突，特殊情况可以自定义）
	RouteName string `json:"routename"`
	//是否为外链（0是 1否）
	IsFrame string `json:"isframe"`
	//是否缓存（0缓存 1不缓存）
	IsCache string `json:"iscache"`
	//类型（M目录 C菜单 F按钮）
	MenuType string `json:"menutype"`
	//显示状态（0显示 1隐藏）
	Visible string `json:"visible"`
	//菜单状态（0正常 1停用）
	Status string `json:"status"`
	// 权限字符串
	Perms string `json:"perms"`
	//菜单图标
	Icon string `json:"icon"`
	//子菜单
	Children []SysMenu `json:"children"`
}
