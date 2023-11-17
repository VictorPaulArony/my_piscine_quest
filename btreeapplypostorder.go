package piscine

/*type TreeNode struct {
	Left, Right *TreeNode
	Data        interface{}
}

func BTreeInsertData(root *TreeNode, data interface{}) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data.(string) < root.Data.(string) {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}

	return root
}
*/
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyPostorder(root.Left, f)
		BTreeApplyPostorder(root.Right, f)
		f(root.Data)
	}
}

/*
func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	BTreeApplyPostorder(root, fmt.Println)
}
*/
