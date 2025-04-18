package sysUser

import (
	"ruoyi-gin/src/ruoyi-common/core/domain/entity/baseEntity"
	"ruoyi-gin/src/ruoyi-domain/sysDept"
	"ruoyi-gin/src/ruoyi-domain/sysRole"
)

type SysUser struct {
	baseEntity.BaseEntity
	UserId      int    `excel:"用户序号" json:"userid"`
	DeptId      int    `excel:"部门编号" json:"deptid"`
	UserName    string `excel:"登录名称" json:"userName"`
	NickName    string `excel:"用户名称" json:"nickName"`
	Email       string `excel:"用户邮箱" json:"email"`
	Phonenumber string `excel:"手机号码" json:"phoneNumber"`
	//"0=男,1=女,2=未知"
	Sex string `excel:"用户性别" json:"sex"`
	//用户头像
	Avatar string `json:"avatar"`
	//密码
	Password string `json:"password"`
	Status   string `excel:"帐号状态" json:"status"`
	//删除标志（0代表存在 2代表删除）`
	DelFlag string `json:"delFlag"`
	LoginIp string `excel:"最后登录IP" json:"loginIp"`
	//LoginDate time.Time `excel:"最后登录时间" json:"loginDate"`
	//部门对象
	Dept sysDept.SysDept `json:"dept"`
	//角色对象
	Role []sysRole.SysRole `json:"role"`
	//角色组
	RoleIds []int `json:"roleIds"`
	//岗位组
	PostIds []int `json:"postIds"`
	//角色ID
	RoleId int `json:"roleId"`
}
