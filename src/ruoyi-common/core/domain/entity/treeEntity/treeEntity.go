package treeEntity

import "ruoyi-gin/src/ruoyi-common/core/domain/entity/baseEntity"

type TreeEntity struct {
	baseEntity.BaseEntity
	//父菜单名称
	ParentName string `json:"parentName"`
	//父菜单ID
	ParentId int `json:"parentId"`
	//显示顺序
	OrderNum int `json:"orderNum"`
	//祖级列表
	Ancestors string `json:"ancestors"`
	//子部门
	Children []interface{} `json:"children"`
}

func NewTreeEntity() *TreeEntity {
	return &TreeEntity{}
}
