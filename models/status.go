package cmd

import (
	"errors"
	"fmt"
)

type Status int

var ErrInvalidStatus = errors.New("invalid status")

const (
	Todo    Status = iota
	Ongoing Status = iota
	Done    Status = iota
)

func (s Status) String() string {
	switch s {
	case Todo:
		return "To-Do"
	case Ongoing:
		return "Ongoing"
	case Done:
		return "Done"
	default:
		return "Unknown"
	}
}

func ParseStatus(in string) (Status, error) {
	switch in {
	case Todo.String():
		return Todo, nil
	case Ongoing.String():
		return Ongoing, nil
	case Done.String():
		return Done, nil
	}
	return Todo, fmt.Errorf("%q is not a valid status: %w", in, ErrInvalidStatus)
}
