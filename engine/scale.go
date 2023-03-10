package engine

import (
	"math"
)

var (
	MajorScaleIntervals = Intervals{
		PerfectUnison,
		MajorSecond,
		MajorThird,
		PerfectFourth,
		PerfectFifth,
		MajorSixth,
		MajorSeventh,
	}

	MinorScaleIntervals = Intervals{
		PerfectUnison,
		MajorSecond,
		MinorThird,
		PerfectFourth,
		PerfectFifth,
		MinorSixth,
		MinorSeventh,
	}

	PentatonicMajorScale = Intervals{
		PerfectUnison,
		MajorSecond,
		MajorThird,
		PerfectFourth,
		PerfectFifth,
		MajorSixth,
		MajorSeventh,
	}

	ChromaticScaleWithSharps = []Note{A, Ax, B, C, Cx, D, Dx, E, F, Fx, G, Gx}
	ChromaticScaleWithFlats  = []Note{A, Bb, B, C, Db, D, Eb, E, F, Gb, G, Ab}
)

type (
	Scale struct {
		Key       Note
		Intervals Intervals
	}
)

func NewMajorScale(key Note) Scale {
	return Scale{
		Key:       key,
		Intervals: MajorScaleIntervals,
	}
}

func NewMinorScale(key Note) Scale {
	return Scale{
		Key:       key,
		Intervals: MinorScaleIntervals,
	}
}

func (s Scale) Notes() []Note {
	if s.Key.IsFlat() {
		return s.NotesWithFlats()
	}

	return s.NotesWithSharps()
}

func (s Scale) NotesWithSharps() []Note {
	return s.notesFromChromaticScale(ChromaticScaleWithSharps)
}

func (s Scale) NotesWithFlats() []Note {
	return s.notesFromChromaticScale(ChromaticScaleWithFlats)
}

func (s Scale) RelativeMinor() Scale {
	return NewMinorScale(s.Notes()[5])
}

func (s Scale) RelativeMajor() Scale {
	return NewMajorScale(s.Notes()[2])
}

func (s Scale) CountSharps() (sharps int) {
	if s.Key.IsFlat() {
		panic("cannot count sharps on a scale with a flat key. Apply the enharmonic equivalent first.")
	}

	for _, note := range s.NotesWithSharps() {
		if note.IsSharp() {
			sharps++
		}
	}

	return
}

func (s Scale) CountFlats() (flats int) {
	if s.Key.IsSharp() {
		panic("cannot count flats on a scale with a sharp key. Apply the enharmonic equivalent first.")
	}

	for _, note := range s.NotesWithFlats() {
		if note.IsFlat() {
			flats++
		}
	}

	return
}

func (s Scale) notesFromChromaticScale(dict []Note) []Note {
	notes := make([]Note, 0, len(s.Intervals))

	notes = append(notes, s.Key)
	for _, interval := range s.Intervals[1:] {
		var note Note
		if (int(interval) + s.Key.Code) < MusicalNotesCount {
			note = dict[int(interval)+s.Key.Code]
		} else {
			note = dict[(int(interval)+s.Key.Code)-MusicalNotesCount]
		}

		notes = append(notes, note)
	}

	return distinctNotes(notes)
}

func distinctNotes(notes []Note) []Note {
	seen := make(map[rune]struct{})

	for i := range notes {
		if _, ok := seen[notes[i].Symbol]; ok {
			if _, ok := EnharmonicEquivalentAccidentals[notes[i]]; ok {
				notes[i] = EnharmonicEquivalentAccidentals[notes[i]]
			} else {
				notes[i-1] = EnharmonicEquivalentAccidentals[notes[i-1]]
				seen[notes[i].Symbol] = struct{}{}
			}
		} else {
			seen[notes[i].Symbol] = struct{}{}
		}
	}

	return notes
}

func (s Scale) DiatonicTriads() []Chord {
	var triads []Chord

	notes := s.Notes()

	for i := range notes {
		root := notes[i]

		third := notes[int(math.Mod(float64(i+2), float64(len(notes))))]
		fifth := notes[int(math.Mod(float64(i+4), float64(len(notes))))]

		triads = append(triads, BuildChord(root, third, fifth))
	}

	return triads
}

func (s Scale) DiatonicQuadrads() []Chord {
	var quadrads []Chord

	notes := s.Notes()

	for i := range notes {
		root := notes[i]
		third := notes[int(math.Mod(float64(i+2), float64(len(notes))))]
		fifth := notes[int(math.Mod(float64(i+4), float64(len(notes))))]
		seventh := notes[int(math.Mod(float64(i+6), float64(len(notes))))]

		quadrads = append(quadrads, BuildChord(root, third, fifth, seventh))
	}

	return quadrads
}
