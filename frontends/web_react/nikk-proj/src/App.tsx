import { useState } from "react";
import heroImg from "./assets/hero.png";
import reactLogo from "./assets/react.svg";
import viteLogo from "./assets/vite.svg";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);
  const [num, setNum] = useState(7);
  const [name, setName] = useState("nikk");
  const [studentList, SetStudentList] = useState([]);
  const [studentDetail, SetStudentDetails] = useState({
    Address: "",
    Age: 0,
    Mark: 0,
  });

  async function fetchStudents() {
    const url = "http://localhost:8000/studentlist";

    try {
      const response = await fetch(url);

      // 2. Always check if the HTTP status code is successful (200-299)
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      // 3. Parse the data stream into JSON
      const data = await response.json();
      SetStudentList(data);
      let students = "";
      for (let i = 0; i < data.length; i++) {
        students = students + data[i] + "<br>";
      }
      // document.getElementById("stud").innerHTML = students;
    } catch (error) {
      // Catch network failures or parsing issues
      console.error("Failed to fetch local data:", error);
    }
  }

  async function fetchStudentDetails(studentName: string) {
    const url =
      "http://localhost:8000/studentDetailsURL" + "?studentName=" + studentName;

    try {
      const response = await fetch(url);

      // 2. Always check if the HTTP status code is successful (200-299)
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }

      // 3. Parse the data stream into JSON
      const data = await response.json();
      SetStudentDetails(data);
    } catch (error) {
      // Catch network failures or parsing issues
      console.error("Failed to fetch local data:", error);
    }
  }

  console.log("studentList ::", studentList);
  console.log("student Details ::", studentDetail);
  return (
    <>
      <section id="center">
        <div>
          <h1>Student List</h1>
        </div>
        <button
          type="button"
          className="counter"
          onClick={() => fetchStudents()}
        >
          fetch students
        </button>
        <div
          style={{
            // background: "green",
            display: "flex",
            flexDirection: "row",
            padding: 10,
          }}
        >
          <div style={{ paddingRight: 20 }}>
            {studentList.map((elemnt) => (
              <div
                onClick={() => fetchStudentDetails(elemnt)}
                style={{
                  borderWidth: "1",
                  borderStyle: "solid",
                  borderColor: "white",
                  paddingLeft: 25,
                  paddingRight: 25,
                  borderRadius: 10,
                  marginBottom: 5,
                }}
                key={elemnt}
              >
                {elemnt}
              </div>
            ))}
          </div>
          <div //details
            style={{
              borderWidth: "1",
              borderStyle: "solid",
              borderColor: "white",
              paddingLeft: 25,
              paddingRight: 25,
              borderRadius: 10,
              height: 100,
            }}
          >
            <p style={{ fontSize: 20, paddingBottom: 5 }}>Student details</p>
            <p> Address : {studentDetail.Address}</p>
            <p> Age : {studentDetail.Age}</p>
            <p> Mark : {studentDetail.Mark}</p>
          </div>
        </div>

        {/* <button
          type="button"
          className="counter"
          onClick={() => setCount((count) => count + 1)}
        >
          Count is {num}
        </button> */}
      </section>

      <div className="ticks"></div>
      <p id="stud"></p>
      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  );
}

export default App;
