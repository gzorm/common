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

func newWinPromotions(db *gorm.DB, opts ...gen.DOOption) winPromotions {
	_winPromotions := winPromotions{}

	_winPromotions.winPromotionsDo.UseDB(db, opts...)
	_winPromotions.winPromotionsDo.UseModel(&model.WinPromotions{})

	tableName := _winPromotions.winPromotionsDo.TableName()
	_winPromotions.ALL = field.NewAsterisk(tableName)
	_winPromotions.ID = field.NewInt64(tableName, "id")
	_winPromotions.Code = field.NewString(tableName, "code")
	_winPromotions.CodeZh = field.NewString(tableName, "code_zh")
	_winPromotions.Amount = field.NewField(tableName, "amount")
	_winPromotions.Balance = field.NewField(tableName, "balance")
	_winPromotions.DescriptZh = field.NewString(tableName, "descript_zh")
	_winPromotions.Img = field.NewString(tableName, "img")
	_winPromotions.Category = field.NewString(tableName, "category")
	_winPromotions.GameType = field.NewInt64(tableName, "game_type")
	_winPromotions.Info = field.NewString(tableName, "info")
	_winPromotions.Descript = field.NewString(tableName, "descript")
	_winPromotions.StartedAt = field.NewInt64(tableName, "started_at")
	_winPromotions.Ladder = field.NewString(tableName, "ladder")
	_winPromotions.PayoutCategory = field.NewInt64(tableName, "payout_category")
	_winPromotions.EndedAt = field.NewInt64(tableName, "ended_at")
	_winPromotions.Sort = field.NewInt64(tableName, "sort")
	_winPromotions.Status = field.NewInt64(tableName, "status")
	_winPromotions.CreatedAt = field.NewInt64(tableName, "created_at")
	_winPromotions.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winPromotions.OperatorName = field.NewString(tableName, "operator_name")
	_winPromotions.Lang = field.NewString(tableName, "lang")

	_winPromotions.fillFieldMap()

	return _winPromotions
}

// winPromotions 优惠活动表
type winPromotions struct {
	winPromotionsDo

	ALL            field.Asterisk
	ID             field.Int64
	Code           field.String // 活动标识:首充优惠-First Deposit Bonus 续充优惠-Second Deposit Bonus 首单包赔-Risk-Free Bet 快乐周末-Happy Weekend Bonus
	CodeZh         field.String // 名称中文
	Amount         field.Field  // 总预算
	Balance        field.Field  // 总预算-剩余金额
	DescriptZh     field.String // 详细描述-中文
	Img            field.String // 图片
	Category       field.String // 类型:1-充值优惠 2-豪礼赠送 3-新活动 4-签到活动
	GameType       field.Int64  // 活动游戏类型，见字典dic_promotion_game_type
	Info           field.String // 补充信息
	Descript       field.String // 详情描述
	StartedAt      field.Int64  // 开始时间
	Ladder         field.String // 新活动阶梯
	PayoutCategory field.Int64  // 派彩类型: 0-自动派彩 1-人工派彩 2-手动派彩
	EndedAt        field.Int64  // 结算时间
	Sort           field.Int64  // 排序(从高到底、ID降序)
	Status         field.Int64  // 状态:1-启用 0-停用
	CreatedAt      field.Int64
	UpdatedAt      field.Int64
	OperatorName   field.String // 操作人姓名
	Lang           field.String // 语言

	fieldMap map[string]field.Expr
}

func (w winPromotions) Table(newTableName string) *winPromotions {
	w.winPromotionsDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winPromotions) As(alias string) *winPromotions {
	w.winPromotionsDo.DO = *(w.winPromotionsDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winPromotions) updateTableName(table string) *winPromotions {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Code = field.NewString(table, "code")
	w.CodeZh = field.NewString(table, "code_zh")
	w.Amount = field.NewField(table, "amount")
	w.Balance = field.NewField(table, "balance")
	w.DescriptZh = field.NewString(table, "descript_zh")
	w.Img = field.NewString(table, "img")
	w.Category = field.NewString(table, "category")
	w.GameType = field.NewInt64(table, "game_type")
	w.Info = field.NewString(table, "info")
	w.Descript = field.NewString(table, "descript")
	w.StartedAt = field.NewInt64(table, "started_at")
	w.Ladder = field.NewString(table, "ladder")
	w.PayoutCategory = field.NewInt64(table, "payout_category")
	w.EndedAt = field.NewInt64(table, "ended_at")
	w.Sort = field.NewInt64(table, "sort")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.OperatorName = field.NewString(table, "operator_name")
	w.Lang = field.NewString(table, "lang")

	w.fillFieldMap()

	return w
}

func (w *winPromotions) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winPromotions) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 21)
	w.fieldMap["id"] = w.ID
	w.fieldMap["code"] = w.Code
	w.fieldMap["code_zh"] = w.CodeZh
	w.fieldMap["amount"] = w.Amount
	w.fieldMap["balance"] = w.Balance
	w.fieldMap["descript_zh"] = w.DescriptZh
	w.fieldMap["img"] = w.Img
	w.fieldMap["category"] = w.Category
	w.fieldMap["game_type"] = w.GameType
	w.fieldMap["info"] = w.Info
	w.fieldMap["descript"] = w.Descript
	w.fieldMap["started_at"] = w.StartedAt
	w.fieldMap["ladder"] = w.Ladder
	w.fieldMap["payout_category"] = w.PayoutCategory
	w.fieldMap["ended_at"] = w.EndedAt
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["operator_name"] = w.OperatorName
	w.fieldMap["lang"] = w.Lang
}

func (w winPromotions) clone(db *gorm.DB) winPromotions {
	w.winPromotionsDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winPromotions) replaceDB(db *gorm.DB) winPromotions {
	w.winPromotionsDo.ReplaceDB(db)
	return w
}

type winPromotionsDo struct{ gen.DO }

type IWinPromotionsDo interface {
	gen.SubQuery
	Debug() IWinPromotionsDo
	WithContext(ctx context.Context) IWinPromotionsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinPromotionsDo
	WriteDB() IWinPromotionsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinPromotionsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinPromotionsDo
	Not(conds ...gen.Condition) IWinPromotionsDo
	Or(conds ...gen.Condition) IWinPromotionsDo
	Select(conds ...field.Expr) IWinPromotionsDo
	Where(conds ...gen.Condition) IWinPromotionsDo
	Order(conds ...field.Expr) IWinPromotionsDo
	Distinct(cols ...field.Expr) IWinPromotionsDo
	Omit(cols ...field.Expr) IWinPromotionsDo
	Join(table schema.Tabler, on ...field.Expr) IWinPromotionsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsDo
	Group(cols ...field.Expr) IWinPromotionsDo
	Having(conds ...gen.Condition) IWinPromotionsDo
	Limit(limit int) IWinPromotionsDo
	Offset(offset int) IWinPromotionsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionsDo
	Unscoped() IWinPromotionsDo
	Create(values ...*model.WinPromotions) error
	CreateInBatches(values []*model.WinPromotions, batchSize int) error
	Save(values ...*model.WinPromotions) error
	First() (*model.WinPromotions, error)
	Take() (*model.WinPromotions, error)
	Last() (*model.WinPromotions, error)
	Find() ([]*model.WinPromotions, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotions, err error)
	FindInBatches(result *[]*model.WinPromotions, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinPromotions) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinPromotionsDo
	Assign(attrs ...field.AssignExpr) IWinPromotionsDo
	Joins(fields ...field.RelationField) IWinPromotionsDo
	Preload(fields ...field.RelationField) IWinPromotionsDo
	FirstOrInit() (*model.WinPromotions, error)
	FirstOrCreate() (*model.WinPromotions, error)
	FindByPage(offset int, limit int) (result []*model.WinPromotions, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinPromotionsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winPromotionsDo) Debug() IWinPromotionsDo {
	return w.withDO(w.DO.Debug())
}

func (w winPromotionsDo) WithContext(ctx context.Context) IWinPromotionsDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winPromotionsDo) ReadDB() IWinPromotionsDo {
	return w.Clauses(dbresolver.Read)
}

func (w winPromotionsDo) WriteDB() IWinPromotionsDo {
	return w.Clauses(dbresolver.Write)
}

func (w winPromotionsDo) Session(config *gorm.Session) IWinPromotionsDo {
	return w.withDO(w.DO.Session(config))
}

func (w winPromotionsDo) Clauses(conds ...clause.Expression) IWinPromotionsDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winPromotionsDo) Returning(value interface{}, columns ...string) IWinPromotionsDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winPromotionsDo) Not(conds ...gen.Condition) IWinPromotionsDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winPromotionsDo) Or(conds ...gen.Condition) IWinPromotionsDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winPromotionsDo) Select(conds ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winPromotionsDo) Where(conds ...gen.Condition) IWinPromotionsDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winPromotionsDo) Order(conds ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winPromotionsDo) Distinct(cols ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winPromotionsDo) Omit(cols ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winPromotionsDo) Join(table schema.Tabler, on ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winPromotionsDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winPromotionsDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winPromotionsDo) Group(cols ...field.Expr) IWinPromotionsDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winPromotionsDo) Having(conds ...gen.Condition) IWinPromotionsDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winPromotionsDo) Limit(limit int) IWinPromotionsDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winPromotionsDo) Offset(offset int) IWinPromotionsDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winPromotionsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinPromotionsDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winPromotionsDo) Unscoped() IWinPromotionsDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winPromotionsDo) Create(values ...*model.WinPromotions) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winPromotionsDo) CreateInBatches(values []*model.WinPromotions, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winPromotionsDo) Save(values ...*model.WinPromotions) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winPromotionsDo) First() (*model.WinPromotions, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotions), nil
	}
}

func (w winPromotionsDo) Take() (*model.WinPromotions, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotions), nil
	}
}

func (w winPromotionsDo) Last() (*model.WinPromotions, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotions), nil
	}
}

func (w winPromotionsDo) Find() ([]*model.WinPromotions, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinPromotions), err
}

func (w winPromotionsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinPromotions, err error) {
	buf := make([]*model.WinPromotions, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winPromotionsDo) FindInBatches(result *[]*model.WinPromotions, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winPromotionsDo) Attrs(attrs ...field.AssignExpr) IWinPromotionsDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winPromotionsDo) Assign(attrs ...field.AssignExpr) IWinPromotionsDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winPromotionsDo) Joins(fields ...field.RelationField) IWinPromotionsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winPromotionsDo) Preload(fields ...field.RelationField) IWinPromotionsDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winPromotionsDo) FirstOrInit() (*model.WinPromotions, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotions), nil
	}
}

func (w winPromotionsDo) FirstOrCreate() (*model.WinPromotions, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinPromotions), nil
	}
}

func (w winPromotionsDo) FindByPage(offset int, limit int) (result []*model.WinPromotions, count int64, err error) {
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

func (w winPromotionsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winPromotionsDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winPromotionsDo) Delete(models ...*model.WinPromotions) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winPromotionsDo) withDO(do gen.Dao) *winPromotionsDo {
	w.DO = *do.(*gen.DO)
	return w
}
