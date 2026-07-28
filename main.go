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
	// result, output2 := calculator(10, 5, "a")
	// fmt.Println(result)
	// fmt.Println(output2)
	// calculator(10, 5, "*")
	// result := login("admin", "golang123")
	// if result == true {
	// 	fmt.Println("login credentials are correct")
	// } else {

	// 	fmt.Println("login credentials are wrong")
	// }
	// first, second := divide(10, 2)
	// fmt.Println("your value is:", first)
	// fmt.Println("your remainder is:", second)

	// a := 5
	// b := 6
	// operation := "*"
	// parameters1 := calculatorInputs{
	// 	firstValue:  a,
	// 	secondValue: b,
	// 	operator:    operation,
	// }
	// parameters2 := calculatorInputs{
	// 	firstValue:  b,
	// 	secondValue: a,
	// 	operator:    "/",
	// 	}
	// 	v1, v2 := calculator2(parameters1)
	// 	fmt.Println(v1, v2)
	// 	v3, v4 := calculator2(parameters2)
	// 	fmt.Println(v3, v4)
	// w := "The Hobbit"
	// x := "J.R.R Tolkien"
	// y := 305
	// z := false
	// bookdetails1 := Book{
	// 	firstchar:   w,
	// 	secondchar:  x,
	// 	thirdvalue:  y,
	// 	fourthvalue: z,
	// }
	// x1 := library(bookdetails1)
	// fmt.Println(x1.fourthvalue)
	p := "Jackson"
	q := 28
	r := 10
	s := true

	studentdetails1 := student{
		studentName:  p,
		studentAge:   q,
		studentClass: r,
		leaderStatus: s,
	}
	a1 := checkLeader(studentdetails1)
	fmt.Println(a1)

}

type student struct {
	studentName  string
	studentAge   int
	studentClass int
	leaderStatus bool
}

func checkLeader(allchar student) bool {
	if allchar.studentName == "Jackson" {
		fmt.Println(allchar.studentName, "is a leader")
	} else {
		fmt.Println(allchar.studentName, "is not a leader")
		allchar.leaderStatus = false
	}
	return allchar.leaderStatus

}

type Book struct {
	firstchar   string
	secondchar  string
	thirdvalue  int
	fourthvalue bool
}

func library(allvalues Book) Book {
	fmt.Println("the", allvalues.firstchar, "was written by", allvalues.secondchar)
	fmt.Println(allvalues.fourthvalue)
	allvalues.fourthvalue = true
	fmt.Println(allvalues.fourthvalue)
	return allvalues
}

type calculatorInputs struct {
	firstValue  int
	secondValue int
	operator    string
}

func calculator2(allInputs calculatorInputs) (int, string) {

	if allInputs.operator == "+" {
		// fmt.Println(a + b)
		return allInputs.firstValue + allInputs.secondValue, "fff"
	}
	if allInputs.operator == "-" {
		// fmt.Println(a - b)
		return allInputs.firstValue - allInputs.secondValue, "ddd"
	}
	if allInputs.operator == "*" {
		// fmt.Println(a * b)
		return allInputs.firstValue * allInputs.secondValue, "ggg"
	}
	if allInputs.operator == "/" {
		// fmt.Println(a / b)
		return allInputs.firstValue / allInputs.secondValue, "sss"
	}

	return 0, "this operation doesnt exist"
}
func divide(a int, b int) (int, int) {
	value := a / b
	remainder := a % b
	return value, remainder

}

func login(userName string, password string) bool {

	if userName == "admin" && password == "golang123" {
		// fmt.Println("true")
		return true
	} else {
		// fmt.Println("false")
		return false

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
