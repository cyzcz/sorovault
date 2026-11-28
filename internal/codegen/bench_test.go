package codegen_test

import (
	"testing"

	zkvote "github.com/soroworks/sorovault/internal/codegen/testdata"
)

// BenchmarkNewClient measures the cost of constructing a generated Go client.
func BenchmarkNewClient(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = zkvote.NewClient(zkvote.ContractID, zkvote.Network)
	}
}
