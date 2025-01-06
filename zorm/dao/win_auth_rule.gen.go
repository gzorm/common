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

func newWinAuthRule(db *gorm.DB, opts ...gen.DOOption) winAuthRule {
	_winAuthRule := winAuthRule{}

	_winAuthRule.winAuthRuleDo.UseDB(db, opts...)
	_winAuthRule.winAuthRuleDo.UseModel(&model.WinAuthRule{})

	tableName := _winAuthRule.winAuthRuleDo.TableName()
	_winAuthRule.ALL = field.NewAsterisk(tableName)
	_winAuthRule.ID = field.NewInt64(tableName, "id")
	_winAuthRule.Icon = field.NewString(tableName, "icon")
	_winAuthRule.MenuName = field.NewString(tableName, "menu_name")
	_winAuthRule.Title = field.NewString(tableName, "title")
	_winAuthRule.Pid = field.NewInt64(tableName, "pid")
	_winAuthRule.IsMenu = field.NewInt64(tableName, "is_menu")
	_winAuthRule.IsRaceMenu = field.NewInt64(tableName, "is_race_menu")
	_winAuthRule.Status = field.NewInt64(tableName, "status")
	_winAuthRule.CreatedAt = field.NewInt64(tableName, "created_at")
	_winAuthRule.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winAuthRule.OperatorName = field.NewString(tableName, "operator_name")

	_winAuthRule.fillFieldMap()

	return _winAuthRule
}

// winAuthRule 菜单表
type winAuthRule struct {
	winAuthRuleDo

	ALL          field.Asterisk
	ID           field.Int64
	Icon         field.String // 图标
	MenuName     field.String // 规则唯一标识Controller/action
	Title        field.String // 菜单名称
	Pid          field.Int64  // 父节点ID
	IsMenu       field.Int64  // 是否主菜单:1-是 0-否
	IsRaceMenu   field.Int64  // 是否按钮:1-是 0-否
	Status       field.Int64  // 状态: 1-启用 0-禁用
	CreatedAt    field.Int64
	UpdatedAt    field.Int64
	OperatorName field.String // 操作人姓名

	fieldMap map[string]field.Expr
}

func (w winAuthRule) Table(newTableName string) *winAuthRule {
	w.winAuthRuleDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAuthRule) As(alias string) *winAuthRule {
	w.winAuthRuleDo.DO = *(w.winAuthRuleDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAuthRule) updateTableName(table string) *winAuthRule {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Icon = field.NewString(table, "icon")
	w.MenuName = field.NewString(table, "menu_name")
	w.Title = field.NewString(table, "title")
	w.Pid = field.NewInt64(table, "pid")
	w.IsMenu = field.NewInt64(table, "is_menu")
	w.IsRaceMenu = field.NewInt64(table, "is_race_menu")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.OperatorName = field.NewString(table, "operator_name")

	w.fillFieldMap()

	return w
}

func (w *winAuthRule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAuthRule) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["icon"] = w.Icon
	w.fieldMap["menu_name"] = w.MenuName
	w.fieldMap["title"] = w.Title
	w.fieldMap["pid"] = w.Pid
	w.fieldMap["is_menu"] = w.IsMenu
	w.fieldMap["is_race_menu"] = w.IsRaceMenu
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["operator_name"] = w.OperatorName
}

func (w winAuthRule) clone(db *gorm.DB) winAuthRule {
	w.winAuthRuleDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAuthRule) replaceDB(db *gorm.DB) winAuthRule {
	w.winAuthRuleDo.ReplaceDB(db)
	return w
}

type winAuthRuleDo struct{ gen.DO }

type IWinAuthRuleDo interface {
	gen.SubQuery
	Debug() IWinAuthRuleDo
	WithContext(ctx context.Context) IWinAuthRuleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinAuthRuleDo
	WriteDB() IWinAuthRuleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinAuthRuleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinAuthRuleDo
	Not(conds ...gen.Condition) IWinAuthRuleDo
	Or(conds ...gen.Condition) IWinAuthRuleDo
	Select(conds ...field.Expr) IWinAuthRuleDo
	Where(conds ...gen.Condition) IWinAuthRuleDo
	Order(conds ...field.Expr) IWinAuthRuleDo
	Distinct(cols ...field.Expr) IWinAuthRuleDo
	Omit(cols ...field.Expr) IWinAuthRuleDo
	Join(table schema.Tabler, on ...field.Expr) IWinAuthRuleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinAuthRuleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinAuthRuleDo
	Group(cols ...field.Expr) IWinAuthRuleDo
	Having(conds ...gen.Condition) IWinAuthRuleDo
	Limit(limit int) IWinAuthRuleDo
	Offset(offset int) IWinAuthRuleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAuthRuleDo
	Unscoped() IWinAuthRuleDo
	Create(values ...*model.WinAuthRule) error
	CreateInBatches(values []*model.WinAuthRule, batchSize int) error
	Save(values ...*model.WinAuthRule) error
	First() (*model.WinAuthRule, error)
	Take() (*model.WinAuthRule, error)
	Last() (*model.WinAuthRule, error)
	Find() ([]*model.WinAuthRule, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAuthRule, err error)
	FindInBatches(result *[]*model.WinAuthRule, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinAuthRule) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinAuthRuleDo
	Assign(attrs ...field.AssignExpr) IWinAuthRuleDo
	Joins(fields ...field.RelationField) IWinAuthRuleDo
	Preload(fields ...field.RelationField) IWinAuthRuleDo
	FirstOrInit() (*model.WinAuthRule, error)
	FirstOrCreate() (*model.WinAuthRule, error)
	FindByPage(offset int, limit int) (result []*model.WinAuthRule, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinAuthRuleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winAuthRuleDo) Debug() IWinAuthRuleDo {
	return w.withDO(w.DO.Debug())
}

func (w winAuthRuleDo) WithContext(ctx context.Context) IWinAuthRuleDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAuthRuleDo) ReadDB() IWinAuthRuleDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAuthRuleDo) WriteDB() IWinAuthRuleDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAuthRuleDo) Session(config *gorm.Session) IWinAuthRuleDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAuthRuleDo) Clauses(conds ...clause.Expression) IWinAuthRuleDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAuthRuleDo) Returning(value interface{}, columns ...string) IWinAuthRuleDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAuthRuleDo) Not(conds ...gen.Condition) IWinAuthRuleDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAuthRuleDo) Or(conds ...gen.Condition) IWinAuthRuleDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAuthRuleDo) Select(conds ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAuthRuleDo) Where(conds ...gen.Condition) IWinAuthRuleDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAuthRuleDo) Order(conds ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAuthRuleDo) Distinct(cols ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAuthRuleDo) Omit(cols ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAuthRuleDo) Join(table schema.Tabler, on ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAuthRuleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAuthRuleDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAuthRuleDo) Group(cols ...field.Expr) IWinAuthRuleDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAuthRuleDo) Having(conds ...gen.Condition) IWinAuthRuleDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAuthRuleDo) Limit(limit int) IWinAuthRuleDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAuthRuleDo) Offset(offset int) IWinAuthRuleDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAuthRuleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAuthRuleDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAuthRuleDo) Unscoped() IWinAuthRuleDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAuthRuleDo) Create(values ...*model.WinAuthRule) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAuthRuleDo) CreateInBatches(values []*model.WinAuthRule, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAuthRuleDo) Save(values ...*model.WinAuthRule) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAuthRuleDo) First() (*model.WinAuthRule, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthRule), nil
	}
}

func (w winAuthRuleDo) Take() (*model.WinAuthRule, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthRule), nil
	}
}

func (w winAuthRuleDo) Last() (*model.WinAuthRule, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthRule), nil
	}
}

func (w winAuthRuleDo) Find() ([]*model.WinAuthRule, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAuthRule), err
}

func (w winAuthRuleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAuthRule, err error) {
	buf := make([]*model.WinAuthRule, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAuthRuleDo) FindInBatches(result *[]*model.WinAuthRule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAuthRuleDo) Attrs(attrs ...field.AssignExpr) IWinAuthRuleDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAuthRuleDo) Assign(attrs ...field.AssignExpr) IWinAuthRuleDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAuthRuleDo) Joins(fields ...field.RelationField) IWinAuthRuleDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAuthRuleDo) Preload(fields ...field.RelationField) IWinAuthRuleDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAuthRuleDo) FirstOrInit() (*model.WinAuthRule, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthRule), nil
	}
}

func (w winAuthRuleDo) FirstOrCreate() (*model.WinAuthRule, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAuthRule), nil
	}
}

func (w winAuthRuleDo) FindByPage(offset int, limit int) (result []*model.WinAuthRule, count int64, err error) {
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

func (w winAuthRuleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAuthRuleDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAuthRuleDo) Delete(models ...*model.WinAuthRule) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAuthRuleDo) withDO(do gen.Dao) *winAuthRuleDo {
	w.DO = *do.(*gen.DO)
	return w
}
