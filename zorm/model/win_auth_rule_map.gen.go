// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinAuthRuleMap = "win_auth_rule_map"

// WinAuthRuleMap 权限映射表（使用路径对应权限ID）
type WinAuthRuleMap struct {
	ID         int64  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	APIPath    string `gorm:"column:api_path;type:varchar(255);comment:权限路径" json:"apiPath"`        // 权限路径
	AuthRuleID int64  `gorm:"column:auth_rule_id;type:int unsigned;comment:权限ID" json:"authRuleId"` // 权限ID
}

// TableName WinAuthRuleMap's table name
func (*WinAuthRuleMap) TableName() string {
	return TableNameWinAuthRuleMap
}