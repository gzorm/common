// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinMonthlyTax = "win_monthly_tax"

// WinMonthlyTax 每月税收统计
type WinMonthlyTax struct {
	ID         int64           `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:ID" json:"id,string"`   // ID
	GameID     int64           `gorm:"column:game_id;type:int;not null;comment:游戏ID" json:"gameId"`                    // 游戏ID
	GameName   string          `gorm:"column:game_name;type:varchar(50);not null;comment:游戏名称" json:"gameName"`        // 游戏名称
	Year       int64           `gorm:"column:year;type:int;not null;comment:统计年" json:"year"`                          // 统计年
	Month      int64           `gorm:"column:month;type:int;not null;comment:统计月" json:"month"`                        // 统计月
	CoinBet    decimal.Decimal `gorm:"column:coin_bet;type:decimal(18,4);not null;comment:投注金额" json:"coinBet"`        // 投注金额
	CoinProfit decimal.Decimal `gorm:"column:coin_profit;type:decimal(18,4);not null;comment:负盈利金额" json:"coinProfit"` // 负盈利金额
	Rate       decimal.Decimal `gorm:"column:rate;type:decimal(11,4);not null;comment:税收比例" json:"rate"`               // 税收比例
	CoinTax    decimal.Decimal `gorm:"column:coin_tax;type:decimal(18,4);not null;comment:税收金额" json:"coinTax"`        // 税收金额
	CreatedAt  int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                // 创建时间
	UpdatedAt  int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                                // 修改时间
}

// TableName WinMonthlyTax's table name
func (*WinMonthlyTax) TableName() string {
	return TableNameWinMonthlyTax
}