package dataloader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	filePath               string = sqlFolder + "/%s"
	infoFilesContentLoaded string = "SQL files content loaded"
	errorCouldNotLoadFile  string = "Could not load file %s: %v"
)

func (loader *DataLoader) loadFilesContent() {
	loader.filesContent = make(map[string][]string)
	for _, fileInfo := range loader.filesInfo {
		filePath := fmt.Sprintf(filePath, fileInfo.Name())
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf(errorCouldNotLoadFile, filePath, err)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		loader.filesContent[fileInfo.Name()] = lines
	}

	loader.filesInfo = nil
	log.Println(infoFilesContentLoaded)
}
