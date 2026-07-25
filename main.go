package main

import "fmt"

func main() {
	fmt.Println("hello")
	printMyName("nikhitha", 22, 12)

}
func printMyName(name string, age int, class int) {
	fmt.Println("name=",name)
	fmt.Println("age=",age)
	fmt.Println("class=",class)
}
