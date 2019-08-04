package main

import "fmt"

type Pegawai struct {
	Nama string
	Usia int
}

func main() {

	var mPeg = make(map[string]Pegawai)

	mPeg["p0001"] = Pegawai{Nama: "Peg 1", Usia: 23}
	fmt.Println(mPeg)

	mPeg["p0001"] = Pegawai{Nama: "Peg 1 - edit", Usia: 24}
	fmt.Println(mPeg)

	mPeg["p0002"] = Pegawai{Nama: "Peg 2", Usia: 29}
	fmt.Println(mPeg)

	// iterasi
	for key, value := range mPeg {
		fmt.Println("NIP:", key, "Nama:", value.Nama, "Usia:", value.Usia)
	}

	// memeriksa ada atau tidak
	_, ok := mPeg["p0002"]
	if ok {
		fmt.Println("Ada pegawai dengan NIP p0002")
	}

	// contoh inisialisasi lagi
	mPeg2 := map[string]Pegawai{
		"p0003": {Nama: "Peg 3", Usia: 34},
		"p0004": {Nama: "Peg 4", Usia: 35},
		"p0005": {Nama: "Peg 5", Usia: 36},
		"p0006": {Nama: "Peg 6", Usia: 37},
		"p0007": {Nama: "Peg 7", Usia: 38},
		"p0008": {Nama: "Peg 8", Usia: 39},
	}
	fmt.Println(mPeg2)
	fmt.Println(mPeg2["p0008"])

	fmt.Println("Jumlah mPeg: ", len(mPeg))
	fmt.Println("Jumlah mPeg2: ", len(mPeg2))

	delete(mPeg2, "p0008")

	fmt.Println("Jumlah mPeg: ", len(mPeg))
	fmt.Println("Jumlah mPeg2: ", len(mPeg2))

	fmt.Println(mPeg2)
	fmt.Println(mPeg2["p0008"])

}
