package piscine

/*type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}
*/
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}

	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}

	return BTreeSearchItem(root.Right, elem)
}

/*
func main() {
	root := &TreeNode{Data: "4"}
	root.Left = &TreeNode{Data: "1"}
	root.Right = &TreeNode{Data: "7"}
	root.Right.Left = &TreeNode{Data: "5"}

	elem := "7"
	result := BTreeSearchItem(root, elem)

	if result != nil {
		fmt.Printf("Found element %s in the tree\n", elem)
	} else {
		fmt.Printf("Element %s not found in the tree\n", elem)
	}
}
*/
