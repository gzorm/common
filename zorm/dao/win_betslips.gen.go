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

func newWinBetslip(db *gorm.DB, opts ...gen.DOOption) winBetslip {
	_winBetslip := winBetslip{}

	_winBetslip.winBetslipDo.UseDB(db, opts...)
	_winBetslip.winBetslipDo.UseModel(&model.WinBetslip{})

	tableName := _winBetslip.winBetslipDo.TableName()
	_winBetslip.ALL = field.NewAsterisk(tableName)
	_winBetslip.ID = field.NewInt64(tableName, "id")
	_winBetslip.RoundID = field.NewString(tableName, "round_id")
	_winBetslip.TransactionID = field.NewString(tableName, "transaction_id")
	_winBetslip.XbStatus = field.NewInt64(tableName, "xb_status")
	_winBetslip.XbUID = field.NewInt64(tableName, "xb_uid")
	_winBetslip.XbUsername = field.NewString(tableName, "xb_username")
	_winBetslip.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winBetslip.XbProfit = field.NewField(tableName, "xb_profit")
	_winBetslip.Stake = field.NewField(tableName, "stake")
	_winBetslip.ValidStake = field.NewField(tableName, "valid_stake")
	_winBetslip.Payout = field.NewField(tableName, "payout")
	_winBetslip.CoinRefund = field.NewField(tableName, "coin_refund")
	_winBetslip.CoinBefore = field.NewField(tableName, "coin_before")
	_winBetslip.GameProviderSubtypeID = field.NewInt64(tableName, "game_provider_subtype_id")
	_winBetslip.GameListID = field.NewInt64(tableName, "game_list_id")
	_winBetslip.GamePagcorID = field.NewInt64(tableName, "game_pagcor_id")
	_winBetslip.GameTypeID = field.NewInt64(tableName, "game_type_id")
	_winBetslip.GameProviderID = field.NewInt64(tableName, "game_provider_id")
	_winBetslip.AmountType = field.NewInt64(tableName, "amount_type")
	_winBetslip.DtStarted = field.NewInt64(tableName, "dt_started")
	_winBetslip.DtCompleted = field.NewInt64(tableName, "dt_completed")
	_winBetslip.WinTransactionID = field.NewString(tableName, "win_transaction_id")
	_winBetslip.CreateTimeStr = field.NewString(tableName, "create_time_str")
	_winBetslip.CreatedAt = field.NewInt64(tableName, "created_at")
	_winBetslip.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_winBetslip.fillFieldMap()

	return _winBetslip
}

type winBetslip struct {
	winBetslipDo

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
	GameTypeID            field.Int64  // 游戏分组id
	GameProviderID        field.Int64  // 游戏供应商id
	AmountType            field.Int64  // 投注方式 1:现金，2:奖金 3:免费旋转 4:活动免费旋转
	DtStarted             field.Int64  // 游戏开始时间
	DtCompleted           field.Int64  // 游戏结束时间
	WinTransactionID      field.String // 开奖交易单号
	CreateTimeStr         field.String // 投注时间
	CreatedAt             field.Int64  // 创建时间
	UpdatedAt             field.Int64  // 更新时间

	fieldMap map[string]field.Expr
}

func (w winBetslip) Table(newTableName string) *winBetslip {
	w.winBetslipDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winBetslip) As(alias string) *winBetslip {
	w.winBetslipDo.DO = *(w.winBetslipDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winBetslip) updateTableName(table string) *winBetslip {
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
	w.GameTypeID = field.NewInt64(table, "game_type_id")
	w.GameProviderID = field.NewInt64(table, "game_provider_id")
	w.AmountType = field.NewInt64(table, "amount_type")
	w.DtStarted = field.NewInt64(table, "dt_started")
	w.DtCompleted = field.NewInt64(table, "dt_completed")
	w.WinTransactionID = field.NewString(table, "win_transaction_id")
	w.CreateTimeStr = field.NewString(table, "create_time_str")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")

	w.fillFieldMap()

	return w
}

func (w *winBetslip) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winBetslip) fillFieldMap() {
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
	w.fieldMap["game_type_id"] = w.GameTypeID
	w.fieldMap["game_provider_id"] = w.GameProviderID
	w.fieldMap["amount_type"] = w.AmountType
	w.fieldMap["dt_started"] = w.DtStarted
	w.fieldMap["dt_completed"] = w.DtCompleted
	w.fieldMap["win_transaction_id"] = w.WinTransactionID
	w.fieldMap["create_time_str"] = w.CreateTimeStr
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
}

func (w winBetslip) clone(db *gorm.DB) winBetslip {
	w.winBetslipDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winBetslip) replaceDB(db *gorm.DB) winBetslip {
	w.winBetslipDo.ReplaceDB(db)
	return w
}

type winBetslipDo struct{ gen.DO }

type IWinBetslipDo interface {
	gen.SubQuery
	Debug() IWinBetslipDo
	WithContext(ctx context.Context) IWinBetslipDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinBetslipDo
	WriteDB() IWinBetslipDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinBetslipDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinBetslipDo
	Not(conds ...gen.Condition) IWinBetslipDo
	Or(conds ...gen.Condition) IWinBetslipDo
	Select(conds ...field.Expr) IWinBetslipDo
	Where(conds ...gen.Condition) IWinBetslipDo
	Order(conds ...field.Expr) IWinBetslipDo
	Distinct(cols ...field.Expr) IWinBetslipDo
	Omit(cols ...field.Expr) IWinBetslipDo
	Join(table schema.Tabler, on ...field.Expr) IWinBetslipDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipDo
	Group(cols ...field.Expr) IWinBetslipDo
	Having(conds ...gen.Condition) IWinBetslipDo
	Limit(limit int) IWinBetslipDo
	Offset(offset int) IWinBetslipDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipDo
	Unscoped() IWinBetslipDo
	Create(values ...*model.WinBetslip) error
	CreateInBatches(values []*model.WinBetslip, batchSize int) error
	Save(values ...*model.WinBetslip) error
	First() (*model.WinBetslip, error)
	Take() (*model.WinBetslip, error)
	Last() (*model.WinBetslip, error)
	Find() ([]*model.WinBetslip, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslip, err error)
	FindInBatches(result *[]*model.WinBetslip, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinBetslip) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinBetslipDo
	Assign(attrs ...field.AssignExpr) IWinBetslipDo
	Joins(fields ...field.RelationField) IWinBetslipDo
	Preload(fields ...field.RelationField) IWinBetslipDo
	FirstOrInit() (*model.WinBetslip, error)
	FirstOrCreate() (*model.WinBetslip, error)
	FindByPage(offset int, limit int) (result []*model.WinBetslip, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinBetslipDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winBetslipDo) Debug() IWinBetslipDo {
	return w.withDO(w.DO.Debug())
}

func (w winBetslipDo) WithContext(ctx context.Context) IWinBetslipDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winBetslipDo) ReadDB() IWinBetslipDo {
	return w.Clauses(dbresolver.Read)
}

func (w winBetslipDo) WriteDB() IWinBetslipDo {
	return w.Clauses(dbresolver.Write)
}

func (w winBetslipDo) Session(config *gorm.Session) IWinBetslipDo {
	return w.withDO(w.DO.Session(config))
}

func (w winBetslipDo) Clauses(conds ...clause.Expression) IWinBetslipDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winBetslipDo) Returning(value interface{}, columns ...string) IWinBetslipDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winBetslipDo) Not(conds ...gen.Condition) IWinBetslipDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winBetslipDo) Or(conds ...gen.Condition) IWinBetslipDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winBetslipDo) Select(conds ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winBetslipDo) Where(conds ...gen.Condition) IWinBetslipDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winBetslipDo) Order(conds ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winBetslipDo) Distinct(cols ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winBetslipDo) Omit(cols ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winBetslipDo) Join(table schema.Tabler, on ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winBetslipDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winBetslipDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winBetslipDo) Group(cols ...field.Expr) IWinBetslipDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winBetslipDo) Having(conds ...gen.Condition) IWinBetslipDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winBetslipDo) Limit(limit int) IWinBetslipDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winBetslipDo) Offset(offset int) IWinBetslipDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winBetslipDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinBetslipDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winBetslipDo) Unscoped() IWinBetslipDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winBetslipDo) Create(values ...*model.WinBetslip) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winBetslipDo) CreateInBatches(values []*model.WinBetslip, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winBetslipDo) Save(values ...*model.WinBetslip) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winBetslipDo) First() (*model.WinBetslip, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslip), nil
	}
}

func (w winBetslipDo) Take() (*model.WinBetslip, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslip), nil
	}
}

func (w winBetslipDo) Last() (*model.WinBetslip, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslip), nil
	}
}

func (w winBetslipDo) Find() ([]*model.WinBetslip, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinBetslip), err
}

func (w winBetslipDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinBetslip, err error) {
	buf := make([]*model.WinBetslip, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winBetslipDo) FindInBatches(result *[]*model.WinBetslip, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winBetslipDo) Attrs(attrs ...field.AssignExpr) IWinBetslipDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winBetslipDo) Assign(attrs ...field.AssignExpr) IWinBetslipDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winBetslipDo) Joins(fields ...field.RelationField) IWinBetslipDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winBetslipDo) Preload(fields ...field.RelationField) IWinBetslipDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winBetslipDo) FirstOrInit() (*model.WinBetslip, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslip), nil
	}
}

func (w winBetslipDo) FirstOrCreate() (*model.WinBetslip, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinBetslip), nil
	}
}

func (w winBetslipDo) FindByPage(offset int, limit int) (result []*model.WinBetslip, count int64, err error) {
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

func (w winBetslipDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winBetslipDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winBetslipDo) Delete(models ...*model.WinBetslip) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winBetslipDo) withDO(do gen.Dao) *winBetslipDo {
	w.DO = *do.(*gen.DO)
	return w
}
