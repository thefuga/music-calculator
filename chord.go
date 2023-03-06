package main

import (
	"fmt"
	"strings"
)

const (
	PerfectUnison = iota
	MinorSecond
	MajorSecond
	MinorThird
	MajorThird
	PerfectFourth
	Tritone
	PerfectFifth
	MinorSixth
	MajorSixth
	MinorSiventh
	MajorSeventh
	Octave
)

const (
	A = iota
	Ax
	B
	C
	Cx
	D
	Dx
	E
	F
	Fx
	G
	Gx
)

var (
	Relationships [12][12]Interval
	MajorTriad = [3]Interval{PerfectUnison, MajorThird, PerfectFifth}
)

type (
	Note  int
	Chord struct {
		Name      string
		Notes []Note
		Intervals []Interval
	}
	Interval int
)

func (note Note) AscendingDistance(to Note) Interval {
	return Relationships[note][to]
}

func (note Note) DescendingDistance(to Note) Interval {
	return Relationships[to][note]
}

func NewChord(notes ...Note) {
	chord := Chord{Notes: notes}
	
	for _, note := range notes {
chord.Intervals 
	}
	switch {

	}
}



func initRelationships() {
	for from := 0; from < 12; from++ {
		for to := 0; to < 12; to++ {
			Relationships[from][to] = calculateInterval(from, to)
		}
	}
}

func calculateInterval(from, to int) Interval {
	if from < to {
		return Interval(to - from)
	}

	if from > to {
		return Interval(-to + from)
	}

	return PerfectUnison
}

func main() {
	initRelationships()
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Printf("%d  ", Relationships[i][j])
		}
		fmt.Println()
	}
}
