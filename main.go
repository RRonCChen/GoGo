package main

import (
	"ron.test/httptest"
	"ron.test/iotest"
)

func main() {
	iotest.PrintFile("./testTxt/test.txt")
	httptest.TestHttpGet("https://google.com")
}
