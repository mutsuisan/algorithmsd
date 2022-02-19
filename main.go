package main

import (
	"fmt"

	"github.com/mutsuisan/algorithmsd/structures"
)

func binarySearchTree() *structures.BinarySearchTree {
	bTree := structures.InitBinarySearchTree(7)
	return bTree
}

func hashTable() {
	hash := structures.InitHashTable(7)
	fmt.Println(hash.GetValues())
	err := hash.Insert("foobar")
	fmt.Println(hash.GetValues())
	if err != nil {
		fmt.Println(err)
	}
	err = hash.Insert("foobar")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash.GetValues())
	err = hash.Insert("hoge")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash.GetValues())
	err = hash.Delete("hoge")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash.GetValues())
}

func main() {
	hashTable()
}
