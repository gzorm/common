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

func newLevelFreeGame(db *gorm.DB, opts ...gen.DOOption) levelFreeGame {
	_levelFreeGame := levelFreeGame{}

	_levelFreeGame.levelFreeGameDo.UseDB(db, opts...)
	_levelFreeGame.levelFreeGameDo.UseModel(&model.LevelFreeGame{})

	tableName := _levelFreeGame.levelFreeGameDo.TableName()
	_levelFreeGame.ALL = field.NewAsterisk(tableName)
	_levelFreeGame.ID = field.NewInt64(tableName, "id")
	_levelFreeGame.LevelID = field.NewInt64(tableName, "level_id")
	_levelFreeGame.GameProvider = field.NewString(tableName, "game_provider")
	_levelFreeGame.GameName = field.NewString(tableName, "game_name")
	_levelFreeGame.SpinNumber = field.NewInt64(tableName, "spin_number")
	_levelFreeGame.BetAmount = field.NewInt64(tableName, "bet_amount")
	_levelFreeGame.ValidTime = field.NewInt64(tableName, "valid_time")
	_levelFreeGame.Status = field.NewInt64(tableName, "status")

	_levelFreeGame.fillFieldMap()

	return _levelFreeGame
}

// levelFreeGame free game
type levelFreeGame struct {
	levelFreeGameDo

	ALL          field.Asterisk
	ID           field.Int64 // id
	LevelID      field.Int64 // level
	GameProvider field.String
	GameName     field.String
	SpinNumber   field.Int64
	BetAmount    field.Int64
	ValidTime    field.Int64
	/*
		游戏开启状态0:关闭，1:开启

	*/
	Status field.Int64

	fieldMap map[string]field.Expr
}

func (l levelFreeGame) Table(newTableName string) *levelFreeGame {
	l.levelFreeGameDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l levelFreeGame) As(alias string) *levelFreeGame {
	l.levelFreeGameDo.DO = *(l.levelFreeGameDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *levelFreeGame) updateTableName(table string) *levelFreeGame {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.LevelID = field.NewInt64(table, "level_id")
	l.GameProvider = field.NewString(table, "game_provider")
	l.GameName = field.NewString(table, "game_name")
	l.SpinNumber = field.NewInt64(table, "spin_number")
	l.BetAmount = field.NewInt64(table, "bet_amount")
	l.ValidTime = field.NewInt64(table, "valid_time")
	l.Status = field.NewInt64(table, "status")

	l.fillFieldMap()

	return l
}

func (l *levelFreeGame) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *levelFreeGame) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 8)
	l.fieldMap["id"] = l.ID
	l.fieldMap["level_id"] = l.LevelID
	l.fieldMap["game_provider"] = l.GameProvider
	l.fieldMap["game_name"] = l.GameName
	l.fieldMap["spin_number"] = l.SpinNumber
	l.fieldMap["bet_amount"] = l.BetAmount
	l.fieldMap["valid_time"] = l.ValidTime
	l.fieldMap["status"] = l.Status
}

func (l levelFreeGame) clone(db *gorm.DB) levelFreeGame {
	l.levelFreeGameDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l levelFreeGame) replaceDB(db *gorm.DB) levelFreeGame {
	l.levelFreeGameDo.ReplaceDB(db)
	return l
}

type levelFreeGameDo struct{ gen.DO }

type ILevelFreeGameDo interface {
	gen.SubQuery
	Debug() ILevelFreeGameDo
	WithContext(ctx context.Context) ILevelFreeGameDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILevelFreeGameDo
	WriteDB() ILevelFreeGameDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILevelFreeGameDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILevelFreeGameDo
	Not(conds ...gen.Condition) ILevelFreeGameDo
	Or(conds ...gen.Condition) ILevelFreeGameDo
	Select(conds ...field.Expr) ILevelFreeGameDo
	Where(conds ...gen.Condition) ILevelFreeGameDo
	Order(conds ...field.Expr) ILevelFreeGameDo
	Distinct(cols ...field.Expr) ILevelFreeGameDo
	Omit(cols ...field.Expr) ILevelFreeGameDo
	Join(table schema.Tabler, on ...field.Expr) ILevelFreeGameDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILevelFreeGameDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILevelFreeGameDo
	Group(cols ...field.Expr) ILevelFreeGameDo
	Having(conds ...gen.Condition) ILevelFreeGameDo
	Limit(limit int) ILevelFreeGameDo
	Offset(offset int) ILevelFreeGameDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILevelFreeGameDo
	Unscoped() ILevelFreeGameDo
	Create(values ...*model.LevelFreeGame) error
	CreateInBatches(values []*model.LevelFreeGame, batchSize int) error
	Save(values ...*model.LevelFreeGame) error
	First() (*model.LevelFreeGame, error)
	Take() (*model.LevelFreeGame, error)
	Last() (*model.LevelFreeGame, error)
	Find() ([]*model.LevelFreeGame, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LevelFreeGame, err error)
	FindInBatches(result *[]*model.LevelFreeGame, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.LevelFreeGame) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILevelFreeGameDo
	Assign(attrs ...field.AssignExpr) ILevelFreeGameDo
	Joins(fields ...field.RelationField) ILevelFreeGameDo
	Preload(fields ...field.RelationField) ILevelFreeGameDo
	FirstOrInit() (*model.LevelFreeGame, error)
	FirstOrCreate() (*model.LevelFreeGame, error)
	FindByPage(offset int, limit int) (result []*model.LevelFreeGame, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILevelFreeGameDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l levelFreeGameDo) Debug() ILevelFreeGameDo {
	return l.withDO(l.DO.Debug())
}

func (l levelFreeGameDo) WithContext(ctx context.Context) ILevelFreeGameDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l levelFreeGameDo) ReadDB() ILevelFreeGameDo {
	return l.Clauses(dbresolver.Read)
}

func (l levelFreeGameDo) WriteDB() ILevelFreeGameDo {
	return l.Clauses(dbresolver.Write)
}

func (l levelFreeGameDo) Session(config *gorm.Session) ILevelFreeGameDo {
	return l.withDO(l.DO.Session(config))
}

func (l levelFreeGameDo) Clauses(conds ...clause.Expression) ILevelFreeGameDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l levelFreeGameDo) Returning(value interface{}, columns ...string) ILevelFreeGameDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l levelFreeGameDo) Not(conds ...gen.Condition) ILevelFreeGameDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l levelFreeGameDo) Or(conds ...gen.Condition) ILevelFreeGameDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l levelFreeGameDo) Select(conds ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l levelFreeGameDo) Where(conds ...gen.Condition) ILevelFreeGameDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l levelFreeGameDo) Order(conds ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l levelFreeGameDo) Distinct(cols ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l levelFreeGameDo) Omit(cols ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l levelFreeGameDo) Join(table schema.Tabler, on ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l levelFreeGameDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l levelFreeGameDo) RightJoin(table schema.Tabler, on ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l levelFreeGameDo) Group(cols ...field.Expr) ILevelFreeGameDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l levelFreeGameDo) Having(conds ...gen.Condition) ILevelFreeGameDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l levelFreeGameDo) Limit(limit int) ILevelFreeGameDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l levelFreeGameDo) Offset(offset int) ILevelFreeGameDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l levelFreeGameDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILevelFreeGameDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l levelFreeGameDo) Unscoped() ILevelFreeGameDo {
	return l.withDO(l.DO.Unscoped())
}

func (l levelFreeGameDo) Create(values ...*model.LevelFreeGame) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l levelFreeGameDo) CreateInBatches(values []*model.LevelFreeGame, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l levelFreeGameDo) Save(values ...*model.LevelFreeGame) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l levelFreeGameDo) First() (*model.LevelFreeGame, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LevelFreeGame), nil
	}
}

func (l levelFreeGameDo) Take() (*model.LevelFreeGame, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LevelFreeGame), nil
	}
}

func (l levelFreeGameDo) Last() (*model.LevelFreeGame, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LevelFreeGame), nil
	}
}

func (l levelFreeGameDo) Find() ([]*model.LevelFreeGame, error) {
	result, err := l.DO.Find()
	return result.([]*model.LevelFreeGame), err
}

func (l levelFreeGameDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LevelFreeGame, err error) {
	buf := make([]*model.LevelFreeGame, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l levelFreeGameDo) FindInBatches(result *[]*model.LevelFreeGame, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l levelFreeGameDo) Attrs(attrs ...field.AssignExpr) ILevelFreeGameDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l levelFreeGameDo) Assign(attrs ...field.AssignExpr) ILevelFreeGameDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l levelFreeGameDo) Joins(fields ...field.RelationField) ILevelFreeGameDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l levelFreeGameDo) Preload(fields ...field.RelationField) ILevelFreeGameDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l levelFreeGameDo) FirstOrInit() (*model.LevelFreeGame, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LevelFreeGame), nil
	}
}

func (l levelFreeGameDo) FirstOrCreate() (*model.LevelFreeGame, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LevelFreeGame), nil
	}
}

func (l levelFreeGameDo) FindByPage(offset int, limit int) (result []*model.LevelFreeGame, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l levelFreeGameDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l levelFreeGameDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l levelFreeGameDo) Delete(models ...*model.LevelFreeGame) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *levelFreeGameDo) withDO(do gen.Dao) *levelFreeGameDo {
	l.DO = *do.(*gen.DO)
	return l
}