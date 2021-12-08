package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// Golang program to illustrate the usage of
// io.TeeReader() function

// Including main package

// Calling main
func Test_io(t *testing.T) {

	// Defining reader using NewReader method
	reader := strings.NewReader("GfG\n")

	// Defining buffer
	var buf bytes.Buffer

	// Calling TeeReader method with its parameters
	r := io.TeeReader(reader, &buf)
	io.MultiReader()
	//r.Read()

	// Calling Copy method with its parameters
	Reader, err := io.Copy(os.Stdout, r)

	// If error is not nil then panics
	if err != nil {
		panic(err)
	}

	// Prints output
	fmt.Printf("n:%v\n", Reader)
}

func demo() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// go tour io.Reader练习题
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

// TODO
//func exercise1() {
//	reader.Validate(MyReader{})
//}

func Test_writeFile(t *testing.T) {
	f, _:= os.OpenFile("./testfile.txt",os.O_WRONLY|os.O_CREATE,0x777)
	f.Write([]byte("1234"))
	f.Write([]byte("abcd"))
}
