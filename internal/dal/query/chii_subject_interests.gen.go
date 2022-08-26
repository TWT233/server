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

func newSubjectCollection(db *gorm.DB) subjectCollection {
	_subjectCollection := subjectCollection{}

	_subjectCollection.subjectCollectionDo.UseDB(db)
	_subjectCollection.subjectCollectionDo.UseModel(&dao.SubjectCollection{})

	tableName := _subjectCollection.subjectCollectionDo.TableName()
	_subjectCollection.ALL = field.NewAsterisk(tableName)
	_subjectCollection.ID = field.NewUint32(tableName, "interest_id")
	_subjectCollection.UserID = field.NewField(tableName, "interest_uid")
	_subjectCollection.SubjectID = field.NewField(tableName, "interest_subject_id")
	_subjectCollection.SubjectType = field.NewUint8(tableName, "interest_subject_type")
	_subjectCollection.Rate = field.NewUint8(tableName, "interest_rate")
	_subjectCollection.Type = field.NewUint8(tableName, "interest_type")
	_subjectCollection.HasComment = field.NewBool(tableName, "interest_has_comment")
	_subjectCollection.Comment = field.NewString(tableName, "interest_comment")
	_subjectCollection.Tag = field.NewString(tableName, "interest_tag")
	_subjectCollection.EpStatus = field.NewUint32(tableName, "interest_ep_status")
	_subjectCollection.VolStatus = field.NewUint32(tableName, "interest_vol_status")
	_subjectCollection.WishTime = field.NewUint32(tableName, "interest_wish_dateline")
	_subjectCollection.DoingTime = field.NewUint32(tableName, "interest_doing_dateline")
	_subjectCollection.DoneTime = field.NewUint32(tableName, "interest_collect_dateline")
	_subjectCollection.OnHoldTime = field.NewUint32(tableName, "interest_on_hold_dateline")
	_subjectCollection.DroppedTime = field.NewUint32(tableName, "interest_dropped_dateline")
	_subjectCollection.CreateIP = field.NewString(tableName, "interest_create_ip")
	_subjectCollection.LastUpdateIP = field.NewString(tableName, "interest_lasttouch_ip")
	_subjectCollection.UpdatedTime = field.NewUint32(tableName, "interest_lasttouch")
	_subjectCollection.Private = field.NewUint8(tableName, "interest_private")

	_subjectCollection.fillFieldMap()

	return _subjectCollection
}

type subjectCollection struct {
	subjectCollectionDo subjectCollectionDo

	ALL          field.Asterisk
	ID           field.Uint32
	UserID       field.Field
	SubjectID    field.Field
	SubjectType  field.Uint8
	Rate         field.Uint8
	Type         field.Uint8
	HasComment   field.Bool
	Comment      field.String
	Tag          field.String
	EpStatus     field.Uint32
	VolStatus    field.Uint32
	WishTime     field.Uint32
	DoingTime    field.Uint32
	DoneTime     field.Uint32
	OnHoldTime   field.Uint32
	DroppedTime  field.Uint32
	CreateIP     field.String
	LastUpdateIP field.String
	UpdatedTime  field.Uint32
	Private      field.Uint8

	fieldMap map[string]field.Expr
}

func (s subjectCollection) Table(newTableName string) *subjectCollection {
	s.subjectCollectionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectCollection) As(alias string) *subjectCollection {
	s.subjectCollectionDo.DO = *(s.subjectCollectionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectCollection) updateTableName(table string) *subjectCollection {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint32(table, "interest_id")
	s.UserID = field.NewField(table, "interest_uid")
	s.SubjectID = field.NewField(table, "interest_subject_id")
	s.SubjectType = field.NewUint8(table, "interest_subject_type")
	s.Rate = field.NewUint8(table, "interest_rate")
	s.Type = field.NewUint8(table, "interest_type")
	s.HasComment = field.NewBool(table, "interest_has_comment")
	s.Comment = field.NewString(table, "interest_comment")
	s.Tag = field.NewString(table, "interest_tag")
	s.EpStatus = field.NewUint32(table, "interest_ep_status")
	s.VolStatus = field.NewUint32(table, "interest_vol_status")
	s.WishTime = field.NewUint32(table, "interest_wish_dateline")
	s.DoingTime = field.NewUint32(table, "interest_doing_dateline")
	s.DoneTime = field.NewUint32(table, "interest_collect_dateline")
	s.OnHoldTime = field.NewUint32(table, "interest_on_hold_dateline")
	s.DroppedTime = field.NewUint32(table, "interest_dropped_dateline")
	s.CreateIP = field.NewString(table, "interest_create_ip")
	s.LastUpdateIP = field.NewString(table, "interest_lasttouch_ip")
	s.UpdatedTime = field.NewUint32(table, "interest_lasttouch")
	s.Private = field.NewUint8(table, "interest_private")

	s.fillFieldMap()

	return s
}

func (s *subjectCollection) WithContext(ctx context.Context) *subjectCollectionDo {
	return s.subjectCollectionDo.WithContext(ctx)
}

func (s subjectCollection) TableName() string { return s.subjectCollectionDo.TableName() }

func (s subjectCollection) Alias() string { return s.subjectCollectionDo.Alias() }

func (s *subjectCollection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectCollection) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 20)
	s.fieldMap["interest_id"] = s.ID
	s.fieldMap["interest_uid"] = s.UserID
	s.fieldMap["interest_subject_id"] = s.SubjectID
	s.fieldMap["interest_subject_type"] = s.SubjectType
	s.fieldMap["interest_rate"] = s.Rate
	s.fieldMap["interest_type"] = s.Type
	s.fieldMap["interest_has_comment"] = s.HasComment
	s.fieldMap["interest_comment"] = s.Comment
	s.fieldMap["interest_tag"] = s.Tag
	s.fieldMap["interest_ep_status"] = s.EpStatus
	s.fieldMap["interest_vol_status"] = s.VolStatus
	s.fieldMap["interest_wish_dateline"] = s.WishTime
	s.fieldMap["interest_doing_dateline"] = s.DoingTime
	s.fieldMap["interest_collect_dateline"] = s.DoneTime
	s.fieldMap["interest_on_hold_dateline"] = s.OnHoldTime
	s.fieldMap["interest_dropped_dateline"] = s.DroppedTime
	s.fieldMap["interest_create_ip"] = s.CreateIP
	s.fieldMap["interest_lasttouch_ip"] = s.LastUpdateIP
	s.fieldMap["interest_lasttouch"] = s.UpdatedTime
	s.fieldMap["interest_private"] = s.Private
}

func (s subjectCollection) clone(db *gorm.DB) subjectCollection {
	s.subjectCollectionDo.ReplaceDB(db)
	return s
}

type subjectCollectionDo struct{ gen.DO }

func (s subjectCollectionDo) Debug() *subjectCollectionDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectCollectionDo) WithContext(ctx context.Context) *subjectCollectionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectCollectionDo) ReadDB() *subjectCollectionDo {
	return s.Clauses(dbresolver.Read)
}

func (s subjectCollectionDo) WriteDB() *subjectCollectionDo {
	return s.Clauses(dbresolver.Write)
}

func (s subjectCollectionDo) Clauses(conds ...clause.Expression) *subjectCollectionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectCollectionDo) Returning(value interface{}, columns ...string) *subjectCollectionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectCollectionDo) Not(conds ...gen.Condition) *subjectCollectionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectCollectionDo) Or(conds ...gen.Condition) *subjectCollectionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectCollectionDo) Select(conds ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectCollectionDo) Where(conds ...gen.Condition) *subjectCollectionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectCollectionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *subjectCollectionDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectCollectionDo) Order(conds ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectCollectionDo) Distinct(cols ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectCollectionDo) Omit(cols ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectCollectionDo) Join(table schema.Tabler, on ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectCollectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectCollectionDo) RightJoin(table schema.Tabler, on ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectCollectionDo) Group(cols ...field.Expr) *subjectCollectionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectCollectionDo) Having(conds ...gen.Condition) *subjectCollectionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectCollectionDo) Limit(limit int) *subjectCollectionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectCollectionDo) Offset(offset int) *subjectCollectionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectCollectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subjectCollectionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectCollectionDo) Unscoped() *subjectCollectionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectCollectionDo) Create(values ...*dao.SubjectCollection) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectCollectionDo) CreateInBatches(values []*dao.SubjectCollection, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectCollectionDo) Save(values ...*dao.SubjectCollection) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectCollectionDo) First() (*dao.SubjectCollection, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectCollection), nil
	}
}

func (s subjectCollectionDo) Take() (*dao.SubjectCollection, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectCollection), nil
	}
}

func (s subjectCollectionDo) Last() (*dao.SubjectCollection, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectCollection), nil
	}
}

func (s subjectCollectionDo) Find() ([]*dao.SubjectCollection, error) {
	result, err := s.DO.Find()
	return result.([]*dao.SubjectCollection), err
}

func (s subjectCollectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.SubjectCollection, err error) {
	buf := make([]*dao.SubjectCollection, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectCollectionDo) FindInBatches(result *[]*dao.SubjectCollection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectCollectionDo) Attrs(attrs ...field.AssignExpr) *subjectCollectionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectCollectionDo) Assign(attrs ...field.AssignExpr) *subjectCollectionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectCollectionDo) Joins(fields ...field.RelationField) *subjectCollectionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectCollectionDo) Preload(fields ...field.RelationField) *subjectCollectionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectCollectionDo) FirstOrInit() (*dao.SubjectCollection, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectCollection), nil
	}
}

func (s subjectCollectionDo) FirstOrCreate() (*dao.SubjectCollection, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectCollection), nil
	}
}

func (s subjectCollectionDo) FindByPage(offset int, limit int) (result []*dao.SubjectCollection, count int64, err error) {
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

func (s subjectCollectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s subjectCollectionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s subjectCollectionDo) Delete(models ...*dao.SubjectCollection) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *subjectCollectionDo) withDO(do gen.Dao) *subjectCollectionDo {
	s.DO = *do.(*gen.DO)
	return s
}
