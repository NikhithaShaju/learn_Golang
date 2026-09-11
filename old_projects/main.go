package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var students []string
var allDetails map[string]studentDetails

type studentDetails struct {
	Address string
	Age     int
	Mark    int
}

type StudentRequest struct {
	StudentName string `json:"sName"`
	param2      string `json:"kjefk"` // Matches the JSON key case
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {
	fmt.Println("server started")
	students = []string{"jack", "alan", "nikitha", "angel", "dixon"}
	allDetails = map[string]studentDetails{
		"jack": {
			Address: "tttt",
			Age:     23,
			Mark:    25,
		},
		"alan": {
			Address: "kkk",
			Age:     21,
			Mark:    27,
		},
		"nikitha": {
			Address: "mmm",
			Age:     22,
			Mark:    28,
		},
		"angel": {
			Address: "llll",
			Age:     20,
			Mark:    29,
		},
		"dixon": {
			Address: "uuuu",
			Age:     26,
			Mark:    20,
		},
	}

	http.HandleFunc("/studentlist", studentListHandler)
	http.HandleFunc("/studentDetailsURL", studDetailsHandlerURL)
	http.HandleFunc("/studentDetailsBody", studDetailsHandlerBody)

	http.ListenAndServe(":8000", nil)

}
func studDetailsHandlerURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request recieved")
	fmt.Println("processing")
	enableCors(w)
	queryParams := r.URL.Query()
	studentName := queryParams.Get("studentName")
	fmt.Println(studentName)
	fmt.Println(allDetails[studentName])
	selectedStudent := allDetails[studentName]
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(selectedStudent); err != nil {
		// 3. Writing Error Responses
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("successfuly processed")

}

func studDetailsHandlerBody(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request recieved")
	fmt.Println("processing")

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	defer r.Body.Close()

	name := StudentRequest{}
	// Decode directly from the stream
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		fmt.Println("some error happened")
	}

	fmt.Println(allDetails[name.StudentName])
	selectedStudent := allDetails[name.StudentName]
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(selectedStudent); err != nil {
		// 3. Writing Error Responses
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("successfuly processed")

}
func studentListHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	fmt.Println("request recieved")

	fmt.Println("processing")
	// Stream the JSON encoding straight into the ResponseWriter
	if err := json.NewEncoder(w).Encode(students); err != nil {
		// 3. Writing Error Responses
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println("successfuly processed")

}

// package main

// import "fmt"

// type LinkedList struct {
// 	Data string
// 	Next *LinkedList
// }

// func (L *LinkedList) Append(nodeData string) {
// 	newNode := createNode(nodeData)

// 	fmt.Println(L)
// 	for L.Next != nil {
// 		L = L.Next
// 	}
// 	L.Next = &newNode
// 	fmt.Println(L)

// }

// func (L *LinkedList) Prepend(nodeData string) {
// 	newNode := createNode(nodeData)

// 	newNode.Next = L.Next
// 	L.Next = &newNode

// 	L.Data, newNode.Data = newNode.Data, L.Data
// }
// func createNode(nodeData string) LinkedList {
// 	newNode := LinkedList{
// 		Data: nodeData,
// 		Next: nil,
// 	}
// 	return newNode

// }
// func (L *LinkedList) Traverse() int {
// 	count := 0
// 	for L != nil {
// 		fmt.Println("index: ", count, L.Data)
// 		count++
// 		L = L.Next
// 	}
// 	fmt.Println("\n")
// 	return count
// }
// func (L *LinkedList) InsertAt(nodeData string, index int) {
// 	if index == 0 {
// 		L.Prepend(nodeData)
// 		return
// 	}
// 	if L.Traverse() == index {
// 		L.Append(nodeData)
// 		return
// 	}
// 	newNode := createNode(nodeData)
// 	for i := 0; i < index-1; i++ {
// 		L = L.Next
// 	}
// 	temp := L.Next
// 	L.Next = &newNode
// 	newNode.Next = temp
// }

// // ["jack", ] ->

// func main() {
// 	L3 := LinkedList{
// 		Data: "alan",
// 		Next: nil,
// 	}
// 	L2 := LinkedList{
// 		Data: "dixon",
// 		Next: &L3,
// 	}
// 	L1 := LinkedList{
// 		Data: "jack",
// 		Next: &L2,
// 	}

// 	// fmt.Println(L1)
// 	L1.Append("raju")
// 	L1.Prepend("ramu")
// 	L1.InsertAt("kichu", 2)

// 	count := L1.Traverse()
// 	fmt.Println(count)

// }

// stack1 := stack.StackArray{
// 	Items: [5]string{},
// 	Size:  0,
// }
// stack1.Push("nikhitha")
// stack1.Push("nithya")
// stack1.Push("angel")
// stack1.Peek()
// stack1.Pop()
// stack1.Peek()

// type StudentJ struct {
// 	name      string
// 	id        int
// 	age       int
// 	chemistry int
// 	maths     int
// }
// type StudentS struct {
// 	name      string
// 	id        int
// 	age       int
// 	chemistry int
// 	maths     int
// }

// func (studDetails *StudentJ) markCalculation() int {
// 	total := studDetails.chemistry + studDetails.maths
// 	return total
// }
// func (studDetails *StudentS) markCalculation() int {
// 	total := studDetails.chemistry + studDetails.maths
// 	studDetails.chemistry = studDetails.chemistry + 5
// 	return total
// }

// func main() {
// 	Student1 := StudentJ{
// 		name:      "nikhitha",
// 		id:        2,
// 		age:       23,
// 		chemistry: 10,
// 		maths:     15,
// 	}
// 	Student2 := StudentS{
// 		name:      "angel",
// 		id:        8,
// 		age:       22,
// 		chemistry: 20,
// 		maths:     10,
// 	}
// 	totalMarks1 := Student1.markCalculation()
// 	fmt.Println(totalMarks1)
// 	totalMarks2 := Student2.markCalculation()
// 	fmt.Println(totalMarks2)
// 	fmt.Println(Student2.chemistry)

// }

// var p *int
// fmt.Println(p)
// a := 20
// p = &a
// *p = 30
// fmt.Println("address", p)
// fmt.Println("value", *p)
// u1 := User{Name: "alan", Age: 24}
// rename(u1)
// reAge(u1)
// fmt.Println(u1)

// var p *int
// fmt.Println(p)
// a := 5
// p = &a
// fmt.Println("address", p)
// fmt.Println("value", *p)

// type User struct {
// 	Name string
// 	Age  int
// }

// func rename(U User) {
// 	U.Name = "jack"
// }
// func reAge(U User) {
// 	U.Age = 28
// }

// a := []int{5, 2, 8, 1, 3}
// 	for i := 0; i < len(a)-1; i++ {
// 		lowest := i
// 		for j := i + 1; j < len(a); j++ {
// 			if a[j] < a[lowest] {
// 				lowest = j
// 			}

// 		}
// 		a[i], a[lowest] = a[lowest], a[i]

// 	}
// 	fmt.Println(a)

// a := []int{5, 2, 8, 1, 3}
// for j := 0; j < len(a); j++ {
// 	for i := 0; i < len(a)-1-j; i++ {
// 		if a[i] > a[i+1] {
// 			a[i], a[i+1] = a[i+1], a[i]

// 		Println()

// 	}

// }
// fmt.Println(a)

// a := 556754
// 	count := 0

// 	for a > 0 {
// 		a = a / 10
// 		count = count + 1

// 	}
// 	fmt.Println(count)

// a := 523
// sum := 0
// for a > 0 {
// 	lastnum := a % 10
// 	sum = sum + lastnum
// 	a = a / 10
// }
// fmt.Println(sum)

// a := 5
// for i := 1; i <= 10; i++ {
// 	fmt.Println(a * i)

// }

// for i := 0; i <= 20; i++ {
// 	if i%2 == 1 {
// 		fmt.Println(i)
// 	}
// }

// for i := 0; i <= 20; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 		}
// 	}

// s := "jackson"
// a := ""
// for i := len(s) - 1; i >= 0; i-- {
// 	a = a + string(s[i])
// }
// fmt.Println(s)
// fmt.Println(a)
