package main

import "fmt"

func main() {

	var nilaiAngka int = 20
	var nilaiHuruf string

	switch nilaiAngka {
	case 90:
		nilaiHuruf = "A"
	case 80:
		nilaiHuruf = "B"
	case 50, 60, 70:
		nilaiHuruf = "C"
	default:
		nilaiHuruf = "D"
	}

	switch {
	case nilaiHuruf == "A":
		fmt.Println("Apik tenan!")
	case nilaiHuruf == "B":
		fmt.Println("Lumayan lah")
	case nilaiHuruf == "C", nilaiHuruf == "D":
		fmt.Println("Lulus sih ... tapi ..")
	case nilaiHuruf == "E":
		fmt.Println("Nangis bombay")
	default:
		fmt.Println("Nilai gak jelas, seperti wajah dosennya!")
	}
	fmt.Printf("Nilai anda =  %s\n", nilaiHuruf)
}
