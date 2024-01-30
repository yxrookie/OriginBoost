package similarity

import (
	"fmt"
	"github.com/yanyiwu/gojieba"
	"math"
	
)

// calculateCosineSimilarity 计算两个向量的余弦相似度
func calculateCosineSimilarity(vector1, vector2 map[string]float64) float64 {
	dotProduct := 0.0
	magnitude1 := 0.0
	magnitude2 := 0.0

	for term, weight1 := range vector1 {
		weight2, exists := vector2[term]
		if exists {
			dotProduct += weight1 * weight2
		}
		magnitude1 += weight1 * weight1
	}

	for _, weight2 := range vector2 {
		magnitude2 += weight2 * weight2
	}

	if magnitude1 == 0 || magnitude2 == 0 {
		return 0.0
	}

	return dotProduct / (math.Sqrt(magnitude1) * math.Sqrt(magnitude2))
}

// calculateTFIDFVector 计算文本的 TF-IDF 向量
func calculateTFIDFVector(text string, seg *gojieba.Jieba) (map[string]float64, error) {
	// 使用结巴分词进行中文分词
	words := seg.Cut(text, true)

	vector := make(map[string]float64)

	for _, word := range words {
		// 计算 TF 值
		if _, exists := vector[word]; exists {
			vector[word]++
		} else {
			vector[word] = 1
		}
	}

	// 计算 TF-IDF 值（这里简化了 IDF 的计算，实际应用中可能需要更复杂的逻辑）
	for term, tf := range vector {
		// 计算 TF-IDF 值
		tfidf := tf * 1.0 // 简化的 IDF 值

		vector[term] = tfidf
	}

	return vector, nil
}

func main() {
	// 示例文本
	text1 := "这是一个中文文本的例子"
	text2 := "这是一个测试文本的例子"

	// 使用结巴分词库
	seg := gojieba.NewJieba()
	defer seg.Free()

	// 计算文本的 TF-IDF 向量
	vector1, err := calculateTFIDFVector(text1, seg)
	if err != nil {
		fmt.Println("Error calculating TF-IDF vector for text1:", err)
		return
	}

	vector2, err := calculateTFIDFVector(text2, seg)
	if err != nil {
		fmt.Println("Error calculating TF-IDF vector for text2:", err)
		return
	}

	// 计算余弦相似度
	cosineSimilarity := calculateCosineSimilarity(vector1, vector2)

	// 打印相似度
	fmt.Printf("相似度: %.4f\n", cosineSimilarity)
}
