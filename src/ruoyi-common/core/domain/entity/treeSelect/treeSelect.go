package treeSelect

import (
	"ruoyi-gin/src/ruoyi-domain/sysDept"
	"ruoyi-gin/src/ruoyi-domain/sysMenu"
)

type TreeSelect struct {
	Id       int          `json:"id"`
	Label    string       `json:"label"`
	Children []TreeSelect `json:"children"`
}

func NewTreeSelect() TreeSelect {
	return TreeSelect{}
}

func NewTreeSelectSysDept(dept *sysDept.SysDept) TreeSelect {
	treeSelect := NewTreeSelect()
	treeSelect.Id = dept.DeptId
	treeSelect.Label = dept.DeptName
	if len(dept.Children) > 0 {
		for _, v := range dept.Children {
			treeSelect.Children = append(treeSelect.Children, NewTreeSelectSysDept(&v))
		}
	}
	return treeSelect
}

func NewTreeSelectSysMenu(menu *sysMenu.SysMenu) TreeSelect {
	treeSelect := NewTreeSelect()
	treeSelect.Id = menu.MenuId
	treeSelect.Label = menu.MenuName
	if len(menu.Children) > 0 {
		for _, v := range menu.Children {
			treeSelect.Children = append(treeSelect.Children, NewTreeSelectSysMenu(&v))
		}
	}

	return treeSelect
}
