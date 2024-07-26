package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type tempDataStore struct {
	keyValStore map[string]string
	mu          sync.RWMutex
}

// This is the index of the last written data to the latest storage file
var currentWrittenIndex int = 0

func (tds *tempDataStore) storeDataToDisk(triggerChannel chan bool) {
	fileCount := 0
	prefix := "./store/store_"
	currentFile := ""
	files, fileCount, err := walkDir()
	if err != nil {
		log.Fatal("Error while reading the storage directory", err)
	}
	if len(files) == 0 {
		fileCount = 00001
		currentFile = prefix + strconv.Itoa(fileCount)
	} else {
		lastFile := prefix + strconv.Itoa(fileCount)
		if files[lastFile] >= int64(maxStorageFileSize) {
			fileCount++
			lastFile = prefix + strconv.Itoa(fileCount)
		}
		currentFile = lastFile
	}

	// Initially set timer for 12 seconds as the sever is expected to be up and running
	storageTimer := time.NewTimer(30 * time.Second)
	for {
		select {
		case <-storageTimer.C:
			writeDataToPersistantStorage(files, &currentFile, fileCount, prefix, tds)
			storageTimer = time.NewTimer(30 * time.Second)
		case <-triggerChannel:
			fmt.Println("Got a kill order, persisting data")
			writeDataToPersistantStorage(files, &currentFile, fileCount, prefix, tds)
			fmt.Println("Sleeping for 5 seconds")
			time.Sleep(5 * time.Second)
			fmt.Println("Exiting")
			os.Exit(1)
		}
	}
}

func writeDataToPersistantStorage(files map[string]int64, currentFile *string, fileCount int, prefix string, tds *tempDataStore) {
	var err error
	if files[*currentFile] >= int64(maxStorageFileSize) {
		fmt.Println("File size exceeded the limit")
		fileCount++
		*currentFile = prefix + strconv.Itoa(fileCount)
		currentWrittenIndex = 0
		if err = tds.writeToFile(*currentFile); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("File size is within the limit")
		if err = tds.writeToFile(*currentFile); err != nil {
			log.Fatal(err)
		}
	}
}

func (tds *tempDataStore) writeToFile(currentFile string) error {
	currentFileHandle, err := os.OpenFile(currentFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error while opning the storage file", err)
	}
	defer currentFileHandle.Close()
	if currentWrittenIndex == 0 {
		for k, v := range tds.keyValStore {
			currentFileHandle.WriteString(k + ":" + v + "\n")
			currentWrittenIndex++
		}
	} else if len(tds.keyValStore) > currentWrittenIndex {
		count := 0
		for k, v := range tds.keyValStore {
			count++
			if count > currentWrittenIndex {
				currentFileHandle.WriteString(k + ":" + v + "\n")
				currentWrittenIndex++
			}
		}
	}

	fmt.Println("Writing the data to file")
	currentFileHandle.Sync()
	// This is done because the buffers and slices are not imediately garbage collected
	return nil
}
