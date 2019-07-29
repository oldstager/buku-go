/*
	showgoenv.go

	Contoh program sederhana untuk menjelaskan
	struktur program Go untuk lebih dari satu
        binary executable

	(c) 2019 - bpdp.name

*/

package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		user   string
		goHome string
		goPath string
	)

	user = os.Getenv("USER")
	goHome = os.Getenv("GOROOT")
	goPath = os.Getenv("GOPATH")

	fmt.Printf("Halo %s", user)
	fmt.Printf("\nAnda menggunakan Go di %s", goHome)
	fmt.Printf("\nGOPATH anda di %s", goPath)
	fmt.Printf("\n")

}
