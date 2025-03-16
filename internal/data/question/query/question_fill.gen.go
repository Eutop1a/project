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

	"helloworld/internal/data/question/model"
)

func newQuestionFill(db *gorm.DB, opts ...gen.DOOption) questionFill {
	_questionFill := questionFill{}

	_questionFill.questionFillDo.UseDB(db, opts...)
	_questionFill.questionFillDo.UseModel(&model.QuestionFill{})

	tableName := _questionFill.questionFillDo.TableName()
	_questionFill.ALL = field.NewAsterisk(tableName)
	_questionFill.ID = field.NewInt64(tableName, "id")
	_questionFill.Answer = field.NewString(tableName, "answer")

	_questionFill.fillFieldMap()

	return _questionFill
}

type questionFill struct {
	questionFillDo

	ALL    field.Asterisk
	ID     field.Int64
	Answer field.String // 支持正则表达式（如 "北京|上海"）

	fieldMap map[string]field.Expr
}

func (q questionFill) Table(newTableName string) *questionFill {
	q.questionFillDo.UseTable(newTableName)
	return q.updateTableName(newTableName)
}

func (q questionFill) As(alias string) *questionFill {
	q.questionFillDo.DO = *(q.questionFillDo.As(alias).(*gen.DO))
	return q.updateTableName(alias)
}

func (q *questionFill) updateTableName(table string) *questionFill {
	q.ALL = field.NewAsterisk(table)
	q.ID = field.NewInt64(table, "id")
	q.Answer = field.NewString(table, "answer")

	q.fillFieldMap()

	return q
}

func (q *questionFill) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := q.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (q *questionFill) fillFieldMap() {
	q.fieldMap = make(map[string]field.Expr, 2)
	q.fieldMap["id"] = q.ID
	q.fieldMap["answer"] = q.Answer
}

func (q questionFill) clone(db *gorm.DB) questionFill {
	q.questionFillDo.ReplaceConnPool(db.Statement.ConnPool)
	return q
}

func (q questionFill) replaceDB(db *gorm.DB) questionFill {
	q.questionFillDo.ReplaceDB(db)
	return q
}

type questionFillDo struct{ gen.DO }

type IQuestionFillDo interface {
	gen.SubQuery
	Debug() IQuestionFillDo
	WithContext(ctx context.Context) IQuestionFillDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IQuestionFillDo
	WriteDB() IQuestionFillDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IQuestionFillDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IQuestionFillDo
	Not(conds ...gen.Condition) IQuestionFillDo
	Or(conds ...gen.Condition) IQuestionFillDo
	Select(conds ...field.Expr) IQuestionFillDo
	Where(conds ...gen.Condition) IQuestionFillDo
	Order(conds ...field.Expr) IQuestionFillDo
	Distinct(cols ...field.Expr) IQuestionFillDo
	Omit(cols ...field.Expr) IQuestionFillDo
	Join(table schema.Tabler, on ...field.Expr) IQuestionFillDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IQuestionFillDo
	RightJoin(table schema.Tabler, on ...field.Expr) IQuestionFillDo
	Group(cols ...field.Expr) IQuestionFillDo
	Having(conds ...gen.Condition) IQuestionFillDo
	Limit(limit int) IQuestionFillDo
	Offset(offset int) IQuestionFillDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IQuestionFillDo
	Unscoped() IQuestionFillDo
	Create(values ...*model.QuestionFill) error
	CreateInBatches(values []*model.QuestionFill, batchSize int) error
	Save(values ...*model.QuestionFill) error
	First() (*model.QuestionFill, error)
	Take() (*model.QuestionFill, error)
	Last() (*model.QuestionFill, error)
	Find() ([]*model.QuestionFill, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QuestionFill, err error)
	FindInBatches(result *[]*model.QuestionFill, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.QuestionFill) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IQuestionFillDo
	Assign(attrs ...field.AssignExpr) IQuestionFillDo
	Joins(fields ...field.RelationField) IQuestionFillDo
	Preload(fields ...field.RelationField) IQuestionFillDo
	FirstOrInit() (*model.QuestionFill, error)
	FirstOrCreate() (*model.QuestionFill, error)
	FindByPage(offset int, limit int) (result []*model.QuestionFill, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IQuestionFillDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (q questionFillDo) Debug() IQuestionFillDo {
	return q.withDO(q.DO.Debug())
}

func (q questionFillDo) WithContext(ctx context.Context) IQuestionFillDo {
	return q.withDO(q.DO.WithContext(ctx))
}

func (q questionFillDo) ReadDB() IQuestionFillDo {
	return q.Clauses(dbresolver.Read)
}

func (q questionFillDo) WriteDB() IQuestionFillDo {
	return q.Clauses(dbresolver.Write)
}

func (q questionFillDo) Session(config *gorm.Session) IQuestionFillDo {
	return q.withDO(q.DO.Session(config))
}

func (q questionFillDo) Clauses(conds ...clause.Expression) IQuestionFillDo {
	return q.withDO(q.DO.Clauses(conds...))
}

func (q questionFillDo) Returning(value interface{}, columns ...string) IQuestionFillDo {
	return q.withDO(q.DO.Returning(value, columns...))
}

func (q questionFillDo) Not(conds ...gen.Condition) IQuestionFillDo {
	return q.withDO(q.DO.Not(conds...))
}

func (q questionFillDo) Or(conds ...gen.Condition) IQuestionFillDo {
	return q.withDO(q.DO.Or(conds...))
}

func (q questionFillDo) Select(conds ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.Select(conds...))
}

func (q questionFillDo) Where(conds ...gen.Condition) IQuestionFillDo {
	return q.withDO(q.DO.Where(conds...))
}

func (q questionFillDo) Order(conds ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.Order(conds...))
}

func (q questionFillDo) Distinct(cols ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.Distinct(cols...))
}

func (q questionFillDo) Omit(cols ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.Omit(cols...))
}

func (q questionFillDo) Join(table schema.Tabler, on ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.Join(table, on...))
}

func (q questionFillDo) LeftJoin(table schema.Tabler, on ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.LeftJoin(table, on...))
}

func (q questionFillDo) RightJoin(table schema.Tabler, on ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.RightJoin(table, on...))
}

func (q questionFillDo) Group(cols ...field.Expr) IQuestionFillDo {
	return q.withDO(q.DO.Group(cols...))
}

func (q questionFillDo) Having(conds ...gen.Condition) IQuestionFillDo {
	return q.withDO(q.DO.Having(conds...))
}

func (q questionFillDo) Limit(limit int) IQuestionFillDo {
	return q.withDO(q.DO.Limit(limit))
}

func (q questionFillDo) Offset(offset int) IQuestionFillDo {
	return q.withDO(q.DO.Offset(offset))
}

func (q questionFillDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IQuestionFillDo {
	return q.withDO(q.DO.Scopes(funcs...))
}

func (q questionFillDo) Unscoped() IQuestionFillDo {
	return q.withDO(q.DO.Unscoped())
}

func (q questionFillDo) Create(values ...*model.QuestionFill) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Create(values)
}

func (q questionFillDo) CreateInBatches(values []*model.QuestionFill, batchSize int) error {
	return q.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (q questionFillDo) Save(values ...*model.QuestionFill) error {
	if len(values) == 0 {
		return nil
	}
	return q.DO.Save(values)
}

func (q questionFillDo) First() (*model.QuestionFill, error) {
	if result, err := q.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionFill), nil
	}
}

func (q questionFillDo) Take() (*model.QuestionFill, error) {
	if result, err := q.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionFill), nil
	}
}

func (q questionFillDo) Last() (*model.QuestionFill, error) {
	if result, err := q.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionFill), nil
	}
}

func (q questionFillDo) Find() ([]*model.QuestionFill, error) {
	result, err := q.DO.Find()
	return result.([]*model.QuestionFill), err
}

func (q questionFillDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.QuestionFill, err error) {
	buf := make([]*model.QuestionFill, 0, batchSize)
	err = q.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (q questionFillDo) FindInBatches(result *[]*model.QuestionFill, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return q.DO.FindInBatches(result, batchSize, fc)
}

func (q questionFillDo) Attrs(attrs ...field.AssignExpr) IQuestionFillDo {
	return q.withDO(q.DO.Attrs(attrs...))
}

func (q questionFillDo) Assign(attrs ...field.AssignExpr) IQuestionFillDo {
	return q.withDO(q.DO.Assign(attrs...))
}

func (q questionFillDo) Joins(fields ...field.RelationField) IQuestionFillDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Joins(_f))
	}
	return &q
}

func (q questionFillDo) Preload(fields ...field.RelationField) IQuestionFillDo {
	for _, _f := range fields {
		q = *q.withDO(q.DO.Preload(_f))
	}
	return &q
}

func (q questionFillDo) FirstOrInit() (*model.QuestionFill, error) {
	if result, err := q.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionFill), nil
	}
}

func (q questionFillDo) FirstOrCreate() (*model.QuestionFill, error) {
	if result, err := q.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.QuestionFill), nil
	}
}

func (q questionFillDo) FindByPage(offset int, limit int) (result []*model.QuestionFill, count int64, err error) {
	result, err = q.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = q.Offset(-1).Limit(-1).Count()
	return
}

func (q questionFillDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = q.Count()
	if err != nil {
		return
	}

	err = q.Offset(offset).Limit(limit).Scan(result)
	return
}

func (q questionFillDo) Scan(result interface{}) (err error) {
	return q.DO.Scan(result)
}

func (q questionFillDo) Delete(models ...*model.QuestionFill) (result gen.ResultInfo, err error) {
	return q.DO.Delete(models)
}

func (q *questionFillDo) withDO(do gen.Dao) *questionFillDo {
	q.DO = *do.(*gen.DO)
	return q
}
