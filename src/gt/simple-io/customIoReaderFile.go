package main

import (
	"fmt"
	"io"
	"os"
)

type Reader struct {
	reader io.Reader
}

func NewReader(reader io.Reader) *Reader {
	return &Reader{reader: reader}
}

func (r *Reader) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if err != nil {
		return n, err
	}

	buf := make([]byte, n)

	for i := 0; i < n; i++ {
		buf[i] = p[i]
	}

	copy(p, buf)
	return n, nil
}

func main() {

	file, err := os.Open("./theFile.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	jsonFile := NewReader(file)
	buf := make([]byte, 1)
	for {
		n, err := jsonFile.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
