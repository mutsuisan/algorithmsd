package structures

import "fmt"

type BinarySearchTree struct {
	root *node
}

type node struct {
	value int
	left  *node
	right *node
}

func InitBinarySearchTree(size int) *BinarySearchTree {
	bTree := BinarySearchTree{}
	return &bTree
}

func (b BinarySearchTree) Insert(v int) {
	if b.root == nil {
		b.root = &node{value: v}
		fmt.Println("root filled")
	}
	b.root.insert(v)
}

func (n node) insert(v int) {
	fmt.Printf("Inserting %d\n", v)
	if v >= n.value {
		fmt.Println("Inserting RIGHT")
		if n.right == nil {
			fmt.Println("NEW")
			n.right = &node{value: v}
			fmt.Println(n)
			return
		}
		n.right.insert(v)
	} else {
		fmt.Println("Inserting LEFT")
		if n.left == nil {
			fmt.Println("NEW")
			n.left = &node{value: v}
			fmt.Println(n)
			return
		}
		n.left.insert(v)
	}
}

func (b BinarySearchTree) Delete() {}

func (b BinarySearchTree) Search(v int) bool {
	return b.root.search(v)
}

func (n node) search(v int) bool {
	if v > n.value {
		return n.right.search(v)
	} else if n.value < v {
		return n.left.search(v)
	}
	return true
}

func (b BinarySearchTree) Max() int {
	var currentNode *node = b.root
	for currentNode == nil {
		currentNode = currentNode.right
	}
	return currentNode.value
}

func (b BinarySearchTree) Min() int {
	var currentNode *node = b.root
	for currentNode == nil {
		currentNode = currentNode.left
	}
	return currentNode.value
}

func (b BinarySearchTree) PrintValuesDesc() {
	fmt.Println(b.root)
	b.root.printValuesDesc()
}

func (n node) printValuesDesc() {
	if n.right != nil {
		n.right.printValuesDesc()
	}
	if n.left != nil {
		n.left.printValuesDesc()
	}
	fmt.Println(n.value)
}
