package baseEntity

import "time"

type BaseEntity struct {
	//搜索值
	SearchValue string `json:"searchValue"`
	//创建者
	CreateBy string `json:"createBy"`
	//创建时间 JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")
	CreateTime time.Time `json:"createTime"`
	//更新者
	UpdateBy string `json:"updateBy"`
	//更新时间 JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")
	UpdateTime time.Time `json:"updateTime"`
	//备注
	Remark string `json:"remark"`
	//请求参数
	Params map[string]interface{} `json:"params"`
}
