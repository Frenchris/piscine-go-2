package piscine

func SortedListMerge(n1 *NodeL, n2 *NodeL) *NodeL {

	return ListSort(NodeMerge(n1, n2))
}

func NodeMerge(n1 *NodeL, n2 *NodeL) *NodeL {

	p := &NodeL{}
	p = nil

	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	p = n1
	for p.Next != nil {
		p = p.Next
	}
	p.Next = n2

	return n1
}
