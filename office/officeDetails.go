package office

import (
	"fmt"
)

type OfficeDetails struct {
	EmpName        string
	EmpId          int
	EmpSalary      int
	empbankdetails string
}

func EmpDetails(office OfficeDetails) {
	fmt.Println(office.EmpSalary)
}
