package problem9

type Node struct {
	value    int
	children []Node
}

func NewNode(val int, children []Node) Node {
	return Node{
		val,
		children,
	}
}

func TreeSum(root *Node, sum int) int {
	if root == nil {
		return 0
	}

	for _, child := range root.children {
		sum = TreeSum(&child, sum)
	}
	return sum + root.value
}
