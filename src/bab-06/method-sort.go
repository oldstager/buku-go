package main

import "fmt"
import "sort"

type TheNumbers struct {
	Numbers []float64
}

func (tn TheNumbers) SortIt(theNums ...float64) []float64 {

	sort.Float64s(theNums)
	return theNums

}

func main() {

	var Angka = TheNumbers{{5.1, 1.2, 6.3, 12.1}}

	fmt.Println(Angka.SortIt())

}
