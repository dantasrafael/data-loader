package dataloader

import (
	"log"

	"github.com/subosito/gotenv"
)

const (
	infoEnvsLoaded      string = "Envs loaded"
	errorLoadingEnvFile string = "Error loading env file: %v"
)

func (loader *DataLoader) loadEnvs() {
	err := gotenv.Load()
	if err != nil {
		log.Fatalf(errorLoadingEnvFile, err)
	}

	log.Println(infoEnvsLoaded)
}
