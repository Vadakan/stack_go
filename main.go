package main

import "fmt"

var stack []int

func crateemptystack() []int {
	return stack
}

func main() {

	empty_stack := crateemptystack()

	fmt.Println("empty stack  :", empty_stack)

}
