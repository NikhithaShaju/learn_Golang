package main

import (
	"project/stack"
)

func main() {
	stack1 := stack.StackArray{
		Items: [5]string{},
		Size:  0,
	}
	stack1.Push("nikhitha")
	stack1.Push("nithya")
	stack1.Push("angel")
	stack1.Peek()
	stack1.Pop()
	stack1.Peek()

}

// type StudentJ struct {
// 	name      string
// 	id        int
// 	age       int
// 	chemistry int
// 	maths     int
// }
// type StudentS struct {
// 	name      string
// 	id        int
// 	age       int
// 	chemistry int
// 	maths     int
// }

// func (studDetails *StudentJ) markCalculation() int {
// 	total := studDetails.chemistry + studDetails.maths
// 	return total
// }
// func (studDetails *StudentS) markCalculation() int {
// 	total := studDetails.chemistry + studDetails.maths
// 	studDetails.chemistry = studDetails.chemistry + 5
// 	return total
// }

// func main() {
// 	Student1 := StudentJ{
// 		name:      "nikhitha",
// 		id:        2,
// 		age:       23,
// 		chemistry: 10,
// 		maths:     15,
// 	}
// 	Student2 := StudentS{
// 		name:      "angel",
// 		id:        8,
// 		age:       22,
// 		chemistry: 20,
// 		maths:     10,
// 	}
// 	totalMarks1 := Student1.markCalculation()
// 	fmt.Println(totalMarks1)
// 	totalMarks2 := Student2.markCalculation()
// 	fmt.Println(totalMarks2)
// 	fmt.Println(Student2.chemistry)

// }

// var p *int
// fmt.Println(p)
// a := 20
// p = &a
// *p = 30
// fmt.Println("address", p)
// fmt.Println("value", *p)
// u1 := User{Name: "alan", Age: 24}
// rename(u1)
// reAge(u1)
// fmt.Println(u1)

// var p *int
// fmt.Println(p)
// a := 5
// p = &a
// fmt.Println("address", p)
// fmt.Println("value", *p)

// type User struct {
// 	Name string
// 	Age  int
// }

// func rename(U User) {
// 	U.Name = "jack"
// }
// func reAge(U User) {
// 	U.Age = 28
// }

// a := []int{5, 2, 8, 1, 3}
// 	for i := 0; i < len(a)-1; i++ {
// 		lowest := i
// 		for j := i + 1; j < len(a); j++ {
// 			if a[j] < a[lowest] {
// 				lowest = j
// 			}

// 		}
// 		a[i], a[lowest] = a[lowest], a[i]

// 	}
// 	fmt.Println(a)

// a := []int{5, 2, 8, 1, 3}
// for j := 0; j < len(a); j++ {
// 	for i := 0; i < len(a)-1-j; i++ {
// 		if a[i] > a[i+1] {
// 			a[i], a[i+1] = a[i+1], a[i]

// 		Println()

// 	}

// }
// fmt.Println(a)

// a := 556754
// 	count := 0

// 	for a > 0 {
// 		a = a / 10
// 		count = count + 1

// 	}
// 	fmt.Println(count)

// a := 523
// sum := 0
// for a > 0 {
// 	lastnum := a % 10
// 	sum = sum + lastnum
// 	a = a / 10
// }
// fmt.Println(sum)

// a := 5
// for i := 1; i <= 10; i++ {
// 	fmt.Println(a * i)

// }

// for i := 0; i <= 20; i++ {
// 	if i%2 == 1 {
// 		fmt.Println(i)
// 	}
// }

// for i := 0; i <= 20; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}

// s := "jackson"
// a := ""
// for i := len(s) - 1; i >= 0; i-- {
// 	a = a + string(s[i])
// }
// fmt.Println(s)
// fmt.Println(a)
