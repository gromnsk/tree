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

var treeIndex map[int]*Node

func addNode (parent *Node, data Data) *Node {
	return &Node{parent, make([]*Node, limit), data}
}

func (tree *Tree) Insert(data Data) (node *Node) {
	treeIndex = make(map[int]*Node)
	if tree.root == nil {
		node = addNode(nil, data)
		treeIndex[data.Id] = node
		tree.root = node

		return
	}

	return
}

func (node *Node) SetAllNodes(results []*Result) {
	var newNode, currentNode *Node
	for _, result := range results {
		if result.Parent != 0 {
			currentNode = node.Search(result.Parent)
		} else {
			currentNode = node
		}
		newNode = currentNode.Insert(result.Data)
		treeIndex[result.Data.Id] = newNode
	}
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
			treeIndex[data.Id] = newNode
			node.child[i] = newNode

			return
		} 
	}

	freeNode := node.findFreeNode()
	newNode = freeNode.Insert(data)

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
	
	n = treeIndex[id]
	if n == nil {
		return nil
	}

	return
}

func (node *Node) GetData() (result *Result) {
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

	return &Result{
		Parent: parentNode,
		Child: childs,
		Data: node.data,
	}
}

func getLevel(nodes []*Node) (results []*Node) {
	for _, value := range nodes {
		for index := 0; index < len(value.child); index++ {
			if (value.child[index] == nil) {
				continue
			}
			results = append(results, value.child[index])
		}
	}

	return
}

func (node *Node) GetNodes(level int) (results []*Result) {
	nodes := make([]*Node, 0)
	nodes = append(nodes, node)
	results = append(results, node.GetData())

	for index := 0; index < level; index++ {
		nodes = getLevel(nodes)
		for _, value := range nodes {
			results = append(results, value.GetData())
		}
	}

	return
}
