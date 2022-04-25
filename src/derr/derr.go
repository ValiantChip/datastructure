package derr

import "fmt"

type IndexOutOfBoundsError struct {
	Message       string
	Bounds        int
	RecievedIndex int
}

func (e IndexOutOfBoundsError) Error() string {
	if e.Message == "" {
		return fmt.Sprintf("index out of bounds error; Index %d out of bounds %d", e.RecievedIndex, e.Bounds)
	}
	return fmt.Sprintf("index out of bounds error; %s", e.Message)
}
