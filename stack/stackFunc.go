package stack

import (
	"fmt"
)

type StackArray struct {
	Items [5]string
	Size  int
}

func (arr *StackArray) Push(element string) {
	if arr.IsFull() {
		fmt.Println("the stack is full")
	}
	arr.Items[arr.Size] = element
	arr.Size += 1
	fmt.Println("element is added to array", arr.Items)

}
func (arr *StackArray) Pop() {
	if arr.IsEmpty() {
		fmt.Println("the stack is empty.cannot pop")
	}
	arr.Size -= 1
	fmt.Println("the last element", arr.Items[arr.Size], "is deleted")
	arr.Items[arr.Size] = ""
	// arr.Size -= 1

}
func (arr StackArray) IsEmpty() bool {
	if arr.Size == 0 {
		fmt.Println("the stack is empty")
		return true
	}
	fmt.Println("the stack is not empty")
	return false
}
func (arr StackArray) IsFull() bool {
	if arr.Size == len(arr.Items) {
		fmt.Println("the stack is full")
		return true
	}
	fmt.Println("the stack still have", 5-arr.Size, "space")
	return false
}
func (arr StackArray) Peek() {
	if arr.IsEmpty() {
		fmt.Println("the stack is empty, cannot peek")
	}
	fmt.Println("the top element is", arr.Items[arr.Size-1])
}
