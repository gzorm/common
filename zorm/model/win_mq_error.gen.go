// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinMqError = "win_mq_error"

// WinMqError mq消费异常表
type WinMqError struct {
	ID            int64  `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:ID主键" json:"id,string"`           // ID主键
	MessageNo     string `gorm:"column:message_no;type:varchar(255);not null;comment:消息ID" json:"messageNo"`                  // 消息ID
	MessageType   string `gorm:"column:message_type;type:varchar(20);not null;comment:消息类型" json:"messageType"`               // 消息类型
	Status        int64  `gorm:"column:status;type:int;not null;comment:0-未处理,1-已处理,2-忽略" json:"status"`                      // 0-未处理,1-已处理,2-忽略
	Data          string `gorm:"column:data;type:text;not null;comment:完整消息体" json:"data"`                                    // 完整消息体
	CreatedAt     int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                             // 创建时间戳
	CreatedUser   string `gorm:"column:created_user;type:varchar(20);not null;default:SYSTEM;comment:创建人" json:"createdUser"` // 创建人
	UpdatedAt     int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                                             // 修改时间戳
	UpdateUser    string `gorm:"column:update_user;type:varchar(20);not null;default:SYSTEM;comment:最后修改人" json:"updateUser"` // 最后修改人
	ExceptionInfo string `gorm:"column:exception_info;type:text;comment:异常信息" json:"exceptionInfo"`                           // 异常信息
}

// TableName WinMqError's table name
func (*WinMqError) TableName() string {
	return TableNameWinMqError
}
