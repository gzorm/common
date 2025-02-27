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

func newWinHelpType(db *gorm.DB, opts ...gen.DOOption) winHelpType {
	_winHelpType := winHelpType{}

	_winHelpType.winHelpTypeDo.UseDB(db, opts...)
	_winHelpType.winHelpTypeDo.UseModel(&model.WinHelpType{})

	tableName := _winHelpType.winHelpTypeDo.TableName()
	_winHelpType.ALL = field.NewAsterisk(tableName)
	_winHelpType.ID = field.NewInt64(tableName, "id")
	_winHelpType.Language = field.NewString(tableName, "language")
	_winHelpType.TypeName = field.NewString(tableName, "type_name")
	_winHelpType.ImageURL = field.NewString(tableName, "image_url")
	_winHelpType.Sort = field.NewInt64(tableName, "sort")
	_winHelpType.Status = field.NewInt64(tableName, "status")
	_winHelpType.CreateBy = field.NewString(tableName, "create_by")
	_winHelpType.UpdateBy = field.NewString(tableName, "update_by")
	_winHelpType.CreatedAt = field.NewInt64(tableName, "created_at")
	_winHelpType.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winHelpType.fillFieldMap()

	return _winHelpType
}

// winHelpType 帮助类型
type winHelpType struct {
	winHelpTypeDo

	ALL       field.Asterisk
	ID        field.Int64  // 主键编号
	Language  field.String // 语言
	TypeName  field.String // 类型名称
	ImageURL  field.String // 图片地址
	Sort      field.Int64  // 排序
	Status    field.Int64  // 状态:1-启用 0-停用
	CreateBy  field.String // 创建者
	UpdateBy  field.String // 更新人
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winHelpType) Table(newTableName string) *winHelpType {
	w.winHelpTypeDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winHelpType) As(alias string) *winHelpType {
	w.winHelpTypeDo.DO = *(w.winHelpTypeDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winHelpType) updateTableName(table string) *winHelpType {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Language = field.NewString(table, "language")
	w.TypeName = field.NewString(table, "type_name")
	w.ImageURL = field.NewString(table, "image_url")
	w.Sort = field.NewInt64(table, "sort")
	w.Status = field.NewInt64(table, "status")
	w.CreateBy = field.NewString(table, "create_by")
	w.UpdateBy = field.NewString(table, "update_by")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winHelpType) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winHelpType) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 10)
	w.fieldMap["id"] = w.ID
	w.fieldMap["language"] = w.Language
	w.fieldMap["type_name"] = w.TypeName
	w.fieldMap["image_url"] = w.ImageURL
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["status"] = w.Status
	w.fieldMap["create_by"] = w.CreateBy
	w.fieldMap["update_by"] = w.UpdateBy
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winHelpType) clone(db *gorm.DB) winHelpType {
	w.winHelpTypeDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winHelpType) replaceDB(db *gorm.DB) winHelpType {
	w.winHelpTypeDo.ReplaceDB(db)
	return w
}

type winHelpTypeDo struct{ gen.DO }

type IWinHelpTypeDo interface {
	gen.SubQuery
	Debug() IWinHelpTypeDo
	WithContext(ctx context.Context) IWinHelpTypeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinHelpTypeDo
	WriteDB() IWinHelpTypeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinHelpTypeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinHelpTypeDo
	Not(conds ...gen.Condition) IWinHelpTypeDo
	Or(conds ...gen.Condition) IWinHelpTypeDo
	Select(conds ...field.Expr) IWinHelpTypeDo
	Where(conds ...gen.Condition) IWinHelpTypeDo
	Order(conds ...field.Expr) IWinHelpTypeDo
	Distinct(cols ...field.Expr) IWinHelpTypeDo
	Omit(cols ...field.Expr) IWinHelpTypeDo
	Join(table schema.Tabler, on ...field.Expr) IWinHelpTypeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinHelpTypeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinHelpTypeDo
	Group(cols ...field.Expr) IWinHelpTypeDo
	Having(conds ...gen.Condition) IWinHelpTypeDo
	Limit(limit int) IWinHelpTypeDo
	Offset(offset int) IWinHelpTypeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinHelpTypeDo
	Unscoped() IWinHelpTypeDo
	Create(values ...*model.WinHelpType) error
	CreateInBatches(values []*model.WinHelpType, batchSize int) error
	Save(values ...*model.WinHelpType) error
	First() (*model.WinHelpType, error)
	Take() (*model.WinHelpType, error)
	Last() (*model.WinHelpType, error)
	Find() ([]*model.WinHelpType, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinHelpType, err error)
	FindInBatches(result *[]*model.WinHelpType, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinHelpType) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinHelpTypeDo
	Assign(attrs ...field.AssignExpr) IWinHelpTypeDo
	Joins(fields ...field.RelationField) IWinHelpTypeDo
	Preload(fields ...field.RelationField) IWinHelpTypeDo
	FirstOrInit() (*model.WinHelpType, error)
	FirstOrCreate() (*model.WinHelpType, error)
	FindByPage(offset int, limit int) (result []*model.WinHelpType, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinHelpTypeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winHelpTypeDo) Debug() IWinHelpTypeDo {
	return w.withDO(w.DO.Debug())
}

func (w winHelpTypeDo) WithContext(ctx context.Context) IWinHelpTypeDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winHelpTypeDo) ReadDB() IWinHelpTypeDo {
	return w.Clauses(dbresolver.Read)
}

func (w winHelpTypeDo) WriteDB() IWinHelpTypeDo {
	return w.Clauses(dbresolver.Write)
}

func (w winHelpTypeDo) Session(config *gorm.Session) IWinHelpTypeDo {
	return w.withDO(w.DO.Session(config))
}

func (w winHelpTypeDo) Clauses(conds ...clause.Expression) IWinHelpTypeDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winHelpTypeDo) Returning(value interface{}, columns ...string) IWinHelpTypeDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winHelpTypeDo) Not(conds ...gen.Condition) IWinHelpTypeDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winHelpTypeDo) Or(conds ...gen.Condition) IWinHelpTypeDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winHelpTypeDo) Select(conds ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winHelpTypeDo) Where(conds ...gen.Condition) IWinHelpTypeDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winHelpTypeDo) Order(conds ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winHelpTypeDo) Distinct(cols ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winHelpTypeDo) Omit(cols ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winHelpTypeDo) Join(table schema.Tabler, on ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winHelpTypeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winHelpTypeDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winHelpTypeDo) Group(cols ...field.Expr) IWinHelpTypeDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winHelpTypeDo) Having(conds ...gen.Condition) IWinHelpTypeDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winHelpTypeDo) Limit(limit int) IWinHelpTypeDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winHelpTypeDo) Offset(offset int) IWinHelpTypeDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winHelpTypeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinHelpTypeDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winHelpTypeDo) Unscoped() IWinHelpTypeDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winHelpTypeDo) Create(values ...*model.WinHelpType) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winHelpTypeDo) CreateInBatches(values []*model.WinHelpType, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winHelpTypeDo) Save(values ...*model.WinHelpType) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winHelpTypeDo) First() (*model.WinHelpType, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpType), nil
	}
}

func (w winHelpTypeDo) Take() (*model.WinHelpType, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpType), nil
	}
}

func (w winHelpTypeDo) Last() (*model.WinHelpType, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpType), nil
	}
}

func (w winHelpTypeDo) Find() ([]*model.WinHelpType, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinHelpType), err
}

func (w winHelpTypeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinHelpType, err error) {
	buf := make([]*model.WinHelpType, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winHelpTypeDo) FindInBatches(result *[]*model.WinHelpType, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winHelpTypeDo) Attrs(attrs ...field.AssignExpr) IWinHelpTypeDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winHelpTypeDo) Assign(attrs ...field.AssignExpr) IWinHelpTypeDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winHelpTypeDo) Joins(fields ...field.RelationField) IWinHelpTypeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winHelpTypeDo) Preload(fields ...field.RelationField) IWinHelpTypeDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winHelpTypeDo) FirstOrInit() (*model.WinHelpType, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpType), nil
	}
}

func (w winHelpTypeDo) FirstOrCreate() (*model.WinHelpType, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinHelpType), nil
	}
}

func (w winHelpTypeDo) FindByPage(offset int, limit int) (result []*model.WinHelpType, count int64, err error) {
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

func (w winHelpTypeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winHelpTypeDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winHelpTypeDo) Delete(models ...*model.WinHelpType) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winHelpTypeDo) withDO(do gen.Dao) *winHelpTypeDo {
	w.DO = *do.(*gen.DO)
	return w
}
