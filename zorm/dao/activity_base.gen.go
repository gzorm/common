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

func newActivityBase(db *gorm.DB, opts ...gen.DOOption) activityBase {
	_activityBase := activityBase{}

	_activityBase.activityBaseDo.UseDB(db, opts...)
	_activityBase.activityBaseDo.UseModel(&model.ActivityBase{})

	tableName := _activityBase.activityBaseDo.TableName()
	_activityBase.ALL = field.NewAsterisk(tableName)
	_activityBase.ID = field.NewInt64(tableName, "id")
	_activityBase.Category = field.NewInt64(tableName, "category")
	_activityBase.ActivityType = field.NewInt64(tableName, "activity_type")
	_activityBase.ActivityTitle = field.NewString(tableName, "activity_title")
	_activityBase.ActivityDescription = field.NewString(tableName, "activity_description")
	_activityBase.DepositType = field.NewInt64(tableName, "deposit_type")
	_activityBase.ActivitySort = field.NewInt64(tableName, "activity_sort")
	_activityBase.IsEnabled = field.NewInt64(tableName, "is_enabled")
	_activityBase.StartDate = field.NewInt64(tableName, "start_date")
	_activityBase.EndDate = field.NewInt64(tableName, "end_date")
	_activityBase.ActivityPlatform = field.NewString(tableName, "activity_platform")
	_activityBase.MinDepositAmount = field.NewField(tableName, "min_deposit_amount")
	_activityBase.MaxDepositAmount = field.NewField(tableName, "max_deposit_amount")
	_activityBase.BonusPercentage = field.NewField(tableName, "bonus_percentage")
	_activityBase.MaxBonusAmount = field.NewField(tableName, "max_bonus_amount")
	_activityBase.MaxActivePerson = field.NewInt64(tableName, "max_active_person")
	_activityBase.GameProviderSubtypeID = field.NewString(tableName, "game_provider_subtype_id")
	_activityBase.GameProviderID = field.NewString(tableName, "game_provider_id")
	_activityBase.GameListID = field.NewString(tableName, "game_list_id")
	_activityBase.WageringRequirement = field.NewField(tableName, "wagering_requirement")
	_activityBase.CreateAt = field.NewInt64(tableName, "create_at")
	_activityBase.UpdateAt = field.NewInt64(tableName, "update_at")
	_activityBase.OpUser = field.NewString(tableName, "op_user")

	_activityBase.fillFieldMap()

	return _activityBase
}

// activityBase 活动基础信息表
type activityBase struct {
	activityBaseDo

	ALL                   field.Asterisk
	ID                    field.Int64  // 活动ID，自增主键
	Category              field.Int64  // 活动类型:1-首存送、2-次存送、3-日存送、4-周存送
	ActivityType          field.Int64  // 活动分类：1-充值送，2-注册送，3-救济金，4-签到，5-特定活动
	ActivityTitle         field.String // 活动标题
	ActivityDescription   field.String // 活动描述
	DepositType           field.Int64  // 存款类型:0-全部、1-首存、2-次存、3-三存、4-其他存款
	ActivitySort          field.Int64  // 活动排序（0-1000）活动排序只针对活动排序内容不涉及弹窗和轮播
	IsEnabled             field.Int64  // 是否启用：1-启用，2-禁用
	StartDate             field.Int64  // 活动开始日期 (YYYYMMDD)
	EndDate               field.Int64  // 活动结束日期 (YYYYMMDD)
	ActivityPlatform      field.String // 活动端多选：H5、WEB、APP
	MinDepositAmount      field.Field  // 活动最小充值金额
	MaxDepositAmount      field.Field  // 活动最大充值金额
	BonusPercentage       field.Field  // 活动奖金比例
	MaxBonusAmount        field.Field  // 最高可获奖金
	MaxActivePerson       field.Int64  // 活动名额限制
	GameProviderSubtypeID field.String // 不能改 游戏类型：0-全部
	GameProviderID        field.String // 不能改 游戏厂商：0-全部
	GameListID            field.String // 不能改 子游戏类型：0-全部
	WageringRequirement   field.Field  // 不能改 活动打码倍数
	CreateAt              field.Int64  // 创建时间
	UpdateAt              field.Int64  // 修改时间
	OpUser                field.String // 操作人

	fieldMap map[string]field.Expr
}

func (a activityBase) Table(newTableName string) *activityBase {
	a.activityBaseDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a activityBase) As(alias string) *activityBase {
	a.activityBaseDo.DO = *(a.activityBaseDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *activityBase) updateTableName(table string) *activityBase {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Category = field.NewInt64(table, "category")
	a.ActivityType = field.NewInt64(table, "activity_type")
	a.ActivityTitle = field.NewString(table, "activity_title")
	a.ActivityDescription = field.NewString(table, "activity_description")
	a.DepositType = field.NewInt64(table, "deposit_type")
	a.ActivitySort = field.NewInt64(table, "activity_sort")
	a.IsEnabled = field.NewInt64(table, "is_enabled")
	a.StartDate = field.NewInt64(table, "start_date")
	a.EndDate = field.NewInt64(table, "end_date")
	a.ActivityPlatform = field.NewString(table, "activity_platform")
	a.MinDepositAmount = field.NewField(table, "min_deposit_amount")
	a.MaxDepositAmount = field.NewField(table, "max_deposit_amount")
	a.BonusPercentage = field.NewField(table, "bonus_percentage")
	a.MaxBonusAmount = field.NewField(table, "max_bonus_amount")
	a.MaxActivePerson = field.NewInt64(table, "max_active_person")
	a.GameProviderSubtypeID = field.NewString(table, "game_provider_subtype_id")
	a.GameProviderID = field.NewString(table, "game_provider_id")
	a.GameListID = field.NewString(table, "game_list_id")
	a.WageringRequirement = field.NewField(table, "wagering_requirement")
	a.CreateAt = field.NewInt64(table, "create_at")
	a.UpdateAt = field.NewInt64(table, "update_at")
	a.OpUser = field.NewString(table, "op_user")

	a.fillFieldMap()

	return a
}

func (a *activityBase) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *activityBase) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 23)
	a.fieldMap["id"] = a.ID
	a.fieldMap["category"] = a.Category
	a.fieldMap["activity_type"] = a.ActivityType
	a.fieldMap["activity_title"] = a.ActivityTitle
	a.fieldMap["activity_description"] = a.ActivityDescription
	a.fieldMap["deposit_type"] = a.DepositType
	a.fieldMap["activity_sort"] = a.ActivitySort
	a.fieldMap["is_enabled"] = a.IsEnabled
	a.fieldMap["start_date"] = a.StartDate
	a.fieldMap["end_date"] = a.EndDate
	a.fieldMap["activity_platform"] = a.ActivityPlatform
	a.fieldMap["min_deposit_amount"] = a.MinDepositAmount
	a.fieldMap["max_deposit_amount"] = a.MaxDepositAmount
	a.fieldMap["bonus_percentage"] = a.BonusPercentage
	a.fieldMap["max_bonus_amount"] = a.MaxBonusAmount
	a.fieldMap["max_active_person"] = a.MaxActivePerson
	a.fieldMap["game_provider_subtype_id"] = a.GameProviderSubtypeID
	a.fieldMap["game_provider_id"] = a.GameProviderID
	a.fieldMap["game_list_id"] = a.GameListID
	a.fieldMap["wagering_requirement"] = a.WageringRequirement
	a.fieldMap["create_at"] = a.CreateAt
	a.fieldMap["update_at"] = a.UpdateAt
	a.fieldMap["op_user"] = a.OpUser
}

func (a activityBase) clone(db *gorm.DB) activityBase {
	a.activityBaseDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a activityBase) replaceDB(db *gorm.DB) activityBase {
	a.activityBaseDo.ReplaceDB(db)
	return a
}

type activityBaseDo struct{ gen.DO }

type IActivityBaseDo interface {
	gen.SubQuery
	Debug() IActivityBaseDo
	WithContext(ctx context.Context) IActivityBaseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IActivityBaseDo
	WriteDB() IActivityBaseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IActivityBaseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IActivityBaseDo
	Not(conds ...gen.Condition) IActivityBaseDo
	Or(conds ...gen.Condition) IActivityBaseDo
	Select(conds ...field.Expr) IActivityBaseDo
	Where(conds ...gen.Condition) IActivityBaseDo
	Order(conds ...field.Expr) IActivityBaseDo
	Distinct(cols ...field.Expr) IActivityBaseDo
	Omit(cols ...field.Expr) IActivityBaseDo
	Join(table schema.Tabler, on ...field.Expr) IActivityBaseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IActivityBaseDo
	RightJoin(table schema.Tabler, on ...field.Expr) IActivityBaseDo
	Group(cols ...field.Expr) IActivityBaseDo
	Having(conds ...gen.Condition) IActivityBaseDo
	Limit(limit int) IActivityBaseDo
	Offset(offset int) IActivityBaseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityBaseDo
	Unscoped() IActivityBaseDo
	Create(values ...*model.ActivityBase) error
	CreateInBatches(values []*model.ActivityBase, batchSize int) error
	Save(values ...*model.ActivityBase) error
	First() (*model.ActivityBase, error)
	Take() (*model.ActivityBase, error)
	Last() (*model.ActivityBase, error)
	Find() ([]*model.ActivityBase, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityBase, err error)
	FindInBatches(result *[]*model.ActivityBase, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ActivityBase) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IActivityBaseDo
	Assign(attrs ...field.AssignExpr) IActivityBaseDo
	Joins(fields ...field.RelationField) IActivityBaseDo
	Preload(fields ...field.RelationField) IActivityBaseDo
	FirstOrInit() (*model.ActivityBase, error)
	FirstOrCreate() (*model.ActivityBase, error)
	FindByPage(offset int, limit int) (result []*model.ActivityBase, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IActivityBaseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a activityBaseDo) Debug() IActivityBaseDo {
	return a.withDO(a.DO.Debug())
}

func (a activityBaseDo) WithContext(ctx context.Context) IActivityBaseDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a activityBaseDo) ReadDB() IActivityBaseDo {
	return a.Clauses(dbresolver.Read)
}

func (a activityBaseDo) WriteDB() IActivityBaseDo {
	return a.Clauses(dbresolver.Write)
}

func (a activityBaseDo) Session(config *gorm.Session) IActivityBaseDo {
	return a.withDO(a.DO.Session(config))
}

func (a activityBaseDo) Clauses(conds ...clause.Expression) IActivityBaseDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a activityBaseDo) Returning(value interface{}, columns ...string) IActivityBaseDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a activityBaseDo) Not(conds ...gen.Condition) IActivityBaseDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a activityBaseDo) Or(conds ...gen.Condition) IActivityBaseDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a activityBaseDo) Select(conds ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a activityBaseDo) Where(conds ...gen.Condition) IActivityBaseDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a activityBaseDo) Order(conds ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a activityBaseDo) Distinct(cols ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a activityBaseDo) Omit(cols ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a activityBaseDo) Join(table schema.Tabler, on ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a activityBaseDo) LeftJoin(table schema.Tabler, on ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a activityBaseDo) RightJoin(table schema.Tabler, on ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a activityBaseDo) Group(cols ...field.Expr) IActivityBaseDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a activityBaseDo) Having(conds ...gen.Condition) IActivityBaseDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a activityBaseDo) Limit(limit int) IActivityBaseDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a activityBaseDo) Offset(offset int) IActivityBaseDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a activityBaseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IActivityBaseDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a activityBaseDo) Unscoped() IActivityBaseDo {
	return a.withDO(a.DO.Unscoped())
}

func (a activityBaseDo) Create(values ...*model.ActivityBase) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a activityBaseDo) CreateInBatches(values []*model.ActivityBase, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a activityBaseDo) Save(values ...*model.ActivityBase) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a activityBaseDo) First() (*model.ActivityBase, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBase), nil
	}
}

func (a activityBaseDo) Take() (*model.ActivityBase, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBase), nil
	}
}

func (a activityBaseDo) Last() (*model.ActivityBase, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBase), nil
	}
}

func (a activityBaseDo) Find() ([]*model.ActivityBase, error) {
	result, err := a.DO.Find()
	return result.([]*model.ActivityBase), err
}

func (a activityBaseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ActivityBase, err error) {
	buf := make([]*model.ActivityBase, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a activityBaseDo) FindInBatches(result *[]*model.ActivityBase, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a activityBaseDo) Attrs(attrs ...field.AssignExpr) IActivityBaseDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a activityBaseDo) Assign(attrs ...field.AssignExpr) IActivityBaseDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a activityBaseDo) Joins(fields ...field.RelationField) IActivityBaseDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a activityBaseDo) Preload(fields ...field.RelationField) IActivityBaseDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a activityBaseDo) FirstOrInit() (*model.ActivityBase, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBase), nil
	}
}

func (a activityBaseDo) FirstOrCreate() (*model.ActivityBase, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ActivityBase), nil
	}
}

func (a activityBaseDo) FindByPage(offset int, limit int) (result []*model.ActivityBase, count int64, err error) {
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

func (a activityBaseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a activityBaseDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a activityBaseDo) Delete(models ...*model.ActivityBase) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *activityBaseDo) withDO(do gen.Dao) *activityBaseDo {
	a.DO = *do.(*gen.DO)
	return a
}
