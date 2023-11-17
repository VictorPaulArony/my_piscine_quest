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

func ListReverse(l *List) {
	prev := (*NodeL)(nil)
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Head = prev
}

/*
package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	var prev *NodeL
	for l.Head != nil {
		l.Head.Next, prev, l.Head = prev, l.Head, l.Head.Next
	}
	l.Head = prev
}
func ListReverse(l *List) {
    // Initialize a pointer to the previous node
    var prev *NodeL = nil

    // Start with the head node
    current := l.Head

    // Traverse the list
    for current != nil {
        // Save the next node in a temporary variable
        next := current.Next

        // Reverse the link by pointing the current node's Next to the previous node
        current.Next = prev

        // Move the pointers forward
        prev = current
        current = next
    }

    // Update the head pointer to the last node (which is now the first node after reversal)
    l.Head = prev
}
*/

/*func main() {
	list := &List{}
	list.Head = &NodeL{Data: "Hello"}
	list.Head.Next = &NodeL{Data: "World"}
	list.Head.Next.Next = &NodeL{Data: "How"}
	list.Head.Next.Next.Next = &NodeL{Data: "Are"}
	list.Head.Next.Next.Next.Next = &NodeL{Data: "You"}

	fmt.Println("Before reverse:")
	PrintList(list)

	ListReverse(list)

	fmt.Println("After reverse:")
	PrintList(list)
}

func PrintList(l *List) {
	current := l.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}*/
