package engine

var (
	Ab = Note{Code: 11, Letter: "A", Flat: true}
	A  = Note{Code: 0, Letter: "A"}
	Ax = Note{Code: 1, Letter: "A", Sharp: true}

	Bb = Note{Code: 1, Letter: "B", Flat: true}
	B  = Note{Code: 2, Letter: "B"}
	Bx = Note{Code: 3, Letter: "B", Sharp: true}

	Cb = Note{Code: 2, Letter: "C", Flat: true}
	C  = Note{Code: 3, Letter: "C"}
	Cx = Note{Code: 4, Letter: "C", Sharp: true}

	Db = Note{Code: 4, Letter: "D", Flat: true}
	D  = Note{Code: 5, Letter: "D"}
	Dx = Note{Code: 6, Letter: "D", Sharp: true}

	Eb = Note{Code: 6, Letter: "E", Flat: true}
	E  = Note{Code: 7, Letter: "E"}
	Ex = Note{Code: 8, Letter: "E", Sharp: true}

	Fb = Note{Code: 7, Letter: "F", Flat: true}
	F  = Note{Code: 8, Letter: "F"}
	Fx = Note{Code: 9, Letter: "F", Sharp: true}

	Gb = Note{Code: 9, Letter: "G", Flat: true}
	G  = Note{Code: 10, Letter: "G"}
	Gx = Note{Code: 11, Letter: "G", Sharp: true}

	EnharmonicEquivalentAccidentals = map[Note]Note{
		B:  Cb,
		C:  Bx,
		E:  Fb,
		F:  Ex,
		Bx: C,
		Cb: B,
		Ex: F,
		Fb: E,
	}

	// TODO remove this
	NoteLetterToCode = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"G": 6,
		"F": 5,
	}
)

type Note struct {
	Code   int
	Letter string
	Sharp  bool
	Flat   bool
}

func (note Note) AscendingDistance(to Note) Interval {
	if to.Code < note.Code {
		return Interval((MusicalNotesCount + to.Code) - note.Code)
	}

	return Interval(to.Code - note.Code)
}

func (note Note) Add(interval Interval) Note {
	if (int(interval) + note.Code) < MusicalNotesCount {
		return ChromaticScaleWithSharps[int(interval)+note.Code]
	}

	return ChromaticScaleWithSharps[(int(interval)+note.Code)-MusicalNotesCount]
}

func (note Note) IsFlat() bool {
	return note.Flat
}

func (note Note) String() string {
	if note.Flat {
		return note.Letter + "b"
	}

	if note.Sharp {
		return note.Letter + "#"
	}

	return note.Letter
}

func (note Note) IsSharp() bool {
	return note.Sharp
}

func (note Note) EnharmonicEquivalent() Note {
	if eqv, ok := EnharmonicEquivalentAccidentals[note]; ok {
		return eqv
	}

	return note
}
