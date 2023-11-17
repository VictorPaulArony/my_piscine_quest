package piscine

/*S
type TreeNode struct {
	Left, Right *TreeNode
	Data        interface{}
}

func BTreeInsertData(root *TreeNode, data interface{}) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}

	return root
}*/

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		BTreeApplyPreorder(root.Left, f)
		BTreeApplyPreorder(root.Right, f)
	}
}

/*
func main() {
	root := &TreeNode{Data: "4"}
	root = BTreeInsertData(root, "1")
	root = BTreeInsertData(root, "7")
	root = BTreeInsertData(root, "5")

	BTreeApplyPreorder(root, func(data interface{}) {
		fmt.Println(data)
	})
}
*/
