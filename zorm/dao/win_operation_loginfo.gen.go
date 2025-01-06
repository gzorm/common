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

func newWinOperationLoginfo(db *gorm.DB, opts ...gen.DOOption) winOperationLoginfo {
	_winOperationLoginfo := winOperationLoginfo{}

	_winOperationLoginfo.winOperationLoginfoDo.UseDB(db, opts...)
	_winOperationLoginfo.winOperationLoginfoDo.UseModel(&model.WinOperationLoginfo{})

	tableName := _winOperationLoginfo.winOperationLoginfoDo.TableName()
	_winOperationLoginfo.ALL = field.NewAsterisk(tableName)
	_winOperationLoginfo.ID = field.NewInt64(tableName, "id")
	_winOperationLoginfo.UID = field.NewInt64(tableName, "uid")
	_winOperationLoginfo.Username = field.NewString(tableName, "username")
	_winOperationLoginfo.URL = field.NewString(tableName, "url")
	_winOperationLoginfo.Method = field.NewString(tableName, "method")
	_winOperationLoginfo.RequestBody = field.NewString(tableName, "request_body")
	_winOperationLoginfo.ResponseStatus = field.NewInt64(tableName, "response_status")
	_winOperationLoginfo.ResponseBody = field.NewString(tableName, "response_body")
	_winOperationLoginfo.Content = field.NewString(tableName, "content")
	_winOperationLoginfo.IP = field.NewString(tableName, "ip")
	_winOperationLoginfo.CreatedAt = field.NewInt64(tableName, "created_at")
	_winOperationLoginfo.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winOperationLoginfo.fillFieldMap()

	return _winOperationLoginfo
}

// winOperationLoginfo 后台操作日志表
type winOperationLoginfo struct {
	winOperationLoginfoDo

	ALL            field.Asterisk
	ID             field.Int64
	UID            field.Int64  // 操作人ID
	Username       field.String // 操作人
	URL            field.String // 请求url
	Method         field.String // 菜单栏
	RequestBody    field.String // 请求体
	ResponseStatus field.Int64
	ResponseBody   field.String
	Content        field.String // 操作内容
	IP             field.String // 操作IP
	CreatedAt      field.Int64  // 操作时间
	UpdatedAt      field.Int64  // 更新时间

	fieldMap map[string]field.Expr
}

func (w winOperationLoginfo) Table(newTableName string) *winOperationLoginfo {
	w.winOperationLoginfoDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winOperationLoginfo) As(alias string) *winOperationLoginfo {
	w.winOperationLoginfoDo.DO = *(w.winOperationLoginfoDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winOperationLoginfo) updateTableName(table string) *winOperationLoginfo {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.URL = field.NewString(table, "url")
	w.Method = field.NewString(table, "method")
	w.RequestBody = field.NewString(table, "request_body")
	w.ResponseStatus = field.NewInt64(table, "response_status")
	w.ResponseBody = field.NewString(table, "response_body")
	w.Content = field.NewString(table, "content")
	w.IP = field.NewString(table, "ip")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winOperationLoginfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winOperationLoginfo) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 12)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["url"] = w.URL
	w.fieldMap["method"] = w.Method
	w.fieldMap["request_body"] = w.RequestBody
	w.fieldMap["response_status"] = w.ResponseStatus
	w.fieldMap["response_body"] = w.ResponseBody
	w.fieldMap["content"] = w.Content
	w.fieldMap["ip"] = w.IP
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winOperationLoginfo) clone(db *gorm.DB) winOperationLoginfo {
	w.winOperationLoginfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winOperationLoginfo) replaceDB(db *gorm.DB) winOperationLoginfo {
	w.winOperationLoginfoDo.ReplaceDB(db)
	return w
}

type winOperationLoginfoDo struct{ gen.DO }

type IWinOperationLoginfoDo interface {
	gen.SubQuery
	Debug() IWinOperationLoginfoDo
	WithContext(ctx context.Context) IWinOperationLoginfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinOperationLoginfoDo
	WriteDB() IWinOperationLoginfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinOperationLoginfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinOperationLoginfoDo
	Not(conds ...gen.Condition) IWinOperationLoginfoDo
	Or(conds ...gen.Condition) IWinOperationLoginfoDo
	Select(conds ...field.Expr) IWinOperationLoginfoDo
	Where(conds ...gen.Condition) IWinOperationLoginfoDo
	Order(conds ...field.Expr) IWinOperationLoginfoDo
	Distinct(cols ...field.Expr) IWinOperationLoginfoDo
	Omit(cols ...field.Expr) IWinOperationLoginfoDo
	Join(table schema.Tabler, on ...field.Expr) IWinOperationLoginfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinOperationLoginfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinOperationLoginfoDo
	Group(cols ...field.Expr) IWinOperationLoginfoDo
	Having(conds ...gen.Condition) IWinOperationLoginfoDo
	Limit(limit int) IWinOperationLoginfoDo
	Offset(offset int) IWinOperationLoginfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinOperationLoginfoDo
	Unscoped() IWinOperationLoginfoDo
	Create(values ...*model.WinOperationLoginfo) error
	CreateInBatches(values []*model.WinOperationLoginfo, batchSize int) error
	Save(values ...*model.WinOperationLoginfo) error
	First() (*model.WinOperationLoginfo, error)
	Take() (*model.WinOperationLoginfo, error)
	Last() (*model.WinOperationLoginfo, error)
	Find() ([]*model.WinOperationLoginfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinOperationLoginfo, err error)
	FindInBatches(result *[]*model.WinOperationLoginfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinOperationLoginfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinOperationLoginfoDo
	Assign(attrs ...field.AssignExpr) IWinOperationLoginfoDo
	Joins(fields ...field.RelationField) IWinOperationLoginfoDo
	Preload(fields ...field.RelationField) IWinOperationLoginfoDo
	FirstOrInit() (*model.WinOperationLoginfo, error)
	FirstOrCreate() (*model.WinOperationLoginfo, error)
	FindByPage(offset int, limit int) (result []*model.WinOperationLoginfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinOperationLoginfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winOperationLoginfoDo) Debug() IWinOperationLoginfoDo {
	return w.withDO(w.DO.Debug())
}

func (w winOperationLoginfoDo) WithContext(ctx context.Context) IWinOperationLoginfoDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winOperationLoginfoDo) ReadDB() IWinOperationLoginfoDo {
	return w.Clauses(dbresolver.Read)
}

func (w winOperationLoginfoDo) WriteDB() IWinOperationLoginfoDo {
	return w.Clauses(dbresolver.Write)
}

func (w winOperationLoginfoDo) Session(config *gorm.Session) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Session(config))
}

func (w winOperationLoginfoDo) Clauses(conds ...clause.Expression) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winOperationLoginfoDo) Returning(value interface{}, columns ...string) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winOperationLoginfoDo) Not(conds ...gen.Condition) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winOperationLoginfoDo) Or(conds ...gen.Condition) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winOperationLoginfoDo) Select(conds ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winOperationLoginfoDo) Where(conds ...gen.Condition) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winOperationLoginfoDo) Order(conds ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winOperationLoginfoDo) Distinct(cols ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winOperationLoginfoDo) Omit(cols ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winOperationLoginfoDo) Join(table schema.Tabler, on ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winOperationLoginfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winOperationLoginfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winOperationLoginfoDo) Group(cols ...field.Expr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winOperationLoginfoDo) Having(conds ...gen.Condition) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winOperationLoginfoDo) Limit(limit int) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winOperationLoginfoDo) Offset(offset int) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winOperationLoginfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winOperationLoginfoDo) Unscoped() IWinOperationLoginfoDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winOperationLoginfoDo) Create(values ...*model.WinOperationLoginfo) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winOperationLoginfoDo) CreateInBatches(values []*model.WinOperationLoginfo, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winOperationLoginfoDo) Save(values ...*model.WinOperationLoginfo) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winOperationLoginfoDo) First() (*model.WinOperationLoginfo, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperationLoginfo), nil
	}
}

func (w winOperationLoginfoDo) Take() (*model.WinOperationLoginfo, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperationLoginfo), nil
	}
}

func (w winOperationLoginfoDo) Last() (*model.WinOperationLoginfo, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperationLoginfo), nil
	}
}

func (w winOperationLoginfoDo) Find() ([]*model.WinOperationLoginfo, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinOperationLoginfo), err
}

func (w winOperationLoginfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinOperationLoginfo, err error) {
	buf := make([]*model.WinOperationLoginfo, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winOperationLoginfoDo) FindInBatches(result *[]*model.WinOperationLoginfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winOperationLoginfoDo) Attrs(attrs ...field.AssignExpr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winOperationLoginfoDo) Assign(attrs ...field.AssignExpr) IWinOperationLoginfoDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winOperationLoginfoDo) Joins(fields ...field.RelationField) IWinOperationLoginfoDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winOperationLoginfoDo) Preload(fields ...field.RelationField) IWinOperationLoginfoDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winOperationLoginfoDo) FirstOrInit() (*model.WinOperationLoginfo, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperationLoginfo), nil
	}
}

func (w winOperationLoginfoDo) FirstOrCreate() (*model.WinOperationLoginfo, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinOperationLoginfo), nil
	}
}

func (w winOperationLoginfoDo) FindByPage(offset int, limit int) (result []*model.WinOperationLoginfo, count int64, err error) {
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

func (w winOperationLoginfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winOperationLoginfoDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winOperationLoginfoDo) Delete(models ...*model.WinOperationLoginfo) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winOperationLoginfoDo) withDO(do gen.Dao) *winOperationLoginfoDo {
	w.DO = *do.(*gen.DO)
	return w
}
