package main

var (
	A  = Note{0, "A"}
	Ax = Note{1, "A#"}
	B  = Note{2, "B"}
	C  = Note{3, "C"}
	Cx = Note{4, "C#"}
	D  = Note{5, "D"}
	Dx = Note{6, "D#"}
	E  = Note{7, "E"}
	F  = Note{8, "F"}
	Fx = Note{9, "F#"}
	G  = Note{10, "G"}
	Gx = Note{11, "G#"}
)

type Note struct {
	Code int
	Name string
}

func (note Note) AscendingDistance(to Note) Interval {
	if to.Code < note.Code {
		return Interval((MusicalNotesCount + to.Code) - note.Code)
	}

	return Interval(to.Code - note.Code)
}
