package orm

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// BaseEntity 对应数据库表基础字段，支持 if_deleted 软删除
type BaseEntity struct {
	Id         int64                 `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	Code       string                `gorm:"column:code;type:varchar(64);uniqueIndex:uk_code_del;not null" json:"code"`
	TenantCode string                `gorm:"column:tenant_code;type:varchar(12);uniqueIndex:uk_code_del;default:'000000'" json:"tenantCode"`
	CreateUser string                `gorm:"column:create_user;type:varchar(64)" json:"createUser"`
	CreateTime time.Time             `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateUser string                `gorm:"column:update_user;type:varchar(64)" json:"updateUser"`
	UpdateTime time.Time             `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	IfDeleted  soft_delete.DeletedAt `gorm:"column:if_deleted;type:smallint;default:0;uniqueIndex:uk_code_del;softDelete:flag,DeletedAtField:DeletedAt" json:"ifDeleted"`
}
