package main

import "fmt"
import "sort"

func DoSort(theNums ...float64) []float64 {

	sort.Float64s(theNums)
	return theNums

}

func main() {

	fmt.Println(DoSort(111, 12, 32, 4, 5, 6, 27, 87, 9, 100))

}
