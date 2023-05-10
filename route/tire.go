package route

import (
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

func (n *node) matchChildren(part string) []*node {
	children := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			children = append(children, child)
		}
	}

	return children
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) >= height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}

		return n
	}

	children := n.matchChildren(parts[height])
	for _, item := range children {
		child := item.search(parts, height+1)
		if child != nil {
			return child
		}
	}

	return nil
}

func (n *node) insert(pattern string, parts []string, height int) {
	if height >= len(parts) {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}
