package domainFactory

import (
	"ruoyi-gin/src/ruoyi-domain/loginBody"
	"ruoyi-gin/src/ruoyi-domain/loginUser"
	"ruoyi-gin/src/ruoyi-domain/registerBody"
	"ruoyi-gin/src/ruoyi-domain/sysDept"
	"ruoyi-gin/src/ruoyi-domain/sysDictData"
	"ruoyi-gin/src/ruoyi-domain/sysDictType"
	"ruoyi-gin/src/ruoyi-domain/sysMenu"
	"ruoyi-gin/src/ruoyi-domain/sysRole"
	"ruoyi-gin/src/ruoyi-domain/sysUser"
	"ruoyi-gin/src/ruoyi-mapper/sysDeptMapper"
	"ruoyi-gin/src/ruoyi-mapper/sysRoleMapper"
	systemMapper "ruoyi-gin/src/ruoyi-mapper/sysUserMapper"
)

func NewLoginBody() *loginBody.LoginBody {
	return &loginBody.LoginBody{}
}

func NewLoginUser(args ...interface{}) *loginUser.LoginUser {
	lu := &loginUser.LoginUser{}
	if len(args) == 2 {
		for i, v := range args {
			switch i {
			case 0:
				if user, ok := v.(sysUser.SysUser); ok {
					lu.User = user
				}
			case 1:
				lu.Permissions = v.([]string)
			}
		}
	}
	if len(args) == 4 {
		for i, v := range args {
			switch i {
			case 0:
				lu.UserId = v.(int)
			case 1:
				lu.DeptId = v.(int)
			case 2:
				if user, ok := v.(sysUser.SysUser); ok {
					lu.User = user
				}
			case 3:
				lu.Permissions = v.([]string)
			}
		}
	}
	return lu
}

func NewRegisterBody() *registerBody.RegisterBody {
	return &registerBody.RegisterBody{}
}

func NewSysDept(deptId int) sysDept.SysDept {
	sd := sysDeptMapper.LoginDept(deptId)
	chlidrenIds := systemMapper.LoginDeptChildrenIds(deptId)

	for _, v := range chlidrenIds {
		sd.Children = append(sd.Children, NewSysDept(v))
	}
	return *sd
}

func NewSysDictData() *sysDictData.SysDictData {
	return &sysDictData.SysDictData{}
}

func NewSysDictType() *sysDictType.SysDictType {
	return &sysDictType.SysDictType{}
}
func NewSysMenu() *sysMenu.SysMenu {
	return &sysMenu.SysMenu{}
}

func NewSysRole() *sysRole.SysRole {
	return &sysRole.SysRole{}
}

func NewSysRoleWithId(id int) *sysRole.SysRole {
	sysrole := NewSysRole()
	sysrole.RoleId = id
	return sysrole
}

func NewSysUser() *sysUser.SysUser {
	return &sysUser.SysUser{}
}

func NewSysUserWithId(userName string) *sysUser.SysUser {

	sysuser := systemMapper.LoginSysUser(userName)
	roleSlice := sysRoleMapper.LoginRoleList(userName)
	for i, _ := range roleSlice {
		sysuser.RoleIds = append(sysuser.RoleIds, roleSlice[i].RoleId)
		menuIds := systemMapper.LoginRoleMenuIds(roleSlice[i].RoleId)
		deptIds := systemMapper.LoginRoleDeptIds(roleSlice[i].RoleId)
		permissions := systemMapper.LoginRolePermissions(roleSlice[i].RoleId)
		roleSlice[i].Permissions = permissions
		roleSlice[i].MenuIds = menuIds
		roleSlice[i].DeptIds = deptIds
	}
	sysuser.Role = roleSlice
	sysuser.Dept = NewSysDept(sysuser.DeptId)
	postIds := systemMapper.LoginSysUserPostIds(sysuser.UserId)
	sysuser.PostIds = postIds
	return sysuser
}
