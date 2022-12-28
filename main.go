package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/bmcculley/mockexample/client"
)

func main() {
	res, err := client.Get("http://static.ex.net/extension-testing/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	var reader io.ReadCloser
	reader = res.Body
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	fmt.Println(buf.String())
}