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

func newWinAdminLoginLog(db *gorm.DB, opts ...gen.DOOption) winAdminLoginLog {
	_winAdminLoginLog := winAdminLoginLog{}

	_winAdminLoginLog.winAdminLoginLogDo.UseDB(db, opts...)
	_winAdminLoginLog.winAdminLoginLogDo.UseModel(&model.WinAdminLoginLog{})

	tableName := _winAdminLoginLog.winAdminLoginLogDo.TableName()
	_winAdminLoginLog.ALL = field.NewAsterisk(tableName)
	_winAdminLoginLog.ID = field.NewInt64(tableName, "id")
	_winAdminLoginLog.UID = field.NewInt64(tableName, "uid")
	_winAdminLoginLog.Username = field.NewString(tableName, "username")
	_winAdminLoginLog.IP = field.NewString(tableName, "ip")
	_winAdminLoginLog.UserAgent = field.NewString(tableName, "user_agent")
	_winAdminLoginLog.Category = field.NewInt64(tableName, "category")
	_winAdminLoginLog.CreatedAt = field.NewInt64(tableName, "created_at")
	_winAdminLoginLog.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winAdminLoginLog.fillFieldMap()

	return _winAdminLoginLog
}

// winAdminLoginLog 后台登录日志表
type winAdminLoginLog struct {
	winAdminLoginLogDo

	ALL       field.Asterisk
	ID        field.Int64
	UID       field.Int64  // UID
	Username  field.String // 用户名
	IP        field.String // IP地址
	UserAgent field.String // User-Agent
	Category  field.Int64  // 类型:0-登出 1-登录
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (w winAdminLoginLog) Table(newTableName string) *winAdminLoginLog {
	w.winAdminLoginLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAdminLoginLog) As(alias string) *winAdminLoginLog {
	w.winAdminLoginLogDo.DO = *(w.winAdminLoginLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAdminLoginLog) updateTableName(table string) *winAdminLoginLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.IP = field.NewString(table, "ip")
	w.UserAgent = field.NewString(table, "user_agent")
	w.Category = field.NewInt64(table, "category")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winAdminLoginLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAdminLoginLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 8)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["ip"] = w.IP
	w.fieldMap["user_agent"] = w.UserAgent
	w.fieldMap["category"] = w.Category
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winAdminLoginLog) clone(db *gorm.DB) winAdminLoginLog {
	w.winAdminLoginLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAdminLoginLog) replaceDB(db *gorm.DB) winAdminLoginLog {
	w.winAdminLoginLogDo.ReplaceDB(db)
	return w
}

type winAdminLoginLogDo struct{ gen.DO }

type IWinAdminLoginLogDo interface {
	gen.SubQuery
	Debug() IWinAdminLoginLogDo
	WithContext(ctx context.Context) IWinAdminLoginLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinAdminLoginLogDo
	WriteDB() IWinAdminLoginLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinAdminLoginLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinAdminLoginLogDo
	Not(conds ...gen.Condition) IWinAdminLoginLogDo
	Or(conds ...gen.Condition) IWinAdminLoginLogDo
	Select(conds ...field.Expr) IWinAdminLoginLogDo
	Where(conds ...gen.Condition) IWinAdminLoginLogDo
	Order(conds ...field.Expr) IWinAdminLoginLogDo
	Distinct(cols ...field.Expr) IWinAdminLoginLogDo
	Omit(cols ...field.Expr) IWinAdminLoginLogDo
	Join(table schema.Tabler, on ...field.Expr) IWinAdminLoginLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinAdminLoginLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinAdminLoginLogDo
	Group(cols ...field.Expr) IWinAdminLoginLogDo
	Having(conds ...gen.Condition) IWinAdminLoginLogDo
	Limit(limit int) IWinAdminLoginLogDo
	Offset(offset int) IWinAdminLoginLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAdminLoginLogDo
	Unscoped() IWinAdminLoginLogDo
	Create(values ...*model.WinAdminLoginLog) error
	CreateInBatches(values []*model.WinAdminLoginLog, batchSize int) error
	Save(values ...*model.WinAdminLoginLog) error
	First() (*model.WinAdminLoginLog, error)
	Take() (*model.WinAdminLoginLog, error)
	Last() (*model.WinAdminLoginLog, error)
	Find() ([]*model.WinAdminLoginLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAdminLoginLog, err error)
	FindInBatches(result *[]*model.WinAdminLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinAdminLoginLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinAdminLoginLogDo
	Assign(attrs ...field.AssignExpr) IWinAdminLoginLogDo
	Joins(fields ...field.RelationField) IWinAdminLoginLogDo
	Preload(fields ...field.RelationField) IWinAdminLoginLogDo
	FirstOrInit() (*model.WinAdminLoginLog, error)
	FirstOrCreate() (*model.WinAdminLoginLog, error)
	FindByPage(offset int, limit int) (result []*model.WinAdminLoginLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinAdminLoginLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winAdminLoginLogDo) Debug() IWinAdminLoginLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winAdminLoginLogDo) WithContext(ctx context.Context) IWinAdminLoginLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAdminLoginLogDo) ReadDB() IWinAdminLoginLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAdminLoginLogDo) WriteDB() IWinAdminLoginLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAdminLoginLogDo) Session(config *gorm.Session) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAdminLoginLogDo) Clauses(conds ...clause.Expression) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAdminLoginLogDo) Returning(value interface{}, columns ...string) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAdminLoginLogDo) Not(conds ...gen.Condition) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAdminLoginLogDo) Or(conds ...gen.Condition) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAdminLoginLogDo) Select(conds ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAdminLoginLogDo) Where(conds ...gen.Condition) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAdminLoginLogDo) Order(conds ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAdminLoginLogDo) Distinct(cols ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAdminLoginLogDo) Omit(cols ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAdminLoginLogDo) Join(table schema.Tabler, on ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAdminLoginLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAdminLoginLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAdminLoginLogDo) Group(cols ...field.Expr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAdminLoginLogDo) Having(conds ...gen.Condition) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAdminLoginLogDo) Limit(limit int) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAdminLoginLogDo) Offset(offset int) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAdminLoginLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAdminLoginLogDo) Unscoped() IWinAdminLoginLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAdminLoginLogDo) Create(values ...*model.WinAdminLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAdminLoginLogDo) CreateInBatches(values []*model.WinAdminLoginLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAdminLoginLogDo) Save(values ...*model.WinAdminLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAdminLoginLogDo) First() (*model.WinAdminLoginLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdminLoginLog), nil
	}
}

func (w winAdminLoginLogDo) Take() (*model.WinAdminLoginLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdminLoginLog), nil
	}
}

func (w winAdminLoginLogDo) Last() (*model.WinAdminLoginLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdminLoginLog), nil
	}
}

func (w winAdminLoginLogDo) Find() ([]*model.WinAdminLoginLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAdminLoginLog), err
}

func (w winAdminLoginLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAdminLoginLog, err error) {
	buf := make([]*model.WinAdminLoginLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAdminLoginLogDo) FindInBatches(result *[]*model.WinAdminLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAdminLoginLogDo) Attrs(attrs ...field.AssignExpr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAdminLoginLogDo) Assign(attrs ...field.AssignExpr) IWinAdminLoginLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAdminLoginLogDo) Joins(fields ...field.RelationField) IWinAdminLoginLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAdminLoginLogDo) Preload(fields ...field.RelationField) IWinAdminLoginLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAdminLoginLogDo) FirstOrInit() (*model.WinAdminLoginLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdminLoginLog), nil
	}
}

func (w winAdminLoginLogDo) FirstOrCreate() (*model.WinAdminLoginLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdminLoginLog), nil
	}
}

func (w winAdminLoginLogDo) FindByPage(offset int, limit int) (result []*model.WinAdminLoginLog, count int64, err error) {
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

func (w winAdminLoginLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAdminLoginLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAdminLoginLogDo) Delete(models ...*model.WinAdminLoginLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAdminLoginLogDo) withDO(do gen.Dao) *winAdminLoginLogDo {
	w.DO = *do.(*gen.DO)
	return w
}
