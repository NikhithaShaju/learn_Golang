package main

import "fmt"

func main() {

	// arr1 := [3]int{1, 2, 3}
	// fmt.Println(arr1)

	// arr2 := []int{1, 2, 3, 4}
	// fmt.Println(arr2)
	// arr2 = append(arr2, 5)
	// fmt.Println(arr2)
	// arr1[2] = 7
	// fmt.Println(arr1)
	// arr2[3] = 5
	// fmt.Println(arr2)
	names := [8]string{"liya", "miya", "riya", "maaya", "fiya", "ammu", "anu", "akhil"}
	// fmt.Println(names[0])
	// fmt.Println(names[4])
	// fmt.Println("size of array", len(names))
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	fmt.Println(names[i])
	// }
	// for i := 2; i < 4; i++ {
	// 	fmt.Println(names[i])
	// }

	// for i := len(names) - 1; i >= 0; i-- {
	// 	fmt.Println(names[i])
	// }
	for i := 0; i < 8; i += 2 {
		fmt.Println(names[i])
	}
	for i := 7; i >= 0; i -= 2 {
		fmt.Println(names[i])

	}

}

type calculate struct {
	value1   int
	value2   int
	operator string
}

func addCalculation(calMains calculate) calculate {
	calMains.operator = "+"
	fmt.Println(calMains.value1 + calMains.value2)
	return calMains

}
func mulCalculation(calMains calculate) calculate {
	calMains.operator = "*"
	fmt.Println(calMains.value1 * calMains.value2)
	return calMains

}
func divCalculation(calMains calculate) calculate {
	calMains.operator = "/"
	fmt.Println(calMains.value1 / calMains.value2)
	return calMains

}
func subCalculation(calMains calculate) calculate {
	calMains.operator = "-"
	fmt.Println(calMains.value1 - calMains.value2)
	return calMains

}

// type player struct {
// 	name  string
// 	power int
// 	life  int
// }

// func decreasePlayerLife(pDetails player) player {
// 	pDetails.life = pDetails.life - 1
// 	fmt.Println("life of", pDetails.name, "decreased by one")

// 	return pDetails
// }
// func increasePlayerLife(pDetails player) player {
// 	pDetails.life = pDetails.life + 1
// 	fmt.Println("life of", pDetails.name, "increased by one")

// 	return pDetails
// }

// type student struct {
// 	name  string
// 	age   int
// 	class int
// }
// count := 10
// fmt.Println(count)
// count = count + 2
// fmt.Println(count)
// count = count - 5
// fmt.Println(count)
// total := 0
// fmt.Println(total)
// sum := total + count
// fmt.Println(sum)
// var a string = "liya"
// var b int = 23
// var c int = 12
// var d bool
// const weeks = 7
// fmt.Println(weeks)
// var e, f, d int = 22, 12, 56
// e++
// f--
// d += 4
// fmt.Println(e, f, d)

// fmt.Println(a, b, c, d,"a")
// int=0
// string=""
// float64=0
// bool=false
// studails := student{
// 	name:  a,
// 	age:   b,
// 	clasets: c,
// fmt.Println("game started")
// fmt.Println("enter your player name")
// playername := ""
// fmt.Scan(&playername)
// fmt.Println("enter how many life you want")
// var playerlife int
// fmt.Scan(&playerlife)

// playerdetails1 := player{
// 	name:  playername,
// 	power: 10,
// 	life:  playerlife,
// }
// fmt.Println("player created")
// fmt.Println("do you want to incraese or decrease life?")
// userDecision := ""
// fmt.Scan(&userDecision)
// if userDecision == "increase" {
// 	playerdetails1 = increasePlayerLife(playerdetails1)
// 	fmt.Println("status of player", playerdetails1.name, "life=", playerdetails1.life)

// } else if userDecision == "decrease" {

// 	playerdetails1 = decreasePlayerLife(playerdetails1)
// 	fmt.Println("status of player", playerdetails1.name, "life=", playerdetails1.life)

// } else {
// 	fmt.Println("type  either increase or decrease")
// }

// // playerdetails1 = decreasePlayerLife(playerdetails1)
// // fmt.Println("status of player", playerdetails1.name, "life=", playerdetails1.life)

// // playerdetails1 = increasePlayerLife(playerdetails1)
// // fmt.Println("status of player", playerdetails1.name, "life=", playerdetails1.life)

// fmt.Println("enter your first value")
// firstValue := 0
// fmt.Scan(&firstValue)
// fmt.Println("enter your seond value ")
// secondValue := 0
// fmt.Scan(&secondValue)
// fmt.Println("enter an operator")
// operator := ""
// fmt.Scan(&operator)

// allCalculation := calculate{
// 	value1:   firstValue,
// 	value2:   secondValue,
// 	operator: operator,
// }
// if operator == "+" {
// 	allCalculation = addCalculation(allCalculation)
// 	fmt.Println(allCalculation.value1, "added with", allCalculation.value2)
// } else if operator == "*" {
// 	allCalculation = mulCalculation(allCalculation)
// 	fmt.Println(allCalculation.value1, "multiplied with", allCalculation.value2)
// } else if operator == "/" {
// 	allCalculation = divCalculation(allCalculation)
// 	fmt.Println(allCalculation.value1, "divided with", allCalculation.value2)
// } else if operator == "-" {
// 	allCalculation = subCalculation(allCalculation)
// 	fmt.Println(allCalculation.value1, "substracted with", allCalculation.value2)
// } else {
// 	fmt.Println("This is a wrong parameter")
// }
// }
