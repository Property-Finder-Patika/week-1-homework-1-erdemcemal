package exer03

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0)
		PopCount(1)
		PopCount(1 << 8)
		PopCount(1<<8 + 1)
		PopCount(1<<8 - 1)
		PopCount(1<<64 - 1)
	}
}

func BenchmarkBitShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitShiftPopCount(0)
		BitShiftPopCount(1)
		BitShiftPopCount(1 << 8)
		BitShiftPopCount(1<<8 + 1)
		BitShiftPopCount(1<<8 - 1)
		BitShiftPopCount(1<<64 - 1)
	}
}
