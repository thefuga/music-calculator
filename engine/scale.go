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
	for _, note := range s.Notes() {
		if note.IsSharp() {
			sharps++
		}
	}

	return
}

func (s Scale) CountFlats() (flats int) {
	for _, note := range s.NotesWithFlats() {
		if note.IsFlat() {
			flats++
		}
	}

	return
}

func (s Scale) notesFromChromaticScale(dict []Note) []Note {
	notes := make([]Note, 0, len(s.Intervals))

	for _, interval := range s.Intervals {
		var note Note
		if (int(interval) + s.Key.Code) < MusicalNotesCount {
			note = dict[int(interval)+s.Key.Code]
		} else {
			note = dict[(int(interval)+s.Key.Code)-MusicalNotesCount]
		}

		notes = append(notes, note)
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
