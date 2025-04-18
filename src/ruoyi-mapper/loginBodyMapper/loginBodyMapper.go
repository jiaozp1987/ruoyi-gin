package loginBodyMapper

import (
	"ruoyi-gin/src/ruoyi-common/core/db"
	"ruoyi-gin/src/ruoyi-domain/loginBody"
)

func LoginAuth(lb *loginBody.LoginBody) *loginBody.LoginBody {
	tmploginBody := &loginBody.LoginBody{}
	db.SqlDB.Raw("select user_name as username from sys_user_go where user_name = ? and password = ?", lb.UserName, lb.Password).Scan(&tmploginBody)
	return tmploginBody
}

func GetRolsByIds(ids []int) []string {
	var roles []string
	db.SqlDB.Raw("select r.role_key from sys_role r where r.del_flag = '0' and  r.role_id in ?", ids).Scan(&roles)
	return roles
}

func GetPermissionsByRoleIds(ids []int) []string {
	var permissions []string
	db.SqlDB.Raw(`
		select m.perms
		from sys_role_menu rm,sys_menu m, sys_role r
		where rm.menu_id = m.menu_id
		and rm.role_id in ?
		and r.role_id =rm.role_id
		and r.del_flag = '0' and  length(trim(m.perms))>0
	`, ids).Scan(&permissions)
	return permissions
}
