package shortener


import (
	"testing"
)

func Test_Gen(t *testing.T) {
	gen := NewGen()
	result, _ := gen.Generate()
	t.Logf("result is %v", result)
}

func Benchmark_Gen(b *testing.B) {
	gen := NewGen()
	gen.Generate()
}