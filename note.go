package main

import (
	"strings"
)

var (
	Ab = Note{Code: 11, Letter: "A", Flat: true}
	A  = Note{Code: 0, Letter: "A"}
	Ax = Note{Code: 1, Letter: "A", Sharp: true}
	Bb = Note{Code: 1, Letter: "B", Flat: true}
	B  = Note{Code: 2, Letter: "B"}
	C  = Note{Code: 3, Letter: "C"}
	Cx = Note{Code: 4, Letter: "C", Sharp: true}
	Db = Note{Code: 4, Letter: "D", Flat: true}
	D  = Note{Code: 5, Letter: "D"}
	Dx = Note{Code: 6, Letter: "D", Sharp: true}
	Eb = Note{Code: 6, Letter: "E", Flat: true}
	E  = Note{Code: 7, Letter: "E"}
	F  = Note{Code: 8, Letter: "F"}
	Fx = Note{Code: 9, Letter: "F", Sharp: true}
	Gb = Note{Code: 9, Letter: "G", Flat: true}
	G  = Note{Code: 10, Letter: "G"}
	Gx = Note{Code: 11, Letter: "G", Sharp: true}

	noteLetterToCode = map[string]int{
		"A":  0,
		"Ax": 1,
		"Bb": 1,
		"B":  2,
		"C":  3,
		"Cx": 4,
		"Db": 4,
		"D":  5,
		"Dx": 6,
		"Eb": 6,
		"E":  7,
		"F":  8,
		"Fx": 9,
		"Gb": 9,
		"G":  10,
		"Ab": 11,
		"Gx": 11,
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
		return chromaticScaleWithSharps[int(interval)+note.Code]
	}

	return chromaticScaleWithSharps[(int(interval)+note.Code)-MusicalNotesCount]
}

func (note Note) IsFlat() bool {
	return strings.Contains(note.Letter, "b")
}

func (note Note) String() string {
	str := note.Letter

	if note.Flat {
		str += "b"
	}

	if note.Sharp {
		str += "#"
	}

	return str
}

func (note Note) IsSharp() bool {
	return strings.Contains(note.Letter, "#")
}
