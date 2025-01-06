// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/gzorm/commons/zorm/model"
)

func newActivityAlertLog(db *gorm.DB, opts ...gen.DOOption) activityAlertLog {
	_activityAlertLog := activityAlertLog{}

	_activityAlertLog.activityAlertLogDo.UseDB(db, opts...)
	_activityAlertLog.activityAlertLogDo.UseModel(&model.ActivityAlertLog{})

	tableName := _activityAlertLog.activityAlertLogDo.TableName()
	_activityAlertLog.ALL = field.NewAsterisk(tableName)
	_activityAlertLog.ID = field.NewInt64(tableName, "id")
	_activityAlertLog.ActivityID = field.NewInt64(tableName, "activity_id")
	_activityAlertLog.UserID = field.NewString(tableName, "user_id")
	_activityAlertLog.AlarmType = field.NewInt64(tableName, "alarm_type")
	_activityAlertLog.TotalBonus = field.NewField(tableName, "total_bonus")
	_activityAlertLog.TodaySumBonus = field.NewField(tableName, "today_sum_bonus")
	_activityAlertLog.TotalBudgetThreshold = field.NewField(tableName, "total_budget_threshold")
	_activityAlertLog.DailyJackpotThreshold = field.NewField(tableName, "daily_jackpot_threshold")
	_activityAlertLog.CreatedAt = field.NewInt64(tableName, "created_at")

	_activityAlertLog.fillFieldMap()

	return _activityAlertLog
}

// activityAlertLog 活动告警日志表
type activityAlertLog struct {
	activityAlertLogDo

	ALL                   field.Asterisk
	ID                    field.Int64  // 主键
	ActivityID            field.Int64  // 活动id
	UserID                field.String // 用户ID
	AlarmType             field.Int64  // 告警类型：1今日超发告警，2总预算超发告警
	TotalBonus            field.Field  // 累计已发放彩金
	TodaySumBonus         field.Field  // 今日已发放彩金
	TotalBudgetThreshold  field.Field  // 活动总预算告警阈值
	DailyJackpotThreshold field.Field  // 单日彩金上限告警阈值
	CreatedAt             field.Int64  // 创建时间

	fieldMap map[string]field.Expr
}

func (a activityAlertLog) Table(newTableName string) *activityAlertLog {
	a.activityAlertLogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activityAlertLog) As(alias string) *activityAlertLog {
	a.activityAlertLogDo.DO = *(a.activityAlertLogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activityAlertLog) updateTableName(table string) *activityAlertLog {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.ActivityID = field.NewInt64(table, "activity_id")
	a.UserID = field.NewString(table, "user_id")
	a.AlarmType = field.NewInt64(table, "alarm_type")
	a.TotalBonus = field.NewField(table, "total_bonus")
	a.TodaySumBonus = field.NewField(table, "today_sum_bonus")
	a.TotalBudgetThreshold = field.NewField(table, "total_budget_threshold")
	a.DailyJackpotThreshold = field.NewField(table, "daily_jackpot_threshold")
	a.CreatedAt = field.NewInt64(table, "created_at")

	a.fillFieldMap()

	return a
}

func (a *activityAlertLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activityAlertLog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["activity_id"] = a.ActivityID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["alarm_type"] = a.AlarmType
	a.fieldMap["total_bonus"] = a.TotalBonus
	a.fieldMap["today_sum_bonus"] = a.TodaySumBonus
	a.fieldMap["total_budget_threshold"] = a.TotalBudgetThreshold
	a.fieldMap["daily_jackpot_threshold"] = a.DailyJackpotThreshold
	a.fieldMap["created_at"] = a.CreatedAt
}

func (a activityAlertLog) clone(db *gorm.DB) activityAlertLog {
	a.activityAlertLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activityAlertLog) replaceDB(db *gorm.DB) activityAlertLog {
	a.activityAlertLogDo.ReplaceDB(db)
	return a
}

type activityAlertLogDo struct{ gen.DO }

type IActivityAlertLogDo interface {
	gen.SubQuery
	Debug() IActivityAlertLogDo
	WithContext(ctx context.Context) IActivityAlertLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IActivityAlertLogDo
	WriteDB() IActivityAlertLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IActivityAlertLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IActivityAlertLogDo
	Not(conds ...gen.Condition) IActivityAlertLogDo
	Or(conds ...gen.Condition) IActivityAlertLogDo
	Select(conds ...field.Expr) IActivityAlertLogDo
	Where(conds ...gen.Condition) IActivityAlertLogDo
	Order(conds ...field.Expr) IActivityAlertLogDo
	Distinct(cols ...field.Expr) IActivityAlertLogDo
	Omit(cols ...field.Expr) IActivityAlertLogDo
	Join(table schema.Tabler, on ...field.Expr) IActivityAlertLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IActivityAlertLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IActivityAlertLogDo
	Group(cols ...field.Expr) IActivityAlertLogDo
	Having(conds ...gen.Condition) IActivityAlertLogDo
	Limit(limit int) IActivityAlertLogDo
	Offset(offset int) IActivityAlertLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityAlertLogDo
	Unscoped() IActivityAlertLogDo
	Create(values ...*model.ActivityAlertLog) error
	CreateInBatches(values []*model.ActivityAlertLog, batchSize int) error
	Save(values ...*model.ActivityAlertLog) error
	First() (*model.ActivityAlertLog, error)
	Take() (*model.ActivityAlertLog, error)
	Last() (*model.ActivityAlertLog, error)
	Find() ([]*model.ActivityAlertLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityAlertLog, err error)
	FindInBatches(result *[]*model.ActivityAlertLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ActivityAlertLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IActivityAlertLogDo
	Assign(attrs ...field.AssignExpr) IActivityAlertLogDo
	Joins(fields ...field.RelationField) IActivityAlertLogDo
	Preload(fields ...field.RelationField) IActivityAlertLogDo
	FirstOrInit() (*model.ActivityAlertLog, error)
	FirstOrCreate() (*model.ActivityAlertLog, error)
	FindByPage(offset int, limit int) (result []*model.ActivityAlertLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IActivityAlertLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a activityAlertLogDo) Debug() IActivityAlertLogDo {
	return a.withDO(a.DO.Debug())
}

func (a activityAlertLogDo) WithContext(ctx context.Context) IActivityAlertLogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityAlertLogDo) ReadDB() IActivityAlertLogDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityAlertLogDo) WriteDB() IActivityAlertLogDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityAlertLogDo) Session(config *gorm.Session) IActivityAlertLogDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityAlertLogDo) Clauses(conds ...clause.Expression) IActivityAlertLogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityAlertLogDo) Returning(value interface{}, columns ...string) IActivityAlertLogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityAlertLogDo) Not(conds ...gen.Condition) IActivityAlertLogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityAlertLogDo) Or(conds ...gen.Condition) IActivityAlertLogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityAlertLogDo) Select(conds ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityAlertLogDo) Where(conds ...gen.Condition) IActivityAlertLogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityAlertLogDo) Order(conds ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityAlertLogDo) Distinct(cols ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityAlertLogDo) Omit(cols ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityAlertLogDo) Join(table schema.Tabler, on ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityAlertLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityAlertLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityAlertLogDo) Group(cols ...field.Expr) IActivityAlertLogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityAlertLogDo) Having(conds ...gen.Condition) IActivityAlertLogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityAlertLogDo) Limit(limit int) IActivityAlertLogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityAlertLogDo) Offset(offset int) IActivityAlertLogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityAlertLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityAlertLogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityAlertLogDo) Unscoped() IActivityAlertLogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityAlertLogDo) Create(values ...*model.ActivityAlertLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityAlertLogDo) CreateInBatches(values []*model.ActivityAlertLog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityAlertLogDo) Save(values ...*model.ActivityAlertLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityAlertLogDo) First() (*model.ActivityAlertLog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityAlertLog), nil
	}
}

func (a activityAlertLogDo) Take() (*model.ActivityAlertLog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityAlertLog), nil
	}
}

func (a activityAlertLogDo) Last() (*model.ActivityAlertLog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityAlertLog), nil
	}
}

func (a activityAlertLogDo) Find() ([]*model.ActivityAlertLog, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActivityAlertLog), err
}

func (a activityAlertLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityAlertLog, err error) {
	buf := make([]*model.ActivityAlertLog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityAlertLogDo) FindInBatches(result *[]*model.ActivityAlertLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityAlertLogDo) Attrs(attrs ...field.AssignExpr) IActivityAlertLogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityAlertLogDo) Assign(attrs ...field.AssignExpr) IActivityAlertLogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityAlertLogDo) Joins(fields ...field.RelationField) IActivityAlertLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityAlertLogDo) Preload(fields ...field.RelationField) IActivityAlertLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityAlertLogDo) FirstOrInit() (*model.ActivityAlertLog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityAlertLog), nil
	}
}

func (a activityAlertLogDo) FirstOrCreate() (*model.ActivityAlertLog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityAlertLog), nil
	}
}

func (a activityAlertLogDo) FindByPage(offset int, limit int) (result []*model.ActivityAlertLog, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a activityAlertLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityAlertLogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityAlertLogDo) Delete(models ...*model.ActivityAlertLog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityAlertLogDo) withDO(do gen.Dao) *activityAlertLogDo {
	a.DO = *do.(*gen.DO)
	return a
}
