package piscine

var nbrNodes int

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}

	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
	return nil
}

// func BTreeNodes(root *TreeNode) int {
// 	fmt.Println(nbrNodes)
// 	if root != nil {
// 		nbrNodes++
// 		BTreeNodes(root.Left)
// 		BTreeNodes(root.Right)
// 	}

// 	count := nbrNodes
// 	nbrNodes = 0
// 	return count
// }
