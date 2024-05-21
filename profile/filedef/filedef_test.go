package filedef_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestToMesgs(t *testing.T) {
	messages := make([]proto.Message, 0, 1)
	records := make([]*mesgdef.Record, 1)
	records[0] = mesgdef.NewRecord(nil).
		SetTimestamp(time.Now())

	filedef.ToMesgs(&messages, nil, mesgnum.Record, records)
	if len(messages) != 1 {
		t.Fatalf("expected 1: got: %d", len(messages))
	}
}

func TestSortMessagesByTimestamp(t *testing.T) {
	now := time.Now()

	messages := []proto.Message{
		0: factory.CreateMesgOnly(mesgnum.FileId).WithFields(
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerDevelopment.Uint16()),
		),
		1: factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
		),
		2: factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
		),
		3: factory.CreateMesgOnly(mesgnum.Record).WithFields(
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(1 * time.Second))),
		),
		4: factory.CreateMesgOnly(mesgnum.Event).WithFields(
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
		),
		5: factory.CreateMesgOnly(mesgnum.Session).WithFields(
			factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
		),
	}

	expected := []proto.Message{
		messages[0],
		messages[1],
		messages[3],
		messages[2],
		messages[4],
		messages[5],
	}

	filedef.SortMessagesByTimestamp(messages)
	if diff := cmp.Diff(messages, expected,
		cmpopts.IgnoreTypes(proto.Value{}), // We only care the ordering
	); diff != "" {
		t.Fatal(diff)
	}
}
