package sysDictData

import "ruoyi-gin/src/ruoyi-common/core/domain/entity/baseEntity"

type SysDictData struct {
	baseEntity.BaseEntity
	DictCode  int    `excel:"字典表面" json:"dictcode"`
	DictSort  int    `excel:"字典排序" json:"dictsort"`
	DictLabel string `excel:"字典标签" json:"dictlabel"`
	DictValue string `excel:"字典键值" json:"dictvalue"`
	DictType  string `excel:"字典类型" json:"dicttype"`
	//样式属性（其他样式扩展）
	CssClass string `json:"cssclass"`
	//表格字典样式
	ListClass string `json:"listclass"`
	//Y=是,N=否
	IsDefault string `excel:"是否默认" json:"isDefault"`
	//0=正常,1=停用
	Status string `excel:"状态" json:"status"`
}
