package funcvarstruct

import "fmt"

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
