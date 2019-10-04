package piscine

func SortListInsert(l *NodeL, data_ref int) *NodeL {
	return ListSort(NodePushBack(l, data_ref))
}
