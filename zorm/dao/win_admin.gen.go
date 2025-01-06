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

func newWinAdmin(db *gorm.DB, opts ...gen.DOOption) winAdmin {
	_winAdmin := winAdmin{}

	_winAdmin.winAdminDo.UseDB(db, opts...)
	_winAdmin.winAdminDo.UseModel(&model.WinAdmin{})

	tableName := _winAdmin.winAdminDo.TableName()
	_winAdmin.ALL = field.NewAsterisk(tableName)
	_winAdmin.ID = field.NewInt64(tableName, "id")
	_winAdmin.Username = field.NewString(tableName, "username")
	_winAdmin.MerchantID = field.NewInt64(tableName, "merchant_id")
	_winAdmin.AgentID = field.NewInt64(tableName, "agent_id")
	_winAdmin.PasswordHash = field.NewString(tableName, "password_hash")
	_winAdmin.AdminGroupID = field.NewInt64(tableName, "admin_group_id")
	_winAdmin.OperateID = field.NewInt64(tableName, "operate_id")
	_winAdmin.Parent = field.NewInt64(tableName, "parent")
	_winAdmin.Secret = field.NewString(tableName, "secret")
	_winAdmin.CodeURL = field.NewString(tableName, "code_url")
	_winAdmin.CodeStatus = field.NewInt64(tableName, "code_status")
	_winAdmin.Status = field.NewInt64(tableName, "status")
	_winAdmin.CreatedAt = field.NewInt64(tableName, "created_at")
	_winAdmin.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_winAdmin.Nickname = field.NewString(tableName, "nickname")
	_winAdmin.HomePath = field.NewString(tableName, "home_path")
	_winAdmin.Mobile = field.NewString(tableName, "mobile")
	_winAdmin.Email = field.NewString(tableName, "email")
	_winAdmin.Avatar = field.NewString(tableName, "avatar")
	_winAdmin.DepartmentID = field.NewInt64(tableName, "department_id")

	_winAdmin.fillFieldMap()

	return _winAdmin
}

// winAdmin 后台管理员账号
type winAdmin struct {
	winAdminDo

	ALL          field.Asterisk
	ID           field.Int64
	Username     field.String // 用户名
	MerchantID   field.Int64  // 商户id
	AgentID      field.Int64  // 代理ID
	PasswordHash field.String // 登录密码
	AdminGroupID field.Int64  // 用户组ID
	OperateID    field.Int64  // 操作人ID
	Parent       field.Int64  // 上级ID
	Secret       field.String // Google密钥
	CodeURL      field.String // google二维码
	CodeStatus   field.Int64  // google绑定验证记录:0=未绑定 ,1=已绑定
	Status       field.Int64  // 状态:10-正常 9-冻结 8-删除
	CreatedAt    field.Int64
	UpdatedAt    field.Int64
	Nickname     field.String // Nickname | 昵称
	HomePath     field.String // The home page that the user enters after logging in | 用户登陆后进入的首页
	Mobile       field.String // Mobile number
	Email        field.String // Email | 邮箱号
	Avatar       field.String // Avatar | 头像路径
	DepartmentID field.Int64  // Department ID | 部门ID

	fieldMap map[string]field.Expr
}

func (w winAdmin) Table(newTableName string) *winAdmin {
	w.winAdminDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w winAdmin) As(alias string) *winAdmin {
	w.winAdminDo.DO = *(w.winAdminDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *winAdmin) updateTableName(table string) *winAdmin {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.Username = field.NewString(table, "username")
	w.MerchantID = field.NewInt64(table, "merchant_id")
	w.AgentID = field.NewInt64(table, "agent_id")
	w.PasswordHash = field.NewString(table, "password_hash")
	w.AdminGroupID = field.NewInt64(table, "admin_group_id")
	w.OperateID = field.NewInt64(table, "operate_id")
	w.Parent = field.NewInt64(table, "parent")
	w.Secret = field.NewString(table, "secret")
	w.CodeURL = field.NewString(table, "code_url")
	w.CodeStatus = field.NewInt64(table, "code_status")
	w.Status = field.NewInt64(table, "status")
	w.CreatedAt = field.NewInt64(table, "created_at")
	w.UpdatedAt = field.NewInt64(table, "updated_at")
	w.Nickname = field.NewString(table, "nickname")
	w.HomePath = field.NewString(table, "home_path")
	w.Mobile = field.NewString(table, "mobile")
	w.Email = field.NewString(table, "email")
	w.Avatar = field.NewString(table, "avatar")
	w.DepartmentID = field.NewInt64(table, "department_id")

	w.fillFieldMap()

	return w
}

func (w *winAdmin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *winAdmin) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 20)
	w.fieldMap["id"] = w.ID
	w.fieldMap["username"] = w.Username
	w.fieldMap["merchant_id"] = w.MerchantID
	w.fieldMap["agent_id"] = w.AgentID
	w.fieldMap["password_hash"] = w.PasswordHash
	w.fieldMap["admin_group_id"] = w.AdminGroupID
	w.fieldMap["operate_id"] = w.OperateID
	w.fieldMap["parent"] = w.Parent
	w.fieldMap["secret"] = w.Secret
	w.fieldMap["code_url"] = w.CodeURL
	w.fieldMap["code_status"] = w.CodeStatus
	w.fieldMap["status"] = w.Status
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["nickname"] = w.Nickname
	w.fieldMap["home_path"] = w.HomePath
	w.fieldMap["mobile"] = w.Mobile
	w.fieldMap["email"] = w.Email
	w.fieldMap["avatar"] = w.Avatar
	w.fieldMap["department_id"] = w.DepartmentID
}

func (w winAdmin) clone(db *gorm.DB) winAdmin {
	w.winAdminDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w winAdmin) replaceDB(db *gorm.DB) winAdmin {
	w.winAdminDo.ReplaceDB(db)
	return w
}

type winAdminDo struct{ gen.DO }

type IWinAdminDo interface {
	gen.SubQuery
	Debug() IWinAdminDo
	WithContext(ctx context.Context) IWinAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWinAdminDo
	WriteDB() IWinAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWinAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWinAdminDo
	Not(conds ...gen.Condition) IWinAdminDo
	Or(conds ...gen.Condition) IWinAdminDo
	Select(conds ...field.Expr) IWinAdminDo
	Where(conds ...gen.Condition) IWinAdminDo
	Order(conds ...field.Expr) IWinAdminDo
	Distinct(cols ...field.Expr) IWinAdminDo
	Omit(cols ...field.Expr) IWinAdminDo
	Join(table schema.Tabler, on ...field.Expr) IWinAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWinAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWinAdminDo
	Group(cols ...field.Expr) IWinAdminDo
	Having(conds ...gen.Condition) IWinAdminDo
	Limit(limit int) IWinAdminDo
	Offset(offset int) IWinAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAdminDo
	Unscoped() IWinAdminDo
	Create(values ...*model.WinAdmin) error
	CreateInBatches(values []*model.WinAdmin, batchSize int) error
	Save(values ...*model.WinAdmin) error
	First() (*model.WinAdmin, error)
	Take() (*model.WinAdmin, error)
	Last() (*model.WinAdmin, error)
	Find() ([]*model.WinAdmin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAdmin, err error)
	FindInBatches(result *[]*model.WinAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.WinAdmin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWinAdminDo
	Assign(attrs ...field.AssignExpr) IWinAdminDo
	Joins(fields ...field.RelationField) IWinAdminDo
	Preload(fields ...field.RelationField) IWinAdminDo
	FirstOrInit() (*model.WinAdmin, error)
	FirstOrCreate() (*model.WinAdmin, error)
	FindByPage(offset int, limit int) (result []*model.WinAdmin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWinAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w winAdminDo) Debug() IWinAdminDo {
	return w.withDO(w.DO.Debug())
}

func (w winAdminDo) WithContext(ctx context.Context) IWinAdminDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w winAdminDo) ReadDB() IWinAdminDo {
	return w.Clauses(dbresolver.Read)
}

func (w winAdminDo) WriteDB() IWinAdminDo {
	return w.Clauses(dbresolver.Write)
}

func (w winAdminDo) Session(config *gorm.Session) IWinAdminDo {
	return w.withDO(w.DO.Session(config))
}

func (w winAdminDo) Clauses(conds ...clause.Expression) IWinAdminDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w winAdminDo) Returning(value interface{}, columns ...string) IWinAdminDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w winAdminDo) Not(conds ...gen.Condition) IWinAdminDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w winAdminDo) Or(conds ...gen.Condition) IWinAdminDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w winAdminDo) Select(conds ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w winAdminDo) Where(conds ...gen.Condition) IWinAdminDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w winAdminDo) Order(conds ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w winAdminDo) Distinct(cols ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w winAdminDo) Omit(cols ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w winAdminDo) Join(table schema.Tabler, on ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w winAdminDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w winAdminDo) RightJoin(table schema.Tabler, on ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w winAdminDo) Group(cols ...field.Expr) IWinAdminDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w winAdminDo) Having(conds ...gen.Condition) IWinAdminDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w winAdminDo) Limit(limit int) IWinAdminDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w winAdminDo) Offset(offset int) IWinAdminDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w winAdminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWinAdminDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w winAdminDo) Unscoped() IWinAdminDo {
	return w.withDO(w.DO.Unscoped())
}

func (w winAdminDo) Create(values ...*model.WinAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w winAdminDo) CreateInBatches(values []*model.WinAdmin, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w winAdminDo) Save(values ...*model.WinAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w winAdminDo) First() (*model.WinAdmin, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdmin), nil
	}
}

func (w winAdminDo) Take() (*model.WinAdmin, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdmin), nil
	}
}

func (w winAdminDo) Last() (*model.WinAdmin, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdmin), nil
	}
}

func (w winAdminDo) Find() ([]*model.WinAdmin, error) {
	result, err := w.DO.Find()
	return result.([]*model.WinAdmin), err
}

func (w winAdminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WinAdmin, err error) {
	buf := make([]*model.WinAdmin, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w winAdminDo) FindInBatches(result *[]*model.WinAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w winAdminDo) Attrs(attrs ...field.AssignExpr) IWinAdminDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w winAdminDo) Assign(attrs ...field.AssignExpr) IWinAdminDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w winAdminDo) Joins(fields ...field.RelationField) IWinAdminDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w winAdminDo) Preload(fields ...field.RelationField) IWinAdminDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w winAdminDo) FirstOrInit() (*model.WinAdmin, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdmin), nil
	}
}

func (w winAdminDo) FirstOrCreate() (*model.WinAdmin, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WinAdmin), nil
	}
}

func (w winAdminDo) FindByPage(offset int, limit int) (result []*model.WinAdmin, count int64, err error) {
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

func (w winAdminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w winAdminDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w winAdminDo) Delete(models ...*model.WinAdmin) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *winAdminDo) withDO(do gen.Dao) *winAdminDo {
	w.DO = *do.(*gen.DO)
	return w
}
