package main

import (
	"fmt"
	"project/office"
	// "project/school"
)

func main() {
	emp1 := office.OfficeDetails{
		EmpName:   "anu",
		EmpId:     405,
		EmpSalary: 50000,
	}
	fmt.Println(emp1)
	office.EmpDetails(emp1)
}

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
