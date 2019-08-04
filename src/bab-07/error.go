package main

import "fmt"
import "errors"

func bagi(angka1, angka2 float64) (float64, error) {
	if angka2 == 0 {
		return -1, errors.New("Pembagian dengan angka 0 dilarang")
	}
	return angka1 / angka2, nil
}

func main() {

	var hasilBagi, err = bagi(23, 3)
	fmt.Println(hasilBagi, err)

	var hasilBagi2, err2 = bagi(23, 0)
	fmt.Println(hasilBagi2, err2)

}
