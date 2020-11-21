package tree

type Node struct {
	char      Char
	leftNode  *Node
	rightNode *Node
}

type Tree struct {
	root *Tree
}

func (t *Tree) Insert() {

}
