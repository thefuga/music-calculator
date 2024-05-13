package fretboard

import (
	"github.com/thefuga/music-calculator/engine"
)

type (
	Chord []int

	Fretboard struct {
		Strings []String
	}

	String struct {
		Tune  engine.Note
		Index int
	}
)

func (fretboard Fretboard) BuildChord(chord engine.Chord, lowerFret, higherFret int) Chord {
	positions := initPositions(len(fretboard.Strings))

	bass, pos := fretboard.FindBass(chord, lowerFret, higherFret)
	if pos < 0 {
		panic("bass not found within this range")
	}
	positions[bass.Index] = pos

	for atFret := lowerFret; atFret <= higherFret; atFret++ {
		for _, atString := range fretboard.Strings[bass.Index+1:] {
			for _, atNote := range chord.Notes {
				if found := atString.PositionEqualsNote(atNote, atFret); found {
					positions[atString.Index] = atFret
					break
				}
			}
		}
	}

	// for _, str := range fretboard.Strings[bass.Index+1:] {
	// 	for _, note := range chord.Notes {
	// 		if pos, found := str.FindNotePosition(note, lower, higher); found {
	// 			positions[str.Index] = pos
	// 			break
	// 		}
	// 	}
	// }

	return positions
}

func initPositions(strings int) Chord {
	var positions = make(Chord, strings)

	for i := range positions {
		positions[i] = -1
	}

	return positions
}

func (fretboard Fretboard) FindBass(chord engine.Chord, lower, higher int) (String, int) {
	for _, str := range fretboard.Strings {
		pos, found := str.FindNotePosition(chord.Root(), lower, higher)
		if !found {
			continue
		}

		return str, pos
	}

	return String{}, -1
}

func (str String) FindNotePosition(note engine.Note, lower, higher int) (int, bool) {
	for at := lower; at <= higher; at++ {
		foundNote := str.Tune.Add(engine.Interval(at))

		if foundNote == note {
			return at, true
		}
	}

	return -1, false
}

func (str String) PositionEqualsNote(note engine.Note, at int) bool {
	return str.Tune.Add(engine.Interval(at)) == note
}

// func (chord Chord) String() string {
// 	display := make([][]string, 6)

// 	for i := 0; i < 6; i++ {
// 		if i == 0 {
// 			display[i] = []string{"---", "---", "---", "---", "---", "---"}
// 		} else {
// 			display[i] = []string{" | ", " | ", " | ", " | ", " | ", " | "}
// 		}
// 	}

// 	for i, pos := range chord {
// 		if pos != -1 {
// 			if pos == 0 {
// 				display[pos][i] = "-O-"
// 			} else {
// 				display[int(math.Mod(float64(pos), 6.0))][i] = " O "
// 			}
// 		}
// 	}

// 	var str string
// 	for i := 0; i < len(display); i++ {
// 		if i > 0 {
// 			str += strconv.Itoa(i)
// 		}
// 		str += "\t" + strings.Join(display[i], "")
// 		str += "\n"

// 		if i > 0 {
// 			str += "\t-----------------\n"
// 		}
// 	}

// 	return str
// }
