package day3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][3]int) *TreeNode {
	var root *TreeNode
	var a, b *TreeNode
	var mp = make(map[int]*TreeNode)
	var rp = make(map[int]int)
	var ok bool
	for _, desc := range descriptions {
		if a, ok = mp[desc[0]]; !ok {
			a = &TreeNode{Val: desc[0]}
			mp[a.Val] = a
		}
		if b, ok = mp[desc[1]]; !ok {
			b = &TreeNode{Val: desc[1]}
			mp[b.Val] = b
		}

		if desc[2] == 1 {
			a.Left = b
		} else {
			a.Right = b
		}
		rp[b.Val] = a.Val

		if root == nil || b.Val == root.Val {
			root = a
		}
	}
	for vr, ok := rp[root.Val]; ok; vr, ok = rp[root.Val] {
		root = mp[vr]
	}
	return root
}
