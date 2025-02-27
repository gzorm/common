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

func newFbSportsMatchStatisticsStageList(db *gorm.DB, opts ...gen.DOOption) fbSportsMatchStatisticsStageList {
	_fbSportsMatchStatisticsStageList := fbSportsMatchStatisticsStageList{}

	_fbSportsMatchStatisticsStageList.fbSportsMatchStatisticsStageListDo.UseDB(db, opts...)
	_fbSportsMatchStatisticsStageList.fbSportsMatchStatisticsStageListDo.UseModel(&model.FbSportsMatchStatisticsStageList{})

	tableName := _fbSportsMatchStatisticsStageList.fbSportsMatchStatisticsStageListDo.TableName()
	_fbSportsMatchStatisticsStageList.ALL = field.NewAsterisk(tableName)
	_fbSportsMatchStatisticsStageList.ID = field.NewInt64(tableName, "id")
	_fbSportsMatchStatisticsStageList.StatisticsID = field.NewInt64(tableName, "statistics_id")
	_fbSportsMatchStatisticsStageList.SlType = field.NewInt64(tableName, "sl_type")
	_fbSportsMatchStatisticsStageList.Description = field.NewString(tableName, "description")
	_fbSportsMatchStatisticsStageList.TotalCount = field.NewInt64(tableName, "total_count")
	_fbSportsMatchStatisticsStageList.SubStageList = field.NewString(tableName, "sub_stage_list")
	_fbSportsMatchStatisticsStageList.CreatedAt = field.NewInt64(tableName, "created_at")
	_fbSportsMatchStatisticsStageList.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_fbSportsMatchStatisticsStageList.fillFieldMap()

	return _fbSportsMatchStatisticsStageList
}

// fbSportsMatchStatisticsStageList 赛事类型表，存储每种类型的赛事统计信息
type fbSportsMatchStatisticsStageList struct {
	fbSportsMatchStatisticsStageListDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键ID
	StatisticsID field.Int64  // 关联主表的 ID（对应主表 fb_sports_match_statistics.id）
	SlType       field.Int64  // 分类类型（对应 data.sl.ty，枚举 match_play_type）
	Description  field.String // 分类描述（对应 data.sl.des）
	TotalCount   field.Int64  // 赛事总数（对应 data.sl.tc）
	SubStageList field.String // 每个类型下的子赛事统计信息（对应 data.sl.ssl，存储为 JSON）
	CreatedAt    field.Int64  // 记录创建时间（时间戳）
	UpdatedAt    field.Int64  // 记录更新时间（时间戳）

	fieldMap map[string]field.Expr
}

func (f fbSportsMatchStatisticsStageList) Table(newTableName string) *fbSportsMatchStatisticsStageList {
	f.fbSportsMatchStatisticsStageListDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fbSportsMatchStatisticsStageList) As(alias string) *fbSportsMatchStatisticsStageList {
	f.fbSportsMatchStatisticsStageListDo.DO = *(f.fbSportsMatchStatisticsStageListDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fbSportsMatchStatisticsStageList) updateTableName(table string) *fbSportsMatchStatisticsStageList {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.StatisticsID = field.NewInt64(table, "statistics_id")
	f.SlType = field.NewInt64(table, "sl_type")
	f.Description = field.NewString(table, "description")
	f.TotalCount = field.NewInt64(table, "total_count")
	f.SubStageList = field.NewString(table, "sub_stage_list")
	f.CreatedAt = field.NewInt64(table, "created_at")
	f.UpdatedAt = field.NewInt64(table, "updated_at")

	f.fillFieldMap()

	return f
}

func (f *fbSportsMatchStatisticsStageList) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fbSportsMatchStatisticsStageList) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 8)
	f.fieldMap["id"] = f.ID
	f.fieldMap["statistics_id"] = f.StatisticsID
	f.fieldMap["sl_type"] = f.SlType
	f.fieldMap["description"] = f.Description
	f.fieldMap["total_count"] = f.TotalCount
	f.fieldMap["sub_stage_list"] = f.SubStageList
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
}

func (f fbSportsMatchStatisticsStageList) clone(db *gorm.DB) fbSportsMatchStatisticsStageList {
	f.fbSportsMatchStatisticsStageListDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fbSportsMatchStatisticsStageList) replaceDB(db *gorm.DB) fbSportsMatchStatisticsStageList {
	f.fbSportsMatchStatisticsStageListDo.ReplaceDB(db)
	return f
}

type fbSportsMatchStatisticsStageListDo struct{ gen.DO }

type IFbSportsMatchStatisticsStageListDo interface {
	gen.SubQuery
	Debug() IFbSportsMatchStatisticsStageListDo
	WithContext(ctx context.Context) IFbSportsMatchStatisticsStageListDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFbSportsMatchStatisticsStageListDo
	WriteDB() IFbSportsMatchStatisticsStageListDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFbSportsMatchStatisticsStageListDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFbSportsMatchStatisticsStageListDo
	Not(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo
	Or(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo
	Select(conds ...field.Expr) IFbSportsMatchStatisticsStageListDo
	Where(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo
	Order(conds ...field.Expr) IFbSportsMatchStatisticsStageListDo
	Distinct(cols ...field.Expr) IFbSportsMatchStatisticsStageListDo
	Omit(cols ...field.Expr) IFbSportsMatchStatisticsStageListDo
	Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsStageListDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsStageListDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsStageListDo
	Group(cols ...field.Expr) IFbSportsMatchStatisticsStageListDo
	Having(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo
	Limit(limit int) IFbSportsMatchStatisticsStageListDo
	Offset(offset int) IFbSportsMatchStatisticsStageListDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchStatisticsStageListDo
	Unscoped() IFbSportsMatchStatisticsStageListDo
	Create(values ...*model.FbSportsMatchStatisticsStageList) error
	CreateInBatches(values []*model.FbSportsMatchStatisticsStageList, batchSize int) error
	Save(values ...*model.FbSportsMatchStatisticsStageList) error
	First() (*model.FbSportsMatchStatisticsStageList, error)
	Take() (*model.FbSportsMatchStatisticsStageList, error)
	Last() (*model.FbSportsMatchStatisticsStageList, error)
	Find() ([]*model.FbSportsMatchStatisticsStageList, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchStatisticsStageList, err error)
	FindInBatches(result *[]*model.FbSportsMatchStatisticsStageList, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FbSportsMatchStatisticsStageList) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFbSportsMatchStatisticsStageListDo
	Assign(attrs ...field.AssignExpr) IFbSportsMatchStatisticsStageListDo
	Joins(fields ...field.RelationField) IFbSportsMatchStatisticsStageListDo
	Preload(fields ...field.RelationField) IFbSportsMatchStatisticsStageListDo
	FirstOrInit() (*model.FbSportsMatchStatisticsStageList, error)
	FirstOrCreate() (*model.FbSportsMatchStatisticsStageList, error)
	FindByPage(offset int, limit int) (result []*model.FbSportsMatchStatisticsStageList, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFbSportsMatchStatisticsStageListDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f fbSportsMatchStatisticsStageListDo) Debug() IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Debug())
}

func (f fbSportsMatchStatisticsStageListDo) WithContext(ctx context.Context) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fbSportsMatchStatisticsStageListDo) ReadDB() IFbSportsMatchStatisticsStageListDo {
	return f.Clauses(dbresolver.Read)
}

func (f fbSportsMatchStatisticsStageListDo) WriteDB() IFbSportsMatchStatisticsStageListDo {
	return f.Clauses(dbresolver.Write)
}

func (f fbSportsMatchStatisticsStageListDo) Session(config *gorm.Session) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Session(config))
}

func (f fbSportsMatchStatisticsStageListDo) Clauses(conds ...clause.Expression) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fbSportsMatchStatisticsStageListDo) Returning(value interface{}, columns ...string) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fbSportsMatchStatisticsStageListDo) Not(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fbSportsMatchStatisticsStageListDo) Or(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fbSportsMatchStatisticsStageListDo) Select(conds ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fbSportsMatchStatisticsStageListDo) Where(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fbSportsMatchStatisticsStageListDo) Order(conds ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fbSportsMatchStatisticsStageListDo) Distinct(cols ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fbSportsMatchStatisticsStageListDo) Omit(cols ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fbSportsMatchStatisticsStageListDo) Join(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fbSportsMatchStatisticsStageListDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fbSportsMatchStatisticsStageListDo) RightJoin(table schema.Tabler, on ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fbSportsMatchStatisticsStageListDo) Group(cols ...field.Expr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fbSportsMatchStatisticsStageListDo) Having(conds ...gen.Condition) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fbSportsMatchStatisticsStageListDo) Limit(limit int) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fbSportsMatchStatisticsStageListDo) Offset(offset int) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fbSportsMatchStatisticsStageListDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fbSportsMatchStatisticsStageListDo) Unscoped() IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fbSportsMatchStatisticsStageListDo) Create(values ...*model.FbSportsMatchStatisticsStageList) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fbSportsMatchStatisticsStageListDo) CreateInBatches(values []*model.FbSportsMatchStatisticsStageList, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fbSportsMatchStatisticsStageListDo) Save(values ...*model.FbSportsMatchStatisticsStageList) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fbSportsMatchStatisticsStageListDo) First() (*model.FbSportsMatchStatisticsStageList, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsStageList), nil
	}
}

func (f fbSportsMatchStatisticsStageListDo) Take() (*model.FbSportsMatchStatisticsStageList, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsStageList), nil
	}
}

func (f fbSportsMatchStatisticsStageListDo) Last() (*model.FbSportsMatchStatisticsStageList, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsStageList), nil
	}
}

func (f fbSportsMatchStatisticsStageListDo) Find() ([]*model.FbSportsMatchStatisticsStageList, error) {
	result, err := f.DO.Find()
	return result.([]*model.FbSportsMatchStatisticsStageList), err
}

func (f fbSportsMatchStatisticsStageListDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FbSportsMatchStatisticsStageList, err error) {
	buf := make([]*model.FbSportsMatchStatisticsStageList, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fbSportsMatchStatisticsStageListDo) FindInBatches(result *[]*model.FbSportsMatchStatisticsStageList, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fbSportsMatchStatisticsStageListDo) Attrs(attrs ...field.AssignExpr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fbSportsMatchStatisticsStageListDo) Assign(attrs ...field.AssignExpr) IFbSportsMatchStatisticsStageListDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fbSportsMatchStatisticsStageListDo) Joins(fields ...field.RelationField) IFbSportsMatchStatisticsStageListDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fbSportsMatchStatisticsStageListDo) Preload(fields ...field.RelationField) IFbSportsMatchStatisticsStageListDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fbSportsMatchStatisticsStageListDo) FirstOrInit() (*model.FbSportsMatchStatisticsStageList, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsStageList), nil
	}
}

func (f fbSportsMatchStatisticsStageListDo) FirstOrCreate() (*model.FbSportsMatchStatisticsStageList, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FbSportsMatchStatisticsStageList), nil
	}
}

func (f fbSportsMatchStatisticsStageListDo) FindByPage(offset int, limit int) (result []*model.FbSportsMatchStatisticsStageList, count int64, err error) {
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

func (f fbSportsMatchStatisticsStageListDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fbSportsMatchStatisticsStageListDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fbSportsMatchStatisticsStageListDo) Delete(models ...*model.FbSportsMatchStatisticsStageList) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fbSportsMatchStatisticsStageListDo) withDO(do gen.Dao) *fbSportsMatchStatisticsStageListDo {
	f.DO = *do.(*gen.DO)
	return f
}
