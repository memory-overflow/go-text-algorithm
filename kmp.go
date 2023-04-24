package textalg

// Kmp 快速字符串匹配算法，时间复杂度 O(n + m), 其中 n 代表模式串的长度，m 代码匹配母串的长度
type Kmp struct {
	pattern string
	next    []int
}

// BuildKmp 构建 kmp
// patternStr 模式串
func BuildKmp(patternStr string) *Kmp {
	kmp := Kmp{
		pattern: patternStr,
	}
	kmp.buildNext(0)
	return &kmp
}

func (k *Kmp) buildNext(start int) {
	r := []rune(k.pattern)
	l := len(r)
	// next 要多构建 1 个长度
	for i := start; i <= l; i++ {
		if i == 0 || i == 1 {
			k.next = append(k.next, 0) // next[0] and next[1] is 0
			continue
		}
		j := k.next[i-1]
		for ; j > 0 && r[i-1] != r[j]; j = k.next[j] {
		}
		if r[i-1] == r[j] {
			j++
		}
		k.next = append(k.next, j)
	}
}

// AppendPatternStr 扩充模式串长度
func (k *Kmp) AppendPatternStr(patternStr string) {
	start := len([]rune(k.pattern)) + 1
	k.pattern += patternStr
	k.buildNext(start)
}

// ResetPatternStr 重置模式串
func (k *Kmp) ResetPatternStr(patternStr string) {
	k.pattern = patternStr
	k.next = nil
	k.buildNext(0)
}

// Search 返回所有 pattern 在 content 中出现的开始位置的 index
func (k Kmp) Search(content string) (indexs []int) {
	rpattern, rcontent := []rune(k.pattern), []rune(content)
	n, m := len(rcontent), len(rpattern)
	if n == 0 || m == 0 {
		return nil
	}
	j := 0
	for i := 0; i < n; i++ {
		for ; j > 0 && rcontent[i] != rpattern[j]; j = k.next[j] {
		}
		if rcontent[i] == rpattern[j] {
			j++
		}
		if j == m {
			indexs = append(indexs, i-m+1)
			j = k.next[j]
		}
	}
	return indexs
}
