package main

import "fmt"
import "strings"

func combineStr(theStr ...string) string {
	return strings.Join(theStr, " ")
}

func doMath(theOperator rune, theNums ...float64) float64 {

	counter := 0.0
	if theOperator == '+' {
		for _, theVal := range theNums {
			counter += theVal
		}
	}
	return counter

}

func main() {

	str1 := "String1"
	str2 := "String2"

	fmt.Println(combineStr(str1, str2))

	fmt.Println(doMath('+', 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}
