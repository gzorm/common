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

func newWinUserLoginLog(db *gorm.DB, opts ...gen.DOOption) winUserLoginLog {
	_winUserLoginLog := winUserLoginLog{}

	_winUserLoginLog.winUserLoginLogDo.UseDB(db, opts...)
	_winUserLoginLog.winUserLoginLogDo.UseModel(&model.WinUserLoginLog{})

	tableName := _winUserLoginLog.winUserLoginLogDo.TableName()
	_winUserLoginLog.ALL = field.NewAsterisk(tableName)
	_winUserLoginLog.ID = field.NewInt64(tableName, "id")
	_winUserLoginLog.UID = field.NewInt64(tableName, "uid")
	_winUserLoginLog.Username = field.NewString(tableName, "username")
	_winUserLoginLog.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winUserLoginLog.Coin = field.NewField(tableName, "coin")
	_winUserLoginLog.IP = field.NewString(tableName, "ip")
	_winUserLoginLog.GameName = field.NewString(tableName, "game_name")
	_winUserLoginLog.URL = field.NewString(tableName, "url")
	_winUserLoginLog.Device = field.NewString(tableName, "device")
	_winUserLoginLog.Category = field.NewInt64(tableName, "category")
	_winUserLoginLog.DeviceID = field.NewString(tableName, "device_id")
	_winUserLoginLog.IsNewDevice = field.NewInt64(tableName, "is_new_device")
	_winUserLoginLog.Remark = field.NewString(tableName, "remark")
	_winUserLoginLog.CreatedAt = field.NewInt64(tableName, "created_at")
	_winUserLoginLog.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winUserLoginLog.fillFieldMap()

	return _winUserLoginLog
}

// winUserLoginLog 用户日志表
type winUserLoginLog struct {
	winUserLoginLogDo

	ALL         field.Asterisk
	ID          field.Int64
	UID         field.Int64  // UID
	Username    field.String // 用户名
	MerchantID  field.Int64  // 商户id
	Coin        field.Field  // 账户余额
	IP          field.String // IP地址
	GameName    field.String // 游戏名称
	URL         field.String // 登录链接
	Device      field.String // 设备:m-手机 d-电脑 ANDROID-安卓 IOS-苹果
	Category    field.Int64  // 类型:0-登出 1-登录 2-进入游戏
	DeviceID    field.String // 设备id
	IsNewDevice field.Int64  // 是新设备id 0-不是 1-是
	Remark      field.String // 备注
	CreatedAt   field.Int64
	UpdatedAt   field.Int64

	fieldMap map[string]field.Expr
}

func (w winUserLoginLog) Table(newTableName string) *winUserLoginLog {
	w.winUserLoginLogDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winUserLoginLog) As(alias string) *winUserLoginLog {
	w.winUserLoginLogDo.DO = *(w.winUserLoginLogDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winUserLoginLog) updateTableName(table string) *winUserLoginLog {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.Coin = field.NewField(table, "coin")
	w.IP = field.NewString(table, "ip")
	w.GameName = field.NewString(table, "game_name")
	w.URL = field.NewString(table, "url")
	w.Device = field.NewString(table, "device")
	w.Category = field.NewInt64(table, "category")
	w.DeviceID = field.NewString(table, "device_id")
	w.IsNewDevice = field.NewInt64(table, "is_new_device")
	w.Remark = field.NewString(table, "remark")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winUserLoginLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winUserLoginLog) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 15)
	w.fieldMap["id"] = w.ID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["coin"] = w.Coin
	w.fieldMap["ip"] = w.IP
	w.fieldMap["game_name"] = w.GameName
	w.fieldMap["url"] = w.URL
	w.fieldMap["device"] = w.Device
	w.fieldMap["category"] = w.Category
	w.fieldMap["device_id"] = w.DeviceID
	w.fieldMap["is_new_device"] = w.IsNewDevice
	w.fieldMap["remark"] = w.Remark
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winUserLoginLog) clone(db *gorm.DB) winUserLoginLog {
	w.winUserLoginLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winUserLoginLog) replaceDB(db *gorm.DB) winUserLoginLog {
	w.winUserLoginLogDo.ReplaceDB(db)
	return w
}

type winUserLoginLogDo struct{ gen.DO }

type IWinUserLoginLogDo interface {
	gen.SubQuery
	Debug() IWinUserLoginLogDo
	WithContext(ctx context.Context) IWinUserLoginLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinUserLoginLogDo
	WriteDB() IWinUserLoginLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinUserLoginLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinUserLoginLogDo
	Not(conds ...gen.Condition) IWinUserLoginLogDo
	Or(conds ...gen.Condition) IWinUserLoginLogDo
	Select(conds ...field.Expr) IWinUserLoginLogDo
	Where(conds ...gen.Condition) IWinUserLoginLogDo
	Order(conds ...field.Expr) IWinUserLoginLogDo
	Distinct(cols ...field.Expr) IWinUserLoginLogDo
	Omit(cols ...field.Expr) IWinUserLoginLogDo
	Join(table schema.Tabler, on ...field.Expr) IWinUserLoginLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLoginLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLoginLogDo
	Group(cols ...field.Expr) IWinUserLoginLogDo
	Having(conds ...gen.Condition) IWinUserLoginLogDo
	Limit(limit int) IWinUserLoginLogDo
	Offset(offset int) IWinUserLoginLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLoginLogDo
	Unscoped() IWinUserLoginLogDo
	Create(values ...*model.WinUserLoginLog) error
	CreateInBatches(values []*model.WinUserLoginLog, batchSize int) error
	Save(values ...*model.WinUserLoginLog) error
	First() (*model.WinUserLoginLog, error)
	Take() (*model.WinUserLoginLog, error)
	Last() (*model.WinUserLoginLog, error)
	Find() ([]*model.WinUserLoginLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLoginLog, err error)
	FindInBatches(result *[]*model.WinUserLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinUserLoginLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinUserLoginLogDo
	Assign(attrs ...field.AssignExpr) IWinUserLoginLogDo
	Joins(fields ...field.RelationField) IWinUserLoginLogDo
	Preload(fields ...field.RelationField) IWinUserLoginLogDo
	FirstOrInit() (*model.WinUserLoginLog, error)
	FirstOrCreate() (*model.WinUserLoginLog, error)
	FindByPage(offset int, limit int) (result []*model.WinUserLoginLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinUserLoginLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winUserLoginLogDo) Debug() IWinUserLoginLogDo {
	return w.withDO(w.DO.Debug())
}

func (w winUserLoginLogDo) WithContext(ctx context.Context) IWinUserLoginLogDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winUserLoginLogDo) ReadDB() IWinUserLoginLogDo {
	return w.Clauses(dbresolver.Read)
}

func (w winUserLoginLogDo) WriteDB() IWinUserLoginLogDo {
	return w.Clauses(dbresolver.Write)
}

func (w winUserLoginLogDo) Session(config *gorm.Session) IWinUserLoginLogDo {
	return w.withDO(w.DO.Session(config))
}

func (w winUserLoginLogDo) Clauses(conds ...clause.Expression) IWinUserLoginLogDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winUserLoginLogDo) Returning(value interface{}, columns ...string) IWinUserLoginLogDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winUserLoginLogDo) Not(conds ...gen.Condition) IWinUserLoginLogDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winUserLoginLogDo) Or(conds ...gen.Condition) IWinUserLoginLogDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winUserLoginLogDo) Select(conds ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winUserLoginLogDo) Where(conds ...gen.Condition) IWinUserLoginLogDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winUserLoginLogDo) Order(conds ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winUserLoginLogDo) Distinct(cols ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winUserLoginLogDo) Omit(cols ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winUserLoginLogDo) Join(table schema.Tabler, on ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winUserLoginLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winUserLoginLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winUserLoginLogDo) Group(cols ...field.Expr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winUserLoginLogDo) Having(conds ...gen.Condition) IWinUserLoginLogDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winUserLoginLogDo) Limit(limit int) IWinUserLoginLogDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winUserLoginLogDo) Offset(offset int) IWinUserLoginLogDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winUserLoginLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinUserLoginLogDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winUserLoginLogDo) Unscoped() IWinUserLoginLogDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winUserLoginLogDo) Create(values ...*model.WinUserLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winUserLoginLogDo) CreateInBatches(values []*model.WinUserLoginLog, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winUserLoginLogDo) Save(values ...*model.WinUserLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winUserLoginLogDo) First() (*model.WinUserLoginLog, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLoginLog), nil
	}
}

func (w winUserLoginLogDo) Take() (*model.WinUserLoginLog, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLoginLog), nil
	}
}

func (w winUserLoginLogDo) Last() (*model.WinUserLoginLog, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLoginLog), nil
	}
}

func (w winUserLoginLogDo) Find() ([]*model.WinUserLoginLog, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinUserLoginLog), err
}

func (w winUserLoginLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinUserLoginLog, err error) {
	buf := make([]*model.WinUserLoginLog, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winUserLoginLogDo) FindInBatches(result *[]*model.WinUserLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winUserLoginLogDo) Attrs(attrs ...field.AssignExpr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winUserLoginLogDo) Assign(attrs ...field.AssignExpr) IWinUserLoginLogDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winUserLoginLogDo) Joins(fields ...field.RelationField) IWinUserLoginLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winUserLoginLogDo) Preload(fields ...field.RelationField) IWinUserLoginLogDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winUserLoginLogDo) FirstOrInit() (*model.WinUserLoginLog, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLoginLog), nil
	}
}

func (w winUserLoginLogDo) FirstOrCreate() (*model.WinUserLoginLog, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinUserLoginLog), nil
	}
}

func (w winUserLoginLogDo) FindByPage(offset int, limit int) (result []*model.WinUserLoginLog, count int64, err error) {
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

func (w winUserLoginLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winUserLoginLogDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winUserLoginLogDo) Delete(models ...*model.WinUserLoginLog) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winUserLoginLogDo) withDO(do gen.Dao) *winUserLoginLogDo {
	w.DO = *do.(*gen.DO)
	return w
}
