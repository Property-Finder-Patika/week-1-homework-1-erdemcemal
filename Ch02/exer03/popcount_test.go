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

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithLoop(0)
		PopCountWithLoop(1)
		PopCountWithLoop(1 << 8)
		PopCountWithLoop(1<<8 + 1)
		PopCountWithLoop(1<<8 - 1)
		PopCountWithLoop(1<<64 - 1)
	}
}
