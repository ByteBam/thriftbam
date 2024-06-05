// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameModuleInfo = "module_info"

// ModuleInfo mapped from table <module_info>
type ModuleInfo struct {
	ID           string    `gorm:"column:id;type:char(19);primaryKey" json:"id"`
	ModuleName   string    `gorm:"column:module_name;type:varchar(50);not null;comment:模块名称" json:"module_name"`               // 模块名称
	BranchID     string    `gorm:"column:branch_id;type:char(19);not null;comment:分支外键" json:"branch_id"`                      // 分支外键
	InterfaceNum int32     `gorm:"column:interface_num;type:int;comment:接口数量" json:"interface_num"`                            // 接口数量
	IsDeleted    int32     `gorm:"column:is_deleted;type:tinyint;not null;comment:是否删除  0 false 1 true" json:"is_deleted"`     // 是否删除  0 false 1 true
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
}

// TableName ModuleInfo's table name
func (*ModuleInfo) TableName() string {
	return TableNameModuleInfo
}
