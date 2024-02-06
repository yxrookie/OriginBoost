package similarity

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/yanyiwu/gojieba"
)

func Sortres(old string, tempstring []string) [][]string {
	res := make([][]string, 0)
	simslice := make([]float64, len(tempstring))
	silmap := make(map[float64]string)

	// 使用结巴分词库
	seg := gojieba.NewJieba()
	defer seg.Free()
	// 计算文本的 TF-IDF 向量
	vectorStart, err := calculateTFIDFVector(old, seg)
	if err != nil {
		fmt.Println("Error calculating TF-IDF vector for text2:", err)
		return res
	}
	for i := 0; i < len(tempstring); i++ {
		vectorNow, err := calculateTFIDFVector(tempstring[i], seg)
		if err != nil {
			fmt.Println("Error calculating TF-IDF vector for text2:", err)
			return res
		}
		// 计算余弦相似度
		cosineSimilarity := calculateCosineSimilarity(vectorStart, vectorNow)
		simslice[i] = cosineSimilarity
		silmap[cosineSimilarity] = tempstring[i]
	}
	sort.Float64s(simslice)

	for i := 0; i < len(simslice); i++ {
		//fmt.Print(simslice[i], " ")
		// 将小数转换为字符串，并保留两位小数
		similarity := strconv.FormatFloat(simslice[i], 'f', 2, 64)
		res = append(res, []string{similarity, silmap[simslice[i]]})
	}
	return res
}
