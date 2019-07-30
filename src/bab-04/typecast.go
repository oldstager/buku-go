package main

import (
	"fmt"
)

var (
	angka1     uint8   = 21
	angka2     uint8   = 17
	angkaFloat float64 = 7.1
)

func main() {
	// ./typecast.go:14:11: cannot use "abc" (type string) as type uint8 in assignment
	//angka1 = "abc"
	fmt.Println(angka1 + angka2)
	// ./typecast.go:15:21: invalid operation: angka1 + angkaFloat (mismatched types uint8 and float64)
	//fmt.Println(angka1 + angkaFloat)
	fmt.Println(float64(angka1) + angkaFloat)
}
