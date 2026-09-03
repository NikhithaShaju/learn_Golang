package queue

import (
	"fmt"
)

type Queue struct {
	Items [5]string
	Size  int
}

func (qu *Queue) Enqueue(element string) {
	if qu.IsFull() {
		fmt.Println("queue is full")
	}
	qu.Items[qu.Size] = element
	qu.Size += 1
	fmt.Println("element is added", element)

}
func (qu *Queue) Dequeue() {
	if qu.IsEmpty() {
		fmt.Println("queue is empty")
	}
	fmt.Println("element is removed", qu.Items[0])
	for i := 0; i < qu.Size-1; i++ {
		qu.Items[i] = qu.Items[i+1]
	}
	qu.Items[qu.Size-1] = ""
	qu.Size -= 1

}
func (qu *Queue) IsFull() bool {
	if qu.Size == len(qu.Items) {
		fmt.Println("queue is full")
		return true
	}
	return false

}
func (qu *Queue) IsEmpty() bool {
	if qu.Size == 0 {
		fmt.Println("queue is empty")
		return true
	}
	return false

}
// func main() {
// 	queue1 := queue.Queue{
// 		Items: [5]string{},
// 		Size:  0,
// 	}
// 	queue1.Enqueue("nikhitha")
// 	queue1.Enqueue("nithya")
// 	queue1.Enqueue("angel")
// 	queue1.Dequeue()
// 	queue1.Dequeue()
// 	fmt.Println(queue1)

// }
