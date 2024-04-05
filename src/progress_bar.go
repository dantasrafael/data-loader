package dataloader

import "fmt"

const (
	progressBarLenght int = 40
)

func showProgressBarWithPercent(file string, percent int) {
	numberOfChars := percent * progressBarLenght / 100

	fmt.Printf("\r%s [", file)
	for index := 0; index < numberOfChars; index++ {
		fmt.Print("#")
	}

	for index := 0; index < progressBarLenght-numberOfChars; index++ {
		fmt.Print(" ")
	}

	fmt.Printf("] %d%%", percent)

	if percent == 100 {
		fmt.Println()
	}
}

func showProgressBarWithChannel(file string, progress <-chan int) {
	for p := range progress {
		showProgressBarWithPercent(file, p)
	}
}
