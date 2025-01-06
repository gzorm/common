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

func newWinCoinRate(db *gorm.DB, opts ...gen.DOOption) winCoinRate {
	_winCoinRate := winCoinRate{}

	_winCoinRate.winCoinRateDo.UseDB(db, opts...)
	_winCoinRate.winCoinRateDo.UseModel(&model.WinCoinRate{})

	tableName := _winCoinRate.winCoinRateDo.TableName()
	_winCoinRate.ALL = field.NewAsterisk(tableName)
	_winCoinRate.ID = field.NewInt64(tableName, "id")
	_winCoinRate.OriginalCurrency = field.NewString(tableName, "original_currency")
	_winCoinRate.TransferCurrency = field.NewString(tableName, "transfer_currency")
	_winCoinRate.Rate = field.NewField(tableName, "rate")
	_winCoinRate.Status = field.NewInt64(tableName, "status")
	_winCoinRate.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinRate.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winCoinRate.OperatorName = field.NewString(tableName, "operator_name")

	_winCoinRate.fillFieldMap()

	return _winCoinRate
}

// winCoinRate 汇率表
type winCoinRate struct {
	winCoinRateDo

	ALL              field.Asterisk
	ID               field.Int64  // ID
	OriginalCurrency field.String // 原始币种
	TransferCurrency field.String // 转换币种
	Rate             field.Field  // 汇率
	Status           field.Int64  // 状态：0-关闭；1-开启
	CreatedAt        field.Int64  // 创建时间
	UpdatedAt        field.Int64  // 修改时间
	OperatorName     field.String // 操作人姓名

	fieldMap map[string]field.Expr
}

func (w winCoinRate) Table(newTableName string) *winCoinRate {
	w.winCoinRateDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinRate) As(alias string) *winCoinRate {
	w.winCoinRateDo.DO = *(w.winCoinRateDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinRate) updateTableName(table string) *winCoinRate {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.OriginalCurrency = field.NewString(table, "original_currency")
	w.TransferCurrency = field.NewString(table, "transfer_currency")
	w.Rate = field.NewField(table, "rate")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.OperatorName = field.NewString(table, "operator_name")

	w.fillFieldMap()

	return w
}

func (w *winCoinRate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinRate) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["original_currency"] = w.OriginalCurrency
	w.fieldMap["transfer_currency"] = w.TransferCurrency
	w.fieldMap["rate"] = w.Rate
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["operator_name"] = w.OperatorName
}

func (w winCoinRate) clone(db *gorm.DB) winCoinRate {
	w.winCoinRateDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinRate) replaceDB(db *gorm.DB) winCoinRate {
	w.winCoinRateDo.ReplaceDB(db)
	return w
}

type winCoinRateDo struct{ gen.DO }

type IWinCoinRateDo interface {
	gen.SubQuery
	Debug() IWinCoinRateDo
	WithContext(ctx context.Context) IWinCoinRateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinRateDo
	WriteDB() IWinCoinRateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinRateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinRateDo
	Not(conds ...gen.Condition) IWinCoinRateDo
	Or(conds ...gen.Condition) IWinCoinRateDo
	Select(conds ...field.Expr) IWinCoinRateDo
	Where(conds ...gen.Condition) IWinCoinRateDo
	Order(conds ...field.Expr) IWinCoinRateDo
	Distinct(cols ...field.Expr) IWinCoinRateDo
	Omit(cols ...field.Expr) IWinCoinRateDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinRateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinRateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinRateDo
	Group(cols ...field.Expr) IWinCoinRateDo
	Having(conds ...gen.Condition) IWinCoinRateDo
	Limit(limit int) IWinCoinRateDo
	Offset(offset int) IWinCoinRateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinRateDo
	Unscoped() IWinCoinRateDo
	Create(values ...*model.WinCoinRate) error
	CreateInBatches(values []*model.WinCoinRate, batchSize int) error
	Save(values ...*model.WinCoinRate) error
	First() (*model.WinCoinRate, error)
	Take() (*model.WinCoinRate, error)
	Last() (*model.WinCoinRate, error)
	Find() ([]*model.WinCoinRate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinRate, err error)
	FindInBatches(result *[]*model.WinCoinRate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinRate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinRateDo
	Assign(attrs ...field.AssignExpr) IWinCoinRateDo
	Joins(fields ...field.RelationField) IWinCoinRateDo
	Preload(fields ...field.RelationField) IWinCoinRateDo
	FirstOrInit() (*model.WinCoinRate, error)
	FirstOrCreate() (*model.WinCoinRate, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinRate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinRateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinRateDo) Debug() IWinCoinRateDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinRateDo) WithContext(ctx context.Context) IWinCoinRateDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinRateDo) ReadDB() IWinCoinRateDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinRateDo) WriteDB() IWinCoinRateDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinRateDo) Session(config *gorm.Session) IWinCoinRateDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinRateDo) Clauses(conds ...clause.Expression) IWinCoinRateDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinRateDo) Returning(value interface{}, columns ...string) IWinCoinRateDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinRateDo) Not(conds ...gen.Condition) IWinCoinRateDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinRateDo) Or(conds ...gen.Condition) IWinCoinRateDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinRateDo) Select(conds ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinRateDo) Where(conds ...gen.Condition) IWinCoinRateDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinRateDo) Order(conds ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinRateDo) Distinct(cols ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinRateDo) Omit(cols ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinRateDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinRateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinRateDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinRateDo) Group(cols ...field.Expr) IWinCoinRateDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinRateDo) Having(conds ...gen.Condition) IWinCoinRateDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinRateDo) Limit(limit int) IWinCoinRateDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinRateDo) Offset(offset int) IWinCoinRateDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinRateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinRateDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinRateDo) Unscoped() IWinCoinRateDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinRateDo) Create(values ...*model.WinCoinRate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinRateDo) CreateInBatches(values []*model.WinCoinRate, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinRateDo) Save(values ...*model.WinCoinRate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinRateDo) First() (*model.WinCoinRate, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRate), nil
	}
}

func (w winCoinRateDo) Take() (*model.WinCoinRate, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRate), nil
	}
}

func (w winCoinRateDo) Last() (*model.WinCoinRate, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRate), nil
	}
}

func (w winCoinRateDo) Find() ([]*model.WinCoinRate, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinRate), err
}

func (w winCoinRateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinRate, err error) {
	buf := make([]*model.WinCoinRate, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinRateDo) FindInBatches(result *[]*model.WinCoinRate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinRateDo) Attrs(attrs ...field.AssignExpr) IWinCoinRateDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinRateDo) Assign(attrs ...field.AssignExpr) IWinCoinRateDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinRateDo) Joins(fields ...field.RelationField) IWinCoinRateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinRateDo) Preload(fields ...field.RelationField) IWinCoinRateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinRateDo) FirstOrInit() (*model.WinCoinRate, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRate), nil
	}
}

func (w winCoinRateDo) FirstOrCreate() (*model.WinCoinRate, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinRate), nil
	}
}

func (w winCoinRateDo) FindByPage(offset int, limit int) (result []*model.WinCoinRate, count int64, err error) {
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

func (w winCoinRateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinRateDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinRateDo) Delete(models ...*model.WinCoinRate) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinRateDo) withDO(do gen.Dao) *winCoinRateDo {
	w.DO = *do.(*gen.DO)
	return w
}
