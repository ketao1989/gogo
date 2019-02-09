package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/ketao1989/gogo/tools"
)

func main() {

	flag.Parse()
	defer glog.Flush()

	var tree = tools.BuildTree()
	fmt.Println(tree.GetVal(), " ", tree.GetLeft().GetVal(), " ", tree.GetRight().GetVal())

	tree.PrintTree()

}
