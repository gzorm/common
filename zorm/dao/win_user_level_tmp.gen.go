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

func newWinUserLevelTmp(db *gorm.DB, opts ...gen.DOOption) winUserLevelTmp {
	_winUserLevelTmp := winUserLevelTmp{}

	_winUserLevelTmp.winUserLevelTmpDo.UseDB(db, opts...)
	_winUserLevelTmp.winUserLevelTmpDo.UseModel(&model.WinUserLevelTmp{})

	tableName := _winUserLevelTmp.winUserLevelTmpDo.TableName()
	_winUserLevelTmp.ALL = field.NewAsterisk(tableName)
	_winUserLevelTmp.ID = field.NewInt64(tableName, "id")
	_winUserLevelTmp.Code = field.NewString(tableName, "code")
	_winUserLevelTmp.Name = field.NewString(tableName, "name")
	_winUserLevelTmp.Icon = field.NewString(tableName, "icon")
	_winUserLevelTmp.DepositChannel = field.NewString(tableName, "deposit_channel")
	_winUserLevelTmp.ScoreUpgradeRate = field.NewInt64(tableName, "score_upgrade_rate")
	_winUserLevelTmp.ScoreUpgradeMax = field.NewInt64(tableName, "score_upgrade_max")
	_winUserLevelTmp.ScoreUpgradeMin = field.NewInt64(tableName, "score_upgrade_min")
	_winUserLevelTmp.ScoreRelegation = field.NewInt64(tableName, "score_relegation")
	_winUserLevelTmp.RelegationDay = field.NewInt64(tableName, "relegation_day")
	_winUserLevelTmp.BetSum = field.NewInt64(tableName, "bet_sum")
	_winUserLevelTmp.DepositSum = field.NewInt64(tableName, "deposit_sum")
	_winUserLevelTmp.WithdrawalCount = field.NewInt64(tableName, "withdrawal_count")
	_winUserLevelTmp.WithdrawalCoin = field.NewInt64(tableName, "withdrawal_coin")
	_winUserLevelTmp.RebateMax = field.NewInt64(tableName, "rebate_max")
	_winUserLevelTmp.UpgradeReward = field.NewInt64(tableName, "upgrade_reward")
	_winUserLevelTmp.BirthdayReward = field.NewInt64(tableName, "birthday_reward")
	_winUserLevelTmp.WeekReward = field.NewInt64(tableName, "week_reward")
	_winUserLevelTmp.MonthReward = field.NewInt64(tableName, "month_reward")
	_winUserLevelTmp.FlowClaim = field.NewInt64(tableName, "flow_claim")
	_winUserLevelTmp.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserLevelTmp.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winUserLevelTmp.UpdatedUser = field.NewString(tableName, "updated_user")

	_winUserLevelTmp.fillFieldMap()

	return _winUserLevelTmp
}

// winUserLevelTmp 会员等级
type winUserLevelTmp struct {
	winUserLevelTmpDo

	ALL              field.Asterisk
	ID               field.Int64
	Code             field.String // 会员等级
	Name             field.String // 等级名称
	Icon             field.String // ICON
	DepositChannel   field.String // 存款通道
	ScoreUpgradeRate field.Int64  // 经验率
	ScoreUpgradeMax  field.Int64  // 最大经验值
	ScoreUpgradeMin  field.Int64  // 最小经验值
	ScoreRelegation  field.Int64  // 保级有效投注
	RelegationDay    field.Int64  // 保级有效天数
	BetSum           field.Int64  // 累计投注
	DepositSum       field.Int64  // 累计存款
	WithdrawalCount  field.Int64  // 每日提款次数
	WithdrawalCoin   field.Int64  // 每日提款额度
	RebateMax        field.Int64  // 每日返水上限
	UpgradeReward    field.Int64  // 升级礼金
	BirthdayReward   field.Int64  // 生日礼金
	WeekReward       field.Int64  // 周礼金
	MonthReward      field.Int64  // 月礼金
	FlowClaim        field.Int64  // 打码倍数
	CreatedAt        field.Int64
	UpdatedAt        field.Int64
	UpdatedUser      field.String // 最后修改人

	fieldMap map[string]field.Expr
}

func (w winUserLevelTmp) Table(newTableName string) *winUserLevelTmp {
	w.winUserLevelTmpDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserLevelTmp) As(alias string) *winUserLevelTmp {
	w.winUserLevelTmpDo.DO = *(w.winUserLevelTmpDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserLevelTmp) updateTableName(table string) *winUserLevelTmp {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Code = field.NewString(table, "code")
	w.Name = field.NewString(table, "name")
	w.Icon = field.NewString(table, "icon")
	w.DepositChannel = field.NewString(table, "deposit_channel")
	w.ScoreUpgradeRate = field.NewInt64(table, "score_upgrade_rate")
	w.ScoreUpgradeMax = field.NewInt64(table, "score_upgrade_max")
	w.ScoreUpgradeMin = field.NewInt64(table, "score_upgrade_min")
	w.ScoreRelegation = field.NewInt64(table, "score_relegation")
	w.RelegationDay = field.NewInt64(table, "relegation_day")
	w.BetSum = field.NewInt64(table, "bet_sum")
	w.DepositSum = field.NewInt64(table, "deposit_sum")
	w.WithdrawalCount = field.NewInt64(table, "withdrawal_count")
	w.WithdrawalCoin = field.NewInt64(table, "withdrawal_coin")
	w.RebateMax = field.NewInt64(table, "rebate_max")
	w.UpgradeReward = field.NewInt64(table, "upgrade_reward")
	w.BirthdayReward = field.NewInt64(table, "birthday_reward")
	w.WeekReward = field.NewInt64(table, "week_reward")
	w.MonthReward = field.NewInt64(table, "month_reward")
	w.FlowClaim = field.NewInt64(table, "flow_claim")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.UpdatedUser = field.NewString(table, "updated_user")

	w.fillFieldMap()

	return w
}

func (w *winUserLevelTmp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserLevelTmp) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 23)
	w.fieldMap["id"] = w.ID
	w.fieldMap["code"] = w.Code
	w.fieldMap["name"] = w.Name
	w.fieldMap["icon"] = w.Icon
	w.fieldMap["deposit_channel"] = w.DepositChannel
	w.fieldMap["score_upgrade_rate"] = w.ScoreUpgradeRate
	w.fieldMap["score_upgrade_max"] = w.ScoreUpgradeMax
	w.fieldMap["score_upgrade_min"] = w.ScoreUpgradeMin
	w.fieldMap["score_relegation"] = w.ScoreRelegation
	w.fieldMap["relegation_day"] = w.RelegationDay
	w.fieldMap["bet_sum"] = w.BetSum
	w.fieldMap["deposit_sum"] = w.DepositSum
	w.fieldMap["withdrawal_count"] = w.WithdrawalCount
	w.fieldMap["withdrawal_coin"] = w.WithdrawalCoin
	w.fieldMap["rebate_max"] = w.RebateMax
	w.fieldMap["upgrade_reward"] = w.UpgradeReward
	w.fieldMap["birthday_reward"] = w.BirthdayReward
	w.fieldMap["week_reward"] = w.WeekReward
	w.fieldMap["month_reward"] = w.MonthReward
	w.fieldMap["flow_claim"] = w.FlowClaim
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["updated_user"] = w.UpdatedUser
}

func (w winUserLevelTmp) clone(db *gorm.DB) winUserLevelTmp {
	w.winUserLevelTmpDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserLevelTmp) replaceDB(db *gorm.DB) winUserLevelTmp {
	w.winUserLevelTmpDo.ReplaceDB(db)
	return w
}

type winUserLevelTmpDo struct{ gen.DO }

type IWinUserLevelTmpDo interface {
	gen.SubQuery
	Debug() IWinUserLevelTmpDo
	WithContext(ctx context.Context) IWinUserLevelTmpDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserLevelTmpDo
	WriteDB() IWinUserLevelTmpDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserLevelTmpDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserLevelTmpDo
	Not(conds ...gen.Condition) IWinUserLevelTmpDo
	Or(conds ...gen.Condition) IWinUserLevelTmpDo
	Select(conds ...field.Expr) IWinUserLevelTmpDo
	Where(conds ...gen.Condition) IWinUserLevelTmpDo
	Order(conds ...field.Expr) IWinUserLevelTmpDo
	Distinct(cols ...field.Expr) IWinUserLevelTmpDo
	Omit(cols ...field.Expr) IWinUserLevelTmpDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserLevelTmpDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelTmpDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelTmpDo
	Group(cols ...field.Expr) IWinUserLevelTmpDo
	Having(conds ...gen.Condition) IWinUserLevelTmpDo
	Limit(limit int) IWinUserLevelTmpDo
	Offset(offset int) IWinUserLevelTmpDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelTmpDo
	Unscoped() IWinUserLevelTmpDo
	Create(values ...*model.WinUserLevelTmp) error
	CreateInBatches(values []*model.WinUserLevelTmp, batchSize int) error
	Save(values ...*model.WinUserLevelTmp) error
	First() (*model.WinUserLevelTmp, error)
	Take() (*model.WinUserLevelTmp, error)
	Last() (*model.WinUserLevelTmp, error)
	Find() ([]*model.WinUserLevelTmp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevelTmp, err error)
	FindInBatches(result *[]*model.WinUserLevelTmp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserLevelTmp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserLevelTmpDo
	Assign(attrs ...field.AssignExpr) IWinUserLevelTmpDo
	Joins(fields ...field.RelationField) IWinUserLevelTmpDo
	Preload(fields ...field.RelationField) IWinUserLevelTmpDo
	FirstOrInit() (*model.WinUserLevelTmp, error)
	FirstOrCreate() (*model.WinUserLevelTmp, error)
	FindByPage(offset int, limit int) (result []*model.WinUserLevelTmp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserLevelTmpDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserLevelTmpDo) Debug() IWinUserLevelTmpDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserLevelTmpDo) WithContext(ctx context.Context) IWinUserLevelTmpDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserLevelTmpDo) ReadDB() IWinUserLevelTmpDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserLevelTmpDo) WriteDB() IWinUserLevelTmpDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserLevelTmpDo) Session(config *gorm.Session) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserLevelTmpDo) Clauses(conds ...clause.Expression) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserLevelTmpDo) Returning(value interface{}, columns ...string) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserLevelTmpDo) Not(conds ...gen.Condition) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserLevelTmpDo) Or(conds ...gen.Condition) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserLevelTmpDo) Select(conds ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserLevelTmpDo) Where(conds ...gen.Condition) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserLevelTmpDo) Order(conds ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserLevelTmpDo) Distinct(cols ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserLevelTmpDo) Omit(cols ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserLevelTmpDo) Join(table schema.Tabler, on ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserLevelTmpDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserLevelTmpDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserLevelTmpDo) Group(cols ...field.Expr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserLevelTmpDo) Having(conds ...gen.Condition) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserLevelTmpDo) Limit(limit int) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserLevelTmpDo) Offset(offset int) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserLevelTmpDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserLevelTmpDo) Unscoped() IWinUserLevelTmpDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserLevelTmpDo) Create(values ...*model.WinUserLevelTmp) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserLevelTmpDo) CreateInBatches(values []*model.WinUserLevelTmp, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserLevelTmpDo) Save(values ...*model.WinUserLevelTmp) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserLevelTmpDo) First() (*model.WinUserLevelTmp, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelTmp), nil
	}
}

func (w winUserLevelTmpDo) Take() (*model.WinUserLevelTmp, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelTmp), nil
	}
}

func (w winUserLevelTmpDo) Last() (*model.WinUserLevelTmp, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelTmp), nil
	}
}

func (w winUserLevelTmpDo) Find() ([]*model.WinUserLevelTmp, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserLevelTmp), err
}

func (w winUserLevelTmpDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevelTmp, err error) {
	buf := make([]*model.WinUserLevelTmp, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserLevelTmpDo) FindInBatches(result *[]*model.WinUserLevelTmp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserLevelTmpDo) Attrs(attrs ...field.AssignExpr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserLevelTmpDo) Assign(attrs ...field.AssignExpr) IWinUserLevelTmpDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserLevelTmpDo) Joins(fields ...field.RelationField) IWinUserLevelTmpDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserLevelTmpDo) Preload(fields ...field.RelationField) IWinUserLevelTmpDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserLevelTmpDo) FirstOrInit() (*model.WinUserLevelTmp, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelTmp), nil
	}
}

func (w winUserLevelTmpDo) FirstOrCreate() (*model.WinUserLevelTmp, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelTmp), nil
	}
}

func (w winUserLevelTmpDo) FindByPage(offset int, limit int) (result []*model.WinUserLevelTmp, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w winUserLevelTmpDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserLevelTmpDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserLevelTmpDo) Delete(models ...*model.WinUserLevelTmp) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserLevelTmpDo) withDO(do gen.Dao) *winUserLevelTmpDo {
	w.DO = *do.(*gen.DO)
	return w
}
