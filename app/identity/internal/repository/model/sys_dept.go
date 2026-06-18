package model

import (
	"zephyr-go/pkg/orm"
)

// SysDept 部门信息表
type SysDept struct {
	orm.BaseEntity
	ParentCode string `gorm:"column:parent_code;type:varchar(64)"`
	DeptName   string `gorm:"column:dept_name;type:varchar(30)"`
	OrderNum   int    `gorm:"column:order_num;type:int;default:0"`
	Status     int16  `gorm:"column:status;type:smallint;default:1"`
	Leaf       int16  `gorm:"column:leaf;type:smallint;default:0"`
}

func (SysDept) TableName() string {
	return "zephyr_sys_dept"
}
