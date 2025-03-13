// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
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

	for i := range records {
		messages = append(messages, records[i].ToMesg(nil))
	}

	if len(messages) != 1 {
		t.Fatalf("expected 1: got: %d", len(messages))
	}
}

func TestSortMessagesByTimestamp(t *testing.T) {
	now := time.Now()

	// Special case:
	// 1. CoursePoint's Timestamp Num is 1
	// 2. Set's Timestamp Num is 254
	messages := []proto.Message{
		0: {Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerDevelopment.Uint16()),
		}},
		1: {Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
		}},
		2: {Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
		}},
		3: {Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(1 * time.Second))),
		}},
		4: {Num: mesgnum.Event, Fields: []proto.Field{
			factory.CreateField(mesgnum.Event, fieldnum.EventTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
		}},
		5: {Num: mesgnum.Session, Fields: []proto.Field{
			factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
		}},
		6: {Num: mesgnum.UserProfile, Fields: []proto.Field{
			factory.CreateField(mesgnum.UserProfile, fieldnum.UserProfileFriendlyName).WithValue("muktihari"),
		}},
		7: {Num: mesgnum.Set, Fields: []proto.Field{
			factory.CreateField(mesgnum.Set, fieldnum.SetTimestamp).WithValue(datetime.ToUint32(now.Add(4 * time.Second))),
		}},
		8: {Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now.Add(2 * time.Second))),
		}},
		9: {Num: mesgnum.CoursePoint, Fields: []proto.Field{
			factory.CreateField(mesgnum.CoursePoint, fieldnum.CoursePointTimestamp).WithValue(datetime.ToUint32(now.Add(3 * time.Second))),
		}},
	}

	expected := []proto.Message{
		messages[0],
		messages[6],
		messages[1],
		messages[3],
		messages[2],
		messages[4],
		messages[5],
		messages[8],
		messages[9],
		messages[7],
	}

	filedef.SortMessagesByTimestamp(messages)
	if diff := cmp.Diff(messages, expected,
		cmpopts.IgnoreTypes(proto.Value{}), // We only care the ordering
	); diff != "" {
		t.Fatal(diff)
	}
}
