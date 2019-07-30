package main

import (
	"fmt"
	"reflect"
)

var (
	hasilPerbandingan bool
	angka1            uint8 = 21
	angka2            uint8 = 17
)

func main() {
	hasilPerbandingan = angka1 < angka2
	fmt.Printf("angka1 = %d\n", angka1)
	fmt.Printf("angka2 = %d\n", angka2)
	fmt.Println(reflect.TypeOf(hasilPerbandingan))
	fmt.Println(hasilPerbandingan)
}
