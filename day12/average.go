package day12

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	var averageOfSubtreeSum func(node *TreeNode) (sum, nnode, found int)
	averageOfSubtreeSum = func(node *TreeNode) (sum, nnode, found int) {
		if node == nil {
			return 0, 0, 0
		} else {
			s, nn, f := averageOfSubtreeSum(node.Left)
			sum, nnode, found = sum+s, nnode+nn, found+f
			s, nn, f = averageOfSubtreeSum(node.Right)
			sum, nnode, found = sum+s, nnode+nn, found+f
		}
		sum, nnode = sum+node.Val, nnode+1
		if sum/nnode == node.Val {
			found++
		}
		return
	}
	_, _, result := averageOfSubtreeSum(root)
	return result
}
