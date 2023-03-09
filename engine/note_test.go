package engine

import "testing"

func TestEnharmonicEquivalents(t *testing.T) {
	testCases := []struct {
		desc       string
		note       Note
		equivalent Note
	}{
		{
			desc:       "B enharmonic equivalent",
			note:       B,
			equivalent: Cb,
		},
		{
			desc:       "C enharmonic equivalent",
			note:       C,
			equivalent: Bx,
		},
		{
			desc:       "E enharmonic equivalent",
			note:       E,
			equivalent: Fb,
		},
		{
			desc:       "F enharmonic equivalent",
			note:       F,
			equivalent: Ex,
		},
		{
			desc:       "G enharmonic equivalent",
			note:       G,
			equivalent: G,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			eqv := tC.note.EnharmonicEquivalent()

			if eqv != tC.equivalent {
				t.Errorf("expected '%v', got '%v'", tC.equivalent, eqv)
			}

			eqv = eqv.EnharmonicEquivalent()
			if eqv != tC.note {
				t.Errorf("expected '%v', got '%v'", tC.note, eqv)
			}
		})
	}
}
