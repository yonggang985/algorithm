package main

import (
	"fmt"
)

func hanoiProblem(num int, left, mid, right string) {
	if num < 1 {
		return
	}
	if num == 1 {
		fmt.Printf("move %v from  %v  to  %v\n", num, left, right)
	} else {
		hanoiProblem(num-1, left, right, mid)
		fmt.Printf("move %v from  %v  to  %v\n", num, left, right)
		hanoiProblem(num-1, mid, left, right)
	}
}

func main() {
	hanoiProblem(20, "L", "M", "R")
}
