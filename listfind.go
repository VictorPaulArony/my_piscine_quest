package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
*/
func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	for current != nil {
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}
	return nil
}

/*
func main() {
	list := &List{}
	list.Head = &NodeL{Data: "Hello"}
	list.Head.Next = &NodeL{Data: "World"}
	list.Head.Next.Next = &NodeL{Data: "How"}
	list.Head.Next.Next.Next = &NodeL{Data: "Are"}
	list.Head.Next.Next.Next.Next = &NodeL{Data: "You"}

	ref := "How"
	result := ListFind(list, ref, CompStr)
	if result != nil {
		fmt.Println("Data found:", *result)
	} else {
		fmt.Println("Data not found")
	}
}*/
