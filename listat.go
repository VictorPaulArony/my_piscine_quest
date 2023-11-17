package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}*/

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}

	current := l
	for i := 0; i < pos; i++ {
		if current == nil {
			return nil
		}
		current = current.Next
	}

	return current
}

/*func main() {
	list := &NodeL{Data: "Hello"}
	list.Next = &NodeL{Data: "World"}
	list.Next.Next = &NodeL{Data: "How"}
	list.Next.Next.Next = &NodeL{Data: "Are"}
	list.Next.Next.Next.Next = &NodeL{Data: "You"}

	node := ListAt(list, 2)
	if node != nil {
		fmt.Println("Data at position 2:", node.Data)
	} else {
		fmt.Println("Error: Invalid position")
	}
}*/
