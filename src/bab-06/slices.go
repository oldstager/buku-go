package main

import "fmt"

func main() {

	distros := [6]string{"arch", "gentoo", "opensuse", "devuan", "debian", "fedora"}

	var myDistros []string = distros[3:5]

	fmt.Println(myDistros)
	fmt.Println(myDistros[0])

	distros[4] = "Ubuntu"

	// slices merupakan referensi ke array. Jika array diubah, slices juga berubah
	// karena menunjuk ke lokasi memory yang sama
	fmt.Println(myDistros)

	// membuat slices dengan cara ini juga bisa:
	b := make([]int, 1, 5)
	fmt.Println(b)

	b[0] = 2
	fmt.Println(len(b))
	fmt.Println(b[0])

}
