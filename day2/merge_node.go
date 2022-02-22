package day2

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	ns := &ListNode{}
	head = head.Next
	for head.Val != 0 {
		ns.Val += head.Val
		head = head.Next
	}
	if head.Next != nil {
		ns.Next = mergeNodes(head)
	}
	return ns
}