// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinCoinAdminTransferM2 = "win_coin_admin_transfer_m2s"

// WinCoinAdminTransferM2 后台调账记录
type WinCoinAdminTransferM2 struct {
	ID         int64           `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	AdminID    int64           `gorm:"column:admin_id;type:int;not null;comment:管理员ID" json:"adminId"`                                // 管理员ID
	Coin       decimal.Decimal `gorm:"column:coin;type:decimal(15,4);not null;default:0.0000;comment:调账金额" json:"coin"`               // 调账金额
	CoinBefore decimal.Decimal `gorm:"column:coin_before;type:decimal(15,4);not null;default:0.0000;comment:调账前金额" json:"coinBefore"` // 调账前金额
	UID        int64           `gorm:"column:uid;type:int;not null;comment:用户ID" json:"uid"`                                          // 用户ID
	Username   string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                         // 用户名
	Category   int64           `gorm:"column:category;type:tinyint;not null;comment:调账原因:0-其他 1-误存调账 2-活动调账" json:"category"`         // 调账原因:0-其他 1-误存调账 2-活动调账
	Mark       string          `gorm:"column:mark;type:varchar(255);not null;comment:调账原因" json:"mark"`                               // 调账原因
	CreatedAt  int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt  int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	FlowClaim  int64           `gorm:"column:flow_claim;type:int;not null;comment:流水倍数" json:"flowClaim"`   // 流水倍数
	Message    string          `gorm:"column:message;type:varchar(500);comment:通知客户信息" json:"message"`      // 通知客户信息
	MerchantID int64           `gorm:"column:merchant_id;type:int;not null;comment:商户id" json:"merchantId"` // 商户id
}

// TableName WinCoinAdminTransferM2's table name
func (*WinCoinAdminTransferM2) TableName() string {
	return TableNameWinCoinAdminTransferM2
}
