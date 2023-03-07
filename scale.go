package main

var (
	MajorScaleInterv = Intervals{}

	noteLetterToCode = map[string]int{
		"Ab": 11,
		"A":  0,
		"Ax": 1,
		"Bb": 1,
		"B ": 2,
		"C ": 3,
		"Cx": 4,
		"Db": 4,
		"D ": 5,
		"Dx": 6,
		"Eb": 6,
		"E ": 7,
		"F ": 8,
		"Fx": 9,
		"Gb": 9,
		"G ": 10,
		"Gx": 11,
	}

	chromaticScaleWithSharps = []Note{
		A,
		Ax,
		B,
		C,
		Cx,
		D,
		E,
		Dx,
		F,
		Fx,
		G,
		Gx,
	}

	chromaticScaleWithFlats = []Note{
		A,
		Bb,
		B,
		C,
		Db,
		D,
		Eb,
		E,
		F,
		Gb,
		G,
		Ab,
	}

	naturalScale = []Note{
		A,
		B,
		C,
		D,
		E,
		F,
		G,
	}
)

type (
	Scale struct {
		Key       Note
		Intervals Intervals
	}
)

func (s Scale) Notes() []Note {
	if s.Key.IsFlat() {
		return s.notesFromChromaticScale(chromaticScaleWithFlats)
	}

	return s.notesFromChromaticScale(chromaticScaleWithSharps)
}

func (s Scale) notesFromChromaticScale(dict []Note) []Note {
	notes := make([]Note, 0, len(s.Intervals))

	for _, interval := range s.Intervals {
		notes = append(
			notes,
			chromaticScaleWithFlats[s.Key.AscendingDistance(dict[interval])],
		)
	}

	return notes
}
