package text

import "github.com/eapache/queue"

// TrieNode ac自动机节点
type TrieNode struct {
	value rune
	next  map[rune]*TrieNode
	fail  *TrieNode
	emit  string
}

func newNode(ch rune) (node *TrieNode) {
	node = new(TrieNode)
	node.value = ch
	node.next = map[rune]*TrieNode{}
	return node
}

// AcTrie ac自动机匹配字符串算法
type AcTrie struct {
	root *TrieNode
}

// Search 返回匹配的字符串
func (ac *AcTrie) Search(s string) (list []string, index []int) {
	node := ac.root
	for i, c := range []rune(s) {
		matched := true
		for {
			_, ok := node.next[c]
			if ok {
				break
			}
			if node.fail == nil {
				matched = false
				node = ac.root
				break
			}
			node = node.fail
		}
		if !matched {
			continue
		}
		node = node.next[c]
		p := node
		for p != nil {
			if p.emit != "" {
				list = append(list, p.emit)
				index = append(index, i+1)
			}
			p = p.fail
		}
	}
	return list, index
}

//BuildAcTrie 构建一个 ac 自动机
func BuildAcTrie(words []string) (acTrie *AcTrie) {
	acTrie = new(AcTrie)
	acTrie.root = newNode(rune('r'))
	for _, word := range words {
		node := acTrie.root
		for _, ch := range []rune(word) {
			if _, ok := node.next[ch]; !ok {
				node.next[ch] = newNode(ch)
			}
			node = node.next[ch]
		}
		node.emit = word
	}
	queue := queue.New()
	queue.Add([]*TrieNode{acTrie.root, nil})
	for queue.Length() > 0 {
		nodeParent := queue.Remove().([]*TrieNode)
		curr, parent := nodeParent[0], nodeParent[1]
		for _, sub := range curr.next {
			queue.Add([]*TrieNode{sub, curr})
		}
		if parent == nil {
			continue
		}
		if parent == acTrie.root {
			curr.fail = acTrie.root
		} else {
			fail := parent.fail
			for fail != nil {
				_, ok := fail.next[curr.value]
				if ok {
					break
				}
				fail = fail.fail
			}
			if fail != nil {
				curr.fail = fail.next[curr.value]
			} else {
				curr.fail = acTrie.root
			}
		}
	}
	return acTrie
}
