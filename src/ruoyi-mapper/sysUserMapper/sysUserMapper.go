package systemMapper

import (
	"ruoyi-gin/src/ruoyi-common/core/db"
	"ruoyi-gin/src/ruoyi-domain/sysUser"
)

func LoginRoleMenuIds(roleid int) []int {
	sql := `
		select
			sm.menu_id
		from sys_role_menu sm, sys_role r
		where sm.role_id = ? and r.del_flag = '0' and  sm.role_id = r.role_id
	`
	var list []int
	db.SqlDB.Raw(sql, roleid).Scan(&list)
	return list
}

func LoginRoleDeptIds(roleid int) []int {
	sql := `
		select sd.dept_id
		from sys_role_dept sd, sys_role r
		where sd.role_id = ? and sd.role_id = r.role_id and r.del_flag = '0'
	`
	var list []int
	db.SqlDB.Raw(sql, roleid).Scan(&list)
	return list
}

func LoginRolePermissions(roleid int) []string {
	sql := `
		select m.perms
		from sys_menu m where m.perms is not null and  m.menu_id in
		(
			select sm.menu_id
			from sys_role_menu sm, sys_role r
			where sm.role_id = ? and r.del_flag = '0' and  sm.role_id = r.role_id
		)
	`
	var list []string
	db.SqlDB.Raw(sql, roleid).Scan(&list)
	return list
}

func LoginDeptChildrenIds(deptId int) []int {
	sql := "select d.dept_id from sys_dept d where d.parent_id =?"
	var list []int
	db.SqlDB.Raw(sql, deptId).Scan(&list)
	return list
}

func LoginSysUser(userName string) *sysUser.SysUser {
	sql := `
		select
			u.user_id as UserId,
			u.dept_id as DeptId,
			u.user_name as UserName,
			u.nick_name as NickName,
			u.email as Email,
			u.phonenumber as Phonenumber,
			u.sex as Sex,
			u.avatar as Avatar,
			u.password as Password,
			u.status as Status,
			u.del_flag as DelFlag,
			u.login_ip as LoginIp
		from sys_user u where user_name = ? 
    `
	sysuser := &sysUser.SysUser{}
	db.SqlDB.Raw(sql, userName).Scan(&sysuser)
	return sysuser

}

func LoginSysUserPostIds(userId int) []int {
	sql := `
		select up.post_id from sys_user_post up where  up.user_id=?
	`
	var list []int
	db.SqlDB.Raw(sql, userId).Scan(&list)
	return list
}
