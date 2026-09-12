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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
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
