package main

// Universal values
const (
	MusicalNotesCount = 12
)

// Intervals
const (
	PerfectUnison = Interval(iota)
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

// Interval Aliases
const (
	AugmentedFifth = MinorSixth
)

const (
	MajorQuality      = "major"
	MinorQuality      = "minor"
	DiminishedQuality = "diminished"
	AugmentedQuality  = "augmented"
)

const (
	regularTriadFormat string = "%s %s"
	sus2TriadFormat    string = "%ssus2"
	sus4TriadFormat    string = "%ssus4"
)
