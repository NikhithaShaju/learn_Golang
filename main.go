package main

import "fmt"

func main() {

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
	fmt.Println("enter your first value")
	firstValue := 0
	fmt.Scan(&firstValue)
	fmt.Println("enter your seond value ")
	secondValue := 0
	fmt.Scan(&secondValue)
	fmt.Println("enter an operator")
	operator := ""
	fmt.Scan(&operator)

	allCalculation := calculate{
		value1:   firstValue,
		value2:   secondValue,
		operator: operator,
	}
	if operator == "+" {
		allCalculation = addCalculation(allCalculation)
		fmt.Println(allCalculation.value1, "added with", allCalculation.value2)
	} else if operator == "*" {
		allCalculation = mulCalculation(allCalculation)
		fmt.Println(allCalculation.value1, "multiplied with", allCalculation.value2)
	} else if operator == "/" {
		allCalculation = divCalculation(allCalculation)
		fmt.Println(allCalculation.value1, "divided with", allCalculation.value2)
	} else if operator == "-" {
		allCalculation = subCalculation(allCalculation)
		fmt.Println(allCalculation.value1, "substracted with", allCalculation.value2)
	} else {
		fmt.Println("This is a wrong parameter")
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
// }
