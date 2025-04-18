package sysDictType

import "ruoyi-gin/src/ruoyi-common/core/domain/entity/baseEntity"

type SysDictType struct {
	baseEntity.BaseEntity
	DictId   int    `excel:"字典主键" json:"dictid"`
	DictName string `excel:"字典名称" json:"dictname"`
	DictType string `excel:"字典类型" json:"dicttype"`
	//状态（0正常 1停用）
	Status string `excel:"状态" json:"status"`
}
