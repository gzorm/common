// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinHelpType = "win_help_type"

// WinHelpType 帮助类型
type WinHelpType struct {
	ID        int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键编号" json:"id,string"` // 主键编号
	Language  string `gorm:"column:language;type:varchar(150);not null;comment:语言" json:"language"`          // 语言
	TypeName  string `gorm:"column:type_name;type:varchar(150);not null;comment:类型名称" json:"typeName"`       // 类型名称
	ImageURL  string `gorm:"column:image_url;type:varchar(500);comment:图片地址" json:"imageUrl"`                // 图片地址
	Sort      int64  `gorm:"column:sort;type:int;comment:排序" json:"sort"`                                    // 排序
	Status    int64  `gorm:"column:status;type:int;comment:状态:1-启用 0-停用" json:"status"`                      // 状态:1-启用 0-停用
	CreateBy  string `gorm:"column:create_by;type:varchar(50);comment:创建者" json:"createBy"`                  // 创建者
	UpdateBy  string `gorm:"column:update_by;type:varchar(50);comment:更新人" json:"updateBy"`                  // 更新人
	CreatedAt int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinHelpType's table name
func (*WinHelpType) TableName() string {
	return TableNameWinHelpType
}
