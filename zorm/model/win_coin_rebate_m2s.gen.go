// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinCoinRebateM2 = "win_coin_rebate_m2s"

// WinCoinRebateM2 返水
type WinCoinRebateM2 struct {
	ID         int64           `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	UID        int64           `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                            // UID
	Username   string          `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`          // 用户名
	LevelID    int64           `gorm:"column:level_id;type:int;not null;comment:会员等级" json:"levelId"`                  // 会员等级
	SettleDate string          `gorm:"column:settle_date;type:varchar(20);not null;comment:结算日期" json:"settleDate"`    // 结算日期
	VaildStake decimal.Decimal `gorm:"column:vaild_stake;type:decimal(15,4);not null;comment:有效投注额" json:"vaildStake"` // 有效投注额
	GroupID    int64           `gorm:"column:group_id;type:int;not null;comment:游戏类型" json:"groupId"`                  // 游戏类型
	RebateRate decimal.Decimal `gorm:"column:rebate_rate;type:decimal(15,4);not null;comment:返水比例" json:"rebateRate"`  // 返水比例
	RabateCoin decimal.Decimal `gorm:"column:rabate_coin;type:decimal(15,4);not null;comment:返水金额" json:"rabateCoin"`  // 返水金额
	Status     int64           `gorm:"column:status;type:int;not null;default:1;comment:状态:1-已发放 0-未发放" json:"status"` // 状态:1-已发放 0-未发放
	CreatedAt  int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt  int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinCoinRebateM2's table name
func (*WinCoinRebateM2) TableName() string {
	return TableNameWinCoinRebateM2
}
