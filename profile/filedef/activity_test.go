package filedef_test

import (
	"fmt"
	"math"
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

func createFieldComparer() cmp.Option {
	return cmp.Comparer(func(field1, field2 proto.Field) bool {
		// Compare float in integer form. since when f is NaN, f != f.
		switch f1 := field1.Value.(type) {
		case float32:
			f2, ok := field2.Value.(float32)
			if !ok {
				break
			}
			return math.Float32bits(f1) == math.Float32bits(f2)
		case float64:
			f2, ok := field2.Value.(float64)
			if !ok {
				break
			}
			return math.Float64bits(f1) == math.Float64bits(f2)
		}
		return cmp.Diff(field1, field2) == ""
	})
}

// incrementSecond will increment v for 1 second to similate incoming message in 1 second sampling.
func incrementSecond(v *time.Time) time.Time {
	*v = v.Add(time.Second)
	return *v
}

func newActivityMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
			fieldnum.FileIdType:        uint8(typedef.FileActivity),
			fieldnum.FileIdTimeCreated: datetime.ToUint32(now),
		}),
		factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
			fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
		}),
		factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
			fieldnum.FieldDescriptionDeveloperDataIndex: uint8(0),
		}),
		factory.CreateMesg(mesgnum.DeviceInfo).WithFieldValues(map[byte]any{
			fieldnum.DeviceInfoManufacturer: uint16(typedef.ManufacturerGarmin),
		}),
		factory.CreateMesg(mesgnum.UserProfile).WithFieldValues(map[byte]any{
			fieldnum.UserProfileFriendlyName: "Mary Jane",
			fieldnum.UserProfileAge:          uint8(21),
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
			fieldnum.LapTimestamp: datetime.ToUint32(now), // intentionally using same timestamp as last message
		}),
		factory.CreateMesg(mesgnum.Session).WithFieldValues(map[byte]any{
			fieldnum.SessionTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.Activity).WithFieldValues(map[byte]any{
			fieldnum.ActivityTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		// Unordered pptional Messages
		factory.CreateMesg(mesgnum.Length).WithFieldValues(map[byte]any{
			fieldnum.LengthAvgSpeed: uint16(1000),
		}),
		factory.CreateMesg(mesgnum.SegmentLap).WithFieldValues(map[byte]any{
			fieldnum.SegmentLapAvgCadence: uint8(100),
		}),
		factory.CreateMesg(mesgnum.ZonesTarget).WithFieldValues(map[byte]any{
			fieldnum.SegmentLapAvgCadence: uint8(100),
		}),
		factory.CreateMesg(mesgnum.Workout).WithFieldValues(map[byte]any{
			fieldnum.WorkoutSessionSport: uint8(typedef.SportCycling),
		}),
		factory.CreateMesg(mesgnum.WorkoutStep).WithFieldValues(map[byte]any{
			fieldnum.WorkoutStepIntensity: uint8(typedef.IntensityActive),
		}),
		factory.CreateMesg(mesgnum.Hr).WithFieldValues(map[byte]any{
			fieldnum.HrTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.Hrv).WithFieldValues(map[byte]any{
			fieldnum.HrvStatusSummaryStatus: uint8(typedef.HrvStatusBalanced),
		}),
		// Unrelated messages
		factory.CreateMesg(mesgnum.BarometerData).WithFieldValues(map[byte]any{
			fieldnum.BarometerDataTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.CoursePoint).WithFieldValues(map[byte]any{
			fieldnum.CoursePointTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
	}
}

func TestActivityCorrectness(t *testing.T) {
	mesgs := newActivityMessageForTest(time.Now())

	activity := filedef.NewActivity(mesgs...)
	if activity.FileId.Type != typedef.FileActivity {
		t.Fatalf("expected: %#v, got: %#v", typedef.FileActivity, activity.FileId.Type)
	}

	fit := activity.ToFit(nil) // use standard factory

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
