package mesgdef

import (
	"testing"

	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
)

func BenchmarkNewRecord(b *testing.B) {
	mesg := factory.CreateMesg(mesgnum.Record)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NewRecord(&mesg)
	}
}

func BenchmarkRecordToMesg(b *testing.B) {
	mesg := factory.CreateMesg(mesgnum.Record)
	record := NewRecord(&mesg)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = record.ToMesg(nil)
	}
}
