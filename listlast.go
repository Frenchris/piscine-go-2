package piscine

func ListLast(l *List) interface{} {

	for l.Head != nil {
		if l.Head.Next == nil {
			return l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return l.Tail
}
