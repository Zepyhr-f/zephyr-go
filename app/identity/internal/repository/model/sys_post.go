package model

import (
	"zephyr-go/pkg/orm"
)

// SysPost 岗位信息表
type SysPost struct {
	orm.BaseEntity
	PostName string `gorm:"column:post_name;type:varchar(50)"`
	OrderNum int    `gorm:"column:order_num;type:int;default:0"`
	Status   int16  `gorm:"column:status;type:smallint;default:1"`
}

func (SysPost) TableName() string {
	return "zephyr_sys_post"
}
