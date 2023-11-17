package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}

/*
func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		return
	}
	n.Next = l.Head
	l.Head = n
}
*/
/*func main() {
	list := &List{}
	ListPushFront(list, "Hello")
	ListPushFront(list, "man")
	ListPushFront(list, "how are you")

	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}*/
