package dataloader

import (
	"log"
	"os"
	"strings"
)

const (
	sqlFolder    string = "./sql"
	sqlExtension string = ".sql"

	infoFilesLoaded            string = "SQL files info loaded"
	errorCouldNotReadSqlFolder string = "Could not read sql folder: %v"
	errorSqlFolderIsEmpty      string = "SQL folder is empty"
)

func (loader *DataLoader) loadFilesInfo() {
	files, err := os.ReadDir(sqlFolder)
	if err != nil {
		log.Fatalf(errorCouldNotReadSqlFolder, err)
	}

	if len(files) == 0 {
		log.Fatal(errorSqlFolderIsEmpty)
	}

	for index, file := range files {
		if !strings.HasSuffix(file.Name(), sqlExtension) {
			files = append(files[:index], files[index+1:]...)
		}
	}

	loader.filesInfo = files
	files = nil
	log.Println(infoFilesLoaded)
}
