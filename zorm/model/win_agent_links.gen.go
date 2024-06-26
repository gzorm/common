// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinAgentLink = "win_agent_links"

// WinAgentLink 代理专属域名
type WinAgentLink struct {
	ID        int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Link      string `gorm:"column:link;type:varchar(256);not null;comment:推广域名" json:"link"`                  // 推广域名
	UID       int64  `gorm:"column:uid;type:int;not null;comment:代理ID" json:"uid"`                             // 代理ID
	Status    int64  `gorm:"column:status;type:tinyint;not null;default:1;comment:状态:1-启用 0-停用" json:"status"` // 状态:1-启用 0-停用
	CreatedAt int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinAgentLink's table name
func (*WinAgentLink) TableName() string {
	return TableNameWinAgentLink
}
