package ids

import (
	"testing"
)

func BenchmarkIds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}
