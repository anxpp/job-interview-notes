package main

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestLc0508(t *testing.T) {
	root := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}},
	}
	fmt.Printf("%v\n", findFrequentTreeSum(root))
}

func findFrequentTreeSum(root *TreeNode) []int {
	m := map[int]int{}
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		node.Val += helper(node.Left) + helper(node.Right)
		m[node.Val]++
		return node.Val
	}
	helper(root)
	var max int
	var ans []int
	for k, v := range m {
		if v == max {
			ans = append(ans, k)
		} else if v > max {
			ans = []int{k}
			max = v
		}
	}
	return ans
}
