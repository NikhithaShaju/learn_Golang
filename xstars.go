package learngolang

import "fmt"

// line1:="*     *" //0,6
//         0123456
// line2:=" *   * " //1,5
//         0123456
// line3:="  * *  " //2,4
// 		   0123456
// line4:="   *   " //3,3
// 		  0123456
// line3:="  * *  " //2,4
//         0123456
// line2:=" *   * " //1,5
//         0123456
// line1:="*     *" //0,6
//         0123456

// 0,6
// 1,5
// 2,4
// 3,3
// 2,4
// 1,5
// 0,6

// 7
func generateSequence(n int) {
	start := 0
	end := n - 1
	for i := 0; i < n; i++ {
		printStarOnIndex(start, end, n)
		start = start + 1
		end = end - 1

	}

}
func printStarOnIndex(a int, b int, count int) {
	starString := ""
	starSymbol := "*"
	dashSymbol := " "

	for i := 0; i < count; i++ { //0,1,2,3,4,5
		if i == a || i == b {
			starString = starString + starSymbol
		} else {
			starString = starString + dashSymbol
		}

	}
	fmt.Println(starString)
}

// 	fmt.Println("*     *")

// 	fmt.Println(" *   * ")
// 	fmt.Println("  * *  ")
// 	fmt.Println("   *   ")
// 	fmt.Println("  * *  ")
// 	fmt.Println(" *   * ")
// 	fmt.Println("*     *")

// 0123456
// 0     6
//  1   5
//   2 4
//    3

// 01234
// 0   4
//  1 3
//   2
//  1 3
// 0   4

// midc:=n-2
