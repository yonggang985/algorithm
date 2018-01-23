package main

import (
	"errors"
	"fmt"
	"github.com/golang-collections/collections/stack"
)

var stackData = stack.New()
var stackMin = stack.New()

//push elem from stack
func push(newNum int) {
	if stackData.Len() == 0 {
		stackMin.Push(newNum)
	} else if newNum <= stackMin.Peek().(int) {
		stackMin.Push(newNum)
	}
	stackData.Push(newNum)
}

//pop elem from stack
func pop() (int, error) {
	if stackData.Len() == 0 {
		return 0, errors.New("pop fail: stack is empty!")
	}
	var value int = stackData.Pop().(int)
	minValue, err := getMin()
	if err != nil {
		return 0, err
	}
	if value == minValue {
		stackMin.Pop()
	}
	return value, nil
}

func getMin() (int, error) {
	if stackMin.Len() == 0 {
		return 0, errors.New("getMin fail: stack is empty!")
	}
	return stackMin.Peek().(int), nil

}

func main() {
	fmt.Printf("begin push\n")
	push(2)
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())
	push(1)
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())
	push(3)
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())
	push(4)
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())

	fmt.Printf("begin getMin\n")
	min, err := getMin()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("min value: %v\n\n", min)

	fmt.Printf("begin push\n")
	push(1)
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())
	push(4)
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())

	fmt.Printf("begin pop\n")
	pop()
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())
	pop()
	fmt.Printf("min: %+v,lenght: %+v,data: %+v,length: %v\n\n", stackMin.Peek().(int), stackMin.Len(), stackData.Peek().(int), stackData.Len())

	fmt.Printf("begin getMin\n")
	min, err = getMin()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("min value: %v\n\n", min)

}
