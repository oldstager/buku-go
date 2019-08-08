package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("daerah istimewa yogyakarta")
	buf := make([]byte, 5)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buf[:n]))

	}

}
