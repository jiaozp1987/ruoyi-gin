package loginUser

import (
	"github.com/golang-jwt/jwt/v4"
	"ruoyi-gin/src/ruoyi-domain/sysUser"
)

type LoginUser struct {
	DeptId     int `json:"deptId"`
	ExpireTime int `json:"expireTime"`
	//LoginTime     int            `json:"loginTime"`
	UserId      int      `json:"userId"`
	Permissions []string `json:"permissions"`
	//Browser       string         `json:"browser"`
	//Ipaddr        string         `json:"ipaddr"`
	//LoginLocation string         `json:"loginLocation"`
	//OS            string         `json:"os"`
	Token string          `json:"token"`
	User  sysUser.SysUser `json:"user"`
	jwt.RegisteredClaims
}
