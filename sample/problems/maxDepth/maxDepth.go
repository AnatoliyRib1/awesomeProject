package maxDepth

import . "sample/utils/binarytree"

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := MaxDepth(root.Left)
	rightHeight := MaxDepth(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
