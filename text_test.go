package text_test

import (
	"testing"

	"github.com/memory-overflow/go-text-algorithm"
)

func TestActrie(t *testing.T) {
	ac := text.BuildAcTrie([]string{"哈哈哈", "234", "dfg"})
	list, index := ac.Search("哈哈哈哈23434dfgdd")
	for i, l := range list {
		t.Log(l, index[i])
	}
}

func TestTextSim(t *testing.T) {
	sim := text.TextSim("编辑距离测试", "测试一下距离")
	if sim != 0 {
		t.Error("Failed")
	}
}

func TestLevenshtein(t *testing.T) {
	dist := text.Levenshtein([]rune("编辑距离测试"), []rune("测试一下距离"))
	t.Logf("dist: %d", dist)
}

func TestSliceSmae(t *testing.T) {
	a := []string{"3", "2", "1"}
	same := text.SliceSame(a, a)
	t.Logf("is same: %v", same)
	// test can not change order of a
	t.Log(a)
}
