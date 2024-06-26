// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	gen2 "github.com/ByteBam/thirftbam/biz/model/gen"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newAuthGiteeUser(db *gorm.DB, opts ...gen.DOOption) authGiteeUser {
	_authGiteeUser := authGiteeUser{}

	_authGiteeUser.authGiteeUserDo.UseDB(db, opts...)
	_authGiteeUser.authGiteeUserDo.UseModel(&gen2.AuthGiteeUser{})

	tableName := _authGiteeUser.authGiteeUserDo.TableName()
	_authGiteeUser.ALL = field.NewAsterisk(tableName)
	_authGiteeUser.GiteeID = field.NewInt64(tableName, "gitee_id")
	_authGiteeUser.GiteeUserName = field.NewString(tableName, "gitee_user_name")
	_authGiteeUser.GiteeUserAvatar = field.NewString(tableName, "gitee_user_avatar")
	_authGiteeUser.GiteeNickName = field.NewString(tableName, "gitee_nick_name")

	_authGiteeUser.fillFieldMap()

	return _authGiteeUser
}

type authGiteeUser struct {
	authGiteeUserDo authGiteeUserDo

	ALL             field.Asterisk
	GiteeID         field.Int64
	GiteeUserName   field.String
	GiteeUserAvatar field.String
	GiteeNickName   field.String

	fieldMap map[string]field.Expr
}

func (a authGiteeUser) Table(newTableName string) *authGiteeUser {
	a.authGiteeUserDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authGiteeUser) As(alias string) *authGiteeUser {
	a.authGiteeUserDo.DO = *(a.authGiteeUserDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authGiteeUser) updateTableName(table string) *authGiteeUser {
	a.ALL = field.NewAsterisk(table)
	a.GiteeID = field.NewInt64(table, "gitee_id")
	a.GiteeUserName = field.NewString(table, "gitee_user_name")
	a.GiteeUserAvatar = field.NewString(table, "gitee_user_avatar")
	a.GiteeNickName = field.NewString(table, "gitee_nick_name")

	a.fillFieldMap()

	return a
}

func (a *authGiteeUser) WithContext(ctx context.Context) IAuthGiteeUserDo {
	return a.authGiteeUserDo.WithContext(ctx)
}

func (a authGiteeUser) TableName() string { return a.authGiteeUserDo.TableName() }

func (a authGiteeUser) Alias() string { return a.authGiteeUserDo.Alias() }

func (a authGiteeUser) Columns(cols ...field.Expr) gen.Columns {
	return a.authGiteeUserDo.Columns(cols...)
}

func (a *authGiteeUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authGiteeUser) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["gitee_id"] = a.GiteeID
	a.fieldMap["gitee_user_name"] = a.GiteeUserName
	a.fieldMap["gitee_user_avatar"] = a.GiteeUserAvatar
	a.fieldMap["gitee_nick_name"] = a.GiteeNickName
}

func (a authGiteeUser) clone(db *gorm.DB) authGiteeUser {
	a.authGiteeUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a authGiteeUser) replaceDB(db *gorm.DB) authGiteeUser {
	a.authGiteeUserDo.ReplaceDB(db)
	return a
}

type authGiteeUserDo struct{ gen.DO }

type IAuthGiteeUserDo interface {
	gen.SubQuery
	Debug() IAuthGiteeUserDo
	WithContext(ctx context.Context) IAuthGiteeUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuthGiteeUserDo
	WriteDB() IAuthGiteeUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuthGiteeUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthGiteeUserDo
	Not(conds ...gen.Condition) IAuthGiteeUserDo
	Or(conds ...gen.Condition) IAuthGiteeUserDo
	Select(conds ...field.Expr) IAuthGiteeUserDo
	Where(conds ...gen.Condition) IAuthGiteeUserDo
	Order(conds ...field.Expr) IAuthGiteeUserDo
	Distinct(cols ...field.Expr) IAuthGiteeUserDo
	Omit(cols ...field.Expr) IAuthGiteeUserDo
	Join(table schema.Tabler, on ...field.Expr) IAuthGiteeUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthGiteeUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthGiteeUserDo
	Group(cols ...field.Expr) IAuthGiteeUserDo
	Having(conds ...gen.Condition) IAuthGiteeUserDo
	Limit(limit int) IAuthGiteeUserDo
	Offset(offset int) IAuthGiteeUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthGiteeUserDo
	Unscoped() IAuthGiteeUserDo
	Create(values ...*gen2.AuthGiteeUser) error
	CreateInBatches(values []*gen2.AuthGiteeUser, batchSize int) error
	Save(values ...*gen2.AuthGiteeUser) error
	First() (*gen2.AuthGiteeUser, error)
	Take() (*gen2.AuthGiteeUser, error)
	Last() (*gen2.AuthGiteeUser, error)
	Find() ([]*gen2.AuthGiteeUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen2.AuthGiteeUser, err error)
	FindInBatches(result *[]*gen2.AuthGiteeUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*gen2.AuthGiteeUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthGiteeUserDo
	Assign(attrs ...field.AssignExpr) IAuthGiteeUserDo
	Joins(fields ...field.RelationField) IAuthGiteeUserDo
	Preload(fields ...field.RelationField) IAuthGiteeUserDo
	FirstOrInit() (*gen2.AuthGiteeUser, error)
	FirstOrCreate() (*gen2.AuthGiteeUser, error)
	FindByPage(offset int, limit int) (result []*gen2.AuthGiteeUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthGiteeUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authGiteeUserDo) Debug() IAuthGiteeUserDo {
	return a.withDO(a.DO.Debug())
}

func (a authGiteeUserDo) WithContext(ctx context.Context) IAuthGiteeUserDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authGiteeUserDo) ReadDB() IAuthGiteeUserDo {
	return a.Clauses(dbresolver.Read)
}

func (a authGiteeUserDo) WriteDB() IAuthGiteeUserDo {
	return a.Clauses(dbresolver.Write)
}

func (a authGiteeUserDo) Session(config *gorm.Session) IAuthGiteeUserDo {
	return a.withDO(a.DO.Session(config))
}

func (a authGiteeUserDo) Clauses(conds ...clause.Expression) IAuthGiteeUserDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authGiteeUserDo) Returning(value interface{}, columns ...string) IAuthGiteeUserDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authGiteeUserDo) Not(conds ...gen.Condition) IAuthGiteeUserDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authGiteeUserDo) Or(conds ...gen.Condition) IAuthGiteeUserDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authGiteeUserDo) Select(conds ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authGiteeUserDo) Where(conds ...gen.Condition) IAuthGiteeUserDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authGiteeUserDo) Order(conds ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authGiteeUserDo) Distinct(cols ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authGiteeUserDo) Omit(cols ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authGiteeUserDo) Join(table schema.Tabler, on ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authGiteeUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authGiteeUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authGiteeUserDo) Group(cols ...field.Expr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authGiteeUserDo) Having(conds ...gen.Condition) IAuthGiteeUserDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authGiteeUserDo) Limit(limit int) IAuthGiteeUserDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authGiteeUserDo) Offset(offset int) IAuthGiteeUserDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authGiteeUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthGiteeUserDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authGiteeUserDo) Unscoped() IAuthGiteeUserDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authGiteeUserDo) Create(values ...*gen2.AuthGiteeUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authGiteeUserDo) CreateInBatches(values []*gen2.AuthGiteeUser, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authGiteeUserDo) Save(values ...*gen2.AuthGiteeUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authGiteeUserDo) First() (*gen2.AuthGiteeUser, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.AuthGiteeUser), nil
	}
}

func (a authGiteeUserDo) Take() (*gen2.AuthGiteeUser, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.AuthGiteeUser), nil
	}
}

func (a authGiteeUserDo) Last() (*gen2.AuthGiteeUser, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.AuthGiteeUser), nil
	}
}

func (a authGiteeUserDo) Find() ([]*gen2.AuthGiteeUser, error) {
	result, err := a.DO.Find()
	return result.([]*gen2.AuthGiteeUser), err
}

func (a authGiteeUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen2.AuthGiteeUser, err error) {
	buf := make([]*gen2.AuthGiteeUser, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authGiteeUserDo) FindInBatches(result *[]*gen2.AuthGiteeUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authGiteeUserDo) Attrs(attrs ...field.AssignExpr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authGiteeUserDo) Assign(attrs ...field.AssignExpr) IAuthGiteeUserDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authGiteeUserDo) Joins(fields ...field.RelationField) IAuthGiteeUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authGiteeUserDo) Preload(fields ...field.RelationField) IAuthGiteeUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authGiteeUserDo) FirstOrInit() (*gen2.AuthGiteeUser, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.AuthGiteeUser), nil
	}
}

func (a authGiteeUserDo) FirstOrCreate() (*gen2.AuthGiteeUser, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.AuthGiteeUser), nil
	}
}

func (a authGiteeUserDo) FindByPage(offset int, limit int) (result []*gen2.AuthGiteeUser, count int64, err error) {
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

func (a authGiteeUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authGiteeUserDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authGiteeUserDo) Delete(models ...*gen2.AuthGiteeUser) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authGiteeUserDo) withDO(do gen.Dao) *authGiteeUserDo {
	a.DO = *do.(*gen.DO)
	return a
}
