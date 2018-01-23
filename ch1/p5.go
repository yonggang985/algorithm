package main

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/stack"
)

var stackPush = stack.New()
var stackPop = stack.New()

func add(value int) {
	stackPush.Push(value)
}

func poll() (int, error) {
	if stackPush.Len() == 0 && stackPop.Len() == 0 {
		return 0, errors.New("poll fail: queue is empty!")
	} else if stackPop.Len() == 0 {
		for stackPush.Len() != 0 {
			stackPop.Push(stackPush.Pop().(int))
		}
	}
	return stackPop.Pop().(int), nil
}

func peek() (int, error) {
	if stackPush.Len() == 0 && stackPop.Len() == 0 {
		return 0, errors.New("peek fail: queue is empty!")
	} else if stackPop.Len() == 0 {
		for stackPush.Len() != 0 {
			stackPop.Push(stackPush.Pop().(int))
		}
	}
	return stackPop.Peek().(int), nil
}

func main() {
	fmt.Printf("begin add\n")
	add(2)
	fmt.Printf("lenght: %+v,push: %+v,length: %v\n", stackPop.Len(), stackPush.Peek().(int), stackPush.Len())
	add(1)
	fmt.Printf("lenght: %+v,push: %+v,length: %v\n", stackPop.Len(), stackPush.Peek().(int), stackPush.Len())
	add(3)
	fmt.Printf("lenght: %+v,push: %+v,length: %v\n", stackPop.Len(), stackPush.Peek().(int), stackPush.Len())
	add(4)
	fmt.Printf("lenght: %+v,push: %+v,length: %v\n\n", stackPop.Len(), stackPush.Peek().(int), stackPush.Len())

	fmt.Printf("begin peek\n")
	p, err := peek()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("peek value: %v\n\n", p)

	fmt.Printf("begin add\n")
	add(1)
	fmt.Printf("pop: %+v,lenght: %+v,push: %+v,length: %v\n", stackPop.Peek().(int), stackPop.Len(), stackPush.Peek().(int), stackPush.Len())
	add(4)
	fmt.Printf("pop: %+v,lenght: %+v,push: %+v,length: %v\n\n", stackPop.Peek().(int), stackPop.Len(), stackPush.Peek().(int), stackPush.Len())

	fmt.Printf("begin poll\n")
	p, err = poll()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("poll value: %v\n", p)

	fmt.Printf("pop: %+v,lenght: %+v,push: %+v,length: %v\n", stackPop.Peek().(int), stackPop.Len(), stackPush.Peek().(int), stackPush.Len())
	p, err = poll()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("poll value: %v\n", p)

	fmt.Printf("pop: %+v,lenght: %+v,push: %+v,length: %v\n\n", stackPop.Peek().(int), stackPop.Len(), stackPush.Peek().(int), stackPush.Len())

	fmt.Printf("begin peek\n")
	p, err = peek()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("peek value: %v\n\n", p)

}
