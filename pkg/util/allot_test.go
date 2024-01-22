package util

import "testing"

func TestValueOrDefault(t *testing.T) {
	testdata := []struct {
		v, def, expected int
	}{
		{v: 1, def: 2, expected: 1},
		{v: 0, def: 2, expected: 2},
	}
	for _, tt := range testdata {
		actual := ValueIf(tt.v != 0, tt.v, tt.def)
		if actual != tt.expected {
			t.Errorf("TestValueOrDefault(%v, %v): expected %v, actual %v", tt.v, tt.def, tt.expected, actual)
		}
	}
}
