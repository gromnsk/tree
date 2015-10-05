package tree

import "fmt"

var limit = 3 // maximum amount of childs for each nodes

type Tree struct {
	root *Node
}

type Node struct {
	parent *Node 
	child []*Node
	data Data
}

type Data struct {
	Id int
}

type Result struct {
	Parent int
	Child []int
	Data Data
}

func addNode (parent *Node, data Data) *Node {
	return &Node{parent, make([]*Node, limit), data}
}

func (tree *Tree) Insert(data Data) (node *Node) {
	if tree.root == nil {
		node = addNode(nil, data)
		tree.root = node

		return
	}

	return
}

func (node *Node) findFreeNode() *Node {
	stack := make([]*Node, 0, 32)
	stack = append(stack, node) // push root to stack
	var n *Node = nil

	for stack != nil {
		n, stack = stack[0], stack[1:] //pop first node

		for index := 0; index < len(n.child); index++ {
			if (n.child[index] == nil) {
				return n
			}
			stack = append(stack, n.child[index]) // push
		}
	}

	return nil
}

func (node *Node) Insert(data Data) (newNode *Node) {
	for i := 0; i< len(node.child); i++ {
		if (node.child[i] == nil) {
			newNode = addNode(node, data)
			node.child[i] = newNode

			return
		} 
	}

	freeNode := node.findFreeNode()
	newNode = freeNode.Insert(data)

	return
}

func (node *Node) Count(c chan int) (counter int) {
	if node == nil {
		return
	}

	c <- 1
	for i := 0; i < len(node.child); i++ {
		node.child[i].Count(c)
	}

	return 
}

func (node *Node) Print() {
	if node == nil {
		return
	}

	fmt.Print(node.data)
	fmt.Print("\n")
	for i := 0; i < len(node.child); i++ {
		node.child[i].Print()
	}
}

func (node *Node) Search(id int) (n *Node) {
	stack := make([]*Node, 0, 32)
	stack = append(stack, node) // push root to stack
	if (node.data.Id == id) {
		return node
	}

	for stack != nil {
		n, stack = stack[0], stack[1:] //shift first node

		for index := 0; index < len(n.child); index++ {
			if (n.child[index] == nil) {
				continue
			}
			if (n.child[index].data.Id == id) {
				return n.child[index]
			}
			stack = append(stack, n.child[index]) // push
		}
	}

	return nil
}

func (node *Node) GetData() (result Result) {
	var childs []int = make([]int, 0)
	for _, value := range node.child {
		if value != nil {
			childs = append(childs, value.data.Id)
		} else {
			childs = append(childs, 0)
		}
	}

	var parentNode int = 0

	if node.parent != nil {
		parentNode = node.parent.data.Id
	}

	return Result{
		Parent: parentNode,
		Child: childs,
		Data: node.data,
	}
}