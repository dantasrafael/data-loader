package dataloader

import (
	"database/sql"
	"io/fs"
)

type DataLoader struct {
	filesInfo       []fs.DirEntry
	database        *sql.DB
	filesContent    map[string][]string
	executionErrors []executionError
}

func New() *DataLoader {
	return &DataLoader{}
}

func (loader *DataLoader) Load() {
	loader.loadFilesInfo()
	loader.loadEnvs()
	loader.loadDatabaseConnection()
	loader.loadFilesContent()
	loader.processFilesContent()
	loader.generateErrorsCsvFile()

	defer loader.database.Close()
}
