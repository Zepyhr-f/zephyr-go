package model

import (
	"zephyr-go/pkg/orm"
)

// SysUser 用户信息表
type SysUser struct {
	orm.BaseEntity
	NickName  string `gorm:"column:nick_name;type:varchar(30)"`
	RealName  string `gorm:"column:real_name;type:varchar(30)"`
	Password  string `gorm:"column:password;type:varchar(100)"`
	Avatar    string `gorm:"column:avatar;type:varchar(100)"`
	Email     string `gorm:"column:email;type:varchar(50)"`
	Phone     string `gorm:"column:phone;type:varchar(11)"`
	Sex       int16  `gorm:"column:sex;type:smallint;default:0"`
	Birthday  string `gorm:"column:birthday;type:varchar(10)"`
	UserType  int16  `gorm:"column:user_type;type:smallint;default:0"`
	Status    int16  `gorm:"column:status;type:smallint;default:1"`
	DeptCode  string `gorm:"column:dept_code;type:varchar(64)"`
	PostCode  string `gorm:"column:post_code;type:varchar(64)"`

	// 关联
	Roles []SysRole `gorm:"many2many:zephyr_sys_user_role;foreignKey:Code;joinForeignKey:user_code;references:Code;joinReferences:role_code"`
	Dept  SysDept   `gorm:"foreignKey:DeptCode;references:Code"`
}

func (SysUser) TableName() string {
	return "zephyr_sys_user"
}
