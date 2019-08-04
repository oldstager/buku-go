package main

import "fmt"

func main() {

	func(angka1, angka2 float64) {
		fmt.Println(angka1 + angka2)
	}(10.0, 20.5)

}
