package iotest

import (
	"fmt"
	"io"
	"os"
)

func PrintFile(fileName string) {
	f, err := os.Open(fileName)

	if err != nil {
		panic("read file occurred error : " + err.Error())
	}
	defer f.Close()

	printReader(f)
}

func printReader(reader io.Reader) (err error) {
	var (
		buffer = make([]byte, 5)
		n      int
	)

	for err == nil {
		n, err = reader.Read(buffer) // n為傳回讀入的位元組數
		fmt.Printf("%v", string(buffer[:n]))
	}

	if err == io.EOF { //結尾 err會等於io.EOF
		err = nil
	}

	return err
}
