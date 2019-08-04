package main

import "fmt"

type Pegawai struct {
	Nama string
	Usia int
}

func (p Pegawai) print() {

	fmt.Println("Nama: ", p.Nama, "Usia: ", p.Usia)

}

func (p *Pegawai) printPointer() {

	fmt.Println("Nama: ", p.Nama, "Usia: ", p.Usia)
}

func main() {

	var pegawaiKantor2 = Pegawai{"Peg kantor 2 - 1", 24}

	pegawaiKantor2.print()

	pPeg := &pegawaiKantor2

	pPeg.printPointer()

}
