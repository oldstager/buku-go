// slightly modified from:
// https://blog.golang.org/defer-panic-and-recover
package main

import "fmt"

func main() {
	f()
	f2()
	fmt.Println("Returned normally from f.")
}

func f() {

	// jika ada panic, maka defer akan dipanggil
	defer func() {
		// mengambil nilai revocer, jika tidak nil (karena panic)
		// maka pada bagian ini akan dilakukan berbagai hal untuk
		// menangani panic tersebut
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")

}

func f2() {

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Will not be executed because of panic")

}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	// rekursif
	g(i + 1)
}
