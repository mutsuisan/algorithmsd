package main

import (
	"fmt"
	"github.com/mutsuisan/algorithmsd/structures"
)

func main() {
	hash := structures.Init(7)
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
	err = hash.Insert("hoge")
	if err != nil {
		fmt.Println(err)
	}
	err = hash.Delete("hoge")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hash.GetValues())
}
