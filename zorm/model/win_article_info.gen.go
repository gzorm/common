// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinArticleInfo = "win_article_info"

// WinArticleInfo 体育资讯表
type WinArticleInfo struct {
	ID        int64  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键编号" json:"id,string"` // 主键编号
	Language  string `gorm:"column:language;type:varchar(150);not null;comment:语言" json:"language"`          // 语言
	Title     string `gorm:"column:title;type:varchar(150);not null;comment:标题" json:"title"`                // 标题
	Sort      int64  `gorm:"column:sort;type:int;comment:排序" json:"sort"`                                    // 排序
	Status    int64  `gorm:"column:status;type:int;comment:状态:1-启用 0-停用" json:"status"`                      // 状态:1-启用 0-停用
	Content   string `gorm:"column:content;type:mediumtext;comment:内容" json:"content"`                       // 内容
	Img       string `gorm:"column:img;type:varchar(355);comment:图片" json:"img"`                             // 图片
	CreateBy  string `gorm:"column:create_by;type:varchar(50);comment:创建者" json:"createBy"`                  // 创建者
	UpdateBy  string `gorm:"column:update_by;type:varchar(50);comment:更新人" json:"updateBy"`                  // 更新人
	CreatedAt int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
}

// TableName WinArticleInfo's table name
func (*WinArticleInfo) TableName() string {
	return TableNameWinArticleInfo
}
