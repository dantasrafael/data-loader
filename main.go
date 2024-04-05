package main

import (
	dataloader "data-loader/src"
	"log"
	"time"
)

const (
	infoStartingDataLoader string = "Starting data loader"
	infoFinishedDataLoader string = "Data loader took %fs"
)

func main() {
	start := time.Now()
	log.Print("\033[H\033[2J")
	log.Println(infoStartingDataLoader)

	dataloader.New().Load()

	log.Printf(infoFinishedDataLoader, time.Since(start).Seconds())
}
