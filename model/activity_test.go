package model

import (
	"testing"
)

func BenchmarkCalcNP(b *testing.B) {
	var tp []UnitData

	for i := 0; i < 10000; i++ {
		tp = append(tp, UnitData{Watts: 180})
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CalcNP(&tp)
	}
}

func BenchmarkParallelCalcNP(b *testing.B) {
	var tp []UnitData

	for i := 0; i < 10000; i++ {
		tp = append(tp, UnitData{Watts: 180})
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ParallelCalcNP(&tp, 5)
	}
}
