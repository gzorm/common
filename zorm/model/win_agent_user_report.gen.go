// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameWinAgentUserReport = "win_agent_user_report"

// WinAgentUserReport 代理专属域名
type WinAgentUserReport struct {
	ID                   int64           `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id,string"`
	UID                  int64           `gorm:"column:uid;type:int;not null;comment:代理ID" json:"uid"` // 代理ID
	Username             string          `gorm:"column:username;type:varchar(255)" json:"username"`
	SupUid1              int64           `gorm:"column:sup_uid_1;type:int;not null;comment:上1级代理" json:"supUid1"`                                                     // 上1级代理
	SupUsername1         string          `gorm:"column:sup_username_1;type:varchar(32);not null;comment:上1级代理" json:"supUsername1"`                                   // 上1级代理
	Coin                 decimal.Decimal `gorm:"column:coin;type:decimal(15,4) unsigned;default:0.0000;comment:账户余额" json:"coin"`                                     // 账户余额
	Status               int64           `gorm:"column:status;type:int;comment:账号状态" json:"status"`                                                                   // 账号状态
	LevelID              int64           `gorm:"column:level_id;type:int;comment:等级序号" json:"levelId"`                                                                // 等级序号
	LevelText            string          `gorm:"column:level_text;type:varchar(255);comment:等级名称" json:"levelText"`                                                   // 等级名称
	TeamNums             int64           `gorm:"column:team_nums;type:int unsigned;comment:团队人数" json:"teamNums"`                                                     // 团队人数
	DirectNums           int64           `gorm:"column:direct_nums;type:int unsigned;comment:直属人数" json:"directNums"`                                                 // 直属人数
	CoinCommission       decimal.Decimal `gorm:"column:coin_commission;type:decimal(15,4) unsigned;default:0.0000;comment:累计返佣" json:"coinCommission"`                // 累计返佣
	CoinBalance          decimal.Decimal `gorm:"column:coin_balance;type:decimal(15,4) unsigned;default:0.0000;comment:团队总余额" json:"coinBalance"`                     // 团队总余额
	CoinDeposit          decimal.Decimal `gorm:"column:coin_deposit;type:decimal(15,4) unsigned;default:0.0000;comment:团队总充值" json:"coinDeposit"`                     // 团队总充值
	FirstDepositSuccCnt  int64           `gorm:"column:first_deposit_succ_cnt;type:int unsigned;comment:团队首存人数" json:"firstDepositSuccCnt"`                           // 团队首存人数
	FirstDepositSuccAmt  decimal.Decimal `gorm:"column:first_deposit_succ_amt;type:decimal(15,4) unsigned;default:0.0000;comment:团队首存金额" json:"firstDepositSuccAmt"`  // 团队首存金额
	DepositSuccCnt       int64           `gorm:"column:deposit_succ_cnt;type:int unsigned;comment:团队总充值次数" json:"depositSuccCnt"`                                     // 团队总充值次数
	DepositSuccUserCnt   int64           `gorm:"column:deposit_succ_user_cnt;type:int unsigned;comment:团队总充值人数" json:"depositSuccUserCnt"`                            // 团队总充值人数
	CoinWithdrawal       decimal.Decimal `gorm:"column:coin_withdrawal;type:decimal(15,4) unsigned;default:0.0000;comment:团队总提现" json:"coinWithdrawal"`               // 团队总提现
	WithdrawSuccCnt      int64           `gorm:"column:withdraw_succ_cnt;type:int unsigned;comment:团队总提现次数" json:"withdrawSuccCnt"`                                   // 团队总提现次数
	WithdrawSuccUserCnt  int64           `gorm:"column:withdraw_succ_user_cnt;type:int unsigned;comment:团队总提现人数" json:"withdrawSuccUserCnt"`                          // 团队总提现人数
	CoinWithdrawalReally decimal.Decimal `gorm:"column:coin_withdrawal_really;type:decimal(15,4) unsigned;default:0.0000;comment:团队实际到账" json:"coinWithdrawalReally"` // 团队实际到账
	CoinAdjust           decimal.Decimal `gorm:"column:coin_adjust;type:decimal(15,4) unsigned;default:0.0000;comment:团队总调账" json:"coinAdjust"`                       // 团队总调账
	CoinBet              decimal.Decimal `gorm:"column:coin_bet;type:decimal(15,4) unsigned;default:0.0000;comment:团队总投注" json:"coinBet"`                             // 团队总投注
	CoinBetBonus         decimal.Decimal `gorm:"column:coin_bet_bonus;type:decimal(15,4);default:0.0000;comment:团队总派彩" json:"coinBetBonus"`                           // 团队总派彩
	CoinProfit           decimal.Decimal `gorm:"column:coin_profit;type:decimal(15,4);default:0.0000;comment:团队总投注盈亏" json:"coinProfit"`                              // 团队总投注盈亏
	OperatorName         string          `gorm:"column:operator_name;type:varchar(255);comment:操作人" json:"operatorName"`                                              // 操作人
	RegisterAt           int64           `gorm:"column:register_at;type:int;comment:注册时间" json:"registerAt"`                                                          // 注册时间
	CreatedAt            int64           `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt            int64           `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinAgentUserReport's table name
func (*WinAgentUserReport) TableName() string {
	return TableNameWinAgentUserReport
}
