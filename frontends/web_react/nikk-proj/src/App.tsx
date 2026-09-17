import { useState } from "react";
import "./App.css";

type StudentDetails = {
  Address: string;
  Age: number;
  Mark: number;
};

function App() {
  const [studentList, setStudentList] = useState<string[]>([]);

  const [selectedStudent, setSelectedStudent] = useState("");

  const [isEditing, setIsEditing] = useState(false);

  const [studentDetail, setStudentDetail] =
    useState<StudentDetails>({
      Address: "",
      Age: 0,
      Mark: 0,
    });

  // Fetch student names
  async function fetchStudents() {
    const url = "http://localhost:8000/studentlist";

    try {
      const response = await fetch(url);

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data: string[] = await response.json();

      setStudentList(data);
    } catch (error) {
      console.error("Failed to fetch students:", error);
    }
  }

  // Fetch selected student's old details
 async function fetchStudentDetails(studentName: string) {
  console.log("Clicked student:", studentName);

  setSelectedStudent(studentName);
  setIsEditing(false);

  const url =
    "http://localhost:8000/studentDetailsURL?studentName=" +
    encodeURIComponent(studentName);

  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const data = await response.json();

    console.log("Old student details:", data);

    setStudentDetail(data);
  } catch (error) {
    console.error("Failed to fetch student details:", error);
  }
}

  // Submit updated details
  async function editStudentDetails() {
    if (selectedStudent === "") {
      alert("Please select a student first");
      return;
    }

    try {
      const response = await fetch(
        "http://localhost:8000/editstudentdetails",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            sName: selectedStudent,
            details: studentDetail,
          }),
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      const data = await response.json();

      alert(data.message || "Student details updated");

      setIsEditing(false);
    } catch (error) {
      console.error("Failed to update student details:", error);
    }
  }

  return (
    <>
      <section id="center">
        <h1>Student List</h1>

        <button
          type="button"
          className="counter"
          onClick={fetchStudents}
        >
          Fetch Students
        </button>

        <div
          style={{
            display: "flex",
            flexDirection: "row",
            padding: 10,
            gap: 20,
          }}
        >
          {/* Student list */}
          <div style={{ paddingRight: 20 }}>
            {studentList.map((student) => (
              <button
                type="button"
                key={student}
                onClick={() => fetchStudentDetails(student)}
                style={{
                  display: "block",
                  width: "120px",
                  padding: "10px",
                  marginBottom: "5px",
                  border: "1px solid white",
                  borderRadius: "10px",
                  background: "transparent",
                  color: "white",
                  cursor: "pointer",
                }}
              >
                {student}
              </button>
            ))}
          </div>

          {/* Student details */}
          <div
            style={{
              borderWidth: "1px",
              borderStyle: "solid",
              borderColor: "white",
              paddingLeft: 25,
              paddingRight: 25,
              paddingBottom: 20,
              borderRadius: 10,
              minWidth: 220,
            }}
          >
            <p style={{ fontSize: 20 }}>
              Student Details
            </p>

            <p>Selected Student: {selectedStudent}</p>

            {isEditing ? (
              <>
                <p>Address:</p>

                <input
                  value={studentDetail.Address}
                  onChange={(e) =>
                    setStudentDetail({
                      ...studentDetail,
                      Address: e.target.value,
                    })
                  }
                />

                <p>Age:</p>

                <input
                  type="number"
                  value={studentDetail.Age}
                  onChange={(e) =>
                    setStudentDetail({
                      ...studentDetail,
                      Age: Number(e.target.value),
                    })
                  }
                />

                <p>Mark:</p>

                <input
                  type="number"
                  value={studentDetail.Mark}
                  onChange={(e) =>
                    setStudentDetail({
                      ...studentDetail,
                      Mark: Number(e.target.value),
                    })
                  }
                />

                <br />
                <br />

                <button onClick={editStudentDetails}>
                  Submit
                </button>

                <button
                  onClick={() => setIsEditing(false)}
                  style={{ marginLeft: 10 }}
                >
                  Cancel
                </button>
              </>
            ) : (
              <>
                <p>Address: {studentDetail.Address}</p>
                <p>Age: {studentDetail.Age}</p>
                <p>Mark: {studentDetail.Mark}</p>

                <button
                  onClick={() => setIsEditing(true)}
                  disabled={selectedStudent === ""}
                >
                  Edit
                </button>
              </>
            )}
          </div>
        </div>
      </section>
    </>
  );
}

export default App;