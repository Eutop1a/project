package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/data/knowledge/model"
	"helloworld/internal/data/knowledge/query"
	"helloworld/internal/domain"
)

type knowledgeRepository struct {
	data   *Data
	logger *log.Helper
}

func NewKnowledgeRepo(data *Data, logger log.Logger) biz.KnowledgeRepository {
	return &knowledgeRepository{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

// CreateKnowledge 创建知识点
func (k *knowledgeRepository) CreateKnowledge(ctx context.Context, req *domain.KnowledgeInfo) (err error) {
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	knowledge := query.Knowledge.WithContext(ctx)
	newKnowledge := &model.Knowledge{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentId,
	}
	err = knowledge.Create(newKnowledge)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// UpdateKnowledge 更新知识点
func (k *knowledgeRepository) UpdateKnowledge(ctx context.Context, req *domain.KnowledgeInfo) (err error) {
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	knowledge := query.Knowledge
	knowledgeDo := query.Knowledge.WithContext(ctx)
	updateData := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"parent_id":   req.ParentId,
	}
	_, err = knowledgeDo.Where(knowledge.ID.Eq(req.Id)).Updates(updateData)
	if err != nil {
		return err
	}
	return tx.Commit()
}

// DeleteKnowledge 删除知识点（软删除）
func (k *knowledgeRepository) DeleteKnowledge(ctx context.Context, id int64) (err error) {
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	knowledge := query.Knowledge
	knowledgeDo := query.Knowledge.WithContext(ctx)
	_, err = knowledgeDo.Where(knowledge.ID.Eq(id)).Delete()
	if err != nil {
		return err
	}
	return tx.Commit()
}

// GetKnowledge 获取单个知识点
func (k *knowledgeRepository) GetKnowledge(ctx context.Context, id int64) (info *domain.KnowledgeInfo, err error) {
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	knowledge := query.Knowledge
	knowledgeDo := query.Knowledge.WithContext(ctx)
	knowledgeModel, err := knowledgeDo.Where(knowledge.ID.Eq(id), knowledge.DeletedAt.IsNull()).First()
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	info = domain.NewKnowledgeInfo(knowledgeModel)
	// 处理父节点信息
	if knowledgeModel.ParentID != 0 {
		info.ParentId = knowledgeModel.ParentID
	}
	return info, nil
}

// ListKnowledge 分页列出知识点
func (k *knowledgeRepository) ListKnowledge(ctx context.Context, pageSize, pageNum int) (infos []*domain.KnowledgeInfo, err error) {
	if pageSize <= 0 {
		pageSize = 10 // 默认值或返回错误
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	tx := query.Q.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	knowledge := query.Knowledge
	knowledgeDo := query.Knowledge.WithContext(ctx)
	list, _, err := knowledgeDo.Select(
		knowledge.ID,
		knowledge.Name,
		knowledge.Description,
		knowledge.ParentID,
		knowledge.CreateTime,
		knowledge.UpdatedAt,
	).Where(
		knowledge.DeletedAt.IsNull(),
	).Order(
		knowledge.CreateTime.Desc(),
	).FindByPage((pageNum-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	infos = domain.NewKnowledgeInfos(list...)
	// 处理父节点信息
	for i, k := range list {
		if k.ParentID != 0 {
			infos[i].ParentId = k.ParentID
		}
	}
	return infos, nil
}
