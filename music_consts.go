package main

// Universal values
const (
	MusicalNotesCount = 12
	HalfStep          = 1
	WholeStep         = 2
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
	MinorSeventh
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
	regularTriadFormat               string = "%s %s"
	sus2TriadFormat                  string = "%ssus2"
	sus4TriadFormat                  string = "%ssus4"
	seventhChordFormat               string = "%s %s 7"
	dominantSeventhChordFormat       string = "%s7"
	halfDiminishedSeventhChordFormat string = "%sm7b5"
	powerChordFormat                        = "%s5"

	// TODO
	// Major w/ added 9th: 1 - 3 - 5 - 9
	// Minor w/ added 9th: 1 - b3 - 5 - 9
	// Major w/ added 6th: 1 - 3 - 5 - 6
	// Minor w/ added 6th: 1 - b3 - 5 - 6
	// Major w/ added 6th & 9th: 1 - 3 - 5 - 6 - 9
	// Minor w/ added 6th & 9th: 1 - b3 - 5 - 6 - 9
	// 7th flat five: 1 - b3 - b5 - b7
	// Diminished 7th: 1 - b3 - b5 - bb7
	// 7th suspended 4th: 1 - 4 - 5 - b7
	// Minor major 7th: 1 - b3 - 5 - 7
	// Major 9th: 1 - 3 - 5 - 7 -  9
	// Minor 9th: 1 - b3 - 5 - b7 - 9
	// Dominant 9th: 1 - 3 - 5 - b7 - 9
	// 9th suspended 4th: 1 - 4 - 5 - b7 - 9
	// Minor 11th: 1 - b3 - 5 - b7 - 9 - 11
	// 11th: 1 - 3 - 5 - b7 - 9 - 11
	// Major 13th: 1 - 3 - 5 - 7 - 9 - 13
	// Minor 13th: 1 - b3 - 5 - b7 - 9 - 11 - 13
	// 13th: 1 - 3 - 5 - b7 - 9 - 13
)
