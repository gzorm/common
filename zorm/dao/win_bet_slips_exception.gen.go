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

func newWinBetSlipsException(db *gorm.DB, opts ...gen.DOOption) winBetSlipsException {
	_winBetSlipsException := winBetSlipsException{}

	_winBetSlipsException.winBetSlipsExceptionDo.UseDB(db, opts...)
	_winBetSlipsException.winBetSlipsExceptionDo.UseModel(&model.WinBetSlipsException{})

	tableName := _winBetSlipsException.winBetSlipsExceptionDo.TableName()
	_winBetSlipsException.ALL = field.NewAsterisk(tableName)
	_winBetSlipsException.ID = field.NewInt64(tableName, "id")
	_winBetSlipsException.GameListID = field.NewInt64(tableName, "game_list_id")
	_winBetSlipsException.Request = field.NewString(tableName, "request")
	_winBetSlipsException.Category = field.NewInt64(tableName, "category")
	_winBetSlipsException.Info = field.NewString(tableName, "info")
	_winBetSlipsException.Status = field.NewInt64(tableName, "status")
	_winBetSlipsException.CreatedAt = field.NewInt64(tableName, "created_at")
	_winBetSlipsException.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winBetSlipsException.fillFieldMap()

	return _winBetSlipsException
}

// winBetSlipsException 拉单异常表
type winBetSlipsException struct {
	winBetSlipsExceptionDo

	ALL        field.Asterisk
	ID         field.Int64
	GameListID field.Int64  // 游戏表ID
	Request    field.String // 请求参数
	Category   field.Int64  // 状态:0-三方异常 1-数据异常
	Info       field.String // 异常信息:三方-返回数据 数据-异常处理
	Status     field.Int64  // 状态:0-未处理 1-处理成功 2-处理失败
	CreatedAt  field.Int64
	UpdatedAt  field.Int64

	fieldMap map[string]field.Expr
}

func (w winBetSlipsException) Table(newTableName string) *winBetSlipsException {
	w.winBetSlipsExceptionDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetSlipsException) As(alias string) *winBetSlipsException {
	w.winBetSlipsExceptionDo.DO = *(w.winBetSlipsExceptionDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetSlipsException) updateTableName(table string) *winBetSlipsException {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.GameListID = field.NewInt64(table, "game_list_id")
	w.Request = field.NewString(table, "request")
	w.Category = field.NewInt64(table, "category")
	w.Info = field.NewString(table, "info")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winBetSlipsException) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetSlipsException) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["game_list_id"] = w.GameListID
	w.fieldMap["request"] = w.Request
	w.fieldMap["category"] = w.Category
	w.fieldMap["info"] = w.Info
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winBetSlipsException) clone(db *gorm.DB) winBetSlipsException {
	w.winBetSlipsExceptionDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetSlipsException) replaceDB(db *gorm.DB) winBetSlipsException {
	w.winBetSlipsExceptionDo.ReplaceDB(db)
	return w
}

type winBetSlipsExceptionDo struct{ gen.DO }

type IWinBetSlipsExceptionDo interface {
	gen.SubQuery
	Debug() IWinBetSlipsExceptionDo
	WithContext(ctx context.Context) IWinBetSlipsExceptionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBetSlipsExceptionDo
	WriteDB() IWinBetSlipsExceptionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBetSlipsExceptionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBetSlipsExceptionDo
	Not(conds ...gen.Condition) IWinBetSlipsExceptionDo
	Or(conds ...gen.Condition) IWinBetSlipsExceptionDo
	Select(conds ...field.Expr) IWinBetSlipsExceptionDo
	Where(conds ...gen.Condition) IWinBetSlipsExceptionDo
	Order(conds ...field.Expr) IWinBetSlipsExceptionDo
	Distinct(cols ...field.Expr) IWinBetSlipsExceptionDo
	Omit(cols ...field.Expr) IWinBetSlipsExceptionDo
	Join(table schema.Tabler, on ...field.Expr) IWinBetSlipsExceptionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetSlipsExceptionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBetSlipsExceptionDo
	Group(cols ...field.Expr) IWinBetSlipsExceptionDo
	Having(conds ...gen.Condition) IWinBetSlipsExceptionDo
	Limit(limit int) IWinBetSlipsExceptionDo
	Offset(offset int) IWinBetSlipsExceptionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetSlipsExceptionDo
	Unscoped() IWinBetSlipsExceptionDo
	Create(values ...*model.WinBetSlipsException) error
	CreateInBatches(values []*model.WinBetSlipsException, batchSize int) error
	Save(values ...*model.WinBetSlipsException) error
	First() (*model.WinBetSlipsException, error)
	Take() (*model.WinBetSlipsException, error)
	Last() (*model.WinBetSlipsException, error)
	Find() ([]*model.WinBetSlipsException, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetSlipsException, err error)
	FindInBatches(result *[]*model.WinBetSlipsException, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBetSlipsException) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBetSlipsExceptionDo
	Assign(attrs ...field.AssignExpr) IWinBetSlipsExceptionDo
	Joins(fields ...field.RelationField) IWinBetSlipsExceptionDo
	Preload(fields ...field.RelationField) IWinBetSlipsExceptionDo
	FirstOrInit() (*model.WinBetSlipsException, error)
	FirstOrCreate() (*model.WinBetSlipsException, error)
	FindByPage(offset int, limit int) (result []*model.WinBetSlipsException, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBetSlipsExceptionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBetSlipsExceptionDo) Debug() IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Debug())
}

func (w winBetSlipsExceptionDo) WithContext(ctx context.Context) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetSlipsExceptionDo) ReadDB() IWinBetSlipsExceptionDo {
	return w.Clauses(dbresolver.Read)
}

func (w winBetSlipsExceptionDo) WriteDB() IWinBetSlipsExceptionDo {
	return w.Clauses(dbresolver.Write)
}

func (w winBetSlipsExceptionDo) Session(config *gorm.Session) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Session(config))
}

func (w winBetSlipsExceptionDo) Clauses(conds ...clause.Expression) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetSlipsExceptionDo) Returning(value interface{}, columns ...string) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetSlipsExceptionDo) Not(conds ...gen.Condition) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetSlipsExceptionDo) Or(conds ...gen.Condition) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetSlipsExceptionDo) Select(conds ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetSlipsExceptionDo) Where(conds ...gen.Condition) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetSlipsExceptionDo) Order(conds ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetSlipsExceptionDo) Distinct(cols ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetSlipsExceptionDo) Omit(cols ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetSlipsExceptionDo) Join(table schema.Tabler, on ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetSlipsExceptionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetSlipsExceptionDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetSlipsExceptionDo) Group(cols ...field.Expr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetSlipsExceptionDo) Having(conds ...gen.Condition) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetSlipsExceptionDo) Limit(limit int) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetSlipsExceptionDo) Offset(offset int) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetSlipsExceptionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetSlipsExceptionDo) Unscoped() IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetSlipsExceptionDo) Create(values ...*model.WinBetSlipsException) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetSlipsExceptionDo) CreateInBatches(values []*model.WinBetSlipsException, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetSlipsExceptionDo) Save(values ...*model.WinBetSlipsException) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetSlipsExceptionDo) First() (*model.WinBetSlipsException, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetSlipsException), nil
	}
}

func (w winBetSlipsExceptionDo) Take() (*model.WinBetSlipsException, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetSlipsException), nil
	}
}

func (w winBetSlipsExceptionDo) Last() (*model.WinBetSlipsException, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetSlipsException), nil
	}
}

func (w winBetSlipsExceptionDo) Find() ([]*model.WinBetSlipsException, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetSlipsException), err
}

func (w winBetSlipsExceptionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetSlipsException, err error) {
	buf := make([]*model.WinBetSlipsException, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetSlipsExceptionDo) FindInBatches(result *[]*model.WinBetSlipsException, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetSlipsExceptionDo) Attrs(attrs ...field.AssignExpr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetSlipsExceptionDo) Assign(attrs ...field.AssignExpr) IWinBetSlipsExceptionDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetSlipsExceptionDo) Joins(fields ...field.RelationField) IWinBetSlipsExceptionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetSlipsExceptionDo) Preload(fields ...field.RelationField) IWinBetSlipsExceptionDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetSlipsExceptionDo) FirstOrInit() (*model.WinBetSlipsException, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetSlipsException), nil
	}
}

func (w winBetSlipsExceptionDo) FirstOrCreate() (*model.WinBetSlipsException, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetSlipsException), nil
	}
}

func (w winBetSlipsExceptionDo) FindByPage(offset int, limit int) (result []*model.WinBetSlipsException, count int64, err error) {
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

func (w winBetSlipsExceptionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetSlipsExceptionDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetSlipsExceptionDo) Delete(models ...*model.WinBetSlipsException) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetSlipsExceptionDo) withDO(do gen.Dao) *winBetSlipsExceptionDo {
	w.DO = *do.(*gen.DO)
	return w
}
