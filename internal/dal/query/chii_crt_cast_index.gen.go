// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newCast(db *gorm.DB) cast {
	_cast := cast{}

	_cast.castDo.UseDB(db)
	_cast.castDo.UseModel(&dao.Cast{})

	tableName := _cast.castDo.TableName()
	_cast.ALL = field.NewAsterisk(tableName)
	_cast.CharacterID = field.NewField(tableName, "crt_id")
	_cast.PersonID = field.NewField(tableName, "prsn_id")
	_cast.SubjectID = field.NewField(tableName, "subject_id")
	_cast.SubjectTypeID = field.NewUint8(tableName, "subject_type_id")
	_cast.Summary = field.NewString(tableName, "summary")
	_cast.Character = castHasOneCharacter{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Character", "dao.Character"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Character.Fields", "dao.PersonField"),
		},
	}

	_cast.Subject = castHasOneSubject{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Subject", "dao.Subject"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Subject.Fields", "dao.SubjectField"),
		},
	}

	_cast.Person = castHasOnePerson{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Person", "dao.Person"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Person.Fields", "dao.PersonField"),
		},
	}

	_cast.fillFieldMap()

	return _cast
}

type cast struct {
	castDo castDo

	ALL           field.Asterisk
	CharacterID   field.Field
	PersonID      field.Field
	SubjectID     field.Field
	SubjectTypeID field.Uint8
	Summary       field.String
	Character     castHasOneCharacter

	Subject castHasOneSubject

	Person castHasOnePerson

	fieldMap map[string]field.Expr
}

func (c cast) Table(newTableName string) *cast {
	c.castDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cast) As(alias string) *cast {
	c.castDo.DO = *(c.castDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cast) updateTableName(table string) *cast {
	c.ALL = field.NewAsterisk(table)
	c.CharacterID = field.NewField(table, "crt_id")
	c.PersonID = field.NewField(table, "prsn_id")
	c.SubjectID = field.NewField(table, "subject_id")
	c.SubjectTypeID = field.NewUint8(table, "subject_type_id")
	c.Summary = field.NewString(table, "summary")

	c.fillFieldMap()

	return c
}

func (c *cast) WithContext(ctx context.Context) *castDo { return c.castDo.WithContext(ctx) }

func (c cast) TableName() string { return c.castDo.TableName() }

func (c cast) Alias() string { return c.castDo.Alias() }

func (c *cast) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cast) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 8)
	c.fieldMap["crt_id"] = c.CharacterID
	c.fieldMap["prsn_id"] = c.PersonID
	c.fieldMap["subject_id"] = c.SubjectID
	c.fieldMap["subject_type_id"] = c.SubjectTypeID
	c.fieldMap["summary"] = c.Summary

}

func (c cast) clone(db *gorm.DB) cast {
	c.castDo.ReplaceDB(db)
	return c
}

type castHasOneCharacter struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a castHasOneCharacter) Where(conds ...field.Expr) *castHasOneCharacter {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a castHasOneCharacter) WithContext(ctx context.Context) *castHasOneCharacter {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a castHasOneCharacter) Model(m *dao.Cast) *castHasOneCharacterTx {
	return &castHasOneCharacterTx{a.db.Model(m).Association(a.Name())}
}

type castHasOneCharacterTx struct{ tx *gorm.Association }

func (a castHasOneCharacterTx) Find() (result *dao.Character, err error) {
	return result, a.tx.Find(&result)
}

func (a castHasOneCharacterTx) Append(values ...*dao.Character) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a castHasOneCharacterTx) Replace(values ...*dao.Character) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a castHasOneCharacterTx) Delete(values ...*dao.Character) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a castHasOneCharacterTx) Clear() error {
	return a.tx.Clear()
}

func (a castHasOneCharacterTx) Count() int64 {
	return a.tx.Count()
}

type castHasOneSubject struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a castHasOneSubject) Where(conds ...field.Expr) *castHasOneSubject {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a castHasOneSubject) WithContext(ctx context.Context) *castHasOneSubject {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a castHasOneSubject) Model(m *dao.Cast) *castHasOneSubjectTx {
	return &castHasOneSubjectTx{a.db.Model(m).Association(a.Name())}
}

type castHasOneSubjectTx struct{ tx *gorm.Association }

func (a castHasOneSubjectTx) Find() (result *dao.Subject, err error) {
	return result, a.tx.Find(&result)
}

func (a castHasOneSubjectTx) Append(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a castHasOneSubjectTx) Replace(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a castHasOneSubjectTx) Delete(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a castHasOneSubjectTx) Clear() error {
	return a.tx.Clear()
}

func (a castHasOneSubjectTx) Count() int64 {
	return a.tx.Count()
}

type castHasOnePerson struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a castHasOnePerson) Where(conds ...field.Expr) *castHasOnePerson {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a castHasOnePerson) WithContext(ctx context.Context) *castHasOnePerson {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a castHasOnePerson) Model(m *dao.Cast) *castHasOnePersonTx {
	return &castHasOnePersonTx{a.db.Model(m).Association(a.Name())}
}

type castHasOnePersonTx struct{ tx *gorm.Association }

func (a castHasOnePersonTx) Find() (result *dao.Person, err error) {
	return result, a.tx.Find(&result)
}

func (a castHasOnePersonTx) Append(values ...*dao.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a castHasOnePersonTx) Replace(values ...*dao.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a castHasOnePersonTx) Delete(values ...*dao.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a castHasOnePersonTx) Clear() error {
	return a.tx.Clear()
}

func (a castHasOnePersonTx) Count() int64 {
	return a.tx.Count()
}

type castDo struct{ gen.DO }

func (c castDo) Debug() *castDo {
	return c.withDO(c.DO.Debug())
}

func (c castDo) WithContext(ctx context.Context) *castDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c castDo) ReadDB() *castDo {
	return c.Clauses(dbresolver.Read)
}

func (c castDo) WriteDB() *castDo {
	return c.Clauses(dbresolver.Write)
}

func (c castDo) Clauses(conds ...clause.Expression) *castDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c castDo) Returning(value interface{}, columns ...string) *castDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c castDo) Not(conds ...gen.Condition) *castDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c castDo) Or(conds ...gen.Condition) *castDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c castDo) Select(conds ...field.Expr) *castDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c castDo) Where(conds ...gen.Condition) *castDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c castDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *castDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c castDo) Order(conds ...field.Expr) *castDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c castDo) Distinct(cols ...field.Expr) *castDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c castDo) Omit(cols ...field.Expr) *castDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c castDo) Join(table schema.Tabler, on ...field.Expr) *castDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c castDo) LeftJoin(table schema.Tabler, on ...field.Expr) *castDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c castDo) RightJoin(table schema.Tabler, on ...field.Expr) *castDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c castDo) Group(cols ...field.Expr) *castDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c castDo) Having(conds ...gen.Condition) *castDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c castDo) Limit(limit int) *castDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c castDo) Offset(offset int) *castDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c castDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *castDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c castDo) Unscoped() *castDo {
	return c.withDO(c.DO.Unscoped())
}

func (c castDo) Create(values ...*dao.Cast) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c castDo) CreateInBatches(values []*dao.Cast, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c castDo) Save(values ...*dao.Cast) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c castDo) First() (*dao.Cast, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Cast), nil
	}
}

func (c castDo) Take() (*dao.Cast, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Cast), nil
	}
}

func (c castDo) Last() (*dao.Cast, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Cast), nil
	}
}

func (c castDo) Find() ([]*dao.Cast, error) {
	result, err := c.DO.Find()
	return result.([]*dao.Cast), err
}

func (c castDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Cast, err error) {
	buf := make([]*dao.Cast, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c castDo) FindInBatches(result *[]*dao.Cast, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c castDo) Attrs(attrs ...field.AssignExpr) *castDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c castDo) Assign(attrs ...field.AssignExpr) *castDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c castDo) Joins(fields ...field.RelationField) *castDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c castDo) Preload(fields ...field.RelationField) *castDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c castDo) FirstOrInit() (*dao.Cast, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Cast), nil
	}
}

func (c castDo) FirstOrCreate() (*dao.Cast, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Cast), nil
	}
}

func (c castDo) FindByPage(offset int, limit int) (result []*dao.Cast, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c castDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c castDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c castDo) Delete(models ...*dao.Cast) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *castDo) withDO(do gen.Dao) *castDo {
	c.DO = *do.(*gen.DO)
	return c
}
