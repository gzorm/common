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

func newWinCoinCommissionM2(db *gorm.DB, opts ...gen.DOOption) winCoinCommissionM2 {
	_winCoinCommissionM2 := winCoinCommissionM2{}

	_winCoinCommissionM2.winCoinCommissionM2Do.UseDB(db, opts...)
	_winCoinCommissionM2.winCoinCommissionM2Do.UseModel(&model.WinCoinCommissionM2{})

	tableName := _winCoinCommissionM2.winCoinCommissionM2Do.TableName()
	_winCoinCommissionM2.ALL = field.NewAsterisk(tableName)
	_winCoinCommissionM2.ID = field.NewInt64(tableName, "id")
	_winCoinCommissionM2.UID = field.NewInt64(tableName, "uid")
	_winCoinCommissionM2.Username = field.NewString(tableName, "username")
	_winCoinCommissionM2.AgentLevel = field.NewInt64(tableName, "agent_level")
	_winCoinCommissionM2.Riqi = field.NewInt64(tableName, "riqi")
	_winCoinCommissionM2.Coin = field.NewField(tableName, "coin")
	_winCoinCommissionM2.SubUID = field.NewInt64(tableName, "sub_uid")
	_winCoinCommissionM2.SubUsername = field.NewString(tableName, "sub_username")
	_winCoinCommissionM2.SubBetTrunover = field.NewField(tableName, "sub_bet_trunover")
	_winCoinCommissionM2.Rate = field.NewField(tableName, "rate")
	_winCoinCommissionM2.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinCommissionM2.Status = field.NewInt64(tableName, "status")
	_winCoinCommissionM2.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinCommissionM2.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winCoinCommissionM2.fillFieldMap()

	return _winCoinCommissionM2
}

// winCoinCommissionM2 佣金表
type winCoinCommissionM2 struct {
	winCoinCommissionM2Do

	ALL            field.Asterisk
	ID             field.Int64
	UID            field.Int64  // 代理UID
	Username       field.String // 用户名
	AgentLevel     field.Int64  // 代理层级
	Riqi           field.Int64  // 佣金时间
	Coin           field.Field  // 佣金金额
	SubUID         field.Int64  // 下级UID
	SubUsername    field.String // 下级用户名
	SubBetTrunover field.Field  // 下级流水总额
	Rate           field.Field  // 佣金比例
	CoinBefore     field.Field  // 即时余额
	Status         field.Int64  // 状态:0-未发放 1-已发放
	CreatedAt      field.Int64
	UpdatedAt      field.Int64

	fieldMap map[string]field.Expr
}

func (w winCoinCommissionM2) Table(newTableName string) *winCoinCommissionM2 {
	w.winCoinCommissionM2Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinCommissionM2) As(alias string) *winCoinCommissionM2 {
	w.winCoinCommissionM2Do.DO = *(w.winCoinCommissionM2Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinCommissionM2) updateTableName(table string) *winCoinCommissionM2 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.AgentLevel = field.NewInt64(table, "agent_level")
	w.Riqi = field.NewInt64(table, "riqi")
	w.Coin = field.NewField(table, "coin")
	w.SubUID = field.NewInt64(table, "sub_uid")
	w.SubUsername = field.NewString(table, "sub_username")
	w.SubBetTrunover = field.NewField(table, "sub_bet_trunover")
	w.Rate = field.NewField(table, "rate")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winCoinCommissionM2) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinCommissionM2) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 14)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["agent_level"] = w.AgentLevel
	w.fieldMap["riqi"] = w.Riqi
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["sub_uid"] = w.SubUID
	w.fieldMap["sub_username"] = w.SubUsername
	w.fieldMap["sub_bet_trunover"] = w.SubBetTrunover
	w.fieldMap["rate"] = w.Rate
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winCoinCommissionM2) clone(db *gorm.DB) winCoinCommissionM2 {
	w.winCoinCommissionM2Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinCommissionM2) replaceDB(db *gorm.DB) winCoinCommissionM2 {
	w.winCoinCommissionM2Do.ReplaceDB(db)
	return w
}

type winCoinCommissionM2Do struct{ gen.DO }

type IWinCoinCommissionM2Do interface {
	gen.SubQuery
	Debug() IWinCoinCommissionM2Do
	WithContext(ctx context.Context) IWinCoinCommissionM2Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinCommissionM2Do
	WriteDB() IWinCoinCommissionM2Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinCommissionM2Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinCommissionM2Do
	Not(conds ...gen.Condition) IWinCoinCommissionM2Do
	Or(conds ...gen.Condition) IWinCoinCommissionM2Do
	Select(conds ...field.Expr) IWinCoinCommissionM2Do
	Where(conds ...gen.Condition) IWinCoinCommissionM2Do
	Order(conds ...field.Expr) IWinCoinCommissionM2Do
	Distinct(cols ...field.Expr) IWinCoinCommissionM2Do
	Omit(cols ...field.Expr) IWinCoinCommissionM2Do
	Join(table schema.Tabler, on ...field.Expr) IWinCoinCommissionM2Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionM2Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionM2Do
	Group(cols ...field.Expr) IWinCoinCommissionM2Do
	Having(conds ...gen.Condition) IWinCoinCommissionM2Do
	Limit(limit int) IWinCoinCommissionM2Do
	Offset(offset int) IWinCoinCommissionM2Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinCommissionM2Do
	Unscoped() IWinCoinCommissionM2Do
	Create(values ...*model.WinCoinCommissionM2) error
	CreateInBatches(values []*model.WinCoinCommissionM2, batchSize int) error
	Save(values ...*model.WinCoinCommissionM2) error
	First() (*model.WinCoinCommissionM2, error)
	Take() (*model.WinCoinCommissionM2, error)
	Last() (*model.WinCoinCommissionM2, error)
	Find() ([]*model.WinCoinCommissionM2, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinCommissionM2, err error)
	FindInBatches(result *[]*model.WinCoinCommissionM2, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinCommissionM2) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinCommissionM2Do
	Assign(attrs ...field.AssignExpr) IWinCoinCommissionM2Do
	Joins(fields ...field.RelationField) IWinCoinCommissionM2Do
	Preload(fields ...field.RelationField) IWinCoinCommissionM2Do
	FirstOrInit() (*model.WinCoinCommissionM2, error)
	FirstOrCreate() (*model.WinCoinCommissionM2, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinCommissionM2, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinCommissionM2Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinCommissionM2Do) Debug() IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Debug())
}

func (w winCoinCommissionM2Do) WithContext(ctx context.Context) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinCommissionM2Do) ReadDB() IWinCoinCommissionM2Do {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinCommissionM2Do) WriteDB() IWinCoinCommissionM2Do {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinCommissionM2Do) Session(config *gorm.Session) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinCommissionM2Do) Clauses(conds ...clause.Expression) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinCommissionM2Do) Returning(value interface{}, columns ...string) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinCommissionM2Do) Not(conds ...gen.Condition) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinCommissionM2Do) Or(conds ...gen.Condition) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinCommissionM2Do) Select(conds ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinCommissionM2Do) Where(conds ...gen.Condition) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinCommissionM2Do) Order(conds ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinCommissionM2Do) Distinct(cols ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinCommissionM2Do) Omit(cols ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinCommissionM2Do) Join(table schema.Tabler, on ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinCommissionM2Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinCommissionM2Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinCommissionM2Do) Group(cols ...field.Expr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinCommissionM2Do) Having(conds ...gen.Condition) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinCommissionM2Do) Limit(limit int) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinCommissionM2Do) Offset(offset int) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinCommissionM2Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinCommissionM2Do) Unscoped() IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinCommissionM2Do) Create(values ...*model.WinCoinCommissionM2) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinCommissionM2Do) CreateInBatches(values []*model.WinCoinCommissionM2, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinCommissionM2Do) Save(values ...*model.WinCoinCommissionM2) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinCommissionM2Do) First() (*model.WinCoinCommissionM2, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionM2), nil
	}
}

func (w winCoinCommissionM2Do) Take() (*model.WinCoinCommissionM2, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionM2), nil
	}
}

func (w winCoinCommissionM2Do) Last() (*model.WinCoinCommissionM2, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionM2), nil
	}
}

func (w winCoinCommissionM2Do) Find() ([]*model.WinCoinCommissionM2, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinCommissionM2), err
}

func (w winCoinCommissionM2Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinCommissionM2, err error) {
	buf := make([]*model.WinCoinCommissionM2, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinCommissionM2Do) FindInBatches(result *[]*model.WinCoinCommissionM2, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinCommissionM2Do) Attrs(attrs ...field.AssignExpr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinCommissionM2Do) Assign(attrs ...field.AssignExpr) IWinCoinCommissionM2Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinCommissionM2Do) Joins(fields ...field.RelationField) IWinCoinCommissionM2Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinCommissionM2Do) Preload(fields ...field.RelationField) IWinCoinCommissionM2Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinCommissionM2Do) FirstOrInit() (*model.WinCoinCommissionM2, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionM2), nil
	}
}

func (w winCoinCommissionM2Do) FirstOrCreate() (*model.WinCoinCommissionM2, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinCommissionM2), nil
	}
}

func (w winCoinCommissionM2Do) FindByPage(offset int, limit int) (result []*model.WinCoinCommissionM2, count int64, err error) {
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

func (w winCoinCommissionM2Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinCommissionM2Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinCommissionM2Do) Delete(models ...*model.WinCoinCommissionM2) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinCommissionM2Do) withDO(do gen.Dao) *winCoinCommissionM2Do {
	w.DO = *do.(*gen.DO)
	return w
}