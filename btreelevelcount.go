package piscine

/*type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/
func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

/*
func main() {
	root := &TreeNode{Data: "4"}
	root.Left = &TreeNode{Data: "2"}
	root.Right = &TreeNode{Data: "6"}
	root.Left.Left = &TreeNode{Data: "1"}
	root.Left.Right = &TreeNode{Data: "3"}

	levelCount := BTreeLevelCount(root)
	fmt.Println("Number of levels in the tree:", levelCount)
}
*/
