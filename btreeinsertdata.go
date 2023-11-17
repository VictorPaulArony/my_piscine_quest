package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data <= root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}

	return root
}

/*
func main() {
	var root *TreeNode
	root = BTreeInsertData(root, "5")
	root = BTreeInsertData(root, "3")
	root = BTreeInsertData(root, "7")
	root = BTreeInsertData(root, "2")
	root = BTreeInsertData(root, "4")
	root = BTreeInsertData(root, "6")
	root = BTreeInsertData(root, "8")

	// Print the tree in-order
	PrintInOrder(root)
}

func PrintInOrder(node *TreeNode) {
	if node != nil {
		PrintInOrder(node.Left)
		fmt.Println(node.Data)
		PrintInOrder(node.Right)
	}
}
*/
