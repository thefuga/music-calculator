package engine

import "testing"

func TestChordRoot(t *testing.T) {
	chord := Chord{Notes: []Note{C, E, G}}

	if root := chord.Root(); root != C {
		t.Errorf("expected '%s', got '%s'", C, root)
	}
}

func TestChordNotesCount(t *testing.T) {
	chord := Chord{Notes: []Note{C, E, G}}

	if count := chord.NotesCount(); count != 3 {
		t.Errorf("expected '%d', got '%d'", 3, count)
	}
}
