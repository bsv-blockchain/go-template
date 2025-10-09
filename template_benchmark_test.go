package template_test

import (
	"testing"

	"github.com/bsv-blockchain/go-template"
)

// BenchmarkGreet benchmarks the Greet function.
func BenchmarkGreet(b *testing.B) {
	for b.Loop() {
		_ = template.Greet("BenchmarkUser")
	}
}
