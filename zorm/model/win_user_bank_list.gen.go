// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserBankList = "win_user_bank_list"

// WinUserBankList 用户银行卡列表
type WinUserBankList struct {
	ID               int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	UID              int64  `gorm:"column:uid;type:int;not null;comment:UID" json:"uid"`                                                                            // UID
	Username         string `gorm:"column:username;type:varchar(32);not null;comment:用户名" json:"username"`                                                          // 用户名
	CategoryCurrency int64  `gorm:"column:category_currency;type:tinyint;not null;comment:货币类型:0-数字货币 1-法币" json:"categoryCurrency"`                                // 货币类型:0-数字货币 1-法币
	CategoryTransfer int64  `gorm:"column:category_transfer;type:tinyint;not null;default:1;comment:转账类型：1-TRC,2-ERC,3-BANK,4-PIX,5-GCASH" json:"categoryTransfer"` // 转账类型：1-TRC,2-ERC,3-BANK,4-PIX,5-GCASH
	BankName         string `gorm:"column:bank_name;type:varchar(255);comment:银行名称" json:"bankName"`                                                                // 银行名称
	Address          string `gorm:"column:address;type:varchar(255);not null;comment:提款地址" json:"address"`                                                          // 提款地址
	BankCode         string `gorm:"column:bank_code;type:varchar(64);comment:银行编码" json:"bankCode"`                                                                 // 银行编码
	Status           int64  `gorm:"column:status;type:tinyint;not null;default:2;comment:状态:1-默认地址(启用) 2-正常启用 3-删除" json:"status"`                                  // 状态:1-默认地址(启用) 2-正常启用 3-删除
	CreatedAt        int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt        int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	CreatedBy        string `gorm:"column:created_by;type:varchar(32);comment:后台创建人" json:"createdBy"`       // 后台创建人
	UpdatedBy        string `gorm:"column:updated_by;type:varchar(32);comment:后台修改人" json:"updatedBy"`       // 后台修改人
	OperatorName     string `gorm:"column:operator_name;type:varchar(32);comment:操作人姓名" json:"operatorName"` // 操作人姓名
}

// TableName WinUserBankList's table name
func (*WinUserBankList) TableName() string {
	return TableNameWinUserBankList
}
