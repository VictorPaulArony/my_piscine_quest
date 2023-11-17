package piscine

/*import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListLast(l *List) interface{} {
	if l.Tail != nil {
		return l.Tail.Data
	}
	return nil
}

/*func main() {
	list := &List{}
	list.Head = &NodeL{Data: "Hello"}
	list.Head.Next = &NodeL{Data: "World"}
	list.Head.Next.Next = &NodeL{Data: "How"}
	list.Head.Next.Next.Next = &NodeL{Data: "Are"}
	list.Head.Next.Next.Next.Next = &NodeL{Data: "You"}

	last := ListLast(list)
	fmt.Println("Last element:", last)
}*/
