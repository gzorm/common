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

func newWinBetslipsDetails1(db *gorm.DB, opts ...gen.DOOption) winBetslipsDetails1 {
	_winBetslipsDetails1 := winBetslipsDetails1{}

	_winBetslipsDetails1.winBetslipsDetails1Do.UseDB(db, opts...)
	_winBetslipsDetails1.winBetslipsDetails1Do.UseModel(&model.WinBetslipsDetails1{})

	tableName := _winBetslipsDetails1.winBetslipsDetails1Do.TableName()
	_winBetslipsDetails1.ALL = field.NewAsterisk(tableName)
	_winBetslipsDetails1.ID = field.NewInt64(tableName, "id")
	_winBetslipsDetails1.XbUID = field.NewInt64(tableName, "xb_uid")
	_winBetslipsDetails1.XbUsername = field.NewString(tableName, "xb_username")
	_winBetslipsDetails1.BetJSON = field.NewString(tableName, "bet_json")
	_winBetslipsDetails1.RewardJSON = field.NewString(tableName, "reward_json")
	_winBetslipsDetails1.RefundJSON = field.NewString(tableName, "refund_json")

	_winBetslipsDetails1.fillFieldMap()

	return _winBetslipsDetails1
}

type winBetslipsDetails1 struct {
	winBetslipsDetails1Do

	ALL        field.Asterisk
	ID         field.Int64  // 主键-同注单表一致
	XbUID      field.Int64  // 对应user表id
	XbUsername field.String // 对应user表username
	BetJSON    field.String // 投注原始json
	RewardJSON field.String // 开彩原始json
	RefundJSON field.String // 退款原始json

	fieldMap map[string]field.Expr
}

func (w winBetslipsDetails1) Table(newTableName string) *winBetslipsDetails1 {
	w.winBetslipsDetails1Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetslipsDetails1) As(alias string) *winBetslipsDetails1 {
	w.winBetslipsDetails1Do.DO = *(w.winBetslipsDetails1Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetslipsDetails1) updateTableName(table string) *winBetslipsDetails1 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.XbUID = field.NewInt64(table, "xb_uid")
	w.XbUsername = field.NewString(table, "xb_username")
	w.BetJSON = field.NewString(table, "bet_json")
	w.RewardJSON = field.NewString(table, "reward_json")
	w.RefundJSON = field.NewString(table, "refund_json")

	w.fillFieldMap()

	return w
}

func (w *winBetslipsDetails1) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetslipsDetails1) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["xb_uid"] = w.XbUID
	w.fieldMap["xb_username"] = w.XbUsername
	w.fieldMap["bet_json"] = w.BetJSON
	w.fieldMap["reward_json"] = w.RewardJSON
	w.fieldMap["refund_json"] = w.RefundJSON
}

func (w winBetslipsDetails1) clone(db *gorm.DB) winBetslipsDetails1 {
	w.winBetslipsDetails1Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetslipsDetails1) replaceDB(db *gorm.DB) winBetslipsDetails1 {
	w.winBetslipsDetails1Do.ReplaceDB(db)
	return w
}

type winBetslipsDetails1Do struct{ gen.DO }

type IWinBetslipsDetails1Do interface {
	gen.SubQuery
	Debug() IWinBetslipsDetails1Do
	WithContext(ctx context.Context) IWinBetslipsDetails1Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBetslipsDetails1Do
	WriteDB() IWinBetslipsDetails1Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBetslipsDetails1Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBetslipsDetails1Do
	Not(conds ...gen.Condition) IWinBetslipsDetails1Do
	Or(conds ...gen.Condition) IWinBetslipsDetails1Do
	Select(conds ...field.Expr) IWinBetslipsDetails1Do
	Where(conds ...gen.Condition) IWinBetslipsDetails1Do
	Order(conds ...field.Expr) IWinBetslipsDetails1Do
	Distinct(cols ...field.Expr) IWinBetslipsDetails1Do
	Omit(cols ...field.Expr) IWinBetslipsDetails1Do
	Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails1Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails1Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails1Do
	Group(cols ...field.Expr) IWinBetslipsDetails1Do
	Having(conds ...gen.Condition) IWinBetslipsDetails1Do
	Limit(limit int) IWinBetslipsDetails1Do
	Offset(offset int) IWinBetslipsDetails1Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDetails1Do
	Unscoped() IWinBetslipsDetails1Do
	Create(values ...*model.WinBetslipsDetails1) error
	CreateInBatches(values []*model.WinBetslipsDetails1, batchSize int) error
	Save(values ...*model.WinBetslipsDetails1) error
	First() (*model.WinBetslipsDetails1, error)
	Take() (*model.WinBetslipsDetails1, error)
	Last() (*model.WinBetslipsDetails1, error)
	Find() ([]*model.WinBetslipsDetails1, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslipsDetails1, err error)
	FindInBatches(result *[]*model.WinBetslipsDetails1, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBetslipsDetails1) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBetslipsDetails1Do
	Assign(attrs ...field.AssignExpr) IWinBetslipsDetails1Do
	Joins(fields ...field.RelationField) IWinBetslipsDetails1Do
	Preload(fields ...field.RelationField) IWinBetslipsDetails1Do
	FirstOrInit() (*model.WinBetslipsDetails1, error)
	FirstOrCreate() (*model.WinBetslipsDetails1, error)
	FindByPage(offset int, limit int) (result []*model.WinBetslipsDetails1, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBetslipsDetails1Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBetslipsDetails1Do) Debug() IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Debug())
}

func (w winBetslipsDetails1Do) WithContext(ctx context.Context) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetslipsDetails1Do) ReadDB() IWinBetslipsDetails1Do {
	return w.Clauses(dbresolver.Read)
}

func (w winBetslipsDetails1Do) WriteDB() IWinBetslipsDetails1Do {
	return w.Clauses(dbresolver.Write)
}

func (w winBetslipsDetails1Do) Session(config *gorm.Session) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Session(config))
}

func (w winBetslipsDetails1Do) Clauses(conds ...clause.Expression) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetslipsDetails1Do) Returning(value interface{}, columns ...string) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetslipsDetails1Do) Not(conds ...gen.Condition) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetslipsDetails1Do) Or(conds ...gen.Condition) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetslipsDetails1Do) Select(conds ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetslipsDetails1Do) Where(conds ...gen.Condition) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetslipsDetails1Do) Order(conds ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetslipsDetails1Do) Distinct(cols ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetslipsDetails1Do) Omit(cols ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetslipsDetails1Do) Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetslipsDetails1Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetslipsDetails1Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetslipsDetails1Do) Group(cols ...field.Expr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetslipsDetails1Do) Having(conds ...gen.Condition) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetslipsDetails1Do) Limit(limit int) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetslipsDetails1Do) Offset(offset int) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetslipsDetails1Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetslipsDetails1Do) Unscoped() IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetslipsDetails1Do) Create(values ...*model.WinBetslipsDetails1) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetslipsDetails1Do) CreateInBatches(values []*model.WinBetslipsDetails1, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetslipsDetails1Do) Save(values ...*model.WinBetslipsDetails1) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetslipsDetails1Do) First() (*model.WinBetslipsDetails1, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails1), nil
	}
}

func (w winBetslipsDetails1Do) Take() (*model.WinBetslipsDetails1, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails1), nil
	}
}

func (w winBetslipsDetails1Do) Last() (*model.WinBetslipsDetails1, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails1), nil
	}
}

func (w winBetslipsDetails1Do) Find() ([]*model.WinBetslipsDetails1, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetslipsDetails1), err
}

func (w winBetslipsDetails1Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslipsDetails1, err error) {
	buf := make([]*model.WinBetslipsDetails1, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetslipsDetails1Do) FindInBatches(result *[]*model.WinBetslipsDetails1, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetslipsDetails1Do) Attrs(attrs ...field.AssignExpr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetslipsDetails1Do) Assign(attrs ...field.AssignExpr) IWinBetslipsDetails1Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetslipsDetails1Do) Joins(fields ...field.RelationField) IWinBetslipsDetails1Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetslipsDetails1Do) Preload(fields ...field.RelationField) IWinBetslipsDetails1Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetslipsDetails1Do) FirstOrInit() (*model.WinBetslipsDetails1, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails1), nil
	}
}

func (w winBetslipsDetails1Do) FirstOrCreate() (*model.WinBetslipsDetails1, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails1), nil
	}
}

func (w winBetslipsDetails1Do) FindByPage(offset int, limit int) (result []*model.WinBetslipsDetails1, count int64, err error) {
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

func (w winBetslipsDetails1Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetslipsDetails1Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetslipsDetails1Do) Delete(models ...*model.WinBetslipsDetails1) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetslipsDetails1Do) withDO(do gen.Dao) *winBetslipsDetails1Do {
	w.DO = *do.(*gen.DO)
	return w
}
