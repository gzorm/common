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

func newAgentDailyReport(db *gorm.DB, opts ...gen.DOOption) agentDailyReport {
	_agentDailyReport := agentDailyReport{}

	_agentDailyReport.agentDailyReportDo.UseDB(db, opts...)
	_agentDailyReport.agentDailyReportDo.UseModel(&model.AgentDailyReport{})

	tableName := _agentDailyReport.agentDailyReportDo.TableName()
	_agentDailyReport.ALL = field.NewAsterisk(tableName)
	_agentDailyReport.ID = field.NewInt64(tableName, "id")
	_agentDailyReport.AgentID = field.NewInt64(tableName, "agent_id")
	_agentDailyReport.Ngr = field.NewField(tableName, "ngr")
	_agentDailyReport.DepositAmount = field.NewField(tableName, "deposit_amount")
	_agentDailyReport.WithdrawalAmount = field.NewField(tableName, "withdrawal_amount")
	_agentDailyReport.BetAmount = field.NewField(tableName, "bet_amount")
	_agentDailyReport.ValidAmount = field.NewField(tableName, "valid_amount")
	_agentDailyReport.PayoutAmount = field.NewField(tableName, "payout_amount")
	_agentDailyReport.TransferDepositAmount = field.NewField(tableName, "transfer_deposit_amount")
	_agentDailyReport.TransferWithdrawalAmount = field.NewField(tableName, "transfer_withdrawal_amount")
	_agentDailyReport.BonusAmount = field.NewField(tableName, "bonus_amount")
	_agentDailyReport.TransferBonusAmount = field.NewField(tableName, "transfer_bonus_amount")
	_agentDailyReport.AllUserCount = field.NewInt64(tableName, "all_user_count")
	_agentDailyReport.DepositCount = field.NewInt64(tableName, "deposit_count")
	_agentDailyReport.FirstDepositCount = field.NewInt64(tableName, "first_deposit_count")
	_agentDailyReport.SecondDepositCount = field.NewInt64(tableName, "second_deposit_count")
	_agentDailyReport.ThirdDepositCount = field.NewInt64(tableName, "third_deposit_count")
	_agentDailyReport.NewUserCount = field.NewInt64(tableName, "new_user_count")
	_agentDailyReport.BonusTypeCount = field.NewInt64(tableName, "bonus_type_count")
	_agentDailyReport.BonusCount = field.NewInt64(tableName, "bonus_count")
	_agentDailyReport.BetCount = field.NewInt64(tableName, "bet_count")
	_agentDailyReport.DepositRatio = field.NewField(tableName, "deposit_ratio")
	_agentDailyReport.FirstDepositAmount = field.NewField(tableName, "first_deposit_amount")
	_agentDailyReport.SecondDepositAmount = field.NewField(tableName, "second_deposit_amount")
	_agentDailyReport.ThirdDepositAmount = field.NewField(tableName, "third_deposit_amount")
	_agentDailyReport.OtherDepositAmount = field.NewField(tableName, "other_deposit_amount")
	_agentDailyReport.SubordinateCount = field.NewInt64(tableName, "subordinate_count")
	_agentDailyReport.NgrPayout = field.NewField(tableName, "ngr_payout")
	_agentDailyReport.NgrSetting = field.NewString(tableName, "ngr_setting")
	_agentDailyReport.ReportDate = field.NewInt64(tableName, "report_date")
	_agentDailyReport.CreatedAt = field.NewInt64(tableName, "created_at")
	_agentDailyReport.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_agentDailyReport.fillFieldMap()

	return _agentDailyReport
}

// agentDailyReport 代理每日報表
type agentDailyReport struct {
	agentDailyReportDo

	ALL                      field.Asterisk
	ID                       field.Int64  // id
	AgentID                  field.Int64  // 代理uid
	Ngr                      field.Field  // ngr=有效投注-派彩-三方稅收-活動支出-parcor稅收
	DepositAmount            field.Field  // 累计存款金额
	WithdrawalAmount         field.Field  // 累计提款金额
	BetAmount                field.Field  // 累计投注金额
	ValidAmount              field.Field  // 累计有效投注金额
	PayoutAmount             field.Field  // 累计派彩金额
	TransferDepositAmount    field.Field  // 人工存款金额
	TransferWithdrawalAmount field.Field  // 人工提款金额
	BonusAmount              field.Field  // 累计发放彩金金额
	TransferBonusAmount      field.Field  // 人工发放彩金金额
	AllUserCount             field.Int64  // 玩家总人数
	DepositCount             field.Int64  // 存款人数
	FirstDepositCount        field.Int64  // 首存人数
	SecondDepositCount       field.Int64  // 次存人数
	ThirdDepositCount        field.Int64  // 三存人数
	NewUserCount             field.Int64  // 新注册人数
	BonusTypeCount           field.Int64  // 彩金类型数量
	BonusCount               field.Int64  // 彩金发放人数
	BetCount                 field.Int64  // 累计投注人数
	DepositRatio             field.Field  // 玩家存款占比
	FirstDepositAmount       field.Field  // 首存金额
	SecondDepositAmount      field.Field  // 次存金额
	ThirdDepositAmount       field.Field  // 三存金额
	OtherDepositAmount       field.Field  // 其它存款金额
	SubordinateCount         field.Int64  // 下級代理人數
	NgrPayout                field.Field  // 代理NGR收益
	NgrSetting               field.String // NGR計算當時的配置
	ReportDate               field.Int64  // 報表日期
	CreatedAt                field.Int64  // 創建時間
	UpdatedAt                field.Int64  // 更新時間

	fieldMap map[string]field.Expr
}

func (a agentDailyReport) Table(newTableName string) *agentDailyReport {
	a.agentDailyReportDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a agentDailyReport) As(alias string) *agentDailyReport {
	a.agentDailyReportDo.DO = *(a.agentDailyReportDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *agentDailyReport) updateTableName(table string) *agentDailyReport {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.AgentID = field.NewInt64(table, "agent_id")
	a.Ngr = field.NewField(table, "ngr")
	a.DepositAmount = field.NewField(table, "deposit_amount")
	a.WithdrawalAmount = field.NewField(table, "withdrawal_amount")
	a.BetAmount = field.NewField(table, "bet_amount")
	a.ValidAmount = field.NewField(table, "valid_amount")
	a.PayoutAmount = field.NewField(table, "payout_amount")
	a.TransferDepositAmount = field.NewField(table, "transfer_deposit_amount")
	a.TransferWithdrawalAmount = field.NewField(table, "transfer_withdrawal_amount")
	a.BonusAmount = field.NewField(table, "bonus_amount")
	a.TransferBonusAmount = field.NewField(table, "transfer_bonus_amount")
	a.AllUserCount = field.NewInt64(table, "all_user_count")
	a.DepositCount = field.NewInt64(table, "deposit_count")
	a.FirstDepositCount = field.NewInt64(table, "first_deposit_count")
	a.SecondDepositCount = field.NewInt64(table, "second_deposit_count")
	a.ThirdDepositCount = field.NewInt64(table, "third_deposit_count")
	a.NewUserCount = field.NewInt64(table, "new_user_count")
	a.BonusTypeCount = field.NewInt64(table, "bonus_type_count")
	a.BonusCount = field.NewInt64(table, "bonus_count")
	a.BetCount = field.NewInt64(table, "bet_count")
	a.DepositRatio = field.NewField(table, "deposit_ratio")
	a.FirstDepositAmount = field.NewField(table, "first_deposit_amount")
	a.SecondDepositAmount = field.NewField(table, "second_deposit_amount")
	a.ThirdDepositAmount = field.NewField(table, "third_deposit_amount")
	a.OtherDepositAmount = field.NewField(table, "other_deposit_amount")
	a.SubordinateCount = field.NewInt64(table, "subordinate_count")
	a.NgrPayout = field.NewField(table, "ngr_payout")
	a.NgrSetting = field.NewString(table, "ngr_setting")
	a.ReportDate = field.NewInt64(table, "report_date")
	a.CreatedAt = field.NewInt64(table, "created_at")
	a.UpdatedAt = field.NewInt64(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *agentDailyReport) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *agentDailyReport) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 32)
	a.fieldMap["id"] = a.ID
	a.fieldMap["agent_id"] = a.AgentID
	a.fieldMap["ngr"] = a.Ngr
	a.fieldMap["deposit_amount"] = a.DepositAmount
	a.fieldMap["withdrawal_amount"] = a.WithdrawalAmount
	a.fieldMap["bet_amount"] = a.BetAmount
	a.fieldMap["valid_amount"] = a.ValidAmount
	a.fieldMap["payout_amount"] = a.PayoutAmount
	a.fieldMap["transfer_deposit_amount"] = a.TransferDepositAmount
	a.fieldMap["transfer_withdrawal_amount"] = a.TransferWithdrawalAmount
	a.fieldMap["bonus_amount"] = a.BonusAmount
	a.fieldMap["transfer_bonus_amount"] = a.TransferBonusAmount
	a.fieldMap["all_user_count"] = a.AllUserCount
	a.fieldMap["deposit_count"] = a.DepositCount
	a.fieldMap["first_deposit_count"] = a.FirstDepositCount
	a.fieldMap["second_deposit_count"] = a.SecondDepositCount
	a.fieldMap["third_deposit_count"] = a.ThirdDepositCount
	a.fieldMap["new_user_count"] = a.NewUserCount
	a.fieldMap["bonus_type_count"] = a.BonusTypeCount
	a.fieldMap["bonus_count"] = a.BonusCount
	a.fieldMap["bet_count"] = a.BetCount
	a.fieldMap["deposit_ratio"] = a.DepositRatio
	a.fieldMap["first_deposit_amount"] = a.FirstDepositAmount
	a.fieldMap["second_deposit_amount"] = a.SecondDepositAmount
	a.fieldMap["third_deposit_amount"] = a.ThirdDepositAmount
	a.fieldMap["other_deposit_amount"] = a.OtherDepositAmount
	a.fieldMap["subordinate_count"] = a.SubordinateCount
	a.fieldMap["ngr_payout"] = a.NgrPayout
	a.fieldMap["ngr_setting"] = a.NgrSetting
	a.fieldMap["report_date"] = a.ReportDate
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a agentDailyReport) clone(db *gorm.DB) agentDailyReport {
	a.agentDailyReportDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a agentDailyReport) replaceDB(db *gorm.DB) agentDailyReport {
	a.agentDailyReportDo.ReplaceDB(db)
	return a
}

type agentDailyReportDo struct{ gen.DO }

type IAgentDailyReportDo interface {
	gen.SubQuery
	Debug() IAgentDailyReportDo
	WithContext(ctx context.Context) IAgentDailyReportDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAgentDailyReportDo
	WriteDB() IAgentDailyReportDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAgentDailyReportDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAgentDailyReportDo
	Not(conds ...gen.Condition) IAgentDailyReportDo
	Or(conds ...gen.Condition) IAgentDailyReportDo
	Select(conds ...field.Expr) IAgentDailyReportDo
	Where(conds ...gen.Condition) IAgentDailyReportDo
	Order(conds ...field.Expr) IAgentDailyReportDo
	Distinct(cols ...field.Expr) IAgentDailyReportDo
	Omit(cols ...field.Expr) IAgentDailyReportDo
	Join(table schema.Tabler, on ...field.Expr) IAgentDailyReportDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAgentDailyReportDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAgentDailyReportDo
	Group(cols ...field.Expr) IAgentDailyReportDo
	Having(conds ...gen.Condition) IAgentDailyReportDo
	Limit(limit int) IAgentDailyReportDo
	Offset(offset int) IAgentDailyReportDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentDailyReportDo
	Unscoped() IAgentDailyReportDo
	Create(values ...*model.AgentDailyReport) error
	CreateInBatches(values []*model.AgentDailyReport, batchSize int) error
	Save(values ...*model.AgentDailyReport) error
	First() (*model.AgentDailyReport, error)
	Take() (*model.AgentDailyReport, error)
	Last() (*model.AgentDailyReport, error)
	Find() ([]*model.AgentDailyReport, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentDailyReport, err error)
	FindInBatches(result *[]*model.AgentDailyReport, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AgentDailyReport) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAgentDailyReportDo
	Assign(attrs ...field.AssignExpr) IAgentDailyReportDo
	Joins(fields ...field.RelationField) IAgentDailyReportDo
	Preload(fields ...field.RelationField) IAgentDailyReportDo
	FirstOrInit() (*model.AgentDailyReport, error)
	FirstOrCreate() (*model.AgentDailyReport, error)
	FindByPage(offset int, limit int) (result []*model.AgentDailyReport, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAgentDailyReportDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a agentDailyReportDo) Debug() IAgentDailyReportDo {
	return a.withDO(a.DO.Debug())
}

func (a agentDailyReportDo) WithContext(ctx context.Context) IAgentDailyReportDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a agentDailyReportDo) ReadDB() IAgentDailyReportDo {
	return a.Clauses(dbresolver.Read)
}

func (a agentDailyReportDo) WriteDB() IAgentDailyReportDo {
	return a.Clauses(dbresolver.Write)
}

func (a agentDailyReportDo) Session(config *gorm.Session) IAgentDailyReportDo {
	return a.withDO(a.DO.Session(config))
}

func (a agentDailyReportDo) Clauses(conds ...clause.Expression) IAgentDailyReportDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a agentDailyReportDo) Returning(value interface{}, columns ...string) IAgentDailyReportDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a agentDailyReportDo) Not(conds ...gen.Condition) IAgentDailyReportDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a agentDailyReportDo) Or(conds ...gen.Condition) IAgentDailyReportDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a agentDailyReportDo) Select(conds ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a agentDailyReportDo) Where(conds ...gen.Condition) IAgentDailyReportDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a agentDailyReportDo) Order(conds ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a agentDailyReportDo) Distinct(cols ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a agentDailyReportDo) Omit(cols ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a agentDailyReportDo) Join(table schema.Tabler, on ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a agentDailyReportDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a agentDailyReportDo) RightJoin(table schema.Tabler, on ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a agentDailyReportDo) Group(cols ...field.Expr) IAgentDailyReportDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a agentDailyReportDo) Having(conds ...gen.Condition) IAgentDailyReportDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a agentDailyReportDo) Limit(limit int) IAgentDailyReportDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a agentDailyReportDo) Offset(offset int) IAgentDailyReportDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a agentDailyReportDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAgentDailyReportDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a agentDailyReportDo) Unscoped() IAgentDailyReportDo {
	return a.withDO(a.DO.Unscoped())
}

func (a agentDailyReportDo) Create(values ...*model.AgentDailyReport) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a agentDailyReportDo) CreateInBatches(values []*model.AgentDailyReport, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a agentDailyReportDo) Save(values ...*model.AgentDailyReport) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a agentDailyReportDo) First() (*model.AgentDailyReport, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentDailyReport), nil
	}
}

func (a agentDailyReportDo) Take() (*model.AgentDailyReport, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentDailyReport), nil
	}
}

func (a agentDailyReportDo) Last() (*model.AgentDailyReport, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentDailyReport), nil
	}
}

func (a agentDailyReportDo) Find() ([]*model.AgentDailyReport, error) {
	result, err := a.DO.Find()
	return result.([]*model.AgentDailyReport), err
}

func (a agentDailyReportDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AgentDailyReport, err error) {
	buf := make([]*model.AgentDailyReport, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a agentDailyReportDo) FindInBatches(result *[]*model.AgentDailyReport, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a agentDailyReportDo) Attrs(attrs ...field.AssignExpr) IAgentDailyReportDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a agentDailyReportDo) Assign(attrs ...field.AssignExpr) IAgentDailyReportDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a agentDailyReportDo) Joins(fields ...field.RelationField) IAgentDailyReportDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a agentDailyReportDo) Preload(fields ...field.RelationField) IAgentDailyReportDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a agentDailyReportDo) FirstOrInit() (*model.AgentDailyReport, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentDailyReport), nil
	}
}

func (a agentDailyReportDo) FirstOrCreate() (*model.AgentDailyReport, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AgentDailyReport), nil
	}
}

func (a agentDailyReportDo) FindByPage(offset int, limit int) (result []*model.AgentDailyReport, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a agentDailyReportDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a agentDailyReportDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a agentDailyReportDo) Delete(models ...*model.AgentDailyReport) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *agentDailyReportDo) withDO(do gen.Dao) *agentDailyReportDo {
	a.DO = *do.(*gen.DO)
	return a
}