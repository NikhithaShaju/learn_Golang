package main

import "fmt"

// "project/school"
func main() {
	a := "ambma"
	for i, j := 0, len(a)-1; i <= j; i, j = i+1, j-1 {
		fmt.Println(i, j)
		if a[i] != a[j] {
			fmt.Println("its not a palindrome")
			return

		}

	}
	fmt.Println("its a palindrome")

}

// a := "ambma"
// le := len(a) - 1
// for i := 0; i < len(a); i++ {
// 	if a[i] != a[le] {
// 		fmt.Println("its not a palindrome")
// 		return
// 	} else {
// 		fmt.Println("its a palindrome")
// 	}
// }

// 	a := []int{4, 5, 1, 6, 3, 7, 2}
// 	for j := 0; j < len(a); j++ {
// 		for i := 0; i < len(a)-1-j; i++ {
// 			if a[i] > a[i+1] {
// 				a[i], a[i+1] = a[i+1], a[i]

// 			}

// 		}

// 	}

// 	fmt.Println(a)

// }

// func main() {
// 	emp1 := office.OfficeDetails{
// 		EmpName:   "anu",
// 		EmpId:     405,
// 		EmpSalary: 50000,
// 	}
// 	fmt.Println(emp1)
// 	office.EmpDetails(emp1)
// }

// lion := animal.AnimalDetailStruct{
// 	Name:       "lion",
// 	Color:      "red",
// 	IsPoisones: true,
// }
// fmt.Println(lion)
// animal.PrintAnimalName(lion)

// }
//
// 	student1 := school.StudentDetails{
// 		Name: "Nikhitha",
// 		Id:   505,
// 	}
// 	fmt.Println(student1)
// 	school.PrintSchoolDetails(student1)
// }
