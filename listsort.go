package piscine

func ListSort(l *NodeL) *NodeL {

	c := l

	arr := NodeToArrayStrings(c)

	SortIntegerTable(arr)

	c = ArrayIntsToNode(arr)

	return c
}

func NodeToArrayStrings(n *NodeL) []int {
	var arr []int

	for n != nil {
		arr = append(arr, (n.Data).(int))
		n = n.Next
	}
	return arr
}

func ArrayIntsToNode(arr []int) *NodeL {

	j := 0
	var m *NodeL

	for i := 0; i < ArrayIntLength(arr); i++ {
		j = arr[i]
		m = NodePushBack(m, j)
	}
	return m
}

func NodePushBack(l *NodeL, data int) *NodeL {
	n := &NodeL{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
