package main

import (
	"log"
	"os"
	"encoding/json"
	"github.com/gocarina/gocsv"
)

// LoadComputerList loads the Computer and return the readed list or an error
func loadComputerList(computerCsvFilePath string) ([]Computer, error) {

	var computers []Computer

	computerCsvFilePointer, computerCsvFileOpenError := os.Open(computerCsvFilePath)
	if computerCsvFileOpenError != nil {
		log.Fatalf("Error on Opening the file")
		return computers, computerCsvFileOpenError
	}

	if computerCsvFileParserError := gocsv.UnmarshalFile(computerCsvFilePointer, &computers); computerCsvFileParserError != nil {
		computerCsvFilePointer.Close()
		return computers, computerCsvFileParserError
	}

	computerCsvFilePointer.Close()
	return computers, nil
}

// Test if Path is a File and it exists
func FileExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}
	return false
}


// LoadComputerList loads the Computer and return the readed list or an error from json file
func loadComputerListFromJson(computerJsonFilePath string) ([]Computer, error) {
	
	var computers []Computer

	computerJsonFilePointer, computerJsonFileOpenError := os.Open(computerJsonFilePath)
	if computerJsonFileOpenError != nil {
		log.Fatalf("Error on Opening the file")
		return computers, computerJsonFileOpenError
	}

	if computerJsonFileParserError := json.NewDecoder(computerJsonFilePointer).Decode(&computers); computerJsonFileParserError != nil {
		computerJsonFilePointer.Close()
		return computers, computerJsonFileParserError
	}

	computerJsonFilePointer.Close()
	return computers, nil
	
}