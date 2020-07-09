package main

import (
	"fmt"
	"log"
	"strings"
)

func debug(v ...interface{}) {
	log.Println(v...)
}

func toString(i interface{}) string {
	switch i.(type) {
	case int:
		return fmt.Sprintf("%v", i)
	case string:
		return fmt.Sprintf("%v", i)
	case bool:
		return fmt.Sprintf("%v", i)
	default:
		return fmt.Sprintf("%p", i)
	}
}

func isMatch(s string, p string) bool {
	begin := new(Node)
	begin.C = '>'
	begin.Size = generatePattern(begin, p, 0)
	debug(begin.String())
	return check(begin, s, 0)
}

type Node struct {
	C        byte
	Parent   *Node
	Children map[byte][]*Node
	End      bool
	Size     int
}

func (n *Node) String() string {
	return n.StringLevel(0, make(map[*Node]bool))
}

func (n *Node) StringLevel(level int, finishNodes map[*Node]bool) string {
	r := make([]string, 0)
	if n.End {
		r = append(r, fmt.Sprintf("  id%s{%v};", toString(n), string(n.C)))
	} else {
		r = append(r, fmt.Sprintf("  id%s(%v);", toString(n), string(n.C)))
	}
	finishNodes[n] = true
	for k, v := range n.Children {
		for _, c := range v {
			if _, ok := finishNodes[c]; !ok {
				r = append(r, c.StringLevel(level+1, finishNodes))
			}
			r = append(r, fmt.Sprintf("  id%s -- %s --> id%s;", toString(n), string(k), toString(c)))
		}
	}
	return strings.Join(r, "\n")
}

func (n *Node) Append(c byte, child *Node) {
	m := n.Children
	if m == nil {
		m = make(map[byte][]*Node)
		n.Children = m
	}
	list := m[c]
	if list == nil {
		list = make([]*Node, 0)
	}
	for _, v := range list {
		if v == child {
			m[c] = list
			return
		}
	}
	list = append(list, child)
	m[c] = list
}

func generatePattern(now *Node, str string, idx int) int {
	if len(str) <= idx {
		now.End = true
		return now.Size
	}
	vnow := now
	switch str[idx] {
	case '*':
		now.Size = 0
		now.Append(now.C, now)
	default:
		node := new(Node)
		node.C = str[idx]
		now.Append(str[idx], node)
		node.Parent = now
		node.Size = 1
		vnow = node
	}
	ret := generatePattern(vnow, str, idx+1)
	if ret == 0 {
		now.End = true
	}
	addParent := now
	for addParent.Parent != nil {
		if addParent.Size == 0 {
			debug(toString(vnow), " -> ", toString(addParent.Parent))
			addParent.Parent.Append(vnow.C, vnow)
			addParent = addParent.Parent
		} else {
			break
		}
	}
	return now.Size + ret
}

func check(now *Node, str string, idx int) bool {
	if len(str) <= idx {
		return now.End
	}
	list := now.Children['.']
	for _, v := range now.Children[str[idx]] {
		list = append(list, v)
	}
	for _, v := range list {
		r := check(v, str, idx+1)
		if r {
			return true
		}
	}
	return false
}

// 作者：qianjigui
// 链接：https://leetcode-cn.com/problems/regular-expression-matching/solution/yi-bu-dao-wei-zhi-jie-an-zheng-ze-biao-da-shi-de-s/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func isMatch2(s, p string) bool {
	if s == p {
		return true
	}
	if len(p) > 1 && p[1] == '*' {
		switch p[0] {
		case '.':
			for i := 0; i <= len(s); i++ {
				if isMatch2(s[i:], p[2:]) {
					return true
				}
			}
		default:
			if isMatch2(s, p[2:]) {
				return true
			}

			for i := 0; i < len(s) && s[i] == p[0]; i++ {
				if isMatch2(s[i+1:], p[2:]) {
					return true
				}
			}

		}
	} else {
		if s == "" || p == "" {
			return false
		}
		if s[0] == p[0] || p[0] == '.' {
			return isMatch2(s[1:], p[1:])
		}
	}

	return false
}
