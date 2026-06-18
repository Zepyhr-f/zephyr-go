package model

import (
	"zephyr-go/pkg/orm"
)

// SysRole 角色信息表
type SysRole struct {
	orm.BaseEntity
	RoleName string `gorm:"column:role_name;type:varchar(30)"`
	OrderNum int    `gorm:"column:order_num;type:int;default:0"`
	Status   int16  `gorm:"column:status;type:smallint;default:1"`

	// 关联
	Menus []SysMenu `gorm:"many2many:zephyr_sys_role_menu;foreignKey:Code;joinForeignKey:role_code;references:Code;joinReferences:menu_code"`
}

func (SysRole) TableName() string {
	return "zephyr_sys_role"
}
