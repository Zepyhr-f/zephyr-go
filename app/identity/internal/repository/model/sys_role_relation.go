package model

// SysRoleMenu 角色和菜单关联表
// 中间表不继承 BaseEntity，只保存业务编码关系。
type SysRoleMenu struct {
	RoleCode string `gorm:"column:role_code;type:varchar(64);primaryKey"`
	MenuCode string `gorm:"column:menu_code;type:varchar(64);primaryKey"`
}

func (SysRoleMenu) TableName() string {
	return "zephyr_sys_role_menu"
}

// SysRoleDept 角色和部门关联表，用于数据权限部门范围。
type SysRoleDept struct {
	RoleCode string `gorm:"column:role_code;type:varchar(64);primaryKey"`
	DeptCode string `gorm:"column:dept_code;type:varchar(64);primaryKey"`
}

func (SysRoleDept) TableName() string {
	return "zephyr_sys_role_dept"
}
