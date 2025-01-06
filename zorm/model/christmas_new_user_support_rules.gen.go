// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameChristmasNewUserSupportRule = "christmas_new_user_support_rules"

// ChristmasNewUserSupportRule 圣诞活动新用户一次助力规则表
type ChristmasNewUserSupportRule struct {
	ID                            int64           `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:自增主键" json:"id,string"`                                                            // 自增主键
	RuleID                        int64           `gorm:"column:rule_id;type:bigint;not null;comment:规则ID" json:"ruleId"`                                                                                     // 规则ID
	ActivityID                    int64           `gorm:"column:activity_id;type:bigint;not null;comment:活动ID" json:"activityId"`                                                                             // 活动ID
	NewOneSupportSuccessRangeFrom decimal.Decimal `gorm:"column:new_one_support_success_range_from;type:decimal(15,4);not null;default:0.0000;comment:新用户一次助力成功充值金额范围起" json:"newOneSupportSuccessRangeFrom"` // 新用户一次助力成功充值金额范围起
	NewOneSupportSuccessRangeTo   decimal.Decimal `gorm:"column:new_one_support_success_range_to;type:decimal(15,4);not null;default:0.0000;comment:新用户一次助力成功充值金额范围止" json:"newOneSupportSuccessRangeTo"`     // 新用户一次助力成功充值金额范围止
	NewOneSupportSuccessPr        decimal.Decimal `gorm:"column:new_one_support_success_pr;type:decimal(15,4);not null;default:0.0000;comment:新用户一次助力成功比例" json:"newOneSupportSuccessPr"`                     // 新用户一次助力成功比例
	CreateAt                      int64           `gorm:"column:create_at;type:int;not null;comment:创建时间" json:"createAt"`                                                                                    // 创建时间
	UpdateAt                      int64           `gorm:"column:update_at;type:int;not null;comment:修改时间" json:"updateAt"`                                                                                    // 修改时间
	OpUser                        string          `gorm:"column:op_user;type:varchar(32);not null;default:system;comment:操作人" json:"opUser"`                                                                  // 操作人
}

// TableName ChristmasNewUserSupportRule's table name
func (*ChristmasNewUserSupportRule) TableName() string {
	return TableNameChristmasNewUserSupportRule
}
