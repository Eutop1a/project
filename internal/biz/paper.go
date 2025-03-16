package biz

import (
	"math"
)

type GenerateRequest struct {
	Types         map[string]int // 题型分布，如 {"choice": 5, "fill": 3}
	Difficulty    float64        // 平均难度（1-5）
	Total         int            // 题目总数
	MaxSimilarity float64        // 最大允许相似度（如 0.3）
}

type PaperQuestion struct {
	ID      int
	Content string
	Vector  []float64 // 题目文本的向量表示（预计算）
}

// 计算余弦相似度
func CosineSimilarity(a, b []float64) float64 {
	dotProduct, normA, normB := 0.0, 0.0, 0.0
	for i := range a {
		dotProduct += a[i] * b[i]
		normA += math.Pow(a[i], 2)
		normB += math.Pow(b[i], 2)
	}
	return dotProduct / (math.Sqrt(normA) * math.Sqrt(normB))
}

// 过滤相似题目（暴力匹配，适用于小规模题库）
func FilterSimilarQuestions(candidates []PaperQuestion, maxSimilarity float64) []PaperQuestion {
	result := make([]PaperQuestion, 0)
	for i, q := range candidates {
		isUnique := true
		for _, existing := range result {
			sim := CosineSimilarity(q.Vector, existing.Vector)
			if sim > maxSimilarity {
				isUnique = false
				break
			}
		}
		if isUnique {
			result = append(result, candidates[i])
		}
	}
	return result
}

//
//type PaperUseCase struct {
//	repo QuestionRepo // 题库数据接口（需在data层实现）
//}
//
//// 生成试卷核心算法
//func (uc *PaperUseCase) GeneratePaper(ctx context.Context, req *GenerateRequest) ([]int, error) {
//	// 1. 获取候选题目（按题型、难度筛选）
//	candidates, err := uc.repo.GetCandidates(ctx, req.Types, req.Difficulty)
//	if err != nil {
//		return nil, err
//	}
//
//	// 2. 相似度过滤
//	filtered := FilterSimilarQuestions(candidates, req.MaxSimilarity)
//
//	// 3. 随机选题（可替换为更复杂的算法）
//	rand.Seed(time.Now().UnixNano())
//	selected := make([]int, 0)
//	for len(selected) < req.Total && len(filtered) > 0 {
//		idx := rand.Intn(len(filtered))
//		selected = append(selected, filtered[idx].ID)
//		filtered = append(filtered[:idx], filtered[idx+1:]...)
//	}
//
//	return selected, nil
//}
//
//// 在biz/paper.go中添加
//func geneticAlgorithm(candidates []PaperQuestion, req GenerateRequest) []int {
//	populationSize := 100
//	maxGenerations := 50
//
//	// 1. 初始化种群
//	population := initPopulation(populationSize, candidates, req.Total)
//
//	for i := 0; i < maxGenerations; i++ {
//		// 2. 计算适应度（题目差异最大化）
//		scores := calculateFitness(population, candidates)
//
//		// 3. 选择、交叉、变异
//		population = evolvePopulation(population, scores)
//	}
//
//	return getBestSolution(population, candidates)
//}
