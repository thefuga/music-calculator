package engine

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
		Type      int
		Name      string
		Notes     []Note
		Intervals Intervals
	}

	ChordFormula string
)

func BuildChord(notes ...Note) Chord {
	if len(notes) < 2 {
		panic("not enough notes")
	}

	chord := Chord{Notes: notes}

	for _, note := range chord.Notes {
		chord.Intervals = append(
			chord.Intervals,
			chord.Root().AscendingDistance(note),
		)
	}

	return chord
}

func (chord Chord) String() string {
	return defaultFormatter.chordFormats[chord.Intervals.Formula()](chord)
}

func (chord Chord) Root() Note {
	return chord.Notes[0]
}

func (chord Chord) Third() Note {
	if len(chord.Intervals) < 3 {
		return NilNote
	}

	return chord.Notes[1]
}

func (chord Chord) Quality() string {
	switch chord.Root().AscendingDistance(chord.Third()) {
	case NilInterval, MajorThird:
		return MajorQuality
	case MinorThird:
		return MinorQuality
	case MajorSecond, PerfectFourth: // TODO check this
		return MajorQuality
	default:
		panic("unknown quality")
	}
}

func (chord Chord) NotesCount() int {
	return len(chord.Notes)
}
