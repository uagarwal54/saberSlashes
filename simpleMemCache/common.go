package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var keyValStore tempDataStore
var maxStorageFileSize int = 1000000000
var triggerChannel = make(chan bool)

func walkDir() (map[string]int64, int, error) {
	entries, err := os.ReadDir("./store")
	if err != nil {
		fmt.Println("Error while reading the directory")
		return nil, 0, err
	}
	files := make(map[string]int64)
	if len(entries) == 0 {
		return files, 0, nil
	}
	fileCount := 0
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			fmt.Println("Error while reading the file info")
			return nil, 0, err
		}
		fileC := strings.Split(e.Name(), "_")[1]
		fileCinI, err := strconv.Atoi(fileC)
		if err != nil {
			log.Fatal("Error while converting the file count to integer", err)
		}
		if fileCinI > fileCount {
			fileCount = fileCinI
		}
		files[e.Name()] = info.Size()
	}
	return files, fileCount, nil
}
