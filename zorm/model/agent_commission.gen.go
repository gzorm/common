// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/shopspring/decimal"

const TableNameAgentCommission = "agent_commission"

// AgentCommission 代理商佣金发放
type AgentCommission struct {
	ID                int64           `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键ID" json:"id,string"`                  // 主键ID
	AgentID           int64           `gorm:"column:agent_id;type:bigint;comment:代理商ID" json:"agentId"`                                           // 代理商ID
	AgentName         string          `gorm:"column:agent_name;type:varchar(255);comment:代理商名称" json:"agentName"`                                 // 代理商名称
	AgentLevelID      int64           `gorm:"column:agent_level_id;type:int;comment:代理商等级序号" json:"agentLevelId"`                                 // 代理商等级序号
	AgentLevelName    string          `gorm:"column:agent_level_name;type:varchar(50);comment:代理商等级名称" json:"agentLevelName"`                     // 代理商等级名称
	CommissionAmount  decimal.Decimal `gorm:"column:commission_amount;type:decimal(18,2);default:0.00;comment:当月佣金额度" json:"commissionAmount"`    // 当月佣金额度
	CommissionMonth   string          `gorm:"column:commission_month;type:varchar(7);comment:月份" json:"commissionMonth"`                          // 月份
	DistributedAmount decimal.Decimal `gorm:"column:distributed_amount;type:decimal(18,2);default:0.00;comment:当前已发放额度" json:"distributedAmount"` // 当前已发放额度
	BalanceAmount     decimal.Decimal `gorm:"column:balance_amount;type:decimal(18,4);default:0.0000;comment:可提现金额" json:"balanceAmount"`         // 可提现金额
	Status            int64           `gorm:"column:status;type:int;comment:状态(未发放=2，已发放=1，不可发放=3)" json:"status"`                                // 状态(未发放=2，已发放=1，不可发放=3)
	LastOperator      string          `gorm:"column:last_operator;type:varchar(100);comment:最后操作人" json:"lastOperator"`                           // 最后操作人
	LastOperationTime int64           `gorm:"column:last_operation_time;type:bigint;comment:最后操作时间" json:"lastOperationTime"`                     // 最后操作时间
	CreateTime        int64           `gorm:"column:create_time;type:bigint;comment:创建时间" json:"createTime"`                                      // 创建时间
}

// TableName AgentCommission's table name
func (*AgentCommission) TableName() string {
	return TableNameAgentCommission
}