// package main

// import "fmt"

// func main() {
// 	fmt.Println("hello")
// 	printMyName("nikhitha", 22, 12)

// }
//
//	func printMyName(name string, age int, class int) {
//		fmt.Println("name=",name)
//		fmt.Println("age=",age)
//		fmt.Println("class=",class)
//	}
package main

import "fmt"

func main() {
	printStudentDetails("nikhitha", 22, 12)
}
func printStudentDetails(name string, age int, class int) {
	for i := 0; i < 5; i++ {
		fmt.Println("name=", name)
	}
	for j := 0; j < 3; j++ {
		fmt.Println("age=", age)
	}
	for k := 0; k < 4; k++ {
		fmt.Println("class=", class)
	}
}
