package main

import (
	"fmt"
)

func main() {

	const mainCodingLang = "Go"
	const kiamatMakinDekat = true

	const angka1, angka2 = 25, 8

	const (
		nomorPegawai = "P001"
		gaji         = 50000000
	)

	const negaraKu string = "Indonesia"

	const gajiBersihSetelahSetorIstri = gaji - 49000000

	fmt.Println(mainCodingLang)
	fmt.Println(kiamatMakinDekat)
	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(nomorPegawai)
	fmt.Println(gaji)
	fmt.Println(negaraKu)
	fmt.Println(gajiBersihSetelahSetorIstri)

	// ./konstanta.go:28:7: cannot assign to gaji
	//gaji = 10000000

}
