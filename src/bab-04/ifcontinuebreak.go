package main

import "fmt"

func main() {
	var a int = 10

	for a < 20 {
		if a == 12 {
			a += 1
			continue
		}
		a++
		if a > 15 {
			break
		}
		fmt.Printf("Nilai a: %d\n", a)
	}
}
