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

func newGameOperationUrl(db *gorm.DB, opts ...gen.DOOption) gameOperationUrl {
	_gameOperationUrl := gameOperationUrl{}

	_gameOperationUrl.gameOperationUrlDo.UseDB(db, opts...)
	_gameOperationUrl.gameOperationUrlDo.UseModel(&model.GameOperationUrl{})

	tableName := _gameOperationUrl.gameOperationUrlDo.TableName()
	_gameOperationUrl.ALL = field.NewAsterisk(tableName)
	_gameOperationUrl.ID = field.NewInt64(tableName, "id")
	_gameOperationUrl.Language = field.NewString(tableName, "language")
	_gameOperationUrl.LeagueName = field.NewString(tableName, "league_name")
	_gameOperationUrl.MatchTime = field.NewInt64(tableName, "match_time")
	_gameOperationUrl.HomeName = field.NewString(tableName, "home_name")
	_gameOperationUrl.HomeIcon = field.NewString(tableName, "home_icon")
	_gameOperationUrl.VisitorName = field.NewString(tableName, "visitor_name")
	_gameOperationUrl.VisitorIcon = field.NewString(tableName, "visitor_icon")
	_gameOperationUrl.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_gameOperationUrl.MatchID = field.NewString(tableName, "match_id")
	_gameOperationUrl.LeagueID = field.NewString(tableName, "league_id")
	_gameOperationUrl.Status = field.NewInt64(tableName, "status")
	_gameOperationUrl.CreatedAt = field.NewInt64(tableName, "created_at")
	_gameOperationUrl.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_gameOperationUrl.CreatedBy = field.NewString(tableName, "created_by")
	_gameOperationUrl.UpdatedBy = field.NewString(tableName, "updated_by")

	_gameOperationUrl.fillFieldMap()

	return _gameOperationUrl
}

// gameOperationUrl 游戏运营内链配置
type gameOperationUrl struct {
	gameOperationUrlDo

	ALL            field.Asterisk
	ID             field.Int64
	Language       field.String // 语言
	LeagueName     field.String // 联赛名称
	MatchTime      field.Int64  // 比赛时间
	HomeName       field.String // 主队名称
	HomeIcon       field.String // 主队图标
	VisitorName    field.String // 客队名称
	VisitorIcon    field.String // 客队图标
	GameProviderID field.Int64  // 游戏平台id
	MatchID        field.String // 比赛id
	LeagueID       field.String // 联赛id
	Status         field.Int64  // 状态: 1-启用 0-停用
	CreatedAt      field.Int64  // 创建时间
	UpdatedAt      field.Int64  // 更新时间
	CreatedBy      field.String // 操作人姓名
	UpdatedBy      field.String // 最后更新人

	fieldMap map[string]field.Expr
}

func (g gameOperationUrl) Table(newTableName string) *gameOperationUrl {
	g.gameOperationUrlDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gameOperationUrl) As(alias string) *gameOperationUrl {
	g.gameOperationUrlDo.DO = *(g.gameOperationUrlDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gameOperationUrl) updateTableName(table string) *gameOperationUrl {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.Language = field.NewString(table, "language")
	g.LeagueName = field.NewString(table, "league_name")
	g.MatchTime = field.NewInt64(table, "match_time")
	g.HomeName = field.NewString(table, "home_name")
	g.HomeIcon = field.NewString(table, "home_icon")
	g.VisitorName = field.NewString(table, "visitor_name")
	g.VisitorIcon = field.NewString(table, "visitor_icon")
	g.GameProviderID = field.NewInt64(table, "game_provider_id")
	g.MatchID = field.NewString(table, "match_id")
	g.LeagueID = field.NewString(table, "league_id")
	g.Status = field.NewInt64(table, "status")
	g.CreatedAt = field.NewInt64(table, "created_at")
	g.UpdatedAt = field.NewInt64(table, "updated_at")
	g.CreatedBy = field.NewString(table, "created_by")
	g.UpdatedBy = field.NewString(table, "updated_by")

	g.fillFieldMap()

	return g
}

func (g *gameOperationUrl) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gameOperationUrl) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 16)
	g.fieldMap["id"] = g.ID
	g.fieldMap["language"] = g.Language
	g.fieldMap["league_name"] = g.LeagueName
	g.fieldMap["match_time"] = g.MatchTime
	g.fieldMap["home_name"] = g.HomeName
	g.fieldMap["home_icon"] = g.HomeIcon
	g.fieldMap["visitor_name"] = g.VisitorName
	g.fieldMap["visitor_icon"] = g.VisitorIcon
	g.fieldMap["game_provider_id"] = g.GameProviderID
	g.fieldMap["match_id"] = g.MatchID
	g.fieldMap["league_id"] = g.LeagueID
	g.fieldMap["status"] = g.Status
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["created_by"] = g.CreatedBy
	g.fieldMap["updated_by"] = g.UpdatedBy
}

func (g gameOperationUrl) clone(db *gorm.DB) gameOperationUrl {
	g.gameOperationUrlDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gameOperationUrl) replaceDB(db *gorm.DB) gameOperationUrl {
	g.gameOperationUrlDo.ReplaceDB(db)
	return g
}

type gameOperationUrlDo struct{ gen.DO }

type IGameOperationUrlDo interface {
	gen.SubQuery
	Debug() IGameOperationUrlDo
	WithContext(ctx context.Context) IGameOperationUrlDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGameOperationUrlDo
	WriteDB() IGameOperationUrlDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGameOperationUrlDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGameOperationUrlDo
	Not(conds ...gen.Condition) IGameOperationUrlDo
	Or(conds ...gen.Condition) IGameOperationUrlDo
	Select(conds ...field.Expr) IGameOperationUrlDo
	Where(conds ...gen.Condition) IGameOperationUrlDo
	Order(conds ...field.Expr) IGameOperationUrlDo
	Distinct(cols ...field.Expr) IGameOperationUrlDo
	Omit(cols ...field.Expr) IGameOperationUrlDo
	Join(table schema.Tabler, on ...field.Expr) IGameOperationUrlDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGameOperationUrlDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGameOperationUrlDo
	Group(cols ...field.Expr) IGameOperationUrlDo
	Having(conds ...gen.Condition) IGameOperationUrlDo
	Limit(limit int) IGameOperationUrlDo
	Offset(offset int) IGameOperationUrlDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGameOperationUrlDo
	Unscoped() IGameOperationUrlDo
	Create(values ...*model.GameOperationUrl) error
	CreateInBatches(values []*model.GameOperationUrl, batchSize int) error
	Save(values ...*model.GameOperationUrl) error
	First() (*model.GameOperationUrl, error)
	Take() (*model.GameOperationUrl, error)
	Last() (*model.GameOperationUrl, error)
	Find() ([]*model.GameOperationUrl, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GameOperationUrl, err error)
	FindInBatches(result *[]*model.GameOperationUrl, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GameOperationUrl) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGameOperationUrlDo
	Assign(attrs ...field.AssignExpr) IGameOperationUrlDo
	Joins(fields ...field.RelationField) IGameOperationUrlDo
	Preload(fields ...field.RelationField) IGameOperationUrlDo
	FirstOrInit() (*model.GameOperationUrl, error)
	FirstOrCreate() (*model.GameOperationUrl, error)
	FindByPage(offset int, limit int) (result []*model.GameOperationUrl, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGameOperationUrlDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g gameOperationUrlDo) Debug() IGameOperationUrlDo {
	return g.withDO(g.DO.Debug())
}

func (g gameOperationUrlDo) WithContext(ctx context.Context) IGameOperationUrlDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gameOperationUrlDo) ReadDB() IGameOperationUrlDo {
	return g.Clauses(dbresolver.Read)
}

func (g gameOperationUrlDo) WriteDB() IGameOperationUrlDo {
	return g.Clauses(dbresolver.Write)
}

func (g gameOperationUrlDo) Session(config *gorm.Session) IGameOperationUrlDo {
	return g.withDO(g.DO.Session(config))
}

func (g gameOperationUrlDo) Clauses(conds ...clause.Expression) IGameOperationUrlDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gameOperationUrlDo) Returning(value interface{}, columns ...string) IGameOperationUrlDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gameOperationUrlDo) Not(conds ...gen.Condition) IGameOperationUrlDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gameOperationUrlDo) Or(conds ...gen.Condition) IGameOperationUrlDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gameOperationUrlDo) Select(conds ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gameOperationUrlDo) Where(conds ...gen.Condition) IGameOperationUrlDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gameOperationUrlDo) Order(conds ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gameOperationUrlDo) Distinct(cols ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gameOperationUrlDo) Omit(cols ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gameOperationUrlDo) Join(table schema.Tabler, on ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gameOperationUrlDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gameOperationUrlDo) RightJoin(table schema.Tabler, on ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gameOperationUrlDo) Group(cols ...field.Expr) IGameOperationUrlDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gameOperationUrlDo) Having(conds ...gen.Condition) IGameOperationUrlDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gameOperationUrlDo) Limit(limit int) IGameOperationUrlDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gameOperationUrlDo) Offset(offset int) IGameOperationUrlDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gameOperationUrlDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGameOperationUrlDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gameOperationUrlDo) Unscoped() IGameOperationUrlDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gameOperationUrlDo) Create(values ...*model.GameOperationUrl) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gameOperationUrlDo) CreateInBatches(values []*model.GameOperationUrl, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gameOperationUrlDo) Save(values ...*model.GameOperationUrl) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gameOperationUrlDo) First() (*model.GameOperationUrl, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameOperationUrl), nil
	}
}

func (g gameOperationUrlDo) Take() (*model.GameOperationUrl, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameOperationUrl), nil
	}
}

func (g gameOperationUrlDo) Last() (*model.GameOperationUrl, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameOperationUrl), nil
	}
}

func (g gameOperationUrlDo) Find() ([]*model.GameOperationUrl, error) {
	result, err := g.DO.Find()
	return result.([]*model.GameOperationUrl), err
}

func (g gameOperationUrlDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GameOperationUrl, err error) {
	buf := make([]*model.GameOperationUrl, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gameOperationUrlDo) FindInBatches(result *[]*model.GameOperationUrl, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gameOperationUrlDo) Attrs(attrs ...field.AssignExpr) IGameOperationUrlDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gameOperationUrlDo) Assign(attrs ...field.AssignExpr) IGameOperationUrlDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gameOperationUrlDo) Joins(fields ...field.RelationField) IGameOperationUrlDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gameOperationUrlDo) Preload(fields ...field.RelationField) IGameOperationUrlDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gameOperationUrlDo) FirstOrInit() (*model.GameOperationUrl, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameOperationUrl), nil
	}
}

func (g gameOperationUrlDo) FirstOrCreate() (*model.GameOperationUrl, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GameOperationUrl), nil
	}
}

func (g gameOperationUrlDo) FindByPage(offset int, limit int) (result []*model.GameOperationUrl, count int64, err error) {
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

func (g gameOperationUrlDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gameOperationUrlDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gameOperationUrlDo) Delete(models ...*model.GameOperationUrl) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gameOperationUrlDo) withDO(do gen.Dao) *gameOperationUrlDo {
	g.DO = *do.(*gen.DO)
	return g
}
