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

func newWinUserGoogleCode(db *gorm.DB, opts ...gen.DOOption) winUserGoogleCode {
	_winUserGoogleCode := winUserGoogleCode{}

	_winUserGoogleCode.winUserGoogleCodeDo.UseDB(db, opts...)
	_winUserGoogleCode.winUserGoogleCodeDo.UseModel(&model.WinUserGoogleCode{})

	tableName := _winUserGoogleCode.winUserGoogleCodeDo.TableName()
	_winUserGoogleCode.ALL = field.NewAsterisk(tableName)
	_winUserGoogleCode.UID = field.NewInt64(tableName, "uid")
	_winUserGoogleCode.GoogleCode = field.NewString(tableName, "google_code")
	_winUserGoogleCode.Status = field.NewInt64(tableName, "status")
	_winUserGoogleCode.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserGoogleCode.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winUserGoogleCode.fillFieldMap()

	return _winUserGoogleCode
}

// winUserGoogleCode 会员谷歌验证码
type winUserGoogleCode struct {
	winUserGoogleCodeDo

	ALL        field.Asterisk
	UID        field.Int64  // 用户id
	GoogleCode field.String // 谷歌令牌
	Status     field.Int64  // 状态:1-开启 0-禁用
	CreatedAt  field.Int64
	UpdatedAt  field.Int64

	fieldMap map[string]field.Expr
}

func (w winUserGoogleCode) Table(newTableName string) *winUserGoogleCode {
	w.winUserGoogleCodeDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserGoogleCode) As(alias string) *winUserGoogleCode {
	w.winUserGoogleCodeDo.DO = *(w.winUserGoogleCodeDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserGoogleCode) updateTableName(table string) *winUserGoogleCode {
	w.ALL = field.NewAsterisk(table)
	w.UID = field.NewInt64(table, "uid")
	w.GoogleCode = field.NewString(table, "google_code")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winUserGoogleCode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserGoogleCode) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 5)
	w.fieldMap["uid"] = w.UID
	w.fieldMap["google_code"] = w.GoogleCode
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winUserGoogleCode) clone(db *gorm.DB) winUserGoogleCode {
	w.winUserGoogleCodeDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserGoogleCode) replaceDB(db *gorm.DB) winUserGoogleCode {
	w.winUserGoogleCodeDo.ReplaceDB(db)
	return w
}

type winUserGoogleCodeDo struct{ gen.DO }

type IWinUserGoogleCodeDo interface {
	gen.SubQuery
	Debug() IWinUserGoogleCodeDo
	WithContext(ctx context.Context) IWinUserGoogleCodeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserGoogleCodeDo
	WriteDB() IWinUserGoogleCodeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserGoogleCodeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserGoogleCodeDo
	Not(conds ...gen.Condition) IWinUserGoogleCodeDo
	Or(conds ...gen.Condition) IWinUserGoogleCodeDo
	Select(conds ...field.Expr) IWinUserGoogleCodeDo
	Where(conds ...gen.Condition) IWinUserGoogleCodeDo
	Order(conds ...field.Expr) IWinUserGoogleCodeDo
	Distinct(cols ...field.Expr) IWinUserGoogleCodeDo
	Omit(cols ...field.Expr) IWinUserGoogleCodeDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserGoogleCodeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserGoogleCodeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserGoogleCodeDo
	Group(cols ...field.Expr) IWinUserGoogleCodeDo
	Having(conds ...gen.Condition) IWinUserGoogleCodeDo
	Limit(limit int) IWinUserGoogleCodeDo
	Offset(offset int) IWinUserGoogleCodeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserGoogleCodeDo
	Unscoped() IWinUserGoogleCodeDo
	Create(values ...*model.WinUserGoogleCode) error
	CreateInBatches(values []*model.WinUserGoogleCode, batchSize int) error
	Save(values ...*model.WinUserGoogleCode) error
	First() (*model.WinUserGoogleCode, error)
	Take() (*model.WinUserGoogleCode, error)
	Last() (*model.WinUserGoogleCode, error)
	Find() ([]*model.WinUserGoogleCode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserGoogleCode, err error)
	FindInBatches(result *[]*model.WinUserGoogleCode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserGoogleCode) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserGoogleCodeDo
	Assign(attrs ...field.AssignExpr) IWinUserGoogleCodeDo
	Joins(fields ...field.RelationField) IWinUserGoogleCodeDo
	Preload(fields ...field.RelationField) IWinUserGoogleCodeDo
	FirstOrInit() (*model.WinUserGoogleCode, error)
	FirstOrCreate() (*model.WinUserGoogleCode, error)
	FindByPage(offset int, limit int) (result []*model.WinUserGoogleCode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserGoogleCodeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserGoogleCodeDo) Debug() IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserGoogleCodeDo) WithContext(ctx context.Context) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserGoogleCodeDo) ReadDB() IWinUserGoogleCodeDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserGoogleCodeDo) WriteDB() IWinUserGoogleCodeDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserGoogleCodeDo) Session(config *gorm.Session) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserGoogleCodeDo) Clauses(conds ...clause.Expression) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserGoogleCodeDo) Returning(value interface{}, columns ...string) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserGoogleCodeDo) Not(conds ...gen.Condition) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserGoogleCodeDo) Or(conds ...gen.Condition) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserGoogleCodeDo) Select(conds ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserGoogleCodeDo) Where(conds ...gen.Condition) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserGoogleCodeDo) Order(conds ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserGoogleCodeDo) Distinct(cols ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserGoogleCodeDo) Omit(cols ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserGoogleCodeDo) Join(table schema.Tabler, on ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserGoogleCodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserGoogleCodeDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserGoogleCodeDo) Group(cols ...field.Expr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserGoogleCodeDo) Having(conds ...gen.Condition) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserGoogleCodeDo) Limit(limit int) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserGoogleCodeDo) Offset(offset int) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserGoogleCodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserGoogleCodeDo) Unscoped() IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserGoogleCodeDo) Create(values ...*model.WinUserGoogleCode) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserGoogleCodeDo) CreateInBatches(values []*model.WinUserGoogleCode, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserGoogleCodeDo) Save(values ...*model.WinUserGoogleCode) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserGoogleCodeDo) First() (*model.WinUserGoogleCode, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserGoogleCode), nil
	}
}

func (w winUserGoogleCodeDo) Take() (*model.WinUserGoogleCode, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserGoogleCode), nil
	}
}

func (w winUserGoogleCodeDo) Last() (*model.WinUserGoogleCode, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserGoogleCode), nil
	}
}

func (w winUserGoogleCodeDo) Find() ([]*model.WinUserGoogleCode, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserGoogleCode), err
}

func (w winUserGoogleCodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserGoogleCode, err error) {
	buf := make([]*model.WinUserGoogleCode, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserGoogleCodeDo) FindInBatches(result *[]*model.WinUserGoogleCode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserGoogleCodeDo) Attrs(attrs ...field.AssignExpr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserGoogleCodeDo) Assign(attrs ...field.AssignExpr) IWinUserGoogleCodeDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserGoogleCodeDo) Joins(fields ...field.RelationField) IWinUserGoogleCodeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserGoogleCodeDo) Preload(fields ...field.RelationField) IWinUserGoogleCodeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserGoogleCodeDo) FirstOrInit() (*model.WinUserGoogleCode, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserGoogleCode), nil
	}
}

func (w winUserGoogleCodeDo) FirstOrCreate() (*model.WinUserGoogleCode, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserGoogleCode), nil
	}
}

func (w winUserGoogleCodeDo) FindByPage(offset int, limit int) (result []*model.WinUserGoogleCode, count int64, err error) {
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

func (w winUserGoogleCodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserGoogleCodeDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserGoogleCodeDo) Delete(models ...*model.WinUserGoogleCode) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserGoogleCodeDo) withDO(do gen.Dao) *winUserGoogleCodeDo {
	w.DO = *do.(*gen.DO)
	return w
}
