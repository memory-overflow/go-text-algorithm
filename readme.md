- [text 模块](#text-模块)
  - [SliceSame](#slicesame)
  - [Aho-Corasick automaton](#aho-corasick-automaton)
  - [计算文本编辑距离](#计算文本编辑距离)
  - [计算文本相似度](#计算文本相似度)

# text 模块
golang 里面的 strings 库已经有了很多丰富的字符串处理功能，但是都是偏向于基础处理。

text模块提供了一些字符串处理相关的算法能力。

## SliceSame
- SliceSame——判断两个字符串数字是否相同。

example: [TestSliceSmae](https://github.com/memory-overflow/go-common-library/blob/main/text/text_test.go#L29)
```go
import (
  "testing"

  "github.com/memory-overflow/go-common-library/text"
)

func TestSliceSmae(t *testing.T) {
	a := []string{"3", "2", "1"}
	same := text.SliceSame(a, a)
	t.Logf("is same: %v", same)
  // test can not change order of a
	t.Log(a)
}
```

## Aho-Corasick automaton
ac 自动机是一种多模式串的匹配算法。

一个常见的例子就是给出 n 个单词，再给出一段包含 m 个字符的文章，让你找出有多少个单词在文章里出现过。

比较容易想到的做法是，调用 n 次 `strings.Contains(s, xxx)`。假设 n 个单词平局长度为 k, 这样处理的算法时间复杂度为 O(n * k * m)。而使用 ac 自动机可以加速上述过程，整体算法时间复杂度只需要 O(n*k + m)。

example: [TestActrie](https://github.com/memory-overflow/go-common-library/blob/main/text/text_test.go#L9)
```go
import (
  "testing"

  "github.com/memory-overflow/go-common-library/text"
)

func TestActrie(t *testing.T) {
  // 在字符串 "哈哈哈哈23434dfgdd" 中找出所有 "哈哈哈", "234"，"dfg" 出现的位置。
  // 使用模式串构建一个 ac 自动机
  ac := text.BuildAcTrie([]string{"哈哈哈", "234", "dfg"})
  // 匹配母串
  list, index := ac.Search("哈哈哈哈23434dfgdd")
  for i, l := range list {
    t.Log(l, index[i])
  }
}
```

## 计算文本编辑距离
编辑距离(Edit Distance)：是一个度量两个字符序列之间差异的字符串度量标准，两个单词之间的编辑距离是将一个单词转换为另一个单词所需的单字符编辑（插入、删除或替换）的最小数量。一般来说，编辑距离越小，两个串的相似度越大。

example: [TestLevenshtein](https://github.com/memory-overflow/go-common-library/blob/main/text/text_test.go#L24)
```go
import (
  "testing"

  "github.com/memory-overflow/go-common-library/text"
)

func TestLevenshtein(t *testing.T) {
	dist := text.Levenshtein([]rune("编辑距离测试"), []rune("测试一下距离"))
	t.Logf("dist: %d", dist)
}
```

## 计算文本相似度
通过编辑距离，计算两个文本的相似度。

example: [TestTextSim](https://github.com/memory-overflow/go-common-library/blob/main/text/text_test.go#L17)
```go
import (
  "testing"

  "github.com/memory-overflow/go-common-library/text"
)

func TestTextSim(t *testing.T) {
	sim := text.TextSim("编辑距离测试", "测试一下距离")
  t.Logf("sim: %f", sim)
}
```