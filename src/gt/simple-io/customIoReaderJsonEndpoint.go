package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	resp, err := http.Get("http://dummy.restapiexample.com/api/v1/employees")

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	jsonEndPoint := NewReader(resp.Body)
	buf := make([]byte, 1)
	for {
		n, err := jsonEndPoint.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
