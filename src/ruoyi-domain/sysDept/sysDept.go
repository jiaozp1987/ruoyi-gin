package sysDept

import (
	"ruoyi-gin/src/ruoyi-common/core/domain/entity/baseEntity"
)

type SysDept struct {
	baseEntity.BaseEntity
	//部门ID
	DeptId int `json:"deptid"`
	//父部门ID
	ParentId int `json:"parentid"`
	//祖级列表
	Ancestors string `json:"ancestors"`
	//部门名称
	DeptName string `json:"deptname"`
	//显示顺序
	OrderNum int `json:"ordernum"`
	//负责人
	Leader string `json:"leader"`
	//联系电话
	Phone string `json:"phone"`
	//邮箱
	Email string `json:"email"`
	//部门状态:0正常,1停用
	Status string `json:"status"`
	//删除标志（0代表存在 2代表删除）
	DelFlag string `json:"delflag"`
	//父部门名称
	ParentName string `json:"parentname"`
	//子部门
	Children []SysDept `json:"children"`
}
