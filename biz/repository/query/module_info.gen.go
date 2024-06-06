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

func newModuleInfo(db *gorm.DB, opts ...gen.DOOption) moduleInfo {
	_moduleInfo := moduleInfo{}

	_moduleInfo.moduleInfoDo.UseDB(db, opts...)
	_moduleInfo.moduleInfoDo.UseModel(&gen2.ModuleInfo{})

	tableName := _moduleInfo.moduleInfoDo.TableName()
	_moduleInfo.ALL = field.NewAsterisk(tableName)
	_moduleInfo.ID = field.NewString(tableName, "id")
	_moduleInfo.ModuleName = field.NewString(tableName, "module_name")
	_moduleInfo.BranchID = field.NewString(tableName, "branch_id")
	_moduleInfo.InterfaceNum = field.NewInt32(tableName, "interface_num")
	_moduleInfo.IsDeleted = field.NewInt32(tableName, "is_deleted")
	_moduleInfo.CreateTime = field.NewTime(tableName, "create_time")
	_moduleInfo.UpdateTime = field.NewTime(tableName, "update_time")

	_moduleInfo.fillFieldMap()

	return _moduleInfo
}

type moduleInfo struct {
	moduleInfoDo moduleInfoDo

	ALL          field.Asterisk
	ID           field.String
	ModuleName   field.String // 模块名称
	BranchID     field.String // 分支外键
	InterfaceNum field.Int32  // 接口数量
	IsDeleted    field.Int32  // 是否删除  0 false 1 true
	CreateTime   field.Time   // 创建时间
	UpdateTime   field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (m moduleInfo) Table(newTableName string) *moduleInfo {
	m.moduleInfoDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m moduleInfo) As(alias string) *moduleInfo {
	m.moduleInfoDo.DO = *(m.moduleInfoDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *moduleInfo) updateTableName(table string) *moduleInfo {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewString(table, "id")
	m.ModuleName = field.NewString(table, "module_name")
	m.BranchID = field.NewString(table, "branch_id")
	m.InterfaceNum = field.NewInt32(table, "interface_num")
	m.IsDeleted = field.NewInt32(table, "is_deleted")
	m.CreateTime = field.NewTime(table, "create_time")
	m.UpdateTime = field.NewTime(table, "update_time")

	m.fillFieldMap()

	return m
}

func (m *moduleInfo) WithContext(ctx context.Context) IModuleInfoDo {
	return m.moduleInfoDo.WithContext(ctx)
}

func (m moduleInfo) TableName() string { return m.moduleInfoDo.TableName() }

func (m moduleInfo) Alias() string { return m.moduleInfoDo.Alias() }

func (m moduleInfo) Columns(cols ...field.Expr) gen.Columns { return m.moduleInfoDo.Columns(cols...) }

func (m *moduleInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *moduleInfo) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 7)
	m.fieldMap["id"] = m.ID
	m.fieldMap["module_name"] = m.ModuleName
	m.fieldMap["branch_id"] = m.BranchID
	m.fieldMap["interface_num"] = m.InterfaceNum
	m.fieldMap["is_deleted"] = m.IsDeleted
	m.fieldMap["create_time"] = m.CreateTime
	m.fieldMap["update_time"] = m.UpdateTime
}

func (m moduleInfo) clone(db *gorm.DB) moduleInfo {
	m.moduleInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m moduleInfo) replaceDB(db *gorm.DB) moduleInfo {
	m.moduleInfoDo.ReplaceDB(db)
	return m
}

type moduleInfoDo struct{ gen.DO }

type IModuleInfoDo interface {
	gen.SubQuery
	Debug() IModuleInfoDo
	WithContext(ctx context.Context) IModuleInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IModuleInfoDo
	WriteDB() IModuleInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IModuleInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IModuleInfoDo
	Not(conds ...gen.Condition) IModuleInfoDo
	Or(conds ...gen.Condition) IModuleInfoDo
	Select(conds ...field.Expr) IModuleInfoDo
	Where(conds ...gen.Condition) IModuleInfoDo
	Order(conds ...field.Expr) IModuleInfoDo
	Distinct(cols ...field.Expr) IModuleInfoDo
	Omit(cols ...field.Expr) IModuleInfoDo
	Join(table schema.Tabler, on ...field.Expr) IModuleInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IModuleInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IModuleInfoDo
	Group(cols ...field.Expr) IModuleInfoDo
	Having(conds ...gen.Condition) IModuleInfoDo
	Limit(limit int) IModuleInfoDo
	Offset(offset int) IModuleInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IModuleInfoDo
	Unscoped() IModuleInfoDo
	Create(values ...*gen2.ModuleInfo) error
	CreateInBatches(values []*gen2.ModuleInfo, batchSize int) error
	Save(values ...*gen2.ModuleInfo) error
	First() (*gen2.ModuleInfo, error)
	Take() (*gen2.ModuleInfo, error)
	Last() (*gen2.ModuleInfo, error)
	Find() ([]*gen2.ModuleInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen2.ModuleInfo, err error)
	FindInBatches(result *[]*gen2.ModuleInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*gen2.ModuleInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IModuleInfoDo
	Assign(attrs ...field.AssignExpr) IModuleInfoDo
	Joins(fields ...field.RelationField) IModuleInfoDo
	Preload(fields ...field.RelationField) IModuleInfoDo
	FirstOrInit() (*gen2.ModuleInfo, error)
	FirstOrCreate() (*gen2.ModuleInfo, error)
	FindByPage(offset int, limit int) (result []*gen2.ModuleInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IModuleInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m moduleInfoDo) Debug() IModuleInfoDo {
	return m.withDO(m.DO.Debug())
}

func (m moduleInfoDo) WithContext(ctx context.Context) IModuleInfoDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m moduleInfoDo) ReadDB() IModuleInfoDo {
	return m.Clauses(dbresolver.Read)
}

func (m moduleInfoDo) WriteDB() IModuleInfoDo {
	return m.Clauses(dbresolver.Write)
}

func (m moduleInfoDo) Session(config *gorm.Session) IModuleInfoDo {
	return m.withDO(m.DO.Session(config))
}

func (m moduleInfoDo) Clauses(conds ...clause.Expression) IModuleInfoDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m moduleInfoDo) Returning(value interface{}, columns ...string) IModuleInfoDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m moduleInfoDo) Not(conds ...gen.Condition) IModuleInfoDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m moduleInfoDo) Or(conds ...gen.Condition) IModuleInfoDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m moduleInfoDo) Select(conds ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m moduleInfoDo) Where(conds ...gen.Condition) IModuleInfoDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m moduleInfoDo) Order(conds ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m moduleInfoDo) Distinct(cols ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m moduleInfoDo) Omit(cols ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m moduleInfoDo) Join(table schema.Tabler, on ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m moduleInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m moduleInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m moduleInfoDo) Group(cols ...field.Expr) IModuleInfoDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m moduleInfoDo) Having(conds ...gen.Condition) IModuleInfoDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m moduleInfoDo) Limit(limit int) IModuleInfoDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m moduleInfoDo) Offset(offset int) IModuleInfoDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m moduleInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IModuleInfoDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m moduleInfoDo) Unscoped() IModuleInfoDo {
	return m.withDO(m.DO.Unscoped())
}

func (m moduleInfoDo) Create(values ...*gen2.ModuleInfo) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m moduleInfoDo) CreateInBatches(values []*gen2.ModuleInfo, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m moduleInfoDo) Save(values ...*gen2.ModuleInfo) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m moduleInfoDo) First() (*gen2.ModuleInfo, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.ModuleInfo), nil
	}
}

func (m moduleInfoDo) Take() (*gen2.ModuleInfo, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.ModuleInfo), nil
	}
}

func (m moduleInfoDo) Last() (*gen2.ModuleInfo, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.ModuleInfo), nil
	}
}

func (m moduleInfoDo) Find() ([]*gen2.ModuleInfo, error) {
	result, err := m.DO.Find()
	return result.([]*gen2.ModuleInfo), err
}

func (m moduleInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*gen2.ModuleInfo, err error) {
	buf := make([]*gen2.ModuleInfo, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m moduleInfoDo) FindInBatches(result *[]*gen2.ModuleInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m moduleInfoDo) Attrs(attrs ...field.AssignExpr) IModuleInfoDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m moduleInfoDo) Assign(attrs ...field.AssignExpr) IModuleInfoDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m moduleInfoDo) Joins(fields ...field.RelationField) IModuleInfoDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m moduleInfoDo) Preload(fields ...field.RelationField) IModuleInfoDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m moduleInfoDo) FirstOrInit() (*gen2.ModuleInfo, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.ModuleInfo), nil
	}
}

func (m moduleInfoDo) FirstOrCreate() (*gen2.ModuleInfo, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*gen2.ModuleInfo), nil
	}
}

func (m moduleInfoDo) FindByPage(offset int, limit int) (result []*gen2.ModuleInfo, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m moduleInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m moduleInfoDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m moduleInfoDo) Delete(models ...*gen2.ModuleInfo) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *moduleInfoDo) withDO(do gen.Dao) *moduleInfoDo {
	m.DO = *do.(*gen.DO)
	return m
}