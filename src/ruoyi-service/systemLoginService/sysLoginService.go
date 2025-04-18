package systemLoginService

import (
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/strutil"
	"ruoyi-gin/src/ruoyi-common/core/securityUtils"

	iniUtils "ruoyi-gin/src/ruoyi-common/core/ini"
	"ruoyi-gin/src/ruoyi-common/core/jwt"
	lb "ruoyi-gin/src/ruoyi-domain/loginBody"
	lu "ruoyi-gin/src/ruoyi-domain/loginUser"
	"ruoyi-gin/src/ruoyi-domain/sysUser"
	"ruoyi-gin/src/ruoyi-mapper/loginBodyMapper"
)

func loginPreCheck(loginBody *lb.LoginBody) bool {
	if strutil.IsBlank(loginBody.UserName) || strutil.IsBlank(loginBody.Password) {
		return false
	}
	return true
}

func Login(loginBody *lb.LoginBody) string {
	if !loginPreCheck(loginBody) {
		return ""
	} else {
		loginBody.Password = cryptor.HmacSha256(loginBody.Password, iniUtils.CRYPTOR_KEY)
		if loginBodyMapper.LoginAuth(loginBody).UserName == loginBody.UserName {
			// 生产token
			token, _ := jwt.GenToken(loginBody)
			return token
		}
		return ""
	}
}

func GetInfo(loginUser *lu.LoginUser) (sysUser.SysUser, []string, []string) {
	su := loginUser.User
	roleIds := su.RoleIds
	roleNames := loginBodyMapper.GetRolsByIds(roleIds)
	permissions := make([]string, 1)
	if securityUtils.IsAdmin(roleIds) {
		permissions[0] = "*:*:*"
	} else {
		permissions = loginBodyMapper.GetPermissionsByRoleIds(roleIds)
	}

	return su, roleNames, permissions
}
