// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	cd                = filepath.Dir(filename)
	testdata          = filepath.Join(cd, "..", "..", "..", "testdata")

	flag int         = os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	perm os.FileMode = 0o777
)

func main() {
	// NOTE: Encoder will always use current profile.Version, if we re-run this program, new generated files
	// 		 might have different FileHeader.ProfileVersion.

	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	ctx := context.Background()

	if err := createValidFitOnlyContainFileId(ctx); err != nil {
		fatalln(err)
	}

	if err := createBigActivityFile(ctx); err != nil {
		fatalln(err)
	}
}

func createValidFitOnlyContainFileId(ctx context.Context) error {
	f, err := os.OpenFile(filepath.Join(testdata, "valid_only_contain_fileid.fit"), flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()

	defer func(begin time.Time) {
		fmt.Printf("took: %s\n", time.Since(begin))
	}(time.Now())

	type Uint16 uint16

	now := datetime.ToTime(uint32(1062766519))
	fit := &proto.FIT{Messages: []proto.Message{
		{Num: mesgnum.FileId, Fields: []proto.Field{
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("something ss"),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerBryton),
			factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.Activity, Fields: []proto.Field{
			factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now)),
		}},
		{Num: mesgnum.Record, Fields: []proto.Field{
			factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(112)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordCadence).WithValue(uint8(80)),
			factory.CreateField(mesgnum.Record, fieldnum.RecordAltitude).WithValue(Uint16((150 + 500) * 5)),
		}},
	}}

	enc := encoder.New(f)
	if err := enc.EncodeWithContext(ctx, fit); err != nil {
		return err
	}

	return nil
}

func fatalln(v any) { fatalf("%v\n", v) }

func fatalf(format string, args ...any) {
	fmt.Printf(format, args...)
	os.Exit(0)
}
