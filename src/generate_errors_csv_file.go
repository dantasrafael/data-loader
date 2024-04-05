package dataloader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	infoNoErrorsFound                        string = "No errors found"
	infoErrorsCsvFileProcessingInitialize    string = "Initializing errors CSV files creation"
	infoErrorsCsvFileCreated                 string = "Errors CSV file created"
	errorCouldNotCreateFile                  string = "Could not create file: %v"
	errorCouldNotCreateHeadersInCsvErrorFile string = "Could not create headers in CSV error file: %v"
	errorFileContent                         string = "file: %s, line: %d, script: %s, error: %s"
)

func (loader *DataLoader) generateErrorsCsvFile() {
	fmt.Println()
	if len(loader.executionErrors) == 0 {
		log.Println(infoNoErrorsFound)
		return
	}

	log.Println(infoErrorsCsvFileProcessingInitialize)
	fileName := fmt.Sprintf("errors_%s.csv", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf(errorCouldNotCreateFile, err)
		loader.printFilesContentAfterError()
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"file", "line", "script", "error"}); err != nil {
		log.Printf(errorCouldNotCreateHeadersInCsvErrorFile, err)
		loader.printFilesContentAfterError()
		return
	}

	for index, error := range loader.executionErrors {
		if err := writer.Write([]string{error.File, strconv.Itoa(error.Line), error.Script, error.Error}); err != nil {
			log.Printf(errorFileContent, error.File, error.Line, error.Script, error.Error)
		}
		showProgressBarWithPercent(fileName, (index+1)*100/len(loader.executionErrors))
	}

	loader.executionErrors = nil
	log.Println(infoErrorsCsvFileCreated)
	fmt.Println()
}

func (loader *DataLoader) printFilesContentAfterError() {
	for _, error := range loader.executionErrors {
		log.Printf(errorFileContent, error.File, error.Line, error.Script, error.Error)
	}

	loader.executionErrors = nil
}
