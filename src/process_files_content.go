package dataloader

import (
	"fmt"
	"log"
	"sync"
)

const (
	infoFilesContentProcessingInitialize string = "Initializing SQL files content processing"
	infoFilesContentProcessed            string = "SQL files content processed"
)

func (loader *DataLoader) processFilesContent() {
	fmt.Println()
	log.Println(infoFilesContentProcessingInitialize)

	for file, scripts := range loader.filesContent {
		progress := make(chan int, len(scripts))
		go showProgressBarWithChannel(file, progress)

		var waitGroup sync.WaitGroup
		executor := func(line int, script string) {
			_, err := loader.database.Exec(script)
			if err != nil {
				loader.executionErrors = append(loader.executionErrors, executionError{file, line, script, err.Error()})
			}

			waitGroup.Done()
		}

		for line, script := range scripts {
			waitGroup.Add(1)

			go executor(line+1, script)

			progress <- (line + 1) * 100 / len(scripts)
		}

		waitGroup.Wait()
		close(progress)
	}

	loader.filesContent = nil
	log.Println(infoFilesContentProcessed)
}
