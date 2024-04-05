package dataloader

type executionError struct {
	File   string
	Line   int
	Script string
	Error  string
}
