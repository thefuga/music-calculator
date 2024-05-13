package main

import (
	"fmt"
	"os"
	"strconv"

	. "github.com/thefuga/music-calculator/engine"
	"github.com/thefuga/music-calculator/engine/fretboard"
)

var NoteLetterToCode = map[string]Note{
	"Ab": Ab,
	"A":  A,
	"A#": Ax,

	"Bb": Bb,
	"B":  B,
	"Bx": Bx,

	"Cb": Cb,
	"C":  C,
	"C#": Cx,

	"Db": Db,
	"D":  D,
	"D#": Dx,

	"Eb": Eb,
	"E":  E,
	"Ex": Ex,

	"Fb": Fb,
	"F":  F,
	"F#": Fx,

	"Gb": Gb,
	"G":  G,
	"G#": Gx,
}

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "chord":
		var notes []Note

		for _, note := range args[1:] {
			notes = append(notes, NoteLetterToCode[note])
		}

		chord := BuildChord(notes...)
		fmt.Println(chord)
	case "guitar-chord":
		fretboard := fretboard.Fretboard{
			Strings: []fretboard.String{
				// {Tune: E, Index: 0},
				// {Tune: A, Index: 1},
				{Tune: D, Index: 0},
				{Tune: G, Index: 1},
				{Tune: B, Index: 2},
				{Tune: E, Index: 3},
			},
		}

		var notes []Note

		for _, note := range args[3:] {
			notes = append(notes, NoteLetterToCode[note])
		}

		chord := BuildChord(notes...)
		from, _ := strconv.Atoi(args[1])
		to, _ := strconv.Atoi(args[2])

		fmt.Println(fretboard.BuildChord(chord, from, to))
	case "triads":
		key := NoteLetterToCode[args[1]]
		scale := NewMajorScale(key)
		fmt.Println(scale.DiatonicTriads())
	case "quadrads":
		key := NoteLetterToCode[args[1]]
		scale := NewMajorScale(key)
		fmt.Println(scale.DiatonicQuadrads())
	case "major-scale":
		key := NoteLetterToCode[args[1]]
		scale := NewMajorScale(key)
		fmt.Println(scale.Notes())
	case "relative-minor":
		key := NoteLetterToCode[args[1]]
		scale := NewMajorScale(key).RelativeMinor()
		fmt.Println(scale.Notes())
	case "minor-scale":
		key := NoteLetterToCode[args[1]]
		scale := NewMinorScale(key)
		fmt.Println(scale.Notes())
	case "relative-major":
		key := NoteLetterToCode[args[1]]
		scale := NewMinorScale(key).RelativeMajor()
		fmt.Println(scale.Notes())
	case "relative-minor-scale":
		key := NoteLetterToCode[args[1]]
		scale := NewMajorScale(key).RelativeMinor()
		fmt.Println(scale.Notes())
	case "relative-major-scale":
		key := NoteLetterToCode[args[1]]
		scale := NewMinorScale(key).RelativeMajor()
		fmt.Println(scale.Notes())
	case "major-pentatonic":
		key := NoteLetterToCode[args[1]]
		scale := NewMajorPentatonicScale(key)
		fmt.Println(scale.Notes())
	case "minor-pentatonic":
		key := NoteLetterToCode[args[1]]
		scale := NewMinorPentatonicScale(key)
		fmt.Println(scale.Notes())
	default:
		fmt.Println("invalid option")
	}
}
