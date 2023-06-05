package httptest

import (
	"fmt"
	"io"
	"net/http"
)

func TestHttpGet(url string) (err error) {
	response, err := http.Get(url)

	if err != nil {
		panic("get url error occurred : " + err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic("get body content error occurred : " + err.Error())
	}

	for b := range body {
		fmt.Print(string(b))
	}

	return nil
}
