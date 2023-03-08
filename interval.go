package main

import (
	"strconv"
	"strings"
)

type (
	Interval  int
	Intervals []Interval
)

func (intervals Intervals) String() string {
	var intervalsString []string

	for _, interval := range intervals {
		intervalsString = append(intervalsString, strconv.Itoa(int(interval)))
	}

	return strings.Join(intervalsString, " ")
}
