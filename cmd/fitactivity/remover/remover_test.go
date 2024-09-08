package remover

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func TestRemove(t *testing.T) {
	tt := []struct {
		name     string
		fit      *proto.FIT
		opts     []Option
		expected *proto.FIT
	}{
		{
			name: "remove unknown messages only",
			fit: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.MfgRangeMin},
					{Num: mesgnum.MfgRangeMax},
					{Num: mesgnum.Record},
				},
			},
			opts: []Option{WithRemoveUnknown()},
			expected: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.Record},
				},
			},
		},
		{
			name: "remove user-defined messages only",
			fit: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.MfgRangeMin},
					{Num: mesgnum.MfgRangeMax},
					{Num: mesgnum.Record},
					{Num: mesgnum.GpsMetadata},
					{Num: mesgnum.Lap},
					{Num: mesgnum.Record},
					{Num: mesgnum.Session},
				},
			},
			opts: []Option{WithRemoveMesgNums(map[typedef.MesgNum]struct{}{
				mesgnum.Record:      {},
				mesgnum.GpsMetadata: {},
			})},
			expected: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.MfgRangeMin},
					{Num: mesgnum.MfgRangeMax},
					{Num: mesgnum.Lap},
					{Num: mesgnum.Session},
				},
			},
		},
		{
			name: "remove developer data only",
			fit: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.DeveloperDataId},
					{Num: mesgnum.FieldDescription},
					{Num: mesgnum.Record, DeveloperFields: make([]proto.DeveloperField, 5)},
					{Num: mesgnum.Lap},
					{Num: mesgnum.Record, DeveloperFields: make([]proto.DeveloperField, 10)},
					{Num: mesgnum.Session},
				},
			},
			opts: []Option{WithRemoveDeveloperData()},
			expected: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.Record},
					{Num: mesgnum.Lap},
					{Num: mesgnum.Record},
					{Num: mesgnum.Session},
				},
			},
		},
		{
			name: "remove unknown messages, user-defined messages, and developerData",
			fit: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.MfgRangeMin},
					{Num: mesgnum.MfgRangeMax},
					{Num: mesgnum.DeveloperDataId},
					{Num: mesgnum.FieldDescription},
					{Num: mesgnum.Record},
					{Num: mesgnum.Lap},
					{Num: mesgnum.Record},
					{Num: mesgnum.Session, DeveloperFields: make([]proto.DeveloperField, 5)},
				},
			},
			opts: []Option{
				WithRemoveUnknown(),
				WithRemoveMesgNums(map[typedef.MesgNum]struct{}{
					mesgnum.Record: {},
				}),
				WithRemoveDeveloperData(),
			},
			expected: &proto.FIT{
				Messages: []proto.Message{
					{Num: mesgnum.FileId},
					{Num: mesgnum.Lap},
					{Num: mesgnum.Session},
				},
			},
		},
	}

	for i, tc := range tt {
		t.Run(fmt.Sprintf("[%d] %s", i, tc.name), func(t *testing.T) {
			Remove(tc.fit, tc.opts...)
			if diff := cmp.Diff(tc.fit, tc.expected); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
