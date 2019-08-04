package main

import "fmt"

func doMath(a, b float64) (float64, float64, float64, float64) {
	return a * b, a / b, a + b, a - b
}

func main() {

	a, b, c, d := doMath(4, 10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	_, nil1, _, _ := doMath(12, 21)

	fmt.Println(nil1)
}
