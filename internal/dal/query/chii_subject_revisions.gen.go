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

func newSubjectRevision(db *gorm.DB) subjectRevision {
	_subjectRevision := subjectRevision{}

	_subjectRevision.subjectRevisionDo.UseDB(db)
	_subjectRevision.subjectRevisionDo.UseModel(&dao.SubjectRevision{})

	tableName := _subjectRevision.subjectRevisionDo.TableName()
	_subjectRevision.ALL = field.NewAsterisk(tableName)
	_subjectRevision.ID = field.NewUint32(tableName, "rev_id")
	_subjectRevision.Type = field.NewUint8(tableName, "rev_type")
	_subjectRevision.SubjectID = field.NewField(tableName, "rev_subject_id")
	_subjectRevision.TypeID = field.NewUint16(tableName, "rev_type_id")
	_subjectRevision.CreatorID = field.NewField(tableName, "rev_creator")
	_subjectRevision.Dateline = field.NewUint32(tableName, "rev_dateline")
	_subjectRevision.Name = field.NewString(tableName, "rev_name")
	_subjectRevision.NameCN = field.NewString(tableName, "rev_name_cn")
	_subjectRevision.FieldInfobox = field.NewString(tableName, "rev_field_infobox")
	_subjectRevision.FieldSummary = field.NewString(tableName, "rev_field_summary")
	_subjectRevision.VoteField = field.NewString(tableName, "rev_vote_field")
	_subjectRevision.FieldEps = field.NewUint32(tableName, "rev_field_eps")
	_subjectRevision.EditSummary = field.NewString(tableName, "rev_edit_summary")
	_subjectRevision.Platform = field.NewUint16(tableName, "rev_platform")
	_subjectRevision.Subject = subjectRevisionBelongsToSubject{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Subject", "dao.Subject"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Subject.Fields", "dao.SubjectField"),
		},
	}

	_subjectRevision.fillFieldMap()

	return _subjectRevision
}

type subjectRevision struct {
	subjectRevisionDo subjectRevisionDo

	ALL          field.Asterisk
	ID           field.Uint32
	Type         field.Uint8
	SubjectID    field.Field
	TypeID       field.Uint16
	CreatorID    field.Field
	Dateline     field.Uint32
	Name         field.String
	NameCN       field.String
	FieldInfobox field.String
	FieldSummary field.String
	VoteField    field.String
	FieldEps     field.Uint32
	EditSummary  field.String
	Platform     field.Uint16
	Subject      subjectRevisionBelongsToSubject

	fieldMap map[string]field.Expr
}

func (s subjectRevision) Table(newTableName string) *subjectRevision {
	s.subjectRevisionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectRevision) As(alias string) *subjectRevision {
	s.subjectRevisionDo.DO = *(s.subjectRevisionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectRevision) updateTableName(table string) *subjectRevision {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint32(table, "rev_id")
	s.Type = field.NewUint8(table, "rev_type")
	s.SubjectID = field.NewField(table, "rev_subject_id")
	s.TypeID = field.NewUint16(table, "rev_type_id")
	s.CreatorID = field.NewField(table, "rev_creator")
	s.Dateline = field.NewUint32(table, "rev_dateline")
	s.Name = field.NewString(table, "rev_name")
	s.NameCN = field.NewString(table, "rev_name_cn")
	s.FieldInfobox = field.NewString(table, "rev_field_infobox")
	s.FieldSummary = field.NewString(table, "rev_field_summary")
	s.VoteField = field.NewString(table, "rev_vote_field")
	s.FieldEps = field.NewUint32(table, "rev_field_eps")
	s.EditSummary = field.NewString(table, "rev_edit_summary")
	s.Platform = field.NewUint16(table, "rev_platform")

	s.fillFieldMap()

	return s
}

func (s *subjectRevision) WithContext(ctx context.Context) *subjectRevisionDo {
	return s.subjectRevisionDo.WithContext(ctx)
}

func (s subjectRevision) TableName() string { return s.subjectRevisionDo.TableName() }

func (s subjectRevision) Alias() string { return s.subjectRevisionDo.Alias() }

func (s *subjectRevision) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectRevision) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["rev_id"] = s.ID
	s.fieldMap["rev_type"] = s.Type
	s.fieldMap["rev_subject_id"] = s.SubjectID
	s.fieldMap["rev_type_id"] = s.TypeID
	s.fieldMap["rev_creator"] = s.CreatorID
	s.fieldMap["rev_dateline"] = s.Dateline
	s.fieldMap["rev_name"] = s.Name
	s.fieldMap["rev_name_cn"] = s.NameCN
	s.fieldMap["rev_field_infobox"] = s.FieldInfobox
	s.fieldMap["rev_field_summary"] = s.FieldSummary
	s.fieldMap["rev_vote_field"] = s.VoteField
	s.fieldMap["rev_field_eps"] = s.FieldEps
	s.fieldMap["rev_edit_summary"] = s.EditSummary
	s.fieldMap["rev_platform"] = s.Platform

}

func (s subjectRevision) clone(db *gorm.DB) subjectRevision {
	s.subjectRevisionDo.ReplaceDB(db)
	return s
}

type subjectRevisionBelongsToSubject struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a subjectRevisionBelongsToSubject) Where(conds ...field.Expr) *subjectRevisionBelongsToSubject {
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

func (a subjectRevisionBelongsToSubject) WithContext(ctx context.Context) *subjectRevisionBelongsToSubject {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a subjectRevisionBelongsToSubject) Model(m *dao.SubjectRevision) *subjectRevisionBelongsToSubjectTx {
	return &subjectRevisionBelongsToSubjectTx{a.db.Model(m).Association(a.Name())}
}

type subjectRevisionBelongsToSubjectTx struct{ tx *gorm.Association }

func (a subjectRevisionBelongsToSubjectTx) Find() (result *dao.Subject, err error) {
	return result, a.tx.Find(&result)
}

func (a subjectRevisionBelongsToSubjectTx) Append(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a subjectRevisionBelongsToSubjectTx) Replace(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a subjectRevisionBelongsToSubjectTx) Delete(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a subjectRevisionBelongsToSubjectTx) Clear() error {
	return a.tx.Clear()
}

func (a subjectRevisionBelongsToSubjectTx) Count() int64 {
	return a.tx.Count()
}

type subjectRevisionDo struct{ gen.DO }

func (s subjectRevisionDo) Debug() *subjectRevisionDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectRevisionDo) WithContext(ctx context.Context) *subjectRevisionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectRevisionDo) ReadDB() *subjectRevisionDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectRevisionDo) WriteDB() *subjectRevisionDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectRevisionDo) Clauses(conds ...clause.Expression) *subjectRevisionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectRevisionDo) Returning(value interface{}, columns ...string) *subjectRevisionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectRevisionDo) Not(conds ...gen.Condition) *subjectRevisionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectRevisionDo) Or(conds ...gen.Condition) *subjectRevisionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectRevisionDo) Select(conds ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectRevisionDo) Where(conds ...gen.Condition) *subjectRevisionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectRevisionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *subjectRevisionDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectRevisionDo) Order(conds ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectRevisionDo) Distinct(cols ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectRevisionDo) Omit(cols ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectRevisionDo) Join(table schema.Tabler, on ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectRevisionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectRevisionDo) RightJoin(table schema.Tabler, on ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectRevisionDo) Group(cols ...field.Expr) *subjectRevisionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectRevisionDo) Having(conds ...gen.Condition) *subjectRevisionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectRevisionDo) Limit(limit int) *subjectRevisionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectRevisionDo) Offset(offset int) *subjectRevisionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectRevisionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subjectRevisionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectRevisionDo) Unscoped() *subjectRevisionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectRevisionDo) Create(values ...*dao.SubjectRevision) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectRevisionDo) CreateInBatches(values []*dao.SubjectRevision, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectRevisionDo) Save(values ...*dao.SubjectRevision) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectRevisionDo) First() (*dao.SubjectRevision, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRevision), nil
	}
}

func (s subjectRevisionDo) Take() (*dao.SubjectRevision, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRevision), nil
	}
}

func (s subjectRevisionDo) Last() (*dao.SubjectRevision, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRevision), nil
	}
}

func (s subjectRevisionDo) Find() ([]*dao.SubjectRevision, error) {
	result, err := s.DO.Find()
	return result.([]*dao.SubjectRevision), err
}

func (s subjectRevisionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.SubjectRevision, err error) {
	buf := make([]*dao.SubjectRevision, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectRevisionDo) FindInBatches(result *[]*dao.SubjectRevision, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectRevisionDo) Attrs(attrs ...field.AssignExpr) *subjectRevisionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectRevisionDo) Assign(attrs ...field.AssignExpr) *subjectRevisionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectRevisionDo) Joins(fields ...field.RelationField) *subjectRevisionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectRevisionDo) Preload(fields ...field.RelationField) *subjectRevisionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectRevisionDo) FirstOrInit() (*dao.SubjectRevision, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRevision), nil
	}
}

func (s subjectRevisionDo) FirstOrCreate() (*dao.SubjectRevision, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRevision), nil
	}
}

func (s subjectRevisionDo) FindByPage(offset int, limit int) (result []*dao.SubjectRevision, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s subjectRevisionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectRevisionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectRevisionDo) Delete(models ...*dao.SubjectRevision) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectRevisionDo) withDO(do gen.Dao) *subjectRevisionDo {
	s.DO = *do.(*gen.DO)
	return s
}
