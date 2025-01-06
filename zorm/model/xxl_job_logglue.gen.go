// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameXxlJobLogglue = "xxl_job_logglue"

// XxlJobLogglue mapped from table <xxl_job_logglue>
type XxlJobLogglue struct {
	ID         int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	JobID      int64     `gorm:"column:job_id;type:int;not null;comment:任务，主键ID" json:"jobId"`                   // 任务，主键ID
	GlueType   string    `gorm:"column:glue_type;type:varchar(50);comment:GLUE类型" json:"glueType"`               // GLUE类型
	GlueSource string    `gorm:"column:glue_source;type:mediumtext;comment:GLUE源代码" json:"glueSource"`           // GLUE源代码
	GlueRemark string    `gorm:"column:glue_remark;type:varchar(128);not null;comment:GLUE备注" json:"glueRemark"` // GLUE备注
	AddTime    time.Time `gorm:"column:add_time;type:datetime" json:"addTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"updateTime"`
}

// TableName XxlJobLogglue's table name
func (*XxlJobLogglue) TableName() string {
	return TableNameXxlJobLogglue
}
