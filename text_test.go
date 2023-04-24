package textalg_test

import (
	"testing"

	textalg "github.com/memory-overflow/go-text-algorithm"
)

func TestActrie(t *testing.T) {
	ac := textalg.BuildAcTrie([]string{"哈哈哈", "234", "dfg"})
	list, index := ac.Search("哈哈哈哈23434dfgdd")
	for i, l := range list {
		t.Log(l, index[i])
	}
}

func TestTextSim(t *testing.T) {
	sim := textalg.TextSim("编辑距离测试", "测试一下距离")
	if sim != 0 {
		t.Error("Failed")
	}
}

func TestLevenshtein(t *testing.T) {
	dist := textalg.Levenshtein([]rune("编辑距离测试"), []rune("测试一下距离"))
	t.Logf("dist: %d", dist)
}

func TestSliceSmae(t *testing.T) {
	a := []string{"3", "2", "1"}
	same := textalg.SliceSame(a, a)
	t.Logf("is same: %v", same)
	// test can not change order of a
	t.Log(a)
}

func TestKmp(t *testing.T) {
	k := textalg.BuildKmp("a")
	indexs := k.Search("aaaaab") // find "a" in "aaaaab"
	t.Log(indexs)
	k.AppendPatternStr("a")
	indexs = k.Search("aaaaab") // find "aa" in "aaaaab"
	t.Log(indexs)
	k.AppendPatternStr("a")
	indexs = k.Search("aaaaab") // find "aaa" in "aaaaab"
	t.Log(indexs)
	k.AppendPatternStr("b")
	indexs = k.Search("aaaaab") // find "aaab" in "aaaaab"
	t.Log(indexs)
	k.AppendPatternStr("b")
	indexs = k.Search("aaaaab") // find "aaabb" in "aaaaab"
	t.Log(indexs)
	k.ResetPatternStr("ab")
	indexs = k.Search("aaaaab") // find "ab" in "aaaaab"
	t.Log(indexs)
}
