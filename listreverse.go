package piscine

func ListReverse(l *List) {

	if l.Head != nil {
		prev := &NodeL{}
		prev = nil
		curr := l.Head
		next := prev

		for curr != nil {
			next = curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		temp := l.Head
		l.Head = l.Tail
		l.Tail = temp

	}
}
