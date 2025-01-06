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

func newWinBetslips5(db *gorm.DB, opts ...gen.DOOption) winBetslips5 {
	_winBetslips5 := winBetslips5{}

	_winBetslips5.winBetslips5Do.UseDB(db, opts...)
	_winBetslips5.winBetslips5Do.UseModel(&model.WinBetslips5{})

	tableName := _winBetslips5.winBetslips5Do.TableName()
	_winBetslips5.ALL = field.NewAsterisk(tableName)
	_winBetslips5.ID = field.NewInt64(tableName, "id")
	_winBetslips5.RoundID = field.NewString(tableName, "round_id")
	_winBetslips5.TransactionID = field.NewString(tableName, "transaction_id")
	_winBetslips5.XbStatus = field.NewInt64(tableName, "xb_status")
	_winBetslips5.XbUID = field.NewInt64(tableName, "xb_uid")
	_winBetslips5.XbUsername = field.NewString(tableName, "xb_username")
	_winBetslips5.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winBetslips5.XbProfit = field.NewField(tableName, "xb_profit")
	_winBetslips5.Stake = field.NewField(tableName, "stake")
	_winBetslips5.ValidStake = field.NewField(tableName, "valid_stake")
	_winBetslips5.Payout = field.NewField(tableName, "payout")
	_winBetslips5.CoinRefund = field.NewField(tableName, "coin_refund")
	_winBetslips5.CoinBefore = field.NewField(tableName, "coin_before")
	_winBetslips5.GameProviderSubtypeID = field.NewInt64(tableName, "game_provider_subtype_id")
	_winBetslips5.GameListID = field.NewInt64(tableName, "game_list_id")
	_winBetslips5.GamePagcorID = field.NewInt64(tableName, "game_pagcor_id")
	_winBetslips5.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_winBetslips5.AmountType = field.NewInt64(tableName, "amount_type")
	_winBetslips5.DtStarted = field.NewInt64(tableName, "dt_started")
	_winBetslips5.DtCompleted = field.NewInt64(tableName, "dt_completed")
	_winBetslips5.WinTransactionID = field.NewString(tableName, "win_transaction_id")
	_winBetslips5.CreateTimeStr = field.NewString(tableName, "create_time_str")
	_winBetslips5.CreatedAt = field.NewInt64(tableName, "created_at")
	_winBetslips5.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winBetslips5.GameTypeID = field.NewInt64(tableName, "game_type_id")

	_winBetslips5.fillFieldMap()

	return _winBetslips5
}

type winBetslips5 struct {
	winBetslips5Do

	ALL                   field.Asterisk
	ID                    field.Int64  // 主键
	RoundID               field.String // 回合id
	TransactionID         field.String // 注单号 对应三方拉单transaction_id
	XbStatus              field.Int64  // 注单状态 1:待开彩  2:完成  3: 退款
	XbUID                 field.Int64  // 对应user表id
	XbUsername            field.String // 对应user表username
	MerchantID            field.Int64  // 商户id
	XbProfit              field.Field  // 盈亏金额
	Stake                 field.Field  // 投注
	ValidStake            field.Field  // 有效投注金额
	Payout                field.Field  // 派彩
	CoinRefund            field.Field  // 退款金额
	CoinBefore            field.Field  // 投注前金额
	GameProviderSubtypeID field.Int64  // 游戏id对应game_provider_subtype表id
	GameListID            field.Int64  // 游戏id对应game_list表id
	GamePagcorID          field.Int64  // pagcor分组id
	GameProviderID        field.Int64  // 游戏供应商id
	AmountType            field.Int64  // 投注方式 1:现金，2:奖金 3:免费旋转 4:活动免费旋转
	DtStarted             field.Int64  // 游戏开始时间
	DtCompleted           field.Int64  // 游戏结束时间
	WinTransactionID      field.String // 开奖交易单号
	CreateTimeStr         field.String // 投注时间
	CreatedAt             field.Int64
	UpdatedAt             field.Int64
	GameTypeID            field.Int64 // 游戏分组id

	fieldMap map[string]field.Expr
}

func (w winBetslips5) Table(newTableName string) *winBetslips5 {
	w.winBetslips5Do.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetslips5) As(alias string) *winBetslips5 {
	w.winBetslips5Do.DO = *(w.winBetslips5Do.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetslips5) updateTableName(table string) *winBetslips5 {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.RoundID = field.NewString(table, "round_id")
	w.TransactionID = field.NewString(table, "transaction_id")
	w.XbStatus = field.NewInt64(table, "xb_status")
	w.XbUID = field.NewInt64(table, "xb_uid")
	w.XbUsername = field.NewString(table, "xb_username")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.XbProfit = field.NewField(table, "xb_profit")
	w.Stake = field.NewField(table, "stake")
	w.ValidStake = field.NewField(table, "valid_stake")
	w.Payout = field.NewField(table, "payout")
	w.CoinRefund = field.NewField(table, "coin_refund")
	w.CoinBefore = field.NewField(table, "coin_before")
	w.GameProviderSubtypeID = field.NewInt64(table, "game_provider_subtype_id")
	w.GameListID = field.NewInt64(table, "game_list_id")
	w.GamePagcorID = field.NewInt64(table, "game_pagcor_id")
	w.GameProviderID = field.NewInt64(table, "game_provider_id")
	w.AmountType = field.NewInt64(table, "amount_type")
	w.DtStarted = field.NewInt64(table, "dt_started")
	w.DtCompleted = field.NewInt64(table, "dt_completed")
	w.WinTransactionID = field.NewString(table, "win_transaction_id")
	w.CreateTimeStr = field.NewString(table, "create_time_str")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.GameTypeID = field.NewInt64(table, "game_type_id")

	w.fillFieldMap()

	return w
}

func (w *winBetslips5) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetslips5) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 25)
	w.fieldMap["id"] = w.ID
	w.fieldMap["round_id"] = w.RoundID
	w.fieldMap["transaction_id"] = w.TransactionID
	w.fieldMap["xb_status"] = w.XbStatus
	w.fieldMap["xb_uid"] = w.XbUID
	w.fieldMap["xb_username"] = w.XbUsername
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["xb_profit"] = w.XbProfit
	w.fieldMap["stake"] = w.Stake
	w.fieldMap["valid_stake"] = w.ValidStake
	w.fieldMap["payout"] = w.Payout
	w.fieldMap["coin_refund"] = w.CoinRefund
	w.fieldMap["coin_before"] = w.CoinBefore
	w.fieldMap["game_provider_subtype_id"] = w.GameProviderSubtypeID
	w.fieldMap["game_list_id"] = w.GameListID
	w.fieldMap["game_pagcor_id"] = w.GamePagcorID
	w.fieldMap["game_provider_id"] = w.GameProviderID
	w.fieldMap["amount_type"] = w.AmountType
	w.fieldMap["dt_started"] = w.DtStarted
	w.fieldMap["dt_completed"] = w.DtCompleted
	w.fieldMap["win_transaction_id"] = w.WinTransactionID
	w.fieldMap["create_time_str"] = w.CreateTimeStr
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["game_type_id"] = w.GameTypeID
}

func (w winBetslips5) clone(db *gorm.DB) winBetslips5 {
	w.winBetslips5Do.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetslips5) replaceDB(db *gorm.DB) winBetslips5 {
	w.winBetslips5Do.ReplaceDB(db)
	return w
}

type winBetslips5Do struct{ gen.DO }

type IWinBetslips5Do interface {
	gen.SubQuery
	Debug() IWinBetslips5Do
	WithContext(ctx context.Context) IWinBetslips5Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBetslips5Do
	WriteDB() IWinBetslips5Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBetslips5Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBetslips5Do
	Not(conds ...gen.Condition) IWinBetslips5Do
	Or(conds ...gen.Condition) IWinBetslips5Do
	Select(conds ...field.Expr) IWinBetslips5Do
	Where(conds ...gen.Condition) IWinBetslips5Do
	Order(conds ...field.Expr) IWinBetslips5Do
	Distinct(cols ...field.Expr) IWinBetslips5Do
	Omit(cols ...field.Expr) IWinBetslips5Do
	Join(table schema.Tabler, on ...field.Expr) IWinBetslips5Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslips5Do
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslips5Do
	Group(cols ...field.Expr) IWinBetslips5Do
	Having(conds ...gen.Condition) IWinBetslips5Do
	Limit(limit int) IWinBetslips5Do
	Offset(offset int) IWinBetslips5Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslips5Do
	Unscoped() IWinBetslips5Do
	Create(values ...*model.WinBetslips5) error
	CreateInBatches(values []*model.WinBetslips5, batchSize int) error
	Save(values ...*model.WinBetslips5) error
	First() (*model.WinBetslips5, error)
	Take() (*model.WinBetslips5, error)
	Last() (*model.WinBetslips5, error)
	Find() ([]*model.WinBetslips5, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslips5, err error)
	FindInBatches(result *[]*model.WinBetslips5, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBetslips5) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBetslips5Do
	Assign(attrs ...field.AssignExpr) IWinBetslips5Do
	Joins(fields ...field.RelationField) IWinBetslips5Do
	Preload(fields ...field.RelationField) IWinBetslips5Do
	FirstOrInit() (*model.WinBetslips5, error)
	FirstOrCreate() (*model.WinBetslips5, error)
	FindByPage(offset int, limit int) (result []*model.WinBetslips5, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBetslips5Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBetslips5Do) Debug() IWinBetslips5Do {
	return w.withDO(w.DO.Debug())
}

func (w winBetslips5Do) WithContext(ctx context.Context) IWinBetslips5Do {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetslips5Do) ReadDB() IWinBetslips5Do {
	return w.Clauses(dbresolver.Read)
}

func (w winBetslips5Do) WriteDB() IWinBetslips5Do {
	return w.Clauses(dbresolver.Write)
}

func (w winBetslips5Do) Session(config *gorm.Session) IWinBetslips5Do {
	return w.withDO(w.DO.Session(config))
}

func (w winBetslips5Do) Clauses(conds ...clause.Expression) IWinBetslips5Do {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetslips5Do) Returning(value interface{}, columns ...string) IWinBetslips5Do {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetslips5Do) Not(conds ...gen.Condition) IWinBetslips5Do {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetslips5Do) Or(conds ...gen.Condition) IWinBetslips5Do {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetslips5Do) Select(conds ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetslips5Do) Where(conds ...gen.Condition) IWinBetslips5Do {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetslips5Do) Order(conds ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetslips5Do) Distinct(cols ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetslips5Do) Omit(cols ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetslips5Do) Join(table schema.Tabler, on ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetslips5Do) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetslips5Do) RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetslips5Do) Group(cols ...field.Expr) IWinBetslips5Do {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetslips5Do) Having(conds ...gen.Condition) IWinBetslips5Do {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetslips5Do) Limit(limit int) IWinBetslips5Do {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetslips5Do) Offset(offset int) IWinBetslips5Do {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetslips5Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslips5Do {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetslips5Do) Unscoped() IWinBetslips5Do {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetslips5Do) Create(values ...*model.WinBetslips5) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetslips5Do) CreateInBatches(values []*model.WinBetslips5, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetslips5Do) Save(values ...*model.WinBetslips5) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetslips5Do) First() (*model.WinBetslips5, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips5), nil
	}
}

func (w winBetslips5Do) Take() (*model.WinBetslips5, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips5), nil
	}
}

func (w winBetslips5Do) Last() (*model.WinBetslips5, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips5), nil
	}
}

func (w winBetslips5Do) Find() ([]*model.WinBetslips5, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetslips5), err
}

func (w winBetslips5Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslips5, err error) {
	buf := make([]*model.WinBetslips5, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetslips5Do) FindInBatches(result *[]*model.WinBetslips5, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetslips5Do) Attrs(attrs ...field.AssignExpr) IWinBetslips5Do {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetslips5Do) Assign(attrs ...field.AssignExpr) IWinBetslips5Do {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetslips5Do) Joins(fields ...field.RelationField) IWinBetslips5Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetslips5Do) Preload(fields ...field.RelationField) IWinBetslips5Do {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetslips5Do) FirstOrInit() (*model.WinBetslips5, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips5), nil
	}
}

func (w winBetslips5Do) FirstOrCreate() (*model.WinBetslips5, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslips5), nil
	}
}

func (w winBetslips5Do) FindByPage(offset int, limit int) (result []*model.WinBetslips5, count int64, err error) {
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

func (w winBetslips5Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetslips5Do) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetslips5Do) Delete(models ...*model.WinBetslips5) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetslips5Do) withDO(do gen.Dao) *winBetslips5Do {
	w.DO = *do.(*gen.DO)
	return w
}
