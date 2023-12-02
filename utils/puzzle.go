package utils

type Puzzle interface {
	Run()
	GetSummary() Summary
}
