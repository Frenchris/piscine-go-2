package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	out := BTreeSearchItem(root, node.Data)

	if out.Parent.Data > rplc.Data {
		out.Parent.Left.Data = rplc.Data
	} else {
		out.Parent.Right.Data = rplc.Data
	}

	return root
}
