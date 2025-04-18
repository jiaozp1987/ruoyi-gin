package sysDeptMapper

import (
	"ruoyi-gin/src/ruoyi-common/core/db"
	"ruoyi-gin/src/ruoyi-domain/sysDept"
)

func LoginDept(deptId int) *sysDept.SysDept {
	sql := `select b.*,s.dept_name as ParentName  from(
				select
					d.dept_id as DeptId,
					d.parent_id as ParentId,
					d.ancestors as Ancestors,
					d.dept_name as DeptName,
					d.order_num as OrderNum,
					d.leader as Leader,
					d.phone as Phone,
					d.email as Email,
					d.status as Status,
					d.del_flag as DelFlag
				from sys_dept d
				where d.dept_id = ?
			) b, sys_dept s
			where b.ParentId = s.dept_id
	`
	tmp := &sysDept.SysDept{}
	db.SqlDB.Raw(sql, deptId).Scan(&tmp)
	return tmp
}
