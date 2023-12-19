package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/muktihari/fit/decoder"
	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/bufferedwriter"
	"github.com/muktihari/fit/profile/filedef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	fin, err := os.Open("../../../testdata/local/klaten-nganjuk.fit")
	// fin, err := os.Open("../../../testdata/big_activity.fit")
	if err != nil {
		panic(err)
	}

	al := filedef.NewListener()
	defer al.Close()

	dec := decoder.New(bufio.NewReader(fin),
		decoder.WithMesgListener(al),
		decoder.WithBroadcastOnly(),
	)

	begin := time.Now()

	_, err = dec.Decode()
	if err != nil {
		panic(err)
	}

	activity := al.File().(*filedef.Activity)

	fmt.Println(activity.FileId.Manufacturer, len(activity.Records))

	fmt.Printf("decode took: %s\n", time.Since(begin))

	/* Do something with the Activity File, for example changing manufacturer and product */
	activity.FileId.Manufacturer = typedef.ManufacturerGarmin
	activity.FileId.Product = uint16(typedef.GarminProductEdge530)

	begin = time.Now()
	// Convert back to RAW Protocol Messages
	fit := activity.ToFit(factory.StandardFactory())

	fmt.Printf("to fit took: %s\n", time.Since(begin))

	fout, err := os.OpenFile("NewActivity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	bw := bufferedwriter.New(fout)
	defer bw.Flush()

	begin = time.Now()

	// Encode FIT to file
	enc := encoder.New(bw)
	if err := enc.Encode(&fit); err != nil {
		panic(err)
	}

	fmt.Printf("encode took: %s\n", time.Since(begin))
}
