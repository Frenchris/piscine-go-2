package piscine

func ListReverse(l *List) {

	if l.Head != nil {
		prev := &NodeL{Data: nil}
		curr := l.Head
		next := &NodeL{Data: nil}

		for curr != nil {
			next = curr.Next
			curr.Next = prev
			prev = curr
			curr = next

		}
		l.Tail = ListAt(prev, ListSize(l)+1)
		l.Tail.Next = nil
		l.Head = prev
	}
}
