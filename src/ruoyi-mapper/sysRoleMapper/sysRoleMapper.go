package sysRoleMapper

import (
	"ruoyi-gin/src/ruoyi-common/core/db"
	"ruoyi-gin/src/ruoyi-domain/sysRole"
)

func LoginRoleList(userName string) []sysRole.SysRole {
	sql := `
		select
			r.role_id as RoleId,
			r.role_name as RoleName,
			r.role_key as RoleKey,
			r.role_sort as RoleSort,
			r.data_scope as DataScope,
			r.menu_check_strictly as  MenuCheckStrictly,
			r.dept_check_strictly as DeptCheckStrictly,
			r.status as Status,
			r.del_flag as  DelFlag,
			r.create_by as CreateBy,
			r.update_by as UpdateBy,
			r.create_time as CreateTime,
			r.update_time as UpdateTime,
			r.remark as Remark
		from sys_user s,sys_user_role sr, sys_role r
		where user_name = ?
		  and s.user_id = sr.user_id
		  and sr.role_id = r.role_id
	`
	var roleSlice []sysRole.SysRole
	db.SqlDB.Raw(sql, userName).Scan(&roleSlice)
	return roleSlice

}
