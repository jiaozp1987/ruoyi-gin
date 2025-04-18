package sysRole

import "ruoyi-gin/src/ruoyi-common/core/domain/entity/baseEntity"

type SysRole struct {
	baseEntity.BaseEntity
	RoleId   int    `excel:"角色序号" json:"roleid"`
	RoleName string `excel:"角色名称" json:"rolename"`
	RoleKey  string `excel:"角色权限" json:"rolekey"`
	RoleSort string `excel:"角色排序" json:"rolesort"`
	//数据范围（1：所有数据权限；2：自定义数据权限；3：本部门数据权限；4：本部门及以下数据权限；5：仅本人数据权限）
	DataScope string `excel:"数据范围" json:"datascope"`
	//菜单树选择项是否关联显示（ 0：父子不互相关联显示 1：父子互相关联显示）
	MenuCheckStrictly bool `json:"menuCheckStrictly"`
	//部门树选择项是否关联显示（0：父子不互相关联显示 1：父子互相关联显示 ）
	DeptCheckStrictly bool `json:"deptCheckStrictly"`
	//角色状态（0正常 1停用）
	Status string `excel:"角色状态" json:"status"`
	//删除标志（0代表存在 2代表删除）
	DelFlag string `json:"delflag"`
	//用户是否存在此角色标识 默认不存在
	Flag bool `json:"flag"`
	//菜单组
	MenuIds []int `json:"menuids"`
	//部门组（数据权限）
	DeptIds []int `json:"deptids"`
	//角色菜单权限
	Permissions []string `json:"permissions"`
}
