package tree_test

import (
	"fmt"
	"testing"

	"github.com/indrasaputra/spenmo/cmd/tree"
)

func TestTreeNode_Find(t *testing.T) {
	t.Run("number exists and found", func(t *testing.T) {
		values := []int{6, 7, 100, 10, 8, 14}

		tree := createTree()
		for _, val := range values {
			found := tree.Find(val)
			if !found {
				t.Error(fmt.Sprintf("val %d should be found!", val))
			}
		}
	})

	t.Run("number doesn't exist and not found", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 9}

		tree := createTree()
		for _, val := range values {
			found := tree.Find(val)
			if found {
				t.Error(fmt.Sprintf("val %d should not be found!", val))
			}
		}
	})
}

func createTree() *tree.TreeNode {
	root := createNode(6)
	root.Left = createNode(7)
	root.Right = createNode(100)
	root.Left.Left = createNode(10)
	root.Left.Right = createNode(8)
	root.Right.Right = createNode(14)

	return root
}

func createNode(val int) *tree.TreeNode {
	return &tree.TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
