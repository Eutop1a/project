// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameQuestionFill = "question_fill"

// QuestionFill mapped from table <question_fill>
type QuestionFill struct {
	ID     int64  `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	Answer string `gorm:"column:answer;type:text;not null;comment:支持正则表达式（如 "北京|上海"）" json:"answer"` // 支持正则表达式（如 "北京|上海"）
}

// TableName QuestionFill's table name
func (*QuestionFill) TableName() string {
	return TableNameQuestionFill
}
