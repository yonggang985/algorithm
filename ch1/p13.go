package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func sortStackByStack(stackSrc *stack.Stack) {
	stackHelp := stack.New()
	for stackSrc.Len() != 0 {
		cur := stackSrc.Pop().(int)
		for stackHelp.Len() != 0 && stackHelp.Peek().(int) < cur {
			stackSrc.Push(stackHelp.Pop().(int))
		}
		stackHelp.Push(cur)
	}
	for stackHelp.Len() != 0 {
		stackSrc.Push(stackHelp.Pop().(int))
	}
}

func main() {
	stackTest := stack.New()
	stackTest.Push(9)
	stackTest.Push(6)
	stackTest.Push(1)
	stackTest.Push(10)
	stackTest.Push(8)
	stackTest.Push(7)
	stackTest.Push(4)
	stackTest.Push(5)
	stackTest.Push(2)
	stackTest.Push(3)

	sortStackByStack(stackTest)
	for stackTest.Len() != 0 {
		pop := stackTest.Pop()
		fmt.Printf("%v ", pop)
	}
	fmt.Println()
}
