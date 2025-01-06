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

func newWinUserLevelRebate(db *gorm.DB, opts ...gen.DOOption) winUserLevelRebate {
	_winUserLevelRebate := winUserLevelRebate{}

	_winUserLevelRebate.winUserLevelRebateDo.UseDB(db, opts...)
	_winUserLevelRebate.winUserLevelRebateDo.UseModel(&model.WinUserLevelRebate{})

	tableName := _winUserLevelRebate.winUserLevelRebateDo.TableName()
	_winUserLevelRebate.ALL = field.NewAsterisk(tableName)
	_winUserLevelRebate.ID = field.NewInt64(tableName, "id")
	_winUserLevelRebate.LevelID = field.NewInt64(tableName, "level_id")
	_winUserLevelRebate.GroupID = field.NewInt64(tableName, "group_id")
	_winUserLevelRebate.RebateRate = field.NewField(tableName, "rebate_rate")
	_winUserLevelRebate.Status = field.NewInt64(tableName, "status")
	_winUserLevelRebate.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserLevelRebate.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winUserLevelRebate.UpdatedUser = field.NewString(tableName, "updated_user")

	_winUserLevelRebate.fillFieldMap()

	return _winUserLevelRebate
}

// winUserLevelRebate 会员等级返水
type winUserLevelRebate struct {
	winUserLevelRebateDo

	ALL         field.Asterisk
	ID          field.Int64 // ID主键
	LevelID     field.Int64 // 会员等级ID
	GroupID     field.Int64 // 类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能 11-table game
	RebateRate  field.Field // 返水比例
	Status      field.Int64 // 状态:1-启用 0-停用
	CreatedAt   field.Int64
	UpdatedAt   field.Int64
	UpdatedUser field.String // 最后修改人

	fieldMap map[string]field.Expr
}

func (w winUserLevelRebate) Table(newTableName string) *winUserLevelRebate {
	w.winUserLevelRebateDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserLevelRebate) As(alias string) *winUserLevelRebate {
	w.winUserLevelRebateDo.DO = *(w.winUserLevelRebateDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserLevelRebate) updateTableName(table string) *winUserLevelRebate {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.LevelID = field.NewInt64(table, "level_id")
	w.GroupID = field.NewInt64(table, "group_id")
	w.RebateRate = field.NewField(table, "rebate_rate")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.UpdatedUser = field.NewString(table, "updated_user")

	w.fillFieldMap()

	return w
}

func (w *winUserLevelRebate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserLevelRebate) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["level_id"] = w.LevelID
	w.fieldMap["group_id"] = w.GroupID
	w.fieldMap["rebate_rate"] = w.RebateRate
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["updated_user"] = w.UpdatedUser
}

func (w winUserLevelRebate) clone(db *gorm.DB) winUserLevelRebate {
	w.winUserLevelRebateDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserLevelRebate) replaceDB(db *gorm.DB) winUserLevelRebate {
	w.winUserLevelRebateDo.ReplaceDB(db)
	return w
}

type winUserLevelRebateDo struct{ gen.DO }

type IWinUserLevelRebateDo interface {
	gen.SubQuery
	Debug() IWinUserLevelRebateDo
	WithContext(ctx context.Context) IWinUserLevelRebateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserLevelRebateDo
	WriteDB() IWinUserLevelRebateDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserLevelRebateDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserLevelRebateDo
	Not(conds ...gen.Condition) IWinUserLevelRebateDo
	Or(conds ...gen.Condition) IWinUserLevelRebateDo
	Select(conds ...field.Expr) IWinUserLevelRebateDo
	Where(conds ...gen.Condition) IWinUserLevelRebateDo
	Order(conds ...field.Expr) IWinUserLevelRebateDo
	Distinct(cols ...field.Expr) IWinUserLevelRebateDo
	Omit(cols ...field.Expr) IWinUserLevelRebateDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserLevelRebateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRebateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRebateDo
	Group(cols ...field.Expr) IWinUserLevelRebateDo
	Having(conds ...gen.Condition) IWinUserLevelRebateDo
	Limit(limit int) IWinUserLevelRebateDo
	Offset(offset int) IWinUserLevelRebateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelRebateDo
	Unscoped() IWinUserLevelRebateDo
	Create(values ...*model.WinUserLevelRebate) error
	CreateInBatches(values []*model.WinUserLevelRebate, batchSize int) error
	Save(values ...*model.WinUserLevelRebate) error
	First() (*model.WinUserLevelRebate, error)
	Take() (*model.WinUserLevelRebate, error)
	Last() (*model.WinUserLevelRebate, error)
	Find() ([]*model.WinUserLevelRebate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevelRebate, err error)
	FindInBatches(result *[]*model.WinUserLevelRebate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserLevelRebate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserLevelRebateDo
	Assign(attrs ...field.AssignExpr) IWinUserLevelRebateDo
	Joins(fields ...field.RelationField) IWinUserLevelRebateDo
	Preload(fields ...field.RelationField) IWinUserLevelRebateDo
	FirstOrInit() (*model.WinUserLevelRebate, error)
	FirstOrCreate() (*model.WinUserLevelRebate, error)
	FindByPage(offset int, limit int) (result []*model.WinUserLevelRebate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserLevelRebateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserLevelRebateDo) Debug() IWinUserLevelRebateDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserLevelRebateDo) WithContext(ctx context.Context) IWinUserLevelRebateDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserLevelRebateDo) ReadDB() IWinUserLevelRebateDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserLevelRebateDo) WriteDB() IWinUserLevelRebateDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserLevelRebateDo) Session(config *gorm.Session) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserLevelRebateDo) Clauses(conds ...clause.Expression) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserLevelRebateDo) Returning(value interface{}, columns ...string) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserLevelRebateDo) Not(conds ...gen.Condition) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserLevelRebateDo) Or(conds ...gen.Condition) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserLevelRebateDo) Select(conds ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserLevelRebateDo) Where(conds ...gen.Condition) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserLevelRebateDo) Order(conds ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserLevelRebateDo) Distinct(cols ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserLevelRebateDo) Omit(cols ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserLevelRebateDo) Join(table schema.Tabler, on ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserLevelRebateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserLevelRebateDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserLevelRebateDo) Group(cols ...field.Expr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserLevelRebateDo) Having(conds ...gen.Condition) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserLevelRebateDo) Limit(limit int) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserLevelRebateDo) Offset(offset int) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserLevelRebateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserLevelRebateDo) Unscoped() IWinUserLevelRebateDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserLevelRebateDo) Create(values ...*model.WinUserLevelRebate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserLevelRebateDo) CreateInBatches(values []*model.WinUserLevelRebate, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserLevelRebateDo) Save(values ...*model.WinUserLevelRebate) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserLevelRebateDo) First() (*model.WinUserLevelRebate, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelRebate), nil
	}
}

func (w winUserLevelRebateDo) Take() (*model.WinUserLevelRebate, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelRebate), nil
	}
}

func (w winUserLevelRebateDo) Last() (*model.WinUserLevelRebate, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelRebate), nil
	}
}

func (w winUserLevelRebateDo) Find() ([]*model.WinUserLevelRebate, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserLevelRebate), err
}

func (w winUserLevelRebateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLevelRebate, err error) {
	buf := make([]*model.WinUserLevelRebate, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserLevelRebateDo) FindInBatches(result *[]*model.WinUserLevelRebate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserLevelRebateDo) Attrs(attrs ...field.AssignExpr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserLevelRebateDo) Assign(attrs ...field.AssignExpr) IWinUserLevelRebateDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserLevelRebateDo) Joins(fields ...field.RelationField) IWinUserLevelRebateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserLevelRebateDo) Preload(fields ...field.RelationField) IWinUserLevelRebateDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserLevelRebateDo) FirstOrInit() (*model.WinUserLevelRebate, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelRebate), nil
	}
}

func (w winUserLevelRebateDo) FirstOrCreate() (*model.WinUserLevelRebate, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLevelRebate), nil
	}
}

func (w winUserLevelRebateDo) FindByPage(offset int, limit int) (result []*model.WinUserLevelRebate, count int64, err error) {
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

func (w winUserLevelRebateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserLevelRebateDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserLevelRebateDo) Delete(models ...*model.WinUserLevelRebate) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserLevelRebateDo) withDO(do gen.Dao) *winUserLevelRebateDo {
	w.DO = *do.(*gen.DO)
	return w
}
