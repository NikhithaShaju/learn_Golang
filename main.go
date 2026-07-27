package main

import "fmt"

func main() {
	// fmt.Println("hello")
	// printMyName("nikhitha", 22, 12)
	// printStudentDetails("nithya", 20, 11, 6, 4, 7)
	// calculateAge(2004)
	// calculateAge(2007)
	// studentGrade("anu", 6)
	// studentGrade("riya", 9)
	// studentGrade("liya", 7)
	// oddOrEven(5)
	// oddOrEven(8)
	// lengthOfName("nithya")
	// lengthOfName("anav")
	result, output2 := calculator(10, 5, "a")
	fmt.Println(result)
	fmt.Println(output2)
	// calculator(10, 5, "*")
	// login("admi", "golang12")

}

func login(userName string, password string) {

	if userName == "admin" && password == "golang123" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
func calculator(a int, b int, operation string) (int, string) {

	if operation == "+" {
		// fmt.Println(a + b)
		return a + b, "fff"
	}
	if operation == "-" {
		// fmt.Println(a - b)
		return a - b, "ddd"
	}
	if operation == "*" {
		// fmt.Println(a * b)
		return a * b, "ggg"
	}
	if operation == "/" {
		// fmt.Println(a / b)
		return a / b, "sss"
	}

	return 0, "this operation doesnt exist"
}

func lengthOfName(name string) {
	count := len(name)
	fmt.Println(count)

}

func oddOrEven(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")

	}

}

func studentGrade(name string, mark int) {
	if mark >= 8 {
		fmt.Println("A grade for", name)
	}
	if mark <= 7 && mark > 5 {
		fmt.Println("B grade for", name)
	}
	if mark <= 5 {
		fmt.Println("C grade for", name)
	}

}

func calculateAge(birthYear int) {
	currentYear := 2026
	age := 0
	age = currentYear - birthYear
	fmt.Println("age=", age)
}

func printMyName(name string, age int, class int) {
	fmt.Println("name=", name)
	fmt.Println("age=", age)
	fmt.Println("class=", class)
}

func printStudentDetails(name string, age int, class int, nameCount int, ageCount int, classCount int) {
	for i := 0; i < nameCount; i++ {
		fmt.Println("name=", name)
	}
	for j := 0; j < ageCount; j++ {
		fmt.Println("age=", age)
	}
	for k := 0; k < classCount; k++ {
		fmt.Println("class=", class)
	}
}
