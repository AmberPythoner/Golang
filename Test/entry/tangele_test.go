package main

import (
	"testing"
)

func TestTrangle(t *testing.T) {
	var tangle_struct = []struct {
		a, b, c int
	}{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
	}
	for _, tt := range tangle_struct {
		if anser := Trangle(tt.a, tt.b); anser != tt.c {
			t.Errorf("%d != input %d", anser, tt.c)
		}
	}
}

func BenchmarkTrangle(b *testing.B) {
	aa, ab, ac := 8, 15, 17
	for i := 0; i < b.N; i++ {
		if anser := Trangle(aa, ab); anser != ac {
			b.Errorf("errors")
		}
	}
}
