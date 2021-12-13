package main

import "fmt"

var stack []int

func crateemptystack() []int {
	return stack
}

func push(inp ...int) {
	stack = append(stack, inp...)
}

func main() {

	empty_stack := crateemptystack()

	fmt.Println("empty stack  :", empty_stack)

}
