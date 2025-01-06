// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserLevel = "win_user_level"

// WinUserLevel 会员等级
type WinUserLevel struct {
	ID               int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	Code             string `gorm:"column:code;type:varchar(16);not null;comment:会员等级" json:"code"`                  // 会员等级
	Name             string `gorm:"column:name;type:varchar(16);not null;comment:等级名称" json:"name"`                  // 等级名称
	Icon             string `gorm:"column:icon;type:varchar(256);not null;comment:ICON" json:"icon"`                 // ICON
	DepositChannel   string `gorm:"column:deposit_channel;type:varchar(511);comment:存款通道" json:"depositChannel"`     // 存款通道
	ScoreRelegation  int64  `gorm:"column:score_relegation;type:int;not null;comment:保级有效投注" json:"scoreRelegation"` // 保级有效投注
	RelegationDay    int64  `gorm:"column:relegation_day;type:int;not null;comment:保级有效天数" json:"relegationDay"`     // 保级有效天数
	BetSum           int64  `gorm:"column:bet_sum;type:int;not null;comment:累计投注" json:"betSum"`                     // 累计投注
	DepositSum       int64  `gorm:"column:deposit_sum;type:int;not null;comment:累计存款" json:"depositSum"`             // 累计存款
	WithdrawalCount  int64  `gorm:"column:withdrawal_count;type:int;not null;comment:每日提款次数" json:"withdrawalCount"` // 每日提款次数
	WithdrawalCoin   int64  `gorm:"column:withdrawal_coin;type:int;not null;comment:每日提款额度" json:"withdrawalCoin"`   // 每日提款额度
	RebateMax        int64  `gorm:"column:rebate_max;type:int;not null;comment:每日返水上限" json:"rebateMax"`             // 每日返水上限
	UpgradeReward    int64  `gorm:"column:upgrade_reward;type:int;not null;comment:升级礼金" json:"upgradeReward"`       // 升级礼金
	BirthdayReward   int64  `gorm:"column:birthday_reward;type:int;not null;comment:生日礼金" json:"birthdayReward"`     // 生日礼金
	WeekReward       int64  `gorm:"column:week_reward;type:int;not null;comment:周礼金" json:"weekReward"`              // 周礼金
	MonthReward      int64  `gorm:"column:month_reward;type:int;not null;comment:月礼金" json:"monthReward"`            // 月礼金
	FlowClaim        int64  `gorm:"column:flow_claim;type:int;comment:打码倍数" json:"flowClaim"`                        // 打码倍数
	CreatedAt        int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt        int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	UpdatedUser      string `gorm:"column:updated_user;type:varchar(30);comment:最后修改人" json:"updatedUser"` // 最后修改人
	ScoreUpgradeMin  int64  `gorm:"column:score_upgrade_min;type:int" json:"scoreUpgradeMin"`
	ScoreUpgradeMax  int64  `gorm:"column:score_upgrade_max;type:int" json:"scoreUpgradeMax"`
	ScoreUpgradeRate int64  `gorm:"column:score_upgrade_rate;type:int;default:1" json:"scoreUpgradeRate"`
}

// TableName WinUserLevel's table name
func (*WinUserLevel) TableName() string {
	return TableNameWinUserLevel
}
