package piscine

/*
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
*/
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current.Left != nil {
		current = current.Left
	}

	return current
}

/*
func main() {
	// Create a sample binary search tree
	root := &TreeNode{Data: "5"}
	root.Left = &TreeNode{Data: "3"}
	root.Right = &TreeNode{Data: "7"}
	root.Left.Left = &TreeNode{Data: "2"}
	root.Left.Right = &TreeNode{Data: "4"}
	root.Right.Left = &TreeNode{Data: "6"}
	root.Right.Right = &TreeNode{Data: "8"}

	// Find the node with the minimum value
	minNode := BTreeMin(root)

	if minNode != nil {
		fmt.Println("Minimum value:", minNode.Data)
	} else {
		fmt.Println("Tree is empty")
	}
}
*/
