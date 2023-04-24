package text

import "sort"

// SliceSame 对于两个列表的值是否一样
func SliceSame(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	tmpa, tempb := []string{}, []string{}
	copy(tmpa, a)
	copy(tempb, b)
	sort.Strings(tmpa)
	sort.Strings(tempb)
	for i := 0; i < len(tmpa); i++ {
		if tmpa[i] != tempb[i] {
			return false
		}
	}
	return true
}
