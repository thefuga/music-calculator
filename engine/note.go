package engine

var (
	NilNote = Note{Code: -1}

	Ab = Note{Code: 11, Symbol: 'A', Flat: true}
	A  = Note{Code: 0, Symbol: 'A'}
	Ax = Note{Code: 1, Symbol: 'A', Sharp: true}

	Bb = Note{Code: 1, Symbol: 'B', Flat: true}
	B  = Note{Code: 2, Symbol: 'B'}
	Bx = Note{Code: 3, Symbol: 'B', Sharp: true}

	Cb = Note{Code: 2, Symbol: 'C', Flat: true}
	C  = Note{Code: 3, Symbol: 'C'}
	Cx = Note{Code: 4, Symbol: 'C', Sharp: true}

	Db = Note{Code: 4, Symbol: 'D', Flat: true}
	D  = Note{Code: 5, Symbol: 'D'}
	Dx = Note{Code: 6, Symbol: 'D', Sharp: true}

	Eb = Note{Code: 6, Symbol: 'E', Flat: true}
	E  = Note{Code: 7, Symbol: 'E'}
	Ex = Note{Code: 8, Symbol: 'E', Sharp: true}

	Fb = Note{Code: 7, Symbol: 'F', Flat: true}
	F  = Note{Code: 8, Symbol: 'F'}
	Fx = Note{Code: 9, Symbol: 'F', Sharp: true}

	Gb = Note{Code: 9, Symbol: 'G', Flat: true}
	G  = Note{Code: 10, Symbol: 'G'}
	Gx = Note{Code: 11, Symbol: 'G', Sharp: true}

	EnharmonicEquivalentAccidentals = map[Note]Note{
		Ax: Bb,
		B:  Cb,
		C:  Bx,
		E:  Fb,
		F:  Ex,
		Bx: C,
		Cb: B,
		Ex: F,
		Fb: E,
	}
)

type Note struct {
	Code   int
	Symbol rune
	Sharp  bool
	Flat   bool
}

func (note Note) AscendingDistance(to Note) Interval {
	if to == NilNote || note == NilNote {
		return NilInterval
	}

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
		return string(note.Symbol) + defaultFormatter.flatNotation
	}

	if note.Sharp {
		return string(note.Symbol) + defaultFormatter.sharpNotation
	}

	return string(note.Symbol)
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
