package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func getAndRemoveLastElement(stack *stack.Stack) int {
	result := stack.Pop().(int)
	if stack.Len() == 0 {
		return result
	} else {
		last := getAndRemoveLastElement(stack)
		stack.Push(result)
		return last
	}
}

func reverse(stack *stack.Stack) {
	if stack.Len() == 0 {
		return
	}
	i := getAndRemoveLastElement(stack)
	reverse(stack)
	stack.Push(i)
}

func main() {
	var stackTest = stack.New()
	stackTest.Push(1)
	stackTest.Push(2)
	stackTest.Push(3)
	reverse(stackTest)
	pop := stackTest.Pop().(int)
	fmt.Printf("one: %v\n", pop)

	pop = stackTest.Pop().(int)
	fmt.Printf("two: %v\n", pop)

	pop = stackTest.Pop().(int)
	fmt.Printf("three: %v\n", pop)
}
