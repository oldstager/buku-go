// taken from: https://www.thegeekstuff.com/2019/03/go-array-examples/
package main

import "fmt"

func main() {

	// String Array
	fmt.Println("1. String Array : ")

	var distros [5]string
	distros[0] = "Ubuntu"
	distros[1] = "CentOS"
	distros[2] = "RedHat"
	distros[3] = "Debian"
	distros[4] = "OpenBSD"

	mydistro := distros[1]
	fmt.Println("mydistro = ", mydistro)
	fmt.Println("distros[2] = ", distros[2])
	fmt.Println("distros = ", distros)
	fmt.Println("Number of distros = ", len(distros))

	// Integer Array (Numbers)
	fmt.Println()
	fmt.Println("2. Integer Array : ")

	var ids [5]int
	ids[0] = 1
	ids[1] = 2
	ids[2] = 3
	ids[3] = 4
	ids[4] = 5

	myid := ids[3]
	fmt.Println("myid = ", myid)
	fmt.Println("ids[2] = ", ids[2])
	fmt.Println("ids = ", ids)
	fmt.Println("Number of ids = ", len(ids))

	// Declare and Initialize Array at the same time
	fmt.Println()
	fmt.Println("3. Declare and Initialize Array at the same time : ")

	os := [3]string{"Linux", "Mac", "Windows"}
	fmt.Println("os = ", os)
	fmt.Println("Number of os = ", len(os))

	fibonacci := [6]int{1, 1, 2, 3, 5, 8}
	fmt.Println("fibonacci = ", fibonacci)

	// Multi-line Array Initialization Syntax
	fmt.Println()
	fmt.Println("4. Multi-line Array Initialization Syntax : ")

	temperature := [3]float64{
		98.5,
		65.5,
		83.2,
	}
	fmt.Println("temperature = ", temperature)

	names := [3]string{
		"John",
		"Jason",
		"Alica",
		// "Rita",
	}
	fmt.Println("names = ", names)

	// Default Values in an Array
	fmt.Println()
	fmt.Println("5. Default Values in an Array : ")

	empIds := [5]int{101, 102, 103}
	fmt.Println("empIds = ", empIds)

	empNames := [5]string{"John", "Jason"}
	fmt.Println("empNames = ", empNames)

	// Loop through Array using For and Range
	fmt.Println()
	fmt.Println("6. Loop through Array using For and Range : ")

	for index, value := range distros {
		fmt.Println(index, " = ", value)
	}

	// Loop through Array using For and Range (Ignore Index)
	fmt.Println()
	fmt.Println("7. Loop through Array using For and Range (Ignore Index) : ")

	total := 0
	for _, value := range ids {
		total = total + value
	}
	fmt.Println("total of all ids = ", total)

	// Initialize an integer array with sequence
	fmt.Println()
	fmt.Println("8. Initialize an integer array with sequence : ")
	var sequence [10]int
	counter := 10
	for index, _ := range sequence {
		sequence[index] = counter
		counter = counter + 5
	}
	fmt.Println("sequence = ", sequence)

	// Multi dimensional array
	fmt.Println()
	fmt.Println("9. Multi dimensional array : ")

	count := 1
	var multi [4][2]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			multi[i][j] = count
			count++
		}
	}
	fmt.Println("Array 4 x 2 : ", multi)

}
