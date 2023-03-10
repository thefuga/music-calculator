package engine

// Universal values
const (
	MusicalNotesCount = 12
	HalfStep          = 1
	WholeStep         = 2
)

// Intervals

const NilInterval = -1

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
