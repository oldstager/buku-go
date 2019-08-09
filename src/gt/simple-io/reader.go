/*
  data source -> io.Reader -> transfer buffer []byte
*/
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// data source
	reader := strings.NewReader("daerah istimewa yogyakarta")

	// deklarasi transfer buffer
	p := make([]byte, 5)

	// looping sampai io.EOF
	for {

		// type Reader interface {
		//   Read(p []byte) (n int, err error)
		// }
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				//fmt.Println(string(p[:n]))
				fmt.Println(p[:n])
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Println(string(p[:n]))
		fmt.Println(p[:n])
	}
}
