package domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"helloworld/internal/data/knowledge/model"
)

type KnowledgeInfo struct {
	Id          int64
	ParentId    int64
	Name        string
	Description string
	CreateTime  *timestamppb.Timestamp
	UpdateTime  *timestamppb.Timestamp
}

func (k *KnowledgeInfo) padding(knowledge *model.Knowledge) {
	k.Id = knowledge.ID
	k.ParentId = knowledge.ParentID
	k.Name = knowledge.Name
	k.Description = knowledge.Description
	k.CreateTime = timestamppb.New(knowledge.CreateTime)
	k.UpdateTime = timestamppb.New(knowledge.UpdatedAt)
}

func NewKnowledgeInfo(knowledge *model.Knowledge) *KnowledgeInfo {
	knowledgeInfo := &KnowledgeInfo{}
	knowledgeInfo.padding(knowledge)
	return knowledgeInfo
}

func NewKnowledgeInfos(knowledge ...*model.Knowledge) []*KnowledgeInfo {
	knowledgeInfos := make([]*KnowledgeInfo, len(knowledge))
	for i, k := range knowledge {
		knowledgeInfos[i] = new(KnowledgeInfo)
		knowledgeInfos[i].padding(k)
	}
	return knowledgeInfos
}
