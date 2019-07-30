package main

import "fmt"

func main() {
	var a int = 200

	if a == 10 {
		fmt.Printf("Nilai a = 10\n")
	} else if a == 20 {
		fmt.Printf("Nilai a = 20\n")
	} else if a == 30 {
		fmt.Printf("Nilai a = 30\n")
	} else {
		fmt.Printf("Semua nilai salah\n")
	}
	fmt.Printf("Nilai dari a adalah: %d\n", a)
}
