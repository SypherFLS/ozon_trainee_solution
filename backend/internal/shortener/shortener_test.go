package shortener

import (
	"testing"
)

func TestGenerateLength(t *testing.T) {
	gen := NewGen()
	result, err := gen.Generate()

	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}
	if len(result) != 10 {
		t.Errorf("Generate() length = %d, want 10", len(result))
	}
}

func TestGenerateUnique(t *testing.T) {
	gen := NewGen()
	result1, _ := gen.Generate()
	result2, _ := gen.Generate()

	if result1 == result2 {
		t.Errorf("Generate() produced duplicates: %s", result1)
	}
}

func TestGenerateMultiple(t *testing.T) {
	gen := NewGen()
	generated := make(map[string]bool)

	for i := 0; i < 100; i++ {
		result, err := gen.Generate()
		if err != nil {
			t.Fatalf("Generate() error on iteration %d: %v", i, err)
		}
		if generated[result] {
			t.Errorf("Generate() produced duplicate on iteration %d: %s", i, result)
		}
		generated[result] = true
	}

	if len(generated) != 100 {
		t.Errorf("Generate() got %d unique strings, want 100", len(generated))
	}
}

func TestGenerateValidCharacters(t *testing.T) {
	gen := NewGen()
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

	for i := 0; i < 50; i++ {
		result, err := gen.Generate()
		if err != nil {
			t.Fatalf("Generate() error: %v", err)
		}

		for _, char := range result {
			found := false
			for _, validChar := range validChars {
				if char == validChar {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Generate() contains invalid character: %c", char)
			}
		}
	}
}

func Benchmark_Gen(b *testing.B) {
	gen := NewGen()
	for i := 0; i < b.N; i++ {
		gen.Generate()
	}
}
