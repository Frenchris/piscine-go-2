package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		rDepth := BTreeLevelCount(root.Right)
		lDepth := BTreeLevelCount(root.Left)

		if lDepth > rDepth {
			return lDepth + 1
		} else {
			return rDepth + 1
		}
	}

}
