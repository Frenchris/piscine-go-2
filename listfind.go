package piscine

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {

	n := l.Head
	for n != nil {
		if CompStr(ref, n.Data) {
			return &n.Data
		}
		n = n.Next
	}
	return nil
}

func CompStr(a, b interface{}) bool {
	return a == b
}
