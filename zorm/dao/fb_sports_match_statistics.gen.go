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

func newFbSportsMatchStatistics(db *gorm.DB, opts ...gen.DOOption) fbSportsMatchStatistics {
	_fbSportsMatchStatistics := fbSportsMatchStatistics{}

	_fbSportsMatchStatistics.fbSportsMatchStatisticsDo.UseDB(db, opts...)
	_fbSportsMatchStatistics.fbSportsMatchStatisticsDo.UseModel(&model.FbSportsMatchStatistics{})

	tableName := _fbSportsMatchStatistics.fbSportsMatchStatisticsDo.TableName()
	_fbSportsMatchStatistics.ALL = field.NewAsterisk(tableName)
	_fbSportsMatchStatistics.ID = field.NewInt64(tableName, "id")
	_fbSportsMatchStatistics.MatchCode = field.NewString(tableName, "match_code")
	_fbSportsMatchStatistics.TotalMatchCount = field.NewInt64(tableName, "total_match_count")
	_fbSportsMatchStatistics.StageList = field.NewString(tableName, "stage_list")
	_fbSportsMatchStatistics.HotMatchTotal = field.NewInt64(tableName, "hot_match_total")
	_fbSportsMatchStatistics.HotLeagueStatistics = field.NewString(tableName, "hot_league_statistics")
	_fbSportsMatchStatistics.MarketTypeStatistics = field.NewString(tableName, "market_type_statistics")
	_fbSportsMatchStatistics.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbSportsMatchStatistics.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_fbSportsMatchStatistics.fillFieldMap()

	return _fbSportsMatchStatistics
}

// fbSportsMatchStatistics 赛事统计主表，存储赛事的总数据和整体统计信息
type fbSportsMatchStatistics struct {
	fbSportsMatchStatisticsDo

	ALL                  field.Asterisk
	ID                   field.Int64  // 主键ID
	MatchCode            field.String // 赛事唯一标识码（用于关联子表）
	TotalMatchCount      field.Int64  // 赛事个数（对应 data.tc）
	StageList            field.String // 赛事类型的场次集合，存储为 JSON 字符串（对应 data.sl）
	HotMatchTotal        field.Int64  // 热门赛事总数，包括竞彩赛事和热门联赛赛事（对应 data.ht）
	HotLeagueStatistics  field.String // 热门联赛赛事统计信息，存储为 JSON 字符串（对应 data.hls）
	MarketTypeStatistics field.String // 盘口组合下赛事统计信息，存储为 JSON 字符串（对应 data.mts）
	CreatedAt            field.Int64  // 记录创建时间（时间戳）
	UpdatedAt            field.Int64  // 记录更新时间（时间戳）

	fieldMap map[string]field.Expr
}

func (f fbSportsMatchStatistics) Table(newTableName string) *fbSportsMatchStatistics {
	f.fbSportsMatchStatisticsDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSportsMatchStatistics) As(alias string) *fbSportsMatchStatistics {
	f.fbSportsMatchStatisticsDo.DO = *(f.fbSportsMatchStatisticsDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSportsMatchStatistics) updateTableName(table string) *fbSportsMatchStatistics {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.MatchCode = field.NewString(table, "match_code")
	f.TotalMatchCount = field.NewInt64(table, "total_match_count")
	f.StageList = field.NewString(table, "stage_list")
	f.HotMatchTotal = field.NewInt64(table, "hot_match_total")
	f.HotLeagueStatistics = field.NewString(table, "hot_league_statistics")
	f.MarketTypeStatistics = field.NewString(table, "market_type_statistics")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fbSportsMatchStatistics) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSportsMatchStatistics) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 9)
	f.fieldMap["id"] = f.ID
	f.fieldMap["match_code"] = f.MatchCode
	f.fieldMap["total_match_count"] = f.TotalMatchCount
	f.fieldMap["stage_list"] = f.StageList
	f.fieldMap["hot_match_total"] = f.HotMatchTotal
	f.fieldMap["hot_league_statistics"] = f.HotLeagueStatistics
	f.fieldMap["market_type_statistics"] = f.MarketTypeStatistics
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fbSportsMatchStatistics) clone(db *gorm.DB) fbSportsMatchStatistics {
	f.fbSportsMatchStatisticsDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSportsMatchStatistics) replaceDB(db *gorm.DB) fbSportsMatchStatistics {
	f.fbSportsMatchStatisticsDo.ReplaceDB(db)
	return f
}

type fbSportsMatchStatisticsDo struct{ gen.DO }

type IFbSportsMatchStatisticsDo interface {
	gen.SubQuery
	Debug() IFbSportsMatchStatisticsDo
	WithContext(ctx context.Context) IFbSportsMatchStatisticsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSportsMatchStatisticsDo
	WriteDB() IFbSportsMatchStatisticsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSportsMatchStatisticsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSportsMatchStatisticsDo
	Not(conds ...gen.Condition) IFbSportsMatchStatisticsDo
	Or(conds ...gen.Condition) IFbSportsMatchStatisticsDo
	Select(conds ...field.Expr) IFbSportsMatchStatisticsDo
	Where(conds ...gen.Condition) IFbSportsMatchStatisticsDo
	Order(conds ...field.Expr) IFbSportsMatchStatisticsDo
	Distinct(cols ...field.Expr) IFbSportsMatchStatisticsDo
	Omit(cols ...field.Expr) IFbSportsMatchStatisticsDo
	Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsDo
	Group(cols ...field.Expr) IFbSportsMatchStatisticsDo
	Having(conds ...gen.Condition) IFbSportsMatchStatisticsDo
	Limit(limit int) IFbSportsMatchStatisticsDo
	Offset(offset int) IFbSportsMatchStatisticsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchStatisticsDo
	Unscoped() IFbSportsMatchStatisticsDo
	Create(values ...*model.FbSportsMatchStatistics) error
	CreateInBatches(values []*model.FbSportsMatchStatistics, batchSize int) error
	Save(values ...*model.FbSportsMatchStatistics) error
	First() (*model.FbSportsMatchStatistics, error)
	Take() (*model.FbSportsMatchStatistics, error)
	Last() (*model.FbSportsMatchStatistics, error)
	Find() ([]*model.FbSportsMatchStatistics, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchStatistics, err error)
	FindInBatches(result *[]*model.FbSportsMatchStatistics, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSportsMatchStatistics) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSportsMatchStatisticsDo
	Assign(attrs ...field.AssignExpr) IFbSportsMatchStatisticsDo
	Joins(fields ...field.RelationField) IFbSportsMatchStatisticsDo
	Preload(fields ...field.RelationField) IFbSportsMatchStatisticsDo
	FirstOrInit() (*model.FbSportsMatchStatistics, error)
	FirstOrCreate() (*model.FbSportsMatchStatistics, error)
	FindByPage(offset int, limit int) (result []*model.FbSportsMatchStatistics, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSportsMatchStatisticsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSportsMatchStatisticsDo) Debug() IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSportsMatchStatisticsDo) WithContext(ctx context.Context) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSportsMatchStatisticsDo) ReadDB() IFbSportsMatchStatisticsDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSportsMatchStatisticsDo) WriteDB() IFbSportsMatchStatisticsDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSportsMatchStatisticsDo) Session(config *gorm.Session) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSportsMatchStatisticsDo) Clauses(conds ...clause.Expression) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSportsMatchStatisticsDo) Returning(value interface{}, columns ...string) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSportsMatchStatisticsDo) Not(conds ...gen.Condition) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSportsMatchStatisticsDo) Or(conds ...gen.Condition) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSportsMatchStatisticsDo) Select(conds ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSportsMatchStatisticsDo) Where(conds ...gen.Condition) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSportsMatchStatisticsDo) Order(conds ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSportsMatchStatisticsDo) Distinct(cols ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSportsMatchStatisticsDo) Omit(cols ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSportsMatchStatisticsDo) Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSportsMatchStatisticsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSportsMatchStatisticsDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSportsMatchStatisticsDo) Group(cols ...field.Expr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSportsMatchStatisticsDo) Having(conds ...gen.Condition) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSportsMatchStatisticsDo) Limit(limit int) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSportsMatchStatisticsDo) Offset(offset int) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSportsMatchStatisticsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSportsMatchStatisticsDo) Unscoped() IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSportsMatchStatisticsDo) Create(values ...*model.FbSportsMatchStatistics) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSportsMatchStatisticsDo) CreateInBatches(values []*model.FbSportsMatchStatistics, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSportsMatchStatisticsDo) Save(values ...*model.FbSportsMatchStatistics) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSportsMatchStatisticsDo) First() (*model.FbSportsMatchStatistics, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatistics), nil
	}
}

func (f fbSportsMatchStatisticsDo) Take() (*model.FbSportsMatchStatistics, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatistics), nil
	}
}

func (f fbSportsMatchStatisticsDo) Last() (*model.FbSportsMatchStatistics, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatistics), nil
	}
}

func (f fbSportsMatchStatisticsDo) Find() ([]*model.FbSportsMatchStatistics, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSportsMatchStatistics), err
}

func (f fbSportsMatchStatisticsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchStatistics, err error) {
	buf := make([]*model.FbSportsMatchStatistics, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSportsMatchStatisticsDo) FindInBatches(result *[]*model.FbSportsMatchStatistics, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSportsMatchStatisticsDo) Attrs(attrs ...field.AssignExpr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSportsMatchStatisticsDo) Assign(attrs ...field.AssignExpr) IFbSportsMatchStatisticsDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSportsMatchStatisticsDo) Joins(fields ...field.RelationField) IFbSportsMatchStatisticsDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSportsMatchStatisticsDo) Preload(fields ...field.RelationField) IFbSportsMatchStatisticsDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSportsMatchStatisticsDo) FirstOrInit() (*model.FbSportsMatchStatistics, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatistics), nil
	}
}

func (f fbSportsMatchStatisticsDo) FirstOrCreate() (*model.FbSportsMatchStatistics, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatistics), nil
	}
}

func (f fbSportsMatchStatisticsDo) FindByPage(offset int, limit int) (result []*model.FbSportsMatchStatistics, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fbSportsMatchStatisticsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSportsMatchStatisticsDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSportsMatchStatisticsDo) Delete(models ...*model.FbSportsMatchStatistics) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSportsMatchStatisticsDo) withDO(do gen.Dao) *fbSportsMatchStatisticsDo {
	f.DO = *do.(*gen.DO)
	return f
}
