package gol

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ParseTreeFromString(text string) *TreeNode {
	parts := strings.Split(text, ",")
	root, _ := parseFromSlice(parts, 0)
	fmt.Printf("Check Parse\nSource = %+v\nParsed =", text)
	for _, num := range InOrderTraversalRecursive(root) {
		fmt.Printf(" %+v", num)
	}
	fmt.Printf("\n")
	return root
}

func parseFromSlice(data []string, start int) (*TreeNode, int) {
	if data[start] == "null" {
		return nil, start
	}
	val, _ := strconv.ParseInt(data[start], 10, 32)
	p := &TreeNode{Val: int(val)}
	if start < len(data)-1 {
		left, end := parseFromSlice(data, start+1)
		p.Left = left
		start = end
	}
	if start < len(data)-1 {
		right, end := parseFromSlice(data, start+1)
		p.Right = right
		start = end
	}
	return p, start
}

func InOrderTraversalRecursive(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = append(result, InOrderTraversalRecursive(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, InOrderTraversalRecursive(root.Right)...)
	}
	return result
}

func InOrderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stores := []*TreeNode{}
	p := root
	isStore := false
	for {
		// fmt.Printf("stores =")
		// for i := range stores {
		// 	fmt.Printf(" %+v", *stores[i])
		// }
		// fmt.Printf("\n")

		if p.Left != nil && !isStore {
			stores = append(stores, p)
			p = p.Left
			isStore = false
			continue
		}
		result = append(result, p.Val)
		if p.Right != nil {
			p = p.Right
			isStore = false
			continue
		}
		if len(stores) < 1 {
			break
		}
		p = stores[len(stores)-1]
		stores = stores[:len(stores)-1]
		isStore = true
	}

	return result
}
