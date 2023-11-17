package piscine

/*
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
*/
func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current.Right != nil {
		current = current.Right
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

	// Find the node with the maximum value
	maxNode := BTreeMax(root)

	if maxNode != nil {
		fmt.Println("Maximum value:", maxNode.Data)
	} else {
		fmt.Println("Tree is empty")
	}
}
*/
