package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListForEach(l *List, f func(*NodeL)) {
	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

/*
func main() {
	list := &List{}
	list.Head = &NodeL{Data: 5}
	list.Head.Next = &NodeL{Data: "Hello"}
	list.Head.Next.Next = &NodeL{Data: 10}

	fmt.Println("Before applying functions:")
	PrintList(list)

	ListForEach(list, Add2_node)
	fmt.Println("After applying Add2_node:")
	PrintList(list)

	ListForEach(list, Subtract3_node)
	fmt.Println("After applying Subtract3_node:")
	PrintList(list)
}

func PrintList(l *List) {
	current := l.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}*/
