package main

import (
	"os"
	"time"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/bufferedwriter"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
)

func main() {
	now := time.Now()

	activity := filedef.NewActivity()

	activity.FileId = *mesgdef.NewFileId(nil).
		SetType(typedef.FileActivity).
		SetTimeCreated(now).
		SetManufacturer(typedef.ManufacturerSuunto).
		SetProduct(56). // Suunto 5 Peak
		SetProductName("Suunto App")

	activity.Records = append(activity.Records,
		mesgdef.NewRecord(nil).
			SetTimestamp(now.Add(1*time.Second)).
			SetSpeed(1000).
			SetCadence(90).
			SetHeartRate(100),
		mesgdef.NewRecord(nil).
			SetTimestamp(now.Add(2*time.Second)).
			SetSpeed(1010).
			SetCadence(100).
			SetHeartRate(110),
	)

	activity.Laps = append(activity.Laps,
		mesgdef.NewLap(nil).
			SetTimestamp(now.Add(3*time.Second)).
			SetStartTime(now.Add(1*time.Second)).
			SetAvgSpeed(1000).
			SetAvgCadence(95).
			SetAvgHeartRate(105),
	)

	activity.Sessions = append(activity.Sessions,
		mesgdef.NewSession(nil).
			SetTimestamp(now.Add(3*time.Second)).
			SetStartTime(now.Add(1*time.Second)).
			SetAvgSpeed(1000).
			SetAvgCadence(95).
			SetAvgHeartRate(105),
	)

	activity.Activity = mesgdef.NewActivity(nil).
		SetType(typedef.ActivityManual).
		SetTimestamp(now.Add(4 * time.Second)).
		SetNumSessions(1)

	// Convert back to FIT protocol messages
	fit := activity.ToFit(factory.StandardFactory())

	f, err := os.OpenFile("NewActivity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	bw := bufferedwriter.New(f)
	defer bw.Flush()

	// Encode FIT to file
	enc := encoder.New(bw)
	if err := enc.Encode(&fit); err != nil {
		panic(err)
	}
}
