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

	"github.com/gzorm/common/zorm/model"
)

func newActivityCondition(db *gorm.DB, opts ...gen.DOOption) activityCondition {
	_activityCondition := activityCondition{}

	_activityCondition.activityConditionDo.UseDB(db, opts...)
	_activityCondition.activityConditionDo.UseModel(&model.ActivityCondition{})

	tableName := _activityCondition.activityConditionDo.TableName()
	_activityCondition.ALL = field.NewAsterisk(tableName)
	_activityCondition.RuleID = field.NewInt64(tableName, "rule_id")
	_activityCondition.ActivityID = field.NewInt64(tableName, "activity_id")
	_activityCondition.ApplicableVipLevel = field.NewInt64(tableName, "applicable_vip_level")
	_activityCondition.MaxClaimPerPeriod = field.NewInt64(tableName, "max_claim_per_period")
	_activityCondition.ClaimIntervalUnit = field.NewInt64(tableName, "claim_interval_unit")
	_activityCondition.DepositChannels = field.NewString(tableName, "deposit_channels")
	_activityCondition.ClaimMethod = field.NewInt64(tableName, "claim_method")
	_activityCondition.IsAgentActive = field.NewInt64(tableName, "is_agent_active")
	_activityCondition.SpecificAgents = field.NewInt64(tableName, "specific_agents")
	_activityCondition.InviteCodes = field.NewString(tableName, "invite_codes")
	_activityCondition.CreateAt = field.NewInt64(tableName, "create_at")
	_activityCondition.UpdateAt = field.NewInt64(tableName, "update_at")
	_activityCondition.OpUser = field.NewString(tableName, "op_user")

	_activityCondition.fillFieldMap()

	return _activityCondition
}

// activityCondition 活动领取条件-谁能领
type activityCondition struct {
	activityConditionDo

	ALL                field.Asterisk
	RuleID             field.Int64  // 规则ID，自增主键
	ActivityID         field.Int64  // 活动ID
	ApplicableVipLevel field.Int64  // 适用VIP等级：0-全部
	MaxClaimPerPeriod  field.Int64  // 可领取次数，0表示不限制
	ClaimIntervalUnit  field.Int64  // 领取次数间隔单位：1-天，2-周，3-月，4-年
	DepositChannels    field.String // 关联存款渠道：0-全部
	ClaimMethod        field.Int64  // 领取方式，0表示自动领取，1表示手动领取
	IsAgentActive      field.Int64  // 是否代理线活动，0表示否，1表示是
	SpecificAgents     field.Int64  // 指定代理线:0表示否，1表示是
	InviteCodes        field.String // 指定代理邀请码,使用,好分割
	CreateAt           field.Int64  // 创建时间
	UpdateAt           field.Int64  // 修改时间
	OpUser             field.String // 操作人

	fieldMap map[string]field.Expr
}

func (a activityCondition) Table(newTableName string) *activityCondition {
	a.activityConditionDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activityCondition) As(alias string) *activityCondition {
	a.activityConditionDo.DO = *(a.activityConditionDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activityCondition) updateTableName(table string) *activityCondition {
	a.ALL = field.NewAsterisk(table)
	a.RuleID = field.NewInt64(table, "rule_id")
	a.ActivityID = field.NewInt64(table, "activity_id")
	a.ApplicableVipLevel = field.NewInt64(table, "applicable_vip_level")
	a.MaxClaimPerPeriod = field.NewInt64(table, "max_claim_per_period")
	a.ClaimIntervalUnit = field.NewInt64(table, "claim_interval_unit")
	a.DepositChannels = field.NewString(table, "deposit_channels")
	a.ClaimMethod = field.NewInt64(table, "claim_method")
	a.IsAgentActive = field.NewInt64(table, "is_agent_active")
	a.SpecificAgents = field.NewInt64(table, "specific_agents")
	a.InviteCodes = field.NewString(table, "invite_codes")
	a.CreateAt = field.NewInt64(table, "create_at")
	a.UpdateAt = field.NewInt64(table, "update_at")
	a.OpUser = field.NewString(table, "op_user")

	a.fillFieldMap()

	return a
}

func (a *activityCondition) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activityCondition) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 13)
	a.fieldMap["rule_id"] = a.RuleID
	a.fieldMap["activity_id"] = a.ActivityID
	a.fieldMap["applicable_vip_level"] = a.ApplicableVipLevel
	a.fieldMap["max_claim_per_period"] = a.MaxClaimPerPeriod
	a.fieldMap["claim_interval_unit"] = a.ClaimIntervalUnit
	a.fieldMap["deposit_channels"] = a.DepositChannels
	a.fieldMap["claim_method"] = a.ClaimMethod
	a.fieldMap["is_agent_active"] = a.IsAgentActive
	a.fieldMap["specific_agents"] = a.SpecificAgents
	a.fieldMap["invite_codes"] = a.InviteCodes
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
	a.fieldMap["op_user"] = a.OpUser
}

func (a activityCondition) clone(db *gorm.DB) activityCondition {
	a.activityConditionDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activityCondition) replaceDB(db *gorm.DB) activityCondition {
	a.activityConditionDo.ReplaceDB(db)
	return a
}

type activityConditionDo struct{ gen.DO }

type IActivityConditionDo interface {
	gen.SubQuery
	Debug() IActivityConditionDo
	WithContext(ctx context.Context) IActivityConditionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IActivityConditionDo
	WriteDB() IActivityConditionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IActivityConditionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IActivityConditionDo
	Not(conds ...gen.Condition) IActivityConditionDo
	Or(conds ...gen.Condition) IActivityConditionDo
	Select(conds ...field.Expr) IActivityConditionDo
	Where(conds ...gen.Condition) IActivityConditionDo
	Order(conds ...field.Expr) IActivityConditionDo
	Distinct(cols ...field.Expr) IActivityConditionDo
	Omit(cols ...field.Expr) IActivityConditionDo
	Join(table schema.Tabler, on ...field.Expr) IActivityConditionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IActivityConditionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IActivityConditionDo
	Group(cols ...field.Expr) IActivityConditionDo
	Having(conds ...gen.Condition) IActivityConditionDo
	Limit(limit int) IActivityConditionDo
	Offset(offset int) IActivityConditionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityConditionDo
	Unscoped() IActivityConditionDo
	Create(values ...*model.ActivityCondition) error
	CreateInBatches(values []*model.ActivityCondition, batchSize int) error
	Save(values ...*model.ActivityCondition) error
	First() (*model.ActivityCondition, error)
	Take() (*model.ActivityCondition, error)
	Last() (*model.ActivityCondition, error)
	Find() ([]*model.ActivityCondition, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityCondition, err error)
	FindInBatches(result *[]*model.ActivityCondition, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ActivityCondition) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IActivityConditionDo
	Assign(attrs ...field.AssignExpr) IActivityConditionDo
	Joins(fields ...field.RelationField) IActivityConditionDo
	Preload(fields ...field.RelationField) IActivityConditionDo
	FirstOrInit() (*model.ActivityCondition, error)
	FirstOrCreate() (*model.ActivityCondition, error)
	FindByPage(offset int, limit int) (result []*model.ActivityCondition, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IActivityConditionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a activityConditionDo) Debug() IActivityConditionDo {
	return a.withDO(a.DO.Debug())
}

func (a activityConditionDo) WithContext(ctx context.Context) IActivityConditionDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityConditionDo) ReadDB() IActivityConditionDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityConditionDo) WriteDB() IActivityConditionDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityConditionDo) Session(config *gorm.Session) IActivityConditionDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityConditionDo) Clauses(conds ...clause.Expression) IActivityConditionDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityConditionDo) Returning(value interface{}, columns ...string) IActivityConditionDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityConditionDo) Not(conds ...gen.Condition) IActivityConditionDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityConditionDo) Or(conds ...gen.Condition) IActivityConditionDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityConditionDo) Select(conds ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityConditionDo) Where(conds ...gen.Condition) IActivityConditionDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityConditionDo) Order(conds ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityConditionDo) Distinct(cols ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityConditionDo) Omit(cols ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityConditionDo) Join(table schema.Tabler, on ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityConditionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityConditionDo) RightJoin(table schema.Tabler, on ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityConditionDo) Group(cols ...field.Expr) IActivityConditionDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityConditionDo) Having(conds ...gen.Condition) IActivityConditionDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityConditionDo) Limit(limit int) IActivityConditionDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityConditionDo) Offset(offset int) IActivityConditionDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityConditionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityConditionDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityConditionDo) Unscoped() IActivityConditionDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityConditionDo) Create(values ...*model.ActivityCondition) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityConditionDo) CreateInBatches(values []*model.ActivityCondition, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityConditionDo) Save(values ...*model.ActivityCondition) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityConditionDo) First() (*model.ActivityCondition, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityCondition), nil
	}
}

func (a activityConditionDo) Take() (*model.ActivityCondition, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityCondition), nil
	}
}

func (a activityConditionDo) Last() (*model.ActivityCondition, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityCondition), nil
	}
}

func (a activityConditionDo) Find() ([]*model.ActivityCondition, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActivityCondition), err
}

func (a activityConditionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityCondition, err error) {
	buf := make([]*model.ActivityCondition, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityConditionDo) FindInBatches(result *[]*model.ActivityCondition, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityConditionDo) Attrs(attrs ...field.AssignExpr) IActivityConditionDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityConditionDo) Assign(attrs ...field.AssignExpr) IActivityConditionDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityConditionDo) Joins(fields ...field.RelationField) IActivityConditionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityConditionDo) Preload(fields ...field.RelationField) IActivityConditionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityConditionDo) FirstOrInit() (*model.ActivityCondition, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityCondition), nil
	}
}

func (a activityConditionDo) FirstOrCreate() (*model.ActivityCondition, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityCondition), nil
	}
}

func (a activityConditionDo) FindByPage(offset int, limit int) (result []*model.ActivityCondition, count int64, err error) {
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

func (a activityConditionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityConditionDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityConditionDo) Delete(models ...*model.ActivityCondition) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityConditionDo) withDO(do gen.Dao) *activityConditionDo {
	a.DO = *do.(*gen.DO)
	return a
}