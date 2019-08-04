package main

import "fmt"

type Pegawai struct {
	Nama string
	Usia int
}

func main() {

	var pegawaiKantor1 Pegawai
	fmt.Println(pegawaiKantor1)

	var pegawaiKantor2 = Pegawai{"Peg kantor 2 - 1", 24}
	fmt.Println(pegawaiKantor2)

	var pegawaiKantor3 = Pegawai{Nama: "Peg kantor 3 - 1", Usia: 54}
	fmt.Println(pegawaiKantor3)

	pegawaiKantor4 := Pegawai{
		Nama: "Peg kantor 4 - 1",
		Usia: 54,
	}
	fmt.Println(pegawaiKantor4)
	fmt.Println(pegawaiKantor4.Nama)
	fmt.Println(pegawaiKantor4.Usia)
	pegawaiKantor4.Nama = "Peg kantor 4 - baru"
	fmt.Println(pegawaiKantor4)

	var pPegawaiKantor4 = &pegawaiKantor4
	var p2PegawaiKantor4 = &pegawaiKantor4

	fmt.Println(*pPegawaiKantor4)
	fmt.Println(*p2PegawaiKantor4)

	pegawaiKantor4.Nama = "Pengganti pegawai kantor 4"

	fmt.Println(*pPegawaiKantor4)
	fmt.Println(*p2PegawaiKantor4)
	fmt.Println((*p2PegawaiKantor4).Nama)

	*pPegawaiKantor4 = Pegawai{Nama: "Pengganti lagi utk kantor 4", Usia: 25}

	fmt.Println(*pPegawaiKantor4)
	fmt.Println(*p2PegawaiKantor4)
	fmt.Println((*p2PegawaiKantor4).Nama)

	// pointer ke struct dengan new:
	pPeg := new(Pegawai)

	pPeg.Nama = "pPeg 1"
	pPeg.Usia = 21

	fmt.Println(pPeg)
	fmt.Println(*pPeg)

}
