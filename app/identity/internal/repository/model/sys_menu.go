package model

import (
	"zephyr-go/pkg/orm"
)

// SysMenu 菜单权限表
type SysMenu struct {
	orm.BaseEntity
	MenuName   string `gorm:"column:menu_name;type:varchar(50)"`
	ParentCode string `gorm:"column:parent_code;type:varchar(64)"`
	OrderNum   int    `gorm:"column:order_num;type:int;default:0"`
	Path       string `gorm:"column:path;type:varchar(200)"`
	Component  string `gorm:"column:component;type:varchar(255)"`
	IsFrame    int16  `gorm:"column:is_frame;type:smallint;default:0"`
	IsCache    int16  `gorm:"column:is_cache;type:smallint;default:0"`
	MenuType   string `gorm:"column:menu_type;type:char(1)"`
	Visible    string `gorm:"column:visible;type:char(1);default:'0'"`
	Status     int16  `gorm:"column:status;type:smallint;default:1"`
	Perms      string `gorm:"column:perms;type:varchar(100)"`
	Icon       string `gorm:"column:icon;type:varchar(100)"`
}

func (SysMenu) TableName() string {
	return "zephyr_sys_menu"
}
