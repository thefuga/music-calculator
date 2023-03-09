package engine

import (
	"fmt"
)

var (
	PowerChordIntervals            = Intervals{PerfectUnison, PerfectFifth}
	MajorTriadIntervals            = Intervals{PerfectUnison, MajorThird, PerfectFifth}
	MinorTriadIntervals            = Intervals{PerfectUnison, MinorThird, PerfectFifth}
	DiminishedTriadIntervals       = Intervals{PerfectUnison, MinorThird, Tritone}
	AugmentedTriadIntervals        = Intervals{PerfectUnison, MajorThird, AugmentedFifth}
	Sus2TriadIntervals             = Intervals{PerfectUnison, MajorSecond, PerfectFifth}
	Sus4TriadIntervals             = Intervals{PerfectUnison, PerfectFourth, PerfectFifth}
	MajorSeventhIntervals          = Intervals{PerfectUnison, MajorThird, PerfectFifth, MajorSeventh}
	MinorSeventhIntervals          = Intervals{PerfectUnison, MinorThird, PerfectFifth, MinorSeventh}
	DominantSeventhIntervals       = Intervals{PerfectUnison, MajorThird, PerfectFifth, MinorSeventh}
	HalfDiminishedSeventhIntervals = Intervals{PerfectUnison, MinorThird, Tritone, MinorSeventh}
)

type (
	Chord struct {
		Name          string
		Notes         []Note
		Quality       string
		Type          int
		Intervals     Intervals
		FormulaFormat string
		formatArgs    []interface{}
	}
)

func NewChord(notes ...Note) Chord {
	if notes == nil {
		panic("empty notes")
	}

	chord := Chord{Notes: notes}

	for _, note := range chord.Notes {
		chord.Intervals = append(
			chord.Intervals,
			chord.Root().AscendingDistance(note),
		)
	}

	switch chord.Intervals.String() {
	case PowerChordIntervals.String():
		chord.FormulaFormat = powerChordFormat
		chord.Quality = MajorQuality
		chord.formatArgs = []interface{}{chord.Root()}
	case MajorTriadIntervals.String():
		chord.FormulaFormat = regularTriadFormat
		chord.Quality = MajorQuality
		chord.formatArgs = []interface{}{chord.Root(), chord.Quality}
	case MinorTriadIntervals.String():
		chord.Quality = MinorQuality
		chord.FormulaFormat = regularTriadFormat
		chord.formatArgs = []interface{}{chord.Root(), chord.Quality}
	case DiminishedTriadIntervals.String():
		chord.Quality = DiminishedQuality
		chord.FormulaFormat = regularTriadFormat
		chord.formatArgs = []interface{}{chord.Root(), chord.Quality}
	case AugmentedTriadIntervals.String():
		chord.Quality = AugmentedQuality
		chord.FormulaFormat = regularTriadFormat
		chord.formatArgs = []interface{}{chord.Root(), chord.Quality}
	case Sus2TriadIntervals.String():
		chord.Quality = MajorQuality // TODO check this
		chord.FormulaFormat = sus2TriadFormat
		chord.formatArgs = []interface{}{chord.Root()}
	case Sus4TriadIntervals.String():
		chord.Quality = MajorQuality // TODO check this
		chord.FormulaFormat = sus4TriadFormat
		chord.formatArgs = []interface{}{chord.Root()}
	case MajorSeventhIntervals.String():
		chord.Quality = MajorQuality
		chord.FormulaFormat = seventhChordFormat
		chord.formatArgs = []interface{}{chord.Root(), chord.Quality}
	case MinorSeventhIntervals.String():
		chord.Quality = MinorQuality
		chord.FormulaFormat = seventhChordFormat
		chord.formatArgs = []interface{}{chord.Root(), chord.Quality}
	case DominantSeventhIntervals.String():
		chord.Quality = MajorQuality
		chord.FormulaFormat = dominantSeventhChordFormat
		chord.formatArgs = []interface{}{chord.Root()}
	case HalfDiminishedSeventhIntervals.String():
		chord.Quality = DiminishedQuality
		chord.FormulaFormat = halfDiminishedSeventhChordFormat
		chord.formatArgs = []interface{}{chord.Root()}
	default:
		panic("unknown formula")
	}

	return chord
}

func (chord Chord) String() string {
	return fmt.Sprintf(
		chord.FormulaFormat,
		chord.formatArgs...,
	)
}

func (chord Chord) Root() Note {
	return chord.Notes[0]
}

func (chord Chord) NotesCount() int {
	return len(chord.Notes)
}
