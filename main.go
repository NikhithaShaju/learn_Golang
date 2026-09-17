package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

// var students []string
// var allDetails map[string]studentDetails
var conn *pgx.Conn
var ctx = context.Background()

type studentDetails struct {
	Address string
	Age     int
	Mark    int
}

type StudentRequest struct {
	StudentName string `json:"sName"`
	// param2      string `json:"kjefk"` // Matches the JSON key case
}
type editStudentDetails struct {
	StudentName string         `json:"sName"`
	Details     studentDetails `json:"details"`
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {

	fmt.Println("server started")

	connStr := "postgres://postgres:nikhitha%402003@localhost:5432/postgres"

	var err error

	conn, err = pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Unable to connect to database: %v\n",
			err,
		)
		os.Exit(1)
	}

	defer conn.Close(ctx)

	var greeting string

	err = conn.QueryRow(
		ctx,
		"SELECT 'Connection successful!'",
	).Scan(&greeting)

	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Query failed: %v\n",
			err,
		)
		os.Exit(1)
	}

	fmt.Println(greeting)

	http.HandleFunc("/studentlist", studentListHandler)
	http.HandleFunc("/studentDetailsURL", studDetailsHandlerURL)
	http.HandleFunc("/studentDetailsBody", studDetailsHandlerBody)
	http.HandleFunc("/editstudentdetails", editStudentDetailsHandler)

	fmt.Println("Server running on http://localhost:8000")

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

	// fmt.Println("server started")
	// students = []string{"jack", "alan", "nikitha", "angel", "dixon"}
	// allDetails = map[string]studentDetails{
	// 	"jack": {
	// 		Address: "tttt",
	// 		Age:     23,
	// 		Mark:    25,
	// 	},
	// 	"alan": {
	// 		Address: "kkk",
	// 		Age:     21,
	// 		Mark:    27,
	// 	},
	// 	"nikitha": {
	// 		Address: "mmm",
	// 		Age:     22,
	// 		Mark:    28,
	// 	},
	// 	"angel": {
	// 		Address: "llll",
	// 		Age:     20,
	// 		Mark:    29,
	// 	},
	// 	"dixon": {
	// 		Address: "uuuu",
	// 		Age:     26,
	// 		Mark:    20,
	// 	},
	// }

}
func studDetailsHandlerURL(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("request received")
	fmt.Println("processing")

	queryParams := r.URL.Query()

	studentName := queryParams.Get("studentName")

	fmt.Println("Selected student:", studentName)

	var selectedStudent studentDetails

	err := conn.QueryRow(
		ctx,
		`SELECT stdaddress, stdage, stdmark
		 FROM stdDetails
		 WHERE stdname = $1`,
		studentName,
	).Scan(
		&selectedStudent.Address,
		&selectedStudent.Age,
		&selectedStudent.Mark,
	)

	if err != nil {
		fmt.Println("Database error:", err)

		http.Error(
			w,
			"Student not found",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(selectedStudent); err != nil {
		http.Error(
			w,
			"Internal Server Error",
			http.StatusInternalServerError,
		)

		return
	}

	fmt.Println("successfully processed")
}

func studDetailsHandlerBody(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("request received")
	fmt.Println("processing")

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	defer r.Body.Close()

	var name StudentRequest

	err := json.NewDecoder(r.Body).Decode(&name)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Selected student:", name.StudentName)

	var selectedStudent studentDetails

	err = conn.QueryRow(
		ctx,
		`SELECT stdaddress, stdage, stdmark
		 FROM stdDetails
		 WHERE stdname = $1`,
		name.StudentName,
	).Scan(
		&selectedStudent.Address,
		&selectedStudent.Age,
		&selectedStudent.Mark,
	)

	if err != nil {
		fmt.Println("Database error:", err)

		http.Error(
			w,
			"Student not found",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(selectedStudent); err != nil {
		http.Error(
			w,
			"Internal Server Error",
			http.StatusInternalServerError,
		)

		return
	}

	fmt.Println("successfully processed")
}

func studentListHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println("request received")
	fmt.Println("processing")

	rows, err := conn.Query(
		ctx,
		"SELECT stdname FROM stdDetails",
	)

	if err != nil {
		fmt.Println("Database query error:", err)

		http.Error(
			w,
			"Failed to fetch students",
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var studentList []string

	for rows.Next() {
		var studentName string

		err := rows.Scan(&studentName)

		if err != nil {
			http.Error(
				w,
				"Failed to read student",
				http.StatusInternalServerError,
			)

			return
		}

		studentList = append(studentList, studentName)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(studentList); err != nil {
		http.Error(
			w,
			"Internal Server Error",
			http.StatusInternalServerError,
		)

		return
	}

	fmt.Println("successfully processed")
}

func editStudentDetailsHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var request editStudentDetails

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Student name:", request.StudentName)
	fmt.Println("Student details:", request.Details)

	result, err := conn.Exec(
		ctx,
		`UPDATE stdDetails
		 SET stdaddress = $1,
		     stdage = $2,
		     stdmark = $3
		 WHERE stdname = $4`,
		request.Details.Address,
		request.Details.Age,
		request.Details.Mark,
		request.StudentName,
	)

	if err != nil {
		fmt.Println("Database update error:", err)

		http.Error(
			w,
			"Failed to update student details",
			http.StatusInternalServerError,
		)

		return
	}

	if result.RowsAffected() == 0 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Student details updated successfully",
	})
}
