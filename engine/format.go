package engine

import "fmt"

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

const (
	defaultFlatSymbol  = "♭"
	defaultSharpSymbol = "♯"
)

var (
	defaultFormatter = chordFormatter{
		flatNotation:  defaultFlatSymbol,
		sharpNotation: defaultSharpSymbol,
		chordFormats: map[ChordFormula]Format{
			PowerChordIntervals.Formula():            powerChordFormat,
			MajorTriadIntervals.Formula():            regularTriadFormat,
			MinorTriadIntervals.Formula():            regularTriadFormat,
			DiminishedTriadIntervals.Formula():       regularTriadFormat,
			AugmentedTriadIntervals.Formula():        regularTriadFormat,
			Sus2TriadIntervals.Formula():             sus2TriadFormat,
			Sus4TriadIntervals.Formula():             sus4TriadFormat,
			MajorSeventhIntervals.Formula():          seventhChordFormat,
			MinorSeventhIntervals.Formula():          seventhChordFormat,
			DominantSeventhIntervals.Formula():       dominantSeventhChordFormat,
			HalfDiminishedSeventhIntervals.Formula(): halfDiminishedSeventhChordFormat,
		},
	}
)

type (
	Format func(chord Chord) string

	chordFormatter struct {
		chordFormats  map[ChordFormula]Format
		sharpNotation string
		flatNotation  string
	}
)

func SetChordFormat(chordType ChordFormula, format Format) chordFormatter {
	return defaultFormatter.setFormat(chordType, format)
}

func (f chordFormatter) setFormat(chordType ChordFormula, format Format) chordFormatter {
	f.chordFormats[chordType] = format
	return f
}

func SetSharpNotation(s string) chordFormatter {
	return defaultFormatter.setSharpNotation(s)
}

func (f chordFormatter) setSharpNotation(s string) chordFormatter {
	f.sharpNotation = s
	return f
}

func SetFlatNotation(s string) chordFormatter {
	return defaultFormatter.setFlatNotation(s)
}

func (f chordFormatter) setFlatNotation(s string) chordFormatter {
	f.flatNotation = s
	return f
}

func regularTriadFormat(chord Chord) string {
	return fmt.Sprintf("%s %s", chord.Root(), chord.Quality())
}

func sus2TriadFormat(chord Chord) string {
	return fmt.Sprintf("%ssus2", chord.Root())
}

func sus4TriadFormat(chord Chord) string {
	return fmt.Sprintf("%ssus4", chord.Root())
}

func seventhChordFormat(chord Chord) string {
	return fmt.Sprintf("%s %s 7", chord.Root(), chord.Quality())
}

func dominantSeventhChordFormat(chord Chord) string {
	return fmt.Sprintf("%s7", chord.Root())
}

func halfDiminishedSeventhChordFormat(chord Chord) string {
	return fmt.Sprintf("%sm7b5", chord.Root())
}

func powerChordFormat(chord Chord) string {
	return fmt.Sprintf("%s5", chord.Root())
}
