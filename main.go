package main

import (
	"fmt"
)

func usingCompose(a int) (int, error) {
	return Compose(
		Add(2),
		TraceInt,
		Multiply(2),
	)(a)
}

func usingTraverse(a []int) ([]int, error) {
	return Traverse(DoStuffToANumber, a)
}

func main() {
	resultCompose, _ := usingCompose(10)
	fmt.Println("Using compose!", resultCompose)
	resultTraverse, _ := usingTraverse([]int{1, 2, 3})
	fmt.Println("using Traverse (with Compose)!", resultTraverse)
}
