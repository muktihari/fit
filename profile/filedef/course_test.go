package filedef_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func newCourseMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
			fieldnum.FileIdType:        uint8(typedef.FileCourse),
			fieldnum.FileIdTimeCreated: datetime.ToUint32(now),
		}),
		factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
			fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
		}),
		factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
			fieldnum.FieldDescriptionDeveloperDataIndex: uint8(0),
		}),
		factory.CreateMesg(mesgnum.Course).WithFieldValues(map[byte]any{
			fieldnum.CourseSport: uint8(typedef.SportRunning),
		}),
		factory.CreateMesg(mesgnum.Event).WithFieldValues(map[byte]any{
			fieldnum.EventTimestamp: datetime.ToUint32(incrementSecond(&now)),
			fieldnum.EventEvent:     uint8(typedef.EventActivity),
			fieldnum.EventEventType: uint8(typedef.EventTypeStart),
		}),
		factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.Event).WithFieldValues(map[byte]any{
			fieldnum.EventTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.Record).WithFieldValues(map[byte]any{
			fieldnum.RecordTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.Event).WithFieldValues(map[byte]any{
			fieldnum.EventTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.Lap).WithFieldValues(map[byte]any{
			fieldnum.LapTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		// Unordered optional Messages
		factory.CreateMesg(mesgnum.CoursePoint).WithFieldValues(map[byte]any{
			fieldnum.CoursePointTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		// Unrelated messages
		factory.CreateMesg(mesgnum.BarometerData).WithFieldValues(map[byte]any{
			fieldnum.BarometerDataTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
	}
}

func TestCourseCorrectness(t *testing.T) {
	mesgs := newCourseMessageForTest(time.Now())

	course := filedef.NewCourse(mesgs...)
	if course.FileId.Type != typedef.FileCourse {
		t.Fatalf("expected: %#v, got: %#v", typedef.FileActivity, course.FileId.Type)
	}

	fit := course.ToFit(nil)

	if diff := cmp.Diff(mesgs, fit.Messages, createFieldComparer()); diff != "" {
		fmt.Println("messages order:")
		for i := range fit.Messages {
			mesg := fit.Messages[i]
			fmt.Printf("%d: %s\n", mesg.Num, mesg.Num)
		}
		fmt.Println("")
		t.Fatal(diff)
	}
}
