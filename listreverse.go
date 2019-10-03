package piscine

func ListReverse(l *List) {

	lastNode := l.Head

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
		l.Head.Next = nil
		l.Tail = lastNode
		l.Head = prev
	}
}
