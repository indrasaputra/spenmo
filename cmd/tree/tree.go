package tree

// TreeNode represents node in tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Find finds a number in the tree.
func (tn *TreeNode) Find(val int) bool {
	node := searchBST(tn, val)
	return node != nil
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if x := searchBST(root.Left, val); x != nil {
		return x
	}
	return searchBST(root.Right, val)
}
