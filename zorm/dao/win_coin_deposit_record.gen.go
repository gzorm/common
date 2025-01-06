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

func newWinCoinDepositRecord(db *gorm.DB, opts ...gen.DOOption) winCoinDepositRecord {
	_winCoinDepositRecord := winCoinDepositRecord{}

	_winCoinDepositRecord.winCoinDepositRecordDo.UseDB(db, opts...)
	_winCoinDepositRecord.winCoinDepositRecordDo.UseModel(&model.WinCoinDepositRecord{})

	tableName := _winCoinDepositRecord.winCoinDepositRecordDo.TableName()
	_winCoinDepositRecord.ALL = field.NewAsterisk(tableName)
	_winCoinDepositRecord.ID = field.NewInt64(tableName, "id")
	_winCoinDepositRecord.OrderID = field.NewString(tableName, "order_id")
	_winCoinDepositRecord.PlatOrderID = field.NewString(tableName, "plat_order_id")
	_winCoinDepositRecord.UID = field.NewInt64(tableName, "uid")
	_winCoinDepositRecord.Username = field.NewString(tableName, "username")
	_winCoinDepositRecord.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winCoinDepositRecord.Code = field.NewString(tableName, "code")
	_winCoinDepositRecord.PlatType = field.NewInt64(tableName, "plat_type")
	_winCoinDepositRecord.PlatName = field.NewString(tableName, "plat_name")
	_winCoinDepositRecord.PlatNickName = field.NewString(tableName, "plat_nick_name")
	_winCoinDepositRecord.CoinBefore = field.NewField(tableName, "coin_before")
	_winCoinDepositRecord.PayAddress = field.NewString(tableName, "pay_address")
	_winCoinDepositRecord.PayAmount = field.NewField(tableName, "pay_amount")
	_winCoinDepositRecord.ExchangeRate = field.NewField(tableName, "exchange_rate")
	_winCoinDepositRecord.RealAmount = field.NewField(tableName, "real_amount")
	_winCoinDepositRecord.Currency = field.NewString(tableName, "currency")
	_winCoinDepositRecord.DepStatus = field.NewInt64(tableName, "dep_status")
	_winCoinDepositRecord.Category = field.NewInt64(tableName, "category")
	_winCoinDepositRecord.CategoryCurrency = field.NewInt64(tableName, "category_currency")
	_winCoinDepositRecord.CategoryTransfer = field.NewInt64(tableName, "category_transfer")
	_winCoinDepositRecord.AdminUID = field.NewInt64(tableName, "admin_uid")
	_winCoinDepositRecord.Mark = field.NewString(tableName, "mark")
	_winCoinDepositRecord.Status = field.NewInt64(tableName, "status")
	_winCoinDepositRecord.ActivityID = field.NewInt64(tableName, "activity_id")
	_winCoinDepositRecord.CreatedAt = field.NewInt64(tableName, "created_at")
	_winCoinDepositRecord.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winCoinDepositRecord.IsCounted = field.NewInt64(tableName, "is_counted")

	_winCoinDepositRecord.fillFieldMap()

	return _winCoinDepositRecord
}

// winCoinDepositRecord 充值
type winCoinDepositRecord struct {
	winCoinDepositRecordDo

	ALL              field.Asterisk
	ID               field.Int64
	OrderID          field.String // 订单号(三方平台用)
	PlatOrderID      field.String // 三方平台订单号
	UID              field.Int64  // UID
	Username         field.String // 用户名
	MerchantID       field.Int64  // 门店ID
	Code             field.String // 支付通道编码
	PlatType         field.Int64  // 通道类型 话费支付=1，银行卡支付=3，钱包支付=5
	PlatName         field.String // 平台名称
	PlatNickName     field.String // 平台自定义名称
	CoinBefore       field.Field  // 充值前金额
	PayAddress       field.String // 加密地址
	PayAmount        field.Field  // 充值金额
	ExchangeRate     field.Field  // 汇率
	RealAmount       field.Field  // 到账金额
	Currency         field.String // 币种
	DepStatus        field.Int64  // 充值标识:1-首充 2-二充 9-其他
	Category         field.Int64  // 类型:0-钱包充值 1-佣金钱包转账充值
	CategoryCurrency field.Int64  // 货币类型:0-数字货币 1-法币
	CategoryTransfer field.Int64  // 转账类型:1-TRC-20 2-ERC-20 3-BANK 4-PIX 5-GCASH
	AdminUID         field.Int64  // 审核ID
	Mark             field.String // 备注
	Status           field.Int64  // 状态: 0-申请中 1-成功 2-失败
	ActivityID       field.Int64  // 参与活动ID
	CreatedAt        field.Int64
	UpdatedAt        field.Int64
	IsCounted        field.Int64 // 是否统计过：1=否，3=是

	fieldMap map[string]field.Expr
}

func (w winCoinDepositRecord) Table(newTableName string) *winCoinDepositRecord {
	w.winCoinDepositRecordDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winCoinDepositRecord) As(alias string) *winCoinDepositRecord {
	w.winCoinDepositRecordDo.DO = *(w.winCoinDepositRecordDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winCoinDepositRecord) updateTableName(table string) *winCoinDepositRecord {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.OrderID = field.NewString(table, "order_id")
	w.PlatOrderID = field.NewString(table, "plat_order_id")
	w.UID = field.NewInt64(table, "uid")
	w.Username = field.NewString(table, "username")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.Code = field.NewString(table, "code")
	w.PlatType = field.NewInt64(table, "plat_type")
	w.PlatName = field.NewString(table, "plat_name")
	w.PlatNickName = field.NewString(table, "plat_nick_name")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.PayAddress = field.NewString(table, "pay_address")
	w.PayAmount = field.NewField(table, "pay_amount")
	w.ExchangeRate = field.NewField(table, "exchange_rate")
	w.RealAmount = field.NewField(table, "real_amount")
	w.Currency = field.NewString(table, "currency")
	w.DepStatus = field.NewInt64(table, "dep_status")
	w.Category = field.NewInt64(table, "category")
	w.CategoryCurrency = field.NewInt64(table, "category_currency")
	w.CategoryTransfer = field.NewInt64(table, "category_transfer")
	w.AdminUID = field.NewInt64(table, "admin_uid")
	w.Mark = field.NewString(table, "mark")
	w.Status = field.NewInt64(table, "status")
	w.ActivityID = field.NewInt64(table, "activity_id")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.IsCounted = field.NewInt64(table, "is_counted")

	w.fillFieldMap()

	return w
}

func (w *winCoinDepositRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winCoinDepositRecord) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 27)
	w.fieldMap["id"] = w.ID
	w.fieldMap["order_id"] = w.OrderID
	w.fieldMap["plat_order_id"] = w.PlatOrderID
	w.fieldMap["uid"] = w.UID
	w.fieldMap["username"] = w.Username
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["code"] = w.Code
	w.fieldMap["plat_type"] = w.PlatType
	w.fieldMap["plat_name"] = w.PlatName
	w.fieldMap["plat_nick_name"] = w.PlatNickName
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["pay_address"] = w.PayAddress
	w.fieldMap["pay_amount"] = w.PayAmount
	w.fieldMap["exchange_rate"] = w.ExchangeRate
	w.fieldMap["real_amount"] = w.RealAmount
	w.fieldMap["currency"] = w.Currency
	w.fieldMap["dep_status"] = w.DepStatus
	w.fieldMap["category"] = w.Category
	w.fieldMap["category_currency"] = w.CategoryCurrency
	w.fieldMap["category_transfer"] = w.CategoryTransfer
	w.fieldMap["admin_uid"] = w.AdminUID
	w.fieldMap["mark"] = w.Mark
	w.fieldMap["status"] = w.Status
	w.fieldMap["activity_id"] = w.ActivityID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["is_counted"] = w.IsCounted
}

func (w winCoinDepositRecord) clone(db *gorm.DB) winCoinDepositRecord {
	w.winCoinDepositRecordDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winCoinDepositRecord) replaceDB(db *gorm.DB) winCoinDepositRecord {
	w.winCoinDepositRecordDo.ReplaceDB(db)
	return w
}

type winCoinDepositRecordDo struct{ gen.DO }

type IWinCoinDepositRecordDo interface {
	gen.SubQuery
	Debug() IWinCoinDepositRecordDo
	WithContext(ctx context.Context) IWinCoinDepositRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinCoinDepositRecordDo
	WriteDB() IWinCoinDepositRecordDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinCoinDepositRecordDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinCoinDepositRecordDo
	Not(conds ...gen.Condition) IWinCoinDepositRecordDo
	Or(conds ...gen.Condition) IWinCoinDepositRecordDo
	Select(conds ...field.Expr) IWinCoinDepositRecordDo
	Where(conds ...gen.Condition) IWinCoinDepositRecordDo
	Order(conds ...field.Expr) IWinCoinDepositRecordDo
	Distinct(cols ...field.Expr) IWinCoinDepositRecordDo
	Omit(cols ...field.Expr) IWinCoinDepositRecordDo
	Join(table schema.Tabler, on ...field.Expr) IWinCoinDepositRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinDepositRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinDepositRecordDo
	Group(cols ...field.Expr) IWinCoinDepositRecordDo
	Having(conds ...gen.Condition) IWinCoinDepositRecordDo
	Limit(limit int) IWinCoinDepositRecordDo
	Offset(offset int) IWinCoinDepositRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinDepositRecordDo
	Unscoped() IWinCoinDepositRecordDo
	Create(values ...*model.WinCoinDepositRecord) error
	CreateInBatches(values []*model.WinCoinDepositRecord, batchSize int) error
	Save(values ...*model.WinCoinDepositRecord) error
	First() (*model.WinCoinDepositRecord, error)
	Take() (*model.WinCoinDepositRecord, error)
	Last() (*model.WinCoinDepositRecord, error)
	Find() ([]*model.WinCoinDepositRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinDepositRecord, err error)
	FindInBatches(result *[]*model.WinCoinDepositRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinCoinDepositRecord) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinCoinDepositRecordDo
	Assign(attrs ...field.AssignExpr) IWinCoinDepositRecordDo
	Joins(fields ...field.RelationField) IWinCoinDepositRecordDo
	Preload(fields ...field.RelationField) IWinCoinDepositRecordDo
	FirstOrInit() (*model.WinCoinDepositRecord, error)
	FirstOrCreate() (*model.WinCoinDepositRecord, error)
	FindByPage(offset int, limit int) (result []*model.WinCoinDepositRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinCoinDepositRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winCoinDepositRecordDo) Debug() IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Debug())
}

func (w winCoinDepositRecordDo) WithContext(ctx context.Context) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winCoinDepositRecordDo) ReadDB() IWinCoinDepositRecordDo {
	return w.Clauses(dbresolver.Read)
}

func (w winCoinDepositRecordDo) WriteDB() IWinCoinDepositRecordDo {
	return w.Clauses(dbresolver.Write)
}

func (w winCoinDepositRecordDo) Session(config *gorm.Session) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Session(config))
}

func (w winCoinDepositRecordDo) Clauses(conds ...clause.Expression) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winCoinDepositRecordDo) Returning(value interface{}, columns ...string) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winCoinDepositRecordDo) Not(conds ...gen.Condition) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winCoinDepositRecordDo) Or(conds ...gen.Condition) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winCoinDepositRecordDo) Select(conds ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winCoinDepositRecordDo) Where(conds ...gen.Condition) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winCoinDepositRecordDo) Order(conds ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winCoinDepositRecordDo) Distinct(cols ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winCoinDepositRecordDo) Omit(cols ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winCoinDepositRecordDo) Join(table schema.Tabler, on ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winCoinDepositRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winCoinDepositRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winCoinDepositRecordDo) Group(cols ...field.Expr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winCoinDepositRecordDo) Having(conds ...gen.Condition) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winCoinDepositRecordDo) Limit(limit int) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winCoinDepositRecordDo) Offset(offset int) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winCoinDepositRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winCoinDepositRecordDo) Unscoped() IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winCoinDepositRecordDo) Create(values ...*model.WinCoinDepositRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winCoinDepositRecordDo) CreateInBatches(values []*model.WinCoinDepositRecord, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winCoinDepositRecordDo) Save(values ...*model.WinCoinDepositRecord) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winCoinDepositRecordDo) First() (*model.WinCoinDepositRecord, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinDepositRecord), nil
	}
}

func (w winCoinDepositRecordDo) Take() (*model.WinCoinDepositRecord, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinDepositRecord), nil
	}
}

func (w winCoinDepositRecordDo) Last() (*model.WinCoinDepositRecord, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinDepositRecord), nil
	}
}

func (w winCoinDepositRecordDo) Find() ([]*model.WinCoinDepositRecord, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinCoinDepositRecord), err
}

func (w winCoinDepositRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinCoinDepositRecord, err error) {
	buf := make([]*model.WinCoinDepositRecord, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winCoinDepositRecordDo) FindInBatches(result *[]*model.WinCoinDepositRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winCoinDepositRecordDo) Attrs(attrs ...field.AssignExpr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winCoinDepositRecordDo) Assign(attrs ...field.AssignExpr) IWinCoinDepositRecordDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winCoinDepositRecordDo) Joins(fields ...field.RelationField) IWinCoinDepositRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winCoinDepositRecordDo) Preload(fields ...field.RelationField) IWinCoinDepositRecordDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winCoinDepositRecordDo) FirstOrInit() (*model.WinCoinDepositRecord, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinDepositRecord), nil
	}
}

func (w winCoinDepositRecordDo) FirstOrCreate() (*model.WinCoinDepositRecord, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinCoinDepositRecord), nil
	}
}

func (w winCoinDepositRecordDo) FindByPage(offset int, limit int) (result []*model.WinCoinDepositRecord, count int64, err error) {
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

func (w winCoinDepositRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winCoinDepositRecordDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winCoinDepositRecordDo) Delete(models ...*model.WinCoinDepositRecord) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winCoinDepositRecordDo) withDO(do gen.Dao) *winCoinDepositRecordDo {
	w.DO = *do.(*gen.DO)
	return w
}
