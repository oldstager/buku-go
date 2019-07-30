package main

import "fmt"

func main() {
	i, j := 42, 2701

	// p berisi memory address dari i
	p := &i
	// tampilkan isi dari memory address p
	fmt.Println(*p)
	// isi dari memory address yang ditunjuk p diubah
	*p = 21
	// implikasinya pada variabel i:
	fmt.Println(i)

	// p berisi memory address dari j
	p = &j
	// isi dari memory address yang ditunjuk p, diubah
	// menjadi isi memoery address yang lama, dibagi 37
	*p = *p / 37
	// implikasinya pada variabel j
	fmt.Println(j) // see the new value of j

	var pa *int

	fmt.Printf("pointer pa dengan tipe %T dan nilai %v\n", pa, pa)

}
