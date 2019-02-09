package main

import (
	"fmt"
	"github.com/ketao1989/gogo/tools"
)

func main() {

	var tree = tools.BuildTree()
	fmt.Println(tree.GetVal(), " ", tree.GetLeft().GetVal(), " ", tree.GetRight().GetVal())

	tree.PrintTree()

}
