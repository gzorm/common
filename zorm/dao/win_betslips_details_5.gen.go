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

func newWinBetslipsDetails5(db *gorm.DB, opts ...gen.DOOption) winBetslipsDetails5 {
	_winBetslipsDetails5 := winBetslipsDetails5{}

	_winBetslipsDetails5.winBetslipsDetails5Do.UseDB(db, opts...)
	_winBetslipsDetails5.winBetslipsDetails5Do.UseModel(&model.WinBetslipsDetails5{})

	tableName := _winBetslipsDetails5.winBetslipsDetails5Do.TableName()
	_winBetslipsDetails5.ALL = field.NewAsterisk(tableName)
	_winBetslipsDetails5.ID = field.NewInt64(tableName, "id")
	_winBetslipsDetails5.XbUID = field.NewInt64(tableName, "xb_uid")
	_winBetslipsDetails5.XbUsername = field.NewString(tableName, "xb_username")
	_winBetslipsDetails5.BetJSON = field.NewString(tableName, "bet_json")
	_winBetslipsDetails5.RewardJSON = field.NewString(tableName, "reward_json")
	_winBetslipsDetails5.RefundJSON = field.NewString(tableName, "refund_json")

	_winBetslipsDetails5.fillFieldMap()

	return _winBetslipsDetails5
}

type winBetslipsDetails5 struct {
	winBetslipsDetails5Do

	ALL        field.Asterisk
	ID         field.Int64  // 主键-同注单表一致
	XbUID      field.Int64  // 对应user表id
	XbUsername field.String // 对应user表username
	BetJSON    field.String // 投注原始json
	RewardJSON field.String // 开彩原始json
	RefundJSON field.String // 退款原始json

	fieldMap map[string]field.Expr
}

func (w winBetslipsDetails5) Table(newTableName string) *winBetslipsDetails5 {
	w.winBetslipsDetails5Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetslipsDetails5) As(alias string) *winBetslipsDetails5 {
	w.winBetslipsDetails5Do.DO = *(w.winBetslipsDetails5Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetslipsDetails5) updateTableName(table string) *winBetslipsDetails5 {
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

func (w *winBetslipsDetails5) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetslipsDetails5) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 6)
	w.fieldMap["id"] = w.ID
	w.fieldMap["xb_uid"] = w.XbUID
	w.fieldMap["xb_username"] = w.XbUsername
	w.fieldMap["bet_json"] = w.BetJSON
	w.fieldMap["reward_json"] = w.RewardJSON
	w.fieldMap["refund_json"] = w.RefundJSON
}

func (w winBetslipsDetails5) clone(db *gorm.DB) winBetslipsDetails5 {
	w.winBetslipsDetails5Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetslipsDetails5) replaceDB(db *gorm.DB) winBetslipsDetails5 {
	w.winBetslipsDetails5Do.ReplaceDB(db)
	return w
}

type winBetslipsDetails5Do struct{ gen.DO }

type IWinBetslipsDetails5Do interface {
	gen.SubQuery
	Debug() IWinBetslipsDetails5Do
	WithContext(ctx context.Context) IWinBetslipsDetails5Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBetslipsDetails5Do
	WriteDB() IWinBetslipsDetails5Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBetslipsDetails5Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBetslipsDetails5Do
	Not(conds ...gen.Condition) IWinBetslipsDetails5Do
	Or(conds ...gen.Condition) IWinBetslipsDetails5Do
	Select(conds ...field.Expr) IWinBetslipsDetails5Do
	Where(conds ...gen.Condition) IWinBetslipsDetails5Do
	Order(conds ...field.Expr) IWinBetslipsDetails5Do
	Distinct(cols ...field.Expr) IWinBetslipsDetails5Do
	Omit(cols ...field.Expr) IWinBetslipsDetails5Do
	Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails5Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails5Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails5Do
	Group(cols ...field.Expr) IWinBetslipsDetails5Do
	Having(conds ...gen.Condition) IWinBetslipsDetails5Do
	Limit(limit int) IWinBetslipsDetails5Do
	Offset(offset int) IWinBetslipsDetails5Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDetails5Do
	Unscoped() IWinBetslipsDetails5Do
	Create(values ...*model.WinBetslipsDetails5) error
	CreateInBatches(values []*model.WinBetslipsDetails5, batchSize int) error
	Save(values ...*model.WinBetslipsDetails5) error
	First() (*model.WinBetslipsDetails5, error)
	Take() (*model.WinBetslipsDetails5, error)
	Last() (*model.WinBetslipsDetails5, error)
	Find() ([]*model.WinBetslipsDetails5, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslipsDetails5, err error)
	FindInBatches(result *[]*model.WinBetslipsDetails5, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBetslipsDetails5) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBetslipsDetails5Do
	Assign(attrs ...field.AssignExpr) IWinBetslipsDetails5Do
	Joins(fields ...field.RelationField) IWinBetslipsDetails5Do
	Preload(fields ...field.RelationField) IWinBetslipsDetails5Do
	FirstOrInit() (*model.WinBetslipsDetails5, error)
	FirstOrCreate() (*model.WinBetslipsDetails5, error)
	FindByPage(offset int, limit int) (result []*model.WinBetslipsDetails5, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBetslipsDetails5Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBetslipsDetails5Do) Debug() IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Debug())
}

func (w winBetslipsDetails5Do) WithContext(ctx context.Context) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetslipsDetails5Do) ReadDB() IWinBetslipsDetails5Do {
	return w.Clauses(dbresolver.Read)
}

func (w winBetslipsDetails5Do) WriteDB() IWinBetslipsDetails5Do {
	return w.Clauses(dbresolver.Write)
}

func (w winBetslipsDetails5Do) Session(config *gorm.Session) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Session(config))
}

func (w winBetslipsDetails5Do) Clauses(conds ...clause.Expression) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetslipsDetails5Do) Returning(value interface{}, columns ...string) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetslipsDetails5Do) Not(conds ...gen.Condition) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetslipsDetails5Do) Or(conds ...gen.Condition) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetslipsDetails5Do) Select(conds ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetslipsDetails5Do) Where(conds ...gen.Condition) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetslipsDetails5Do) Order(conds ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetslipsDetails5Do) Distinct(cols ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetslipsDetails5Do) Omit(cols ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetslipsDetails5Do) Join(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetslipsDetails5Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetslipsDetails5Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetslipsDetails5Do) Group(cols ...field.Expr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetslipsDetails5Do) Having(conds ...gen.Condition) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetslipsDetails5Do) Limit(limit int) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetslipsDetails5Do) Offset(offset int) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetslipsDetails5Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetslipsDetails5Do) Unscoped() IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetslipsDetails5Do) Create(values ...*model.WinBetslipsDetails5) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetslipsDetails5Do) CreateInBatches(values []*model.WinBetslipsDetails5, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetslipsDetails5Do) Save(values ...*model.WinBetslipsDetails5) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetslipsDetails5Do) First() (*model.WinBetslipsDetails5, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails5), nil
	}
}

func (w winBetslipsDetails5Do) Take() (*model.WinBetslipsDetails5, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails5), nil
	}
}

func (w winBetslipsDetails5Do) Last() (*model.WinBetslipsDetails5, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails5), nil
	}
}

func (w winBetslipsDetails5Do) Find() ([]*model.WinBetslipsDetails5, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetslipsDetails5), err
}

func (w winBetslipsDetails5Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslipsDetails5, err error) {
	buf := make([]*model.WinBetslipsDetails5, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetslipsDetails5Do) FindInBatches(result *[]*model.WinBetslipsDetails5, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetslipsDetails5Do) Attrs(attrs ...field.AssignExpr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetslipsDetails5Do) Assign(attrs ...field.AssignExpr) IWinBetslipsDetails5Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetslipsDetails5Do) Joins(fields ...field.RelationField) IWinBetslipsDetails5Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetslipsDetails5Do) Preload(fields ...field.RelationField) IWinBetslipsDetails5Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetslipsDetails5Do) FirstOrInit() (*model.WinBetslipsDetails5, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails5), nil
	}
}

func (w winBetslipsDetails5Do) FirstOrCreate() (*model.WinBetslipsDetails5, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslipsDetails5), nil
	}
}

func (w winBetslipsDetails5Do) FindByPage(offset int, limit int) (result []*model.WinBetslipsDetails5, count int64, err error) {
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

func (w winBetslipsDetails5Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetslipsDetails5Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetslipsDetails5Do) Delete(models ...*model.WinBetslipsDetails5) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetslipsDetails5Do) withDO(do gen.Dao) *winBetslipsDetails5Do {
	w.DO = *do.(*gen.DO)
	return w
}
