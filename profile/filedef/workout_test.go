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

func newWorkoutMessageForTest(now time.Time) []proto.Message {
	return []proto.Message{
		factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
			fieldnum.FileIdType:        uint8(typedef.FileWorkout),
			fieldnum.FileIdTimeCreated: datetime.ToUint32(now),
		}),
		factory.CreateMesg(mesgnum.DeveloperDataId).WithFieldValues(map[byte]any{
			fieldnum.DeveloperDataIdDeveloperDataIndex: uint8(0),
		}),
		factory.CreateMesg(mesgnum.FieldDescription).WithFieldValues(map[byte]any{
			fieldnum.FieldDescriptionDeveloperDataIndex: uint8(0),
		}),
		factory.CreateMesg(mesgnum.Workout).WithFieldValues(map[byte]any{
			fieldnum.WorkoutSessionSport: uint8(typedef.SportSwimming),
		}),
		factory.CreateMesg(mesgnum.WorkoutStep).WithFieldValues(map[byte]any{
			fieldnum.WorkoutStepEquipment: uint8(typedef.WorkoutEquipmentSwimFins),
		}),
		factory.CreateMesg(mesgnum.WorkoutStep).WithFieldValues(map[byte]any{
			fieldnum.WorkoutStepEquipment: uint8(typedef.WorkoutEquipmentSwimSnorkel),
		}),
		// Unrelated messages
		factory.CreateMesg(mesgnum.CoursePoint).WithFieldValues(map[byte]any{
			fieldnum.CoursePointTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
		factory.CreateMesg(mesgnum.BarometerData).WithFieldValues(map[byte]any{
			fieldnum.BarometerDataTimestamp: datetime.ToUint32(incrementSecond(&now)),
		}),
	}
}

func TestWorkoutCorrectness(t *testing.T) {
	mesgs := newWorkoutMessageForTest(time.Now())

	workout := filedef.NewWorkout(mesgs...)
	if workout.FileId.Type != typedef.FileWorkout {
		t.Fatalf("expected: %#v, got: %#v", typedef.FileActivity, workout.FileId.Type)
	}

	fit := workout.ToFit(nil)

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
