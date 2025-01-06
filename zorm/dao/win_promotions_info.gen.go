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

func newWinPromotionsInfo(db *gorm.DB, opts ...gen.DOOption) winPromotionsInfo {
	_winPromotionsInfo := winPromotionsInfo{}

	_winPromotionsInfo.winPromotionsInfoDo.UseDB(db, opts...)
	_winPromotionsInfo.winPromotionsInfoDo.UseModel(&model.WinPromotionsInfo{})

	tableName := _winPromotionsInfo.winPromotionsInfoDo.TableName()
	_winPromotionsInfo.ALL = field.NewAsterisk(tableName)
	_winPromotionsInfo.ID = field.NewInt64(tableName, "id")
	_winPromotionsInfo.PromotionID = field.NewInt64(tableName, "promotion_id")
	_winPromotionsInfo.Code = field.NewString(tableName, "code")
	_winPromotionsInfo.CodeZh = field.NewString(tableName, "code_zh")
	_winPromotionsInfo.LangCode = field.NewString(tableName, "lang_code")
	_winPromotionsInfo.Title = field.NewString(tableName, "title")
	_winPromotionsInfo.Img = field.NewString(tableName, "img")
	_winPromotionsInfo.Details = field.NewString(tableName, "details")
	_winPromotionsInfo.CreatedAt = field.NewInt64(tableName, "created_at")

	_winPromotionsInfo.fillFieldMap()

	return _winPromotionsInfo
}

// winPromotionsInfo 活动信息表
type winPromotionsInfo struct {
	winPromotionsInfoDo

	ALL         field.Asterisk
	ID          field.Int64  // 主键
	PromotionID field.Int64  // 签到活动ID
	Code        field.String // 活动标记:首充彩金,二充彩金等
	CodeZh      field.String // 活动标记-中文名称
	LangCode    field.String // 语言
	Title       field.String // 标题
	Img         field.String // 活动图片
	Details     field.String // 详情描述
	CreatedAt   field.Int64  // 创建时间

	fieldMap map[string]field.Expr
}

func (w winPromotionsInfo) Table(newTableName string) *winPromotionsInfo {
	w.winPromotionsInfoDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winPromotionsInfo) As(alias string) *winPromotionsInfo {
	w.winPromotionsInfoDo.DO = *(w.winPromotionsInfoDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winPromotionsInfo) updateTableName(table string) *winPromotionsInfo {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.PromotionID = field.NewInt64(table, "promotion_id")
	w.Code = field.NewString(table, "code")
	w.CodeZh = field.NewString(table, "code_zh")
	w.LangCode = field.NewString(table, "lang_code")
	w.Title = field.NewString(table, "title")
	w.Img = field.NewString(table, "img")
	w.Details = field.NewString(table, "details")
	w.CreatedAt = field.NewInt64(table, "created_at")

	w.fillFieldMap()

	return w
}

func (w *winPromotionsInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winPromotionsInfo) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["promotion_id"] = w.PromotionID
	w.fieldMap["code"] = w.Code
	w.fieldMap["code_zh"] = w.CodeZh
	w.fieldMap["lang_code"] = w.LangCode
	w.fieldMap["title"] = w.Title
	w.fieldMap["img"] = w.Img
	w.fieldMap["details"] = w.Details
	w.fieldMap["created_at"] = w.CreatedAt
}

func (w winPromotionsInfo) clone(db *gorm.DB) winPromotionsInfo {
	w.winPromotionsInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winPromotionsInfo) replaceDB(db *gorm.DB) winPromotionsInfo {
	w.winPromotionsInfoDo.ReplaceDB(db)
	return w
}

type winPromotionsInfoDo struct{ gen.DO }

type IWinPromotionsInfoDo interface {
	gen.SubQuery
	Debug() IWinPromotionsInfoDo
	WithContext(ctx context.Context) IWinPromotionsInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinPromotionsInfoDo
	WriteDB() IWinPromotionsInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinPromotionsInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinPromotionsInfoDo
	Not(conds ...gen.Condition) IWinPromotionsInfoDo
	Or(conds ...gen.Condition) IWinPromotionsInfoDo
	Select(conds ...field.Expr) IWinPromotionsInfoDo
	Where(conds ...gen.Condition) IWinPromotionsInfoDo
	Order(conds ...field.Expr) IWinPromotionsInfoDo
	Distinct(cols ...field.Expr) IWinPromotionsInfoDo
	Omit(cols ...field.Expr) IWinPromotionsInfoDo
	Join(table schema.Tabler, on ...field.Expr) IWinPromotionsInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsInfoDo
	Group(cols ...field.Expr) IWinPromotionsInfoDo
	Having(conds ...gen.Condition) IWinPromotionsInfoDo
	Limit(limit int) IWinPromotionsInfoDo
	Offset(offset int) IWinPromotionsInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionsInfoDo
	Unscoped() IWinPromotionsInfoDo
	Create(values ...*model.WinPromotionsInfo) error
	CreateInBatches(values []*model.WinPromotionsInfo, batchSize int) error
	Save(values ...*model.WinPromotionsInfo) error
	First() (*model.WinPromotionsInfo, error)
	Take() (*model.WinPromotionsInfo, error)
	Last() (*model.WinPromotionsInfo, error)
	Find() ([]*model.WinPromotionsInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotionsInfo, err error)
	FindInBatches(result *[]*model.WinPromotionsInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinPromotionsInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinPromotionsInfoDo
	Assign(attrs ...field.AssignExpr) IWinPromotionsInfoDo
	Joins(fields ...field.RelationField) IWinPromotionsInfoDo
	Preload(fields ...field.RelationField) IWinPromotionsInfoDo
	FirstOrInit() (*model.WinPromotionsInfo, error)
	FirstOrCreate() (*model.WinPromotionsInfo, error)
	FindByPage(offset int, limit int) (result []*model.WinPromotionsInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinPromotionsInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winPromotionsInfoDo) Debug() IWinPromotionsInfoDo {
	return w.withDO(w.DO.Debug())
}

func (w winPromotionsInfoDo) WithContext(ctx context.Context) IWinPromotionsInfoDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winPromotionsInfoDo) ReadDB() IWinPromotionsInfoDo {
	return w.Clauses(dbresolver.Read)
}

func (w winPromotionsInfoDo) WriteDB() IWinPromotionsInfoDo {
	return w.Clauses(dbresolver.Write)
}

func (w winPromotionsInfoDo) Session(config *gorm.Session) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Session(config))
}

func (w winPromotionsInfoDo) Clauses(conds ...clause.Expression) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winPromotionsInfoDo) Returning(value interface{}, columns ...string) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winPromotionsInfoDo) Not(conds ...gen.Condition) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winPromotionsInfoDo) Or(conds ...gen.Condition) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winPromotionsInfoDo) Select(conds ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winPromotionsInfoDo) Where(conds ...gen.Condition) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winPromotionsInfoDo) Order(conds ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winPromotionsInfoDo) Distinct(cols ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winPromotionsInfoDo) Omit(cols ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winPromotionsInfoDo) Join(table schema.Tabler, on ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winPromotionsInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winPromotionsInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winPromotionsInfoDo) Group(cols ...field.Expr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winPromotionsInfoDo) Having(conds ...gen.Condition) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winPromotionsInfoDo) Limit(limit int) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winPromotionsInfoDo) Offset(offset int) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winPromotionsInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winPromotionsInfoDo) Unscoped() IWinPromotionsInfoDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winPromotionsInfoDo) Create(values ...*model.WinPromotionsInfo) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winPromotionsInfoDo) CreateInBatches(values []*model.WinPromotionsInfo, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winPromotionsInfoDo) Save(values ...*model.WinPromotionsInfo) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winPromotionsInfoDo) First() (*model.WinPromotionsInfo, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionsInfo), nil
	}
}

func (w winPromotionsInfoDo) Take() (*model.WinPromotionsInfo, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionsInfo), nil
	}
}

func (w winPromotionsInfoDo) Last() (*model.WinPromotionsInfo, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionsInfo), nil
	}
}

func (w winPromotionsInfoDo) Find() ([]*model.WinPromotionsInfo, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinPromotionsInfo), err
}

func (w winPromotionsInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotionsInfo, err error) {
	buf := make([]*model.WinPromotionsInfo, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winPromotionsInfoDo) FindInBatches(result *[]*model.WinPromotionsInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winPromotionsInfoDo) Attrs(attrs ...field.AssignExpr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winPromotionsInfoDo) Assign(attrs ...field.AssignExpr) IWinPromotionsInfoDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winPromotionsInfoDo) Joins(fields ...field.RelationField) IWinPromotionsInfoDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winPromotionsInfoDo) Preload(fields ...field.RelationField) IWinPromotionsInfoDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winPromotionsInfoDo) FirstOrInit() (*model.WinPromotionsInfo, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionsInfo), nil
	}
}

func (w winPromotionsInfoDo) FirstOrCreate() (*model.WinPromotionsInfo, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotionsInfo), nil
	}
}

func (w winPromotionsInfoDo) FindByPage(offset int, limit int) (result []*model.WinPromotionsInfo, count int64, err error) {
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

func (w winPromotionsInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winPromotionsInfoDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winPromotionsInfoDo) Delete(models ...*model.WinPromotionsInfo) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winPromotionsInfoDo) withDO(do gen.Dao) *winPromotionsInfoDo {
	w.DO = *do.(*gen.DO)
	return w
}
