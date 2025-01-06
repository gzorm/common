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

func newWinFrontGameClassifyCms(db *gorm.DB, opts ...gen.DOOption) winFrontGameClassifyCms {
	_winFrontGameClassifyCms := winFrontGameClassifyCms{}

	_winFrontGameClassifyCms.winFrontGameClassifyCmsDo.UseDB(db, opts...)
	_winFrontGameClassifyCms.winFrontGameClassifyCmsDo.UseModel(&model.WinFrontGameClassifyCms{})

	tableName := _winFrontGameClassifyCms.winFrontGameClassifyCmsDo.TableName()
	_winFrontGameClassifyCms.ALL = field.NewAsterisk(tableName)
	_winFrontGameClassifyCms.ID = field.NewInt64(tableName, "id")
	_winFrontGameClassifyCms.Sort = field.NewInt64(tableName, "sort")
	_winFrontGameClassifyCms.Lang = field.NewString(tableName, "lang")
	_winFrontGameClassifyCms.CName = field.NewString(tableName, "c_name")
	_winFrontGameClassifyCms.LogoURL = field.NewString(tableName, "logo_url")
	_winFrontGameClassifyCms.Enable = field.NewInt64(tableName, "enable")
	_winFrontGameClassifyCms.CreatedAt = field.NewInt64(tableName, "created_at")
	_winFrontGameClassifyCms.CreateUser = field.NewString(tableName, "create_user")
	_winFrontGameClassifyCms.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winFrontGameClassifyCms.UpdateUser = field.NewString(tableName, "update_user")
	_winFrontGameClassifyCms.OperatorName = field.NewString(tableName, "operator_name")

	_winFrontGameClassifyCms.fillFieldMap()

	return _winFrontGameClassifyCms
}

// winFrontGameClassifyCms 首页游戏分类页面配置表
type winFrontGameClassifyCms struct {
	winFrontGameClassifyCmsDo

	ALL          field.Asterisk
	ID           field.Int64
	Sort         field.Int64  // 排序
	Lang         field.String // 语言
	CName        field.String // 分类名称
	LogoURL      field.String // 图片地址
	Enable       field.Int64  // 状态:0关闭,1启用
	CreatedAt    field.Int64  // 创建时间
	CreateUser   field.String // 创建人
	UpdatedAt    field.Int64  // 修改人
	UpdateUser   field.String // 修改人
	OperatorName field.String // 操作人姓名

	fieldMap map[string]field.Expr
}

func (w winFrontGameClassifyCms) Table(newTableName string) *winFrontGameClassifyCms {
	w.winFrontGameClassifyCmsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winFrontGameClassifyCms) As(alias string) *winFrontGameClassifyCms {
	w.winFrontGameClassifyCmsDo.DO = *(w.winFrontGameClassifyCmsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winFrontGameClassifyCms) updateTableName(table string) *winFrontGameClassifyCms {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Sort = field.NewInt64(table, "sort")
	w.Lang = field.NewString(table, "lang")
	w.CName = field.NewString(table, "c_name")
	w.LogoURL = field.NewString(table, "logo_url")
	w.Enable = field.NewInt64(table, "enable")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.CreateUser = field.NewString(table, "create_user")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.UpdateUser = field.NewString(table, "update_user")
	w.OperatorName = field.NewString(table, "operator_name")

	w.fillFieldMap()

	return w
}

func (w *winFrontGameClassifyCms) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winFrontGameClassifyCms) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["lang"] = w.Lang
	w.fieldMap["c_name"] = w.CName
	w.fieldMap["logo_url"] = w.LogoURL
	w.fieldMap["enable"] = w.Enable
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["create_user"] = w.CreateUser
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["update_user"] = w.UpdateUser
	w.fieldMap["operator_name"] = w.OperatorName
}

func (w winFrontGameClassifyCms) clone(db *gorm.DB) winFrontGameClassifyCms {
	w.winFrontGameClassifyCmsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winFrontGameClassifyCms) replaceDB(db *gorm.DB) winFrontGameClassifyCms {
	w.winFrontGameClassifyCmsDo.ReplaceDB(db)
	return w
}

type winFrontGameClassifyCmsDo struct{ gen.DO }

type IWinFrontGameClassifyCmsDo interface {
	gen.SubQuery
	Debug() IWinFrontGameClassifyCmsDo
	WithContext(ctx context.Context) IWinFrontGameClassifyCmsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinFrontGameClassifyCmsDo
	WriteDB() IWinFrontGameClassifyCmsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinFrontGameClassifyCmsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinFrontGameClassifyCmsDo
	Not(conds ...gen.Condition) IWinFrontGameClassifyCmsDo
	Or(conds ...gen.Condition) IWinFrontGameClassifyCmsDo
	Select(conds ...field.Expr) IWinFrontGameClassifyCmsDo
	Where(conds ...gen.Condition) IWinFrontGameClassifyCmsDo
	Order(conds ...field.Expr) IWinFrontGameClassifyCmsDo
	Distinct(cols ...field.Expr) IWinFrontGameClassifyCmsDo
	Omit(cols ...field.Expr) IWinFrontGameClassifyCmsDo
	Join(table schema.Tabler, on ...field.Expr) IWinFrontGameClassifyCmsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinFrontGameClassifyCmsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinFrontGameClassifyCmsDo
	Group(cols ...field.Expr) IWinFrontGameClassifyCmsDo
	Having(conds ...gen.Condition) IWinFrontGameClassifyCmsDo
	Limit(limit int) IWinFrontGameClassifyCmsDo
	Offset(offset int) IWinFrontGameClassifyCmsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinFrontGameClassifyCmsDo
	Unscoped() IWinFrontGameClassifyCmsDo
	Create(values ...*model.WinFrontGameClassifyCms) error
	CreateInBatches(values []*model.WinFrontGameClassifyCms, batchSize int) error
	Save(values ...*model.WinFrontGameClassifyCms) error
	First() (*model.WinFrontGameClassifyCms, error)
	Take() (*model.WinFrontGameClassifyCms, error)
	Last() (*model.WinFrontGameClassifyCms, error)
	Find() ([]*model.WinFrontGameClassifyCms, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinFrontGameClassifyCms, err error)
	FindInBatches(result *[]*model.WinFrontGameClassifyCms, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinFrontGameClassifyCms) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinFrontGameClassifyCmsDo
	Assign(attrs ...field.AssignExpr) IWinFrontGameClassifyCmsDo
	Joins(fields ...field.RelationField) IWinFrontGameClassifyCmsDo
	Preload(fields ...field.RelationField) IWinFrontGameClassifyCmsDo
	FirstOrInit() (*model.WinFrontGameClassifyCms, error)
	FirstOrCreate() (*model.WinFrontGameClassifyCms, error)
	FindByPage(offset int, limit int) (result []*model.WinFrontGameClassifyCms, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinFrontGameClassifyCmsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winFrontGameClassifyCmsDo) Debug() IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Debug())
}

func (w winFrontGameClassifyCmsDo) WithContext(ctx context.Context) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winFrontGameClassifyCmsDo) ReadDB() IWinFrontGameClassifyCmsDo {
	return w.Clauses(dbresolver.Read)
}

func (w winFrontGameClassifyCmsDo) WriteDB() IWinFrontGameClassifyCmsDo {
	return w.Clauses(dbresolver.Write)
}

func (w winFrontGameClassifyCmsDo) Session(config *gorm.Session) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Session(config))
}

func (w winFrontGameClassifyCmsDo) Clauses(conds ...clause.Expression) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winFrontGameClassifyCmsDo) Returning(value interface{}, columns ...string) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winFrontGameClassifyCmsDo) Not(conds ...gen.Condition) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winFrontGameClassifyCmsDo) Or(conds ...gen.Condition) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winFrontGameClassifyCmsDo) Select(conds ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winFrontGameClassifyCmsDo) Where(conds ...gen.Condition) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winFrontGameClassifyCmsDo) Order(conds ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winFrontGameClassifyCmsDo) Distinct(cols ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winFrontGameClassifyCmsDo) Omit(cols ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winFrontGameClassifyCmsDo) Join(table schema.Tabler, on ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winFrontGameClassifyCmsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winFrontGameClassifyCmsDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winFrontGameClassifyCmsDo) Group(cols ...field.Expr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winFrontGameClassifyCmsDo) Having(conds ...gen.Condition) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winFrontGameClassifyCmsDo) Limit(limit int) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winFrontGameClassifyCmsDo) Offset(offset int) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winFrontGameClassifyCmsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winFrontGameClassifyCmsDo) Unscoped() IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winFrontGameClassifyCmsDo) Create(values ...*model.WinFrontGameClassifyCms) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winFrontGameClassifyCmsDo) CreateInBatches(values []*model.WinFrontGameClassifyCms, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winFrontGameClassifyCmsDo) Save(values ...*model.WinFrontGameClassifyCms) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winFrontGameClassifyCmsDo) First() (*model.WinFrontGameClassifyCms, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFrontGameClassifyCms), nil
	}
}

func (w winFrontGameClassifyCmsDo) Take() (*model.WinFrontGameClassifyCms, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFrontGameClassifyCms), nil
	}
}

func (w winFrontGameClassifyCmsDo) Last() (*model.WinFrontGameClassifyCms, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFrontGameClassifyCms), nil
	}
}

func (w winFrontGameClassifyCmsDo) Find() ([]*model.WinFrontGameClassifyCms, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinFrontGameClassifyCms), err
}

func (w winFrontGameClassifyCmsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinFrontGameClassifyCms, err error) {
	buf := make([]*model.WinFrontGameClassifyCms, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winFrontGameClassifyCmsDo) FindInBatches(result *[]*model.WinFrontGameClassifyCms, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winFrontGameClassifyCmsDo) Attrs(attrs ...field.AssignExpr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winFrontGameClassifyCmsDo) Assign(attrs ...field.AssignExpr) IWinFrontGameClassifyCmsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winFrontGameClassifyCmsDo) Joins(fields ...field.RelationField) IWinFrontGameClassifyCmsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winFrontGameClassifyCmsDo) Preload(fields ...field.RelationField) IWinFrontGameClassifyCmsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winFrontGameClassifyCmsDo) FirstOrInit() (*model.WinFrontGameClassifyCms, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFrontGameClassifyCms), nil
	}
}

func (w winFrontGameClassifyCmsDo) FirstOrCreate() (*model.WinFrontGameClassifyCms, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFrontGameClassifyCms), nil
	}
}

func (w winFrontGameClassifyCmsDo) FindByPage(offset int, limit int) (result []*model.WinFrontGameClassifyCms, count int64, err error) {
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

func (w winFrontGameClassifyCmsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winFrontGameClassifyCmsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winFrontGameClassifyCmsDo) Delete(models ...*model.WinFrontGameClassifyCms) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winFrontGameClassifyCmsDo) withDO(do gen.Dao) *winFrontGameClassifyCmsDo {
	w.DO = *do.(*gen.DO)
	return w
}
