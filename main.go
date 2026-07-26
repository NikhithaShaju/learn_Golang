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
	lengthOfName("nithya")
	lengthOfName("anav")

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
