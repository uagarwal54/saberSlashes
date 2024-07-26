package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error while opening the file", err)
	}
	var data []byte
	data, err = io.ReadAll(f)
	if err != nil {
		fmt.Println("Error while reading the file", err)
	}
	stringifiedData := string(data)
	slice := strings.Split(stringifiedData, "\n")
	for i, v := range slice {
		fmt.Println(i, v)
	}

}
