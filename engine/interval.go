package engine

import (
	"strconv"
	"strings"
)

type (
	Interval  int
	Intervals []Interval
)

func (intervals Intervals) Formula() ChordFormula {
	var intervalsString []string

	for _, interval := range intervals {
		intervalsString = append(intervalsString, strconv.Itoa(int(interval)))
	}

	return ChordFormula(strings.Join(intervalsString, " "))

}

func (intervals Intervals) String() string {
	return string(intervals.Formula())
}
