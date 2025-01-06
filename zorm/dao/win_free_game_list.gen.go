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

func newWinFreeGameList(db *gorm.DB, opts ...gen.DOOption) winFreeGameList {
	_winFreeGameList := winFreeGameList{}

	_winFreeGameList.winFreeGameListDo.UseDB(db, opts...)
	_winFreeGameList.winFreeGameListDo.UseModel(&model.WinFreeGameList{})

	tableName := _winFreeGameList.winFreeGameListDo.TableName()
	_winFreeGameList.ALL = field.NewAsterisk(tableName)
	_winFreeGameList.ID = field.NewInt64(tableName, "id")
	_winFreeGameList.Code = field.NewString(tableName, "code")
	_winFreeGameList.Name = field.NewString(tableName, "name")
	_winFreeGameList.GameProvider = field.NewString(tableName, "game_provider")
	_winFreeGameList.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_winFreeGameList.LatestIcon = field.NewString(tableName, "latest_icon")
	_winFreeGameList.CreateAt = field.NewInt64(tableName, "create_at")
	_winFreeGameList.UpdateAt = field.NewInt64(tableName, "update_at")
	_winFreeGameList.UpdateBy = field.NewString(tableName, "update_by")
	_winFreeGameList.Status = field.NewInt64(tableName, "status")
	_winFreeGameList.GameProviderSubtypeID = field.NewInt64(tableName, "game_provider_subtype_id")

	_winFreeGameList.fillFieldMap()

	return _winFreeGameList
}

// winFreeGameList 免费游戏列表
type winFreeGameList struct {
	winFreeGameListDo

	ALL                   field.Asterisk
	ID                    field.Int64
	Code                  field.String // 启动游戏编码
	Name                  field.String // 游戏名称
	GameProvider          field.String // 游戏供应商
	GameProviderID        field.Int64  // 游戏供应商id
	LatestIcon            field.String // 新版游戏图片
	CreateAt              field.Int64  // 创建时间
	UpdateAt              field.Int64  // 修改时间
	UpdateBy              field.String // 操作人
	Status                field.Int64  // 0关闭1展示
	GameProviderSubtypeID field.Int64

	fieldMap map[string]field.Expr
}

func (w winFreeGameList) Table(newTableName string) *winFreeGameList {
	w.winFreeGameListDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winFreeGameList) As(alias string) *winFreeGameList {
	w.winFreeGameListDo.DO = *(w.winFreeGameListDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winFreeGameList) updateTableName(table string) *winFreeGameList {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Code = field.NewString(table, "code")
	w.Name = field.NewString(table, "name")
	w.GameProvider = field.NewString(table, "game_provider")
	w.GameProviderID = field.NewInt64(table, "game_provider_id")
	w.LatestIcon = field.NewString(table, "latest_icon")
	w.CreateAt = field.NewInt64(table, "create_at")
	w.UpdateAt = field.NewInt64(table, "update_at")
	w.UpdateBy = field.NewString(table, "update_by")
	w.Status = field.NewInt64(table, "status")
	w.GameProviderSubtypeID = field.NewInt64(table, "game_provider_subtype_id")

	w.fillFieldMap()

	return w
}

func (w *winFreeGameList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winFreeGameList) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["code"] = w.Code
	w.fieldMap["name"] = w.Name
	w.fieldMap["game_provider"] = w.GameProvider
	w.fieldMap["game_provider_id"] = w.GameProviderID
	w.fieldMap["latest_icon"] = w.LatestIcon
	w.fieldMap["create_at"] = w.CreateAt
	w.fieldMap["update_at"] = w.UpdateAt
	w.fieldMap["update_by"] = w.UpdateBy
	w.fieldMap["status"] = w.Status
	w.fieldMap["game_provider_subtype_id"] = w.GameProviderSubtypeID
}

func (w winFreeGameList) clone(db *gorm.DB) winFreeGameList {
	w.winFreeGameListDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winFreeGameList) replaceDB(db *gorm.DB) winFreeGameList {
	w.winFreeGameListDo.ReplaceDB(db)
	return w
}

type winFreeGameListDo struct{ gen.DO }

type IWinFreeGameListDo interface {
	gen.SubQuery
	Debug() IWinFreeGameListDo
	WithContext(ctx context.Context) IWinFreeGameListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinFreeGameListDo
	WriteDB() IWinFreeGameListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinFreeGameListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinFreeGameListDo
	Not(conds ...gen.Condition) IWinFreeGameListDo
	Or(conds ...gen.Condition) IWinFreeGameListDo
	Select(conds ...field.Expr) IWinFreeGameListDo
	Where(conds ...gen.Condition) IWinFreeGameListDo
	Order(conds ...field.Expr) IWinFreeGameListDo
	Distinct(cols ...field.Expr) IWinFreeGameListDo
	Omit(cols ...field.Expr) IWinFreeGameListDo
	Join(table schema.Tabler, on ...field.Expr) IWinFreeGameListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinFreeGameListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinFreeGameListDo
	Group(cols ...field.Expr) IWinFreeGameListDo
	Having(conds ...gen.Condition) IWinFreeGameListDo
	Limit(limit int) IWinFreeGameListDo
	Offset(offset int) IWinFreeGameListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinFreeGameListDo
	Unscoped() IWinFreeGameListDo
	Create(values ...*model.WinFreeGameList) error
	CreateInBatches(values []*model.WinFreeGameList, batchSize int) error
	Save(values ...*model.WinFreeGameList) error
	First() (*model.WinFreeGameList, error)
	Take() (*model.WinFreeGameList, error)
	Last() (*model.WinFreeGameList, error)
	Find() ([]*model.WinFreeGameList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinFreeGameList, err error)
	FindInBatches(result *[]*model.WinFreeGameList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinFreeGameList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinFreeGameListDo
	Assign(attrs ...field.AssignExpr) IWinFreeGameListDo
	Joins(fields ...field.RelationField) IWinFreeGameListDo
	Preload(fields ...field.RelationField) IWinFreeGameListDo
	FirstOrInit() (*model.WinFreeGameList, error)
	FirstOrCreate() (*model.WinFreeGameList, error)
	FindByPage(offset int, limit int) (result []*model.WinFreeGameList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinFreeGameListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winFreeGameListDo) Debug() IWinFreeGameListDo {
	return w.withDO(w.DO.Debug())
}

func (w winFreeGameListDo) WithContext(ctx context.Context) IWinFreeGameListDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winFreeGameListDo) ReadDB() IWinFreeGameListDo {
	return w.Clauses(dbresolver.Read)
}

func (w winFreeGameListDo) WriteDB() IWinFreeGameListDo {
	return w.Clauses(dbresolver.Write)
}

func (w winFreeGameListDo) Session(config *gorm.Session) IWinFreeGameListDo {
	return w.withDO(w.DO.Session(config))
}

func (w winFreeGameListDo) Clauses(conds ...clause.Expression) IWinFreeGameListDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winFreeGameListDo) Returning(value interface{}, columns ...string) IWinFreeGameListDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winFreeGameListDo) Not(conds ...gen.Condition) IWinFreeGameListDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winFreeGameListDo) Or(conds ...gen.Condition) IWinFreeGameListDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winFreeGameListDo) Select(conds ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winFreeGameListDo) Where(conds ...gen.Condition) IWinFreeGameListDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winFreeGameListDo) Order(conds ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winFreeGameListDo) Distinct(cols ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winFreeGameListDo) Omit(cols ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winFreeGameListDo) Join(table schema.Tabler, on ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winFreeGameListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winFreeGameListDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winFreeGameListDo) Group(cols ...field.Expr) IWinFreeGameListDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winFreeGameListDo) Having(conds ...gen.Condition) IWinFreeGameListDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winFreeGameListDo) Limit(limit int) IWinFreeGameListDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winFreeGameListDo) Offset(offset int) IWinFreeGameListDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winFreeGameListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinFreeGameListDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winFreeGameListDo) Unscoped() IWinFreeGameListDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winFreeGameListDo) Create(values ...*model.WinFreeGameList) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winFreeGameListDo) CreateInBatches(values []*model.WinFreeGameList, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winFreeGameListDo) Save(values ...*model.WinFreeGameList) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winFreeGameListDo) First() (*model.WinFreeGameList, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFreeGameList), nil
	}
}

func (w winFreeGameListDo) Take() (*model.WinFreeGameList, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFreeGameList), nil
	}
}

func (w winFreeGameListDo) Last() (*model.WinFreeGameList, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFreeGameList), nil
	}
}

func (w winFreeGameListDo) Find() ([]*model.WinFreeGameList, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinFreeGameList), err
}

func (w winFreeGameListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinFreeGameList, err error) {
	buf := make([]*model.WinFreeGameList, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winFreeGameListDo) FindInBatches(result *[]*model.WinFreeGameList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winFreeGameListDo) Attrs(attrs ...field.AssignExpr) IWinFreeGameListDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winFreeGameListDo) Assign(attrs ...field.AssignExpr) IWinFreeGameListDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winFreeGameListDo) Joins(fields ...field.RelationField) IWinFreeGameListDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winFreeGameListDo) Preload(fields ...field.RelationField) IWinFreeGameListDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winFreeGameListDo) FirstOrInit() (*model.WinFreeGameList, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFreeGameList), nil
	}
}

func (w winFreeGameListDo) FirstOrCreate() (*model.WinFreeGameList, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinFreeGameList), nil
	}
}

func (w winFreeGameListDo) FindByPage(offset int, limit int) (result []*model.WinFreeGameList, count int64, err error) {
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

func (w winFreeGameListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winFreeGameListDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winFreeGameListDo) Delete(models ...*model.WinFreeGameList) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winFreeGameListDo) withDO(do gen.Dao) *winFreeGameListDo {
	w.DO = *do.(*gen.DO)
	return w
}
