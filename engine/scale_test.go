package engine

import (
	"reflect"
	"testing"
)

func TestTestScaleConstructors(t *testing.T) {
	cases := []struct {
		name        string
		key         Note
		constructor func(Note) Scale
		expected    Intervals
	}{
		{
			"C major scale",
			C,
			NewMajorScale,
			MajorScaleIntervals,
		},
		{
			"C minor scale",
			C,
			NewMinorScale,
			MinorScaleIntervals,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			scale := tc.constructor(tc.key)

			eq := reflect.DeepEqual(
				[7]Interval(scale.Intervals),
				[7]Interval(tc.expected),
			)

			if !eq {
				t.Errorf("expected '%v', got '%v'", scale.Intervals, MajorScaleIntervals)
			}
		})
	}
}

func TestScaleRelativeMinor(t *testing.T) {
	testCases := []struct {
		desc     string
		majorKey Note
		minorKey Note
	}{
		{
			desc:     "C relative minor",
			majorKey: C,
			minorKey: A,
		},
		{
			desc:     "C# relative minor",
			majorKey: Cx,
			minorKey: Ax,
		},
		{
			desc:     "A relative minor",
			majorKey: A,
			minorKey: Fx,
		},
		{
			desc:     "F relative minor",
			majorKey: Fx,
			minorKey: Dx,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			majorScale := NewMajorScale(tC.majorKey)
			minorScale := majorScale.RelativeMinor()

			if minorScale.Key != tC.minorKey {
				t.Errorf("expected '%v', got '%v'", tC.minorKey, minorScale.Key)
			}

			intervalEq := reflect.DeepEqual(
				[7]Interval(minorScale.Intervals),
				[7]Interval(MinorScaleIntervals),
			)

			if !intervalEq {
				t.Errorf("expected '%v', got '%v'", MinorScaleIntervals, minorScale.Intervals)
			}
		})
	}
}

func TestScaleRelativeMajor(t *testing.T) {
	testCases := []struct {
		desc     string
		minorKey Note
		majorKey Note
	}{
		{
			desc:     "C relative major",
			minorKey: A,
			majorKey: C,
		},
		{
			desc:     "C# relative major",
			minorKey: Ax,
			majorKey: Cx,
		},
		{
			desc:     "A relative major",
			minorKey: Fx,
			majorKey: A,
		},
		{
			desc:     "F relative major",
			minorKey: Dx,
			majorKey: Fx,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			minorScale := NewMinorScale(tC.minorKey)
			majorScale := minorScale.RelativeMajor()

			if minorScale.Key != tC.minorKey {
				t.Errorf("expected '%v', got '%v'", tC.majorKey, majorScale.Key)
			}

			intervalEq := reflect.DeepEqual(
				[7]Interval(majorScale.Intervals),
				[7]Interval(MajorScaleIntervals),
			)

			if !intervalEq {
				t.Errorf("expected '%v', got '%v'", MajorScaleIntervals, majorScale.Intervals)
			}
		})
	}
}

func TestMajorScaleCountSharps(t *testing.T) {
	testCases := []struct {
		desc           string
		key            Note
		expectedSharps int
	}{
		{
			desc:           "C sharps count",
			key:            C,
			expectedSharps: 0,
		},
		{
			desc:           "G sharps count",
			key:            G,
			expectedSharps: 1,
		},
		{
			desc:           "B sharps count",
			key:            B,
			expectedSharps: 5,
		},
		// TODO fix this case
		// {
		// 	desc:           "Gb sharps count",
		// 	key:            Gb,
		// 	expectedSharps: 6,
		// },
		// {
		// 	desc:           "F# sharps count",
		// 	key:            Fx,
		// 	expectedSharps: 6,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			scale := NewMajorScale(tC.key)

			if sharpsCount := scale.CountSharps(); sharpsCount != tC.expectedSharps {
				t.Errorf("expected '%v', got '%v'", tC.expectedSharps, sharpsCount)
			}
		})
	}
}

func TestMajorScaleCountFlats(t *testing.T) {
	testCases := []struct {
		desc          string
		key           Note
		expectedFlats int
	}{
		{
			desc:          "C flats count",
			key:           C,
			expectedFlats: 0,
		},
		{
			desc:          "F flats count",
			key:           F,
			expectedFlats: 1,
		},
		{
			desc:          "F flats count",
			key:           Cx,
			expectedFlats: 5,
		},
		{
			desc:          "F flats count",
			key:           Db,
			expectedFlats: 5,
		},
		// TODO fix this case
		// {
		// 	desc:          "Fx flats count",
		// 	key:           Fx,
		// 	expectedFlats: 6,
		// },
		// {
		// 	desc:          "Gb flats count",
		// 	key:           Gb,
		// 	expectedFlats: 6,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			scale := NewMajorScale(tC.key)

			if flatsCount := scale.CountFlats(); flatsCount != tC.expectedFlats {
				t.Errorf("expected '%v', got '%v'", tC.expectedFlats, flatsCount)
			}
		})
	}
}
