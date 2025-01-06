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

func newGameListTmp(db *gorm.DB, opts ...gen.DOOption) gameListTmp {
	_gameListTmp := gameListTmp{}

	_gameListTmp.gameListTmpDo.UseDB(db, opts...)
	_gameListTmp.gameListTmpDo.UseModel(&model.GameListTmp{})

	tableName := _gameListTmp.gameListTmpDo.TableName()
	_gameListTmp.ALL = field.NewAsterisk(tableName)
	_gameListTmp.ID = field.NewInt64(tableName, "id")
	_gameListTmp.Code = field.NewString(tableName, "code")
	_gameListTmp.GameProviderSubtypeID = field.NewInt64(tableName, "game_provider_subtype_id")
	_gameListTmp.GamePagcorID = field.NewInt64(tableName, "game_pagcor_id")
	_gameListTmp.GameTypeID = field.NewInt64(tableName, "game_type_id")
	_gameListTmp.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_gameListTmp.GameStartParam = field.NewString(tableName, "game_start_param")
	_gameListTmp.Sort = field.NewInt64(tableName, "sort")
	_gameListTmp.Status = field.NewInt64(tableName, "status")
	_gameListTmp.Name = field.NewString(tableName, "name")
	_gameListTmp.OriginalIcon = field.NewString(tableName, "original_icon")
	_gameListTmp.LatestIcon = field.NewString(tableName, "latest_icon")
	_gameListTmp.IsNew = field.NewInt64(tableName, "is_new")
	_gameListTmp.FavoriteStar = field.NewInt64(tableName, "favorite_star")
	_gameListTmp.HotStar = field.NewInt64(tableName, "hot_star")
	_gameListTmp.Device = field.NewInt64(tableName, "device")
	_gameListTmp.CreatedAt = field.NewInt64(tableName, "created_at")
	_gameListTmp.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_gameListTmp.CreatedBy = field.NewString(tableName, "created_by")
	_gameListTmp.UpdatedBy = field.NewString(tableName, "updated_by")
	_gameListTmp.Maintenance = field.NewString(tableName, "maintenance")

	_gameListTmp.fillFieldMap()

	return _gameListTmp
}

// gameListTmp 游戏全量表
type gameListTmp struct {
	gameListTmpDo

	ALL                   field.Asterisk
	ID                    field.Int64
	Code                  field.String // 启动游戏编码
	GameProviderSubtypeID field.Int64  // 关联game_provider_subtype表ID
	GamePagcorID          field.Int64  // pagcor类型id
	GameTypeID            field.Int64  // 游戏类型id
	GameProviderID        field.Int64  // 游戏供应商id
	GameStartParam        field.String // 特殊游戏三方启动参数，如:elbet
	Sort                  field.Int64  // 排序: 从低到高
	Status                field.Int64  // 状态: 1-启用 0-停用
	Name                  field.String // 简体名称
	OriginalIcon          field.String // 英文图片
	LatestIcon            field.String // 新版游戏图片
	IsNew                 field.Int64  // 是否新游戏:1-是 0-否
	FavoriteStar          field.Int64  // 收藏值
	HotStar               field.Int64  // 热度
	Device                field.Int64  // 设备:0-all 1-pc 2-h5
	CreatedAt             field.Int64
	UpdatedAt             field.Int64
	CreatedBy             field.String // 操作人姓名
	UpdatedBy             field.String // 最后更新人
	Maintenance           field.String // 维护时间

	fieldMap map[string]field.Expr
}

func (g gameListTmp) Table(newTableName string) *gameListTmp {
	g.gameListTmpDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gameListTmp) As(alias string) *gameListTmp {
	g.gameListTmpDo.DO = *(g.gameListTmpDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gameListTmp) updateTableName(table string) *gameListTmp {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.Code = field.NewString(table, "code")
	g.GameProviderSubtypeID = field.NewInt64(table, "game_provider_subtype_id")
	g.GamePagcorID = field.NewInt64(table, "game_pagcor_id")
	g.GameTypeID = field.NewInt64(table, "game_type_id")
	g.GameProviderID = field.NewInt64(table, "game_provider_id")
	g.GameStartParam = field.NewString(table, "game_start_param")
	g.Sort = field.NewInt64(table, "sort")
	g.Status = field.NewInt64(table, "status")
	g.Name = field.NewString(table, "name")
	g.OriginalIcon = field.NewString(table, "original_icon")
	g.LatestIcon = field.NewString(table, "latest_icon")
	g.IsNew = field.NewInt64(table, "is_new")
	g.FavoriteStar = field.NewInt64(table, "favorite_star")
	g.HotStar = field.NewInt64(table, "hot_star")
	g.Device = field.NewInt64(table, "device")
	g.CreatedAt = field.NewInt64(table, "created_at")
	g.UpdatedAt = field.NewInt64(table, "updated_at")
	g.CreatedBy = field.NewString(table, "created_by")
	g.UpdatedBy = field.NewString(table, "updated_by")
	g.Maintenance = field.NewString(table, "maintenance")

	g.fillFieldMap()

	return g
}

func (g *gameListTmp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gameListTmp) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 21)
	g.fieldMap["id"] = g.ID
	g.fieldMap["code"] = g.Code
	g.fieldMap["game_provider_subtype_id"] = g.GameProviderSubtypeID
	g.fieldMap["game_pagcor_id"] = g.GamePagcorID
	g.fieldMap["game_type_id"] = g.GameTypeID
	g.fieldMap["game_provider_id"] = g.GameProviderID
	g.fieldMap["game_start_param"] = g.GameStartParam
	g.fieldMap["sort"] = g.Sort
	g.fieldMap["status"] = g.Status
	g.fieldMap["name"] = g.Name
	g.fieldMap["original_icon"] = g.OriginalIcon
	g.fieldMap["latest_icon"] = g.LatestIcon
	g.fieldMap["is_new"] = g.IsNew
	g.fieldMap["favorite_star"] = g.FavoriteStar
	g.fieldMap["hot_star"] = g.HotStar
	g.fieldMap["device"] = g.Device
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["created_by"] = g.CreatedBy
	g.fieldMap["updated_by"] = g.UpdatedBy
	g.fieldMap["maintenance"] = g.Maintenance
}

func (g gameListTmp) clone(db *gorm.DB) gameListTmp {
	g.gameListTmpDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gameListTmp) replaceDB(db *gorm.DB) gameListTmp {
	g.gameListTmpDo.ReplaceDB(db)
	return g
}

type gameListTmpDo struct{ gen.DO }

type IGameListTmpDo interface {
	gen.SubQuery
	Debug() IGameListTmpDo
	WithContext(ctx context.Context) IGameListTmpDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGameListTmpDo
	WriteDB() IGameListTmpDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGameListTmpDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGameListTmpDo
	Not(conds ...gen.Condition) IGameListTmpDo
	Or(conds ...gen.Condition) IGameListTmpDo
	Select(conds ...field.Expr) IGameListTmpDo
	Where(conds ...gen.Condition) IGameListTmpDo
	Order(conds ...field.Expr) IGameListTmpDo
	Distinct(cols ...field.Expr) IGameListTmpDo
	Omit(cols ...field.Expr) IGameListTmpDo
	Join(table schema.Tabler, on ...field.Expr) IGameListTmpDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGameListTmpDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGameListTmpDo
	Group(cols ...field.Expr) IGameListTmpDo
	Having(conds ...gen.Condition) IGameListTmpDo
	Limit(limit int) IGameListTmpDo
	Offset(offset int) IGameListTmpDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGameListTmpDo
	Unscoped() IGameListTmpDo
	Create(values ...*model.GameListTmp) error
	CreateInBatches(values []*model.GameListTmp, batchSize int) error
	Save(values ...*model.GameListTmp) error
	First() (*model.GameListTmp, error)
	Take() (*model.GameListTmp, error)
	Last() (*model.GameListTmp, error)
	Find() ([]*model.GameListTmp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GameListTmp, err error)
	FindInBatches(result *[]*model.GameListTmp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GameListTmp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGameListTmpDo
	Assign(attrs ...field.AssignExpr) IGameListTmpDo
	Joins(fields ...field.RelationField) IGameListTmpDo
	Preload(fields ...field.RelationField) IGameListTmpDo
	FirstOrInit() (*model.GameListTmp, error)
	FirstOrCreate() (*model.GameListTmp, error)
	FindByPage(offset int, limit int) (result []*model.GameListTmp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGameListTmpDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g gameListTmpDo) Debug() IGameListTmpDo {
	return g.withDO(g.DO.Debug())
}

func (g gameListTmpDo) WithContext(ctx context.Context) IGameListTmpDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gameListTmpDo) ReadDB() IGameListTmpDo {
	return g.Clauses(dbresolver.Read)
}

func (g gameListTmpDo) WriteDB() IGameListTmpDo {
	return g.Clauses(dbresolver.Write)
}

func (g gameListTmpDo) Session(config *gorm.Session) IGameListTmpDo {
	return g.withDO(g.DO.Session(config))
}

func (g gameListTmpDo) Clauses(conds ...clause.Expression) IGameListTmpDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gameListTmpDo) Returning(value interface{}, columns ...string) IGameListTmpDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gameListTmpDo) Not(conds ...gen.Condition) IGameListTmpDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gameListTmpDo) Or(conds ...gen.Condition) IGameListTmpDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gameListTmpDo) Select(conds ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gameListTmpDo) Where(conds ...gen.Condition) IGameListTmpDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gameListTmpDo) Order(conds ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gameListTmpDo) Distinct(cols ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gameListTmpDo) Omit(cols ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gameListTmpDo) Join(table schema.Tabler, on ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gameListTmpDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gameListTmpDo) RightJoin(table schema.Tabler, on ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gameListTmpDo) Group(cols ...field.Expr) IGameListTmpDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gameListTmpDo) Having(conds ...gen.Condition) IGameListTmpDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gameListTmpDo) Limit(limit int) IGameListTmpDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gameListTmpDo) Offset(offset int) IGameListTmpDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gameListTmpDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGameListTmpDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gameListTmpDo) Unscoped() IGameListTmpDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gameListTmpDo) Create(values ...*model.GameListTmp) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gameListTmpDo) CreateInBatches(values []*model.GameListTmp, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gameListTmpDo) Save(values ...*model.GameListTmp) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gameListTmpDo) First() (*model.GameListTmp, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameListTmp), nil
	}
}

func (g gameListTmpDo) Take() (*model.GameListTmp, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameListTmp), nil
	}
}

func (g gameListTmpDo) Last() (*model.GameListTmp, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameListTmp), nil
	}
}

func (g gameListTmpDo) Find() ([]*model.GameListTmp, error) {
	result, err := g.DO.Find()
	return result.([]*model.GameListTmp), err
}

func (g gameListTmpDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GameListTmp, err error) {
	buf := make([]*model.GameListTmp, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gameListTmpDo) FindInBatches(result *[]*model.GameListTmp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gameListTmpDo) Attrs(attrs ...field.AssignExpr) IGameListTmpDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gameListTmpDo) Assign(attrs ...field.AssignExpr) IGameListTmpDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gameListTmpDo) Joins(fields ...field.RelationField) IGameListTmpDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gameListTmpDo) Preload(fields ...field.RelationField) IGameListTmpDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gameListTmpDo) FirstOrInit() (*model.GameListTmp, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameListTmp), nil
	}
}

func (g gameListTmpDo) FirstOrCreate() (*model.GameListTmp, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameListTmp), nil
	}
}

func (g gameListTmpDo) FindByPage(offset int, limit int) (result []*model.GameListTmp, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gameListTmpDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gameListTmpDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gameListTmpDo) Delete(models ...*model.GameListTmp) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gameListTmpDo) withDO(do gen.Dao) *gameListTmpDo {
	g.DO = *do.(*gen.DO)
	return g
}
