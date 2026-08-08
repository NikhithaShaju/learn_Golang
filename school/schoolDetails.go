package school

import "fmt"

type StudentDetails struct {
	Name    string
	Id      int
	Class   int
	address string
	Mark    int
}

func PrintSchoolDetails(school StudentDetails) {
	fmt.Println(school.Name)

}
