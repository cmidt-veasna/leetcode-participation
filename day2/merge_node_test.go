package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func mergeCaseInHelper(is []int) *ListNode {
	head := &ListNode{Val: is[0]}
	ns := head
	for i := 1; i < len(is); i++ {
		ns.Next = &ListNode{Val: is[i]}
		ns = ns.Next
	}
	return head
}

func mergeCaseOutHelper(head *ListNode) []int {
	result := make([]int, 0, 100)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeNode(t *testing.T) {
	var testCases = [][][]int{
		{{0, 3, 1, 0, 4, 5, 2, 0}, {4, 11}},
		{{0, 1, 0, 3, 0, 2, 2, 0}, {1, 3, 4}},
	}
	for i, tc := range testCases {
		t.Logf("TestMergeNode #%d: %v", i+1, tc[0])
		var result = mergeNodes(mergeCaseInHelper(tc[0]))
		assert.Equal(t, tc[1], mergeCaseOutHelper(result))
	}
}
