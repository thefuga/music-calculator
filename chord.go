package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	MajorTriadIntervals      = Intervals{PerfectUnison, MajorThird, PerfectFifth}
	MinorTriadIntervals      = Intervals{PerfectUnison, MinorThird, PerfectFifth}
	DiminishedTriadIntervals = Intervals{PerfectUnison, MinorThird, Tritone}
	AugmentedTriadIntervals  = Intervals{PerfectUnison, MajorThird, AugmentedFifth}
	Sus2TriadIntervals       = Intervals{PerfectUnison, MajorSecond, PerfectFifth}
	Sus4TriadIntervals       = Intervals{PerfectUnison, PerfectFourth, PerfectFifth}
)

type (
	Interval  int
	Intervals []Interval

	Chord interface {
		String() string
		Root() Note
	}

	chord struct {
		Name          string
		Notes         []Note
		Quality       string
		Type          int
		Intervals     Intervals
		FormulaFormat string
	}

	regularTriadChord struct {
		chord
	}

	susTriadChord struct {
		chord
	}
)

func (note Note) String() string {
	return note.Name
}

func NewChord(notes ...Note) Chord {
	if notes == nil {
		panic("empty notes")
	}

	baseChord := chord{Notes: notes}

	for _, note := range baseChord.Notes {
		baseChord.Intervals = append(
			baseChord.Intervals,
			baseChord.Root().AscendingDistance(note),
		)
	}

	switch baseChord.Intervals.String() {
	case MajorTriadIntervals.String():
		baseChord.FormulaFormat = regularTriadFormat
		baseChord.Quality = MajorQuality

		return regularTriadChord{baseChord}
	case MinorTriadIntervals.String():
		baseChord.Quality = MinorQuality
		baseChord.FormulaFormat = regularTriadFormat

		return regularTriadChord{baseChord}
	case DiminishedTriadIntervals.String():
		baseChord.Quality = DiminishedQuality
		baseChord.FormulaFormat = regularTriadFormat

		return regularTriadChord{baseChord}
	case AugmentedTriadIntervals.String():
		baseChord.Quality = AugmentedQuality
		baseChord.FormulaFormat = regularTriadFormat

		return regularTriadChord{baseChord}
	case Sus2TriadIntervals.String():
		baseChord.Quality = MajorQuality // TODO check this
		baseChord.FormulaFormat = sus2TriadFormat

		return susTriadChord{baseChord}
	case Sus4TriadIntervals.String():
		baseChord.Quality = MajorQuality // TODO check this
		baseChord.FormulaFormat = sus4TriadFormat

		return susTriadChord{baseChord}
	default:
		panic("unknown formula")
	}
}

func (chord regularTriadChord) String() string {
	return fmt.Sprintf(
		chord.FormulaFormat,
		chord.Root().String(),
		chord.Quality,
	)
}

func (chord susTriadChord) String() string {
	return fmt.Sprintf(
		chord.FormulaFormat,
		chord.Root().String(),
	)
}

func (chord chord) Root() Note {
	return chord.Notes[0]
}

func (chord chord) NotesCount() int {
	return len(chord.Notes)
}

func (intervals Intervals) String() string {
	var intervalsString []string
	for _, interval := range intervals {
		intervalsString = append(intervalsString, strconv.Itoa(int(interval)))
	}

	return strings.Join(intervalsString, " ")
}
