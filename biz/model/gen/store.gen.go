// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"time"
)

const TableNameStore = "store"

// Store mapped from table <store>
type Store struct {
	ID            string    `gorm:"column:id;type:char(19);primaryKey" json:"id"`
	Name          string    `gorm:"column:name;type:varchar(20);not null;comment:仓库名称" json:"name"`                                      // 仓库名称
	Level         int32     `gorm:"column:level;type:tinyint;not null;comment:仓库等级" json:"level"`                                        // 仓库等级
	PocID         int64     `gorm:"column:poc_id;type:bigint;not null;comment:仓库负责人Id" json:"poc_id"`                                    // 仓库负责人Id
	IsDeleted     int32     `gorm:"column:is_deleted;type:tinyint;not null;comment:是否删除  0 false 1 true" json:"is_deleted"`              // 是否删除  0 false 1 true
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;comment:修改时间" json:"update_time"`          // 修改时间
	CreateTime    time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	StoreLink     string    `gorm:"column:store_link;type:varchar(128);not null;comment:仓库地址" json:"store_link"`                         // 仓库地址
	StorePlatform string    `gorm:"column:store_platform;type:char(8);not null;comment:第三方平台" json:"store_platform"`                     // 第三方平台
	GitID         string    `gorm:"column:git_id;type:char(32);not null;uniqueIndex:git_id,priority:1;comment:第三方平台的仓库id" json:"git_id"` // 第三方平台的仓库id
	Isbind        int32     `gorm:"column:isbind;type:tinyint;not null;comment:是否绑定了webhook" json:"isbind"`                              // 是否绑定了webhook
}

// TableName Store's table name
func (*Store) TableName() string {
	return TableNameStore
}
