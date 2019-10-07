package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		nd := &TreeNode{}
		nd = BTreeMin(root.Right)

		root.Data = nd.Data

		root.Right = BTreeDeleteNode(root.Right, nd)
	}
	return root
}
