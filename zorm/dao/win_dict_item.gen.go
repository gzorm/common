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

func newWinDictItem(db *gorm.DB, opts ...gen.DOOption) winDictItem {
	_winDictItem := winDictItem{}

	_winDictItem.winDictItemDo.UseDB(db, opts...)
	_winDictItem.winDictItemDo.UseModel(&model.WinDictItem{})

	tableName := _winDictItem.winDictItemDo.TableName()
	_winDictItem.ALL = field.NewAsterisk(tableName)
	_winDictItem.ID = field.NewInt64(tableName, "id")
	_winDictItem.Code = field.NewString(tableName, "code")
	_winDictItem.Title = field.NewString(tableName, "title")
	_winDictItem.TitleEn = field.NewString(tableName, "title_en")
	_winDictItem.TitleAr = field.NewString(tableName, "title_ar")
	_winDictItem.Remark = field.NewString(tableName, "remark")
	_winDictItem.Sort = field.NewInt64(tableName, "sort")
	_winDictItem.ReferID = field.NewInt64(tableName, "refer_id")
	_winDictItem.Status = field.NewInt64(tableName, "status")
	_winDictItem.IsShow = field.NewInt64(tableName, "is_show")
	_winDictItem.CreatedAt = field.NewInt64(tableName, "created_at")
	_winDictItem.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winDictItem.fillFieldMap()

	return _winDictItem
}

// winDictItem 字典项表
type winDictItem struct {
	winDictItemDo

	ALL       field.Asterisk
	ID        field.Int64
	Code      field.String // 字典码
	Title     field.String // 字典名称
	TitleEn   field.String // 英文
	TitleAr   field.String // 阿拉伯文
	Remark    field.String // 备注
	Sort      field.Int64  // 排序:从高到低
	ReferID   field.Int64  // 字典表ID
	Status    field.Int64  // 状态:1-启用 0-禁用
	IsShow    field.Int64  // 类型:0-全部 1-前端 2-后台
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winDictItem) Table(newTableName string) *winDictItem {
	w.winDictItemDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winDictItem) As(alias string) *winDictItem {
	w.winDictItemDo.DO = *(w.winDictItemDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winDictItem) updateTableName(table string) *winDictItem {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Code = field.NewString(table, "code")
	w.Title = field.NewString(table, "title")
	w.TitleEn = field.NewString(table, "title_en")
	w.TitleAr = field.NewString(table, "title_ar")
	w.Remark = field.NewString(table, "remark")
	w.Sort = field.NewInt64(table, "sort")
	w.ReferID = field.NewInt64(table, "refer_id")
	w.Status = field.NewInt64(table, "status")
	w.IsShow = field.NewInt64(table, "is_show")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winDictItem) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winDictItem) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 12)
	w.fieldMap["id"] = w.ID
	w.fieldMap["code"] = w.Code
	w.fieldMap["title"] = w.Title
	w.fieldMap["title_en"] = w.TitleEn
	w.fieldMap["title_ar"] = w.TitleAr
	w.fieldMap["remark"] = w.Remark
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["refer_id"] = w.ReferID
	w.fieldMap["status"] = w.Status
	w.fieldMap["is_show"] = w.IsShow
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winDictItem) clone(db *gorm.DB) winDictItem {
	w.winDictItemDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winDictItem) replaceDB(db *gorm.DB) winDictItem {
	w.winDictItemDo.ReplaceDB(db)
	return w
}

type winDictItemDo struct{ gen.DO }

type IWinDictItemDo interface {
	gen.SubQuery
	Debug() IWinDictItemDo
	WithContext(ctx context.Context) IWinDictItemDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinDictItemDo
	WriteDB() IWinDictItemDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinDictItemDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinDictItemDo
	Not(conds ...gen.Condition) IWinDictItemDo
	Or(conds ...gen.Condition) IWinDictItemDo
	Select(conds ...field.Expr) IWinDictItemDo
	Where(conds ...gen.Condition) IWinDictItemDo
	Order(conds ...field.Expr) IWinDictItemDo
	Distinct(cols ...field.Expr) IWinDictItemDo
	Omit(cols ...field.Expr) IWinDictItemDo
	Join(table schema.Tabler, on ...field.Expr) IWinDictItemDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinDictItemDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinDictItemDo
	Group(cols ...field.Expr) IWinDictItemDo
	Having(conds ...gen.Condition) IWinDictItemDo
	Limit(limit int) IWinDictItemDo
	Offset(offset int) IWinDictItemDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinDictItemDo
	Unscoped() IWinDictItemDo
	Create(values ...*model.WinDictItem) error
	CreateInBatches(values []*model.WinDictItem, batchSize int) error
	Save(values ...*model.WinDictItem) error
	First() (*model.WinDictItem, error)
	Take() (*model.WinDictItem, error)
	Last() (*model.WinDictItem, error)
	Find() ([]*model.WinDictItem, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinDictItem, err error)
	FindInBatches(result *[]*model.WinDictItem, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinDictItem) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinDictItemDo
	Assign(attrs ...field.AssignExpr) IWinDictItemDo
	Joins(fields ...field.RelationField) IWinDictItemDo
	Preload(fields ...field.RelationField) IWinDictItemDo
	FirstOrInit() (*model.WinDictItem, error)
	FirstOrCreate() (*model.WinDictItem, error)
	FindByPage(offset int, limit int) (result []*model.WinDictItem, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinDictItemDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winDictItemDo) Debug() IWinDictItemDo {
	return w.withDO(w.DO.Debug())
}

func (w winDictItemDo) WithContext(ctx context.Context) IWinDictItemDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winDictItemDo) ReadDB() IWinDictItemDo {
	return w.Clauses(dbresolver.Read)
}

func (w winDictItemDo) WriteDB() IWinDictItemDo {
	return w.Clauses(dbresolver.Write)
}

func (w winDictItemDo) Session(config *gorm.Session) IWinDictItemDo {
	return w.withDO(w.DO.Session(config))
}

func (w winDictItemDo) Clauses(conds ...clause.Expression) IWinDictItemDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winDictItemDo) Returning(value interface{}, columns ...string) IWinDictItemDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winDictItemDo) Not(conds ...gen.Condition) IWinDictItemDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winDictItemDo) Or(conds ...gen.Condition) IWinDictItemDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winDictItemDo) Select(conds ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winDictItemDo) Where(conds ...gen.Condition) IWinDictItemDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winDictItemDo) Order(conds ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winDictItemDo) Distinct(cols ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winDictItemDo) Omit(cols ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winDictItemDo) Join(table schema.Tabler, on ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winDictItemDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winDictItemDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winDictItemDo) Group(cols ...field.Expr) IWinDictItemDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winDictItemDo) Having(conds ...gen.Condition) IWinDictItemDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winDictItemDo) Limit(limit int) IWinDictItemDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winDictItemDo) Offset(offset int) IWinDictItemDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winDictItemDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinDictItemDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winDictItemDo) Unscoped() IWinDictItemDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winDictItemDo) Create(values ...*model.WinDictItem) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winDictItemDo) CreateInBatches(values []*model.WinDictItem, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winDictItemDo) Save(values ...*model.WinDictItem) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winDictItemDo) First() (*model.WinDictItem, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDictItem), nil
	}
}

func (w winDictItemDo) Take() (*model.WinDictItem, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDictItem), nil
	}
}

func (w winDictItemDo) Last() (*model.WinDictItem, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDictItem), nil
	}
}

func (w winDictItemDo) Find() ([]*model.WinDictItem, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinDictItem), err
}

func (w winDictItemDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinDictItem, err error) {
	buf := make([]*model.WinDictItem, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winDictItemDo) FindInBatches(result *[]*model.WinDictItem, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winDictItemDo) Attrs(attrs ...field.AssignExpr) IWinDictItemDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winDictItemDo) Assign(attrs ...field.AssignExpr) IWinDictItemDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winDictItemDo) Joins(fields ...field.RelationField) IWinDictItemDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winDictItemDo) Preload(fields ...field.RelationField) IWinDictItemDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winDictItemDo) FirstOrInit() (*model.WinDictItem, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDictItem), nil
	}
}

func (w winDictItemDo) FirstOrCreate() (*model.WinDictItem, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinDictItem), nil
	}
}

func (w winDictItemDo) FindByPage(offset int, limit int) (result []*model.WinDictItem, count int64, err error) {
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

func (w winDictItemDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winDictItemDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winDictItemDo) Delete(models ...*model.WinDictItem) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winDictItemDo) withDO(do gen.Dao) *winDictItemDo {
	w.DO = *do.(*gen.DO)
	return w
}