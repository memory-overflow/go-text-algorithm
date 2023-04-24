package text

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Levenshtein 文本编辑距离
func Levenshtein(word1, word2 []rune) int {
	/*
		计算编辑距离
		Args
				word1: 第一个文本
				word2: 第二个文本
		Returns
				两个文本的编辑距离
	*/
	if len(word1) == 0 || len(word2) == 0 {
		return max(len(word1), len(word2))
	}
	tmp := []int{}
	for i := 0; i < len(word2)+1; i++ {
		tmp = append(tmp, i)
	}
	value := 0
	for i := range word1 {
		tmp[0] = i + 1
		last := i
		for j := range word2 {
			if word1[i] == word2[j] {
				value = last
			} else {
				value = 1 + min(last, min(tmp[j], tmp[j+1]))
			}
			last = tmp[j+1]
			tmp[j+1] = value
		}
	}
	return value
}

// TextSim 计算文本的相识度
func TextSim(str1, str2 string) float32 {
	// 需要把 string 转换成 rune
	s1 := []rune(str1)
	s2 := []rune(str2)
	if len(s1) == 0 && len(s2) == 0 {
		return 1.0
	}
	n := Levenshtein(s1, s2)
	maxn := max(len(s1), len(s2))
	l1 := 1.0 - float32(n)/float32(maxn)
	return l1
}
