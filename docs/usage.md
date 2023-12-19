# Usage

Table of Contents:

1. [Dicoding](./usage.md#Decoding)
   - [Decode RAW Protocol Messages](#Decode-RAW-Protocol-Messages)
   - [Decode to Common File Types](#Decode-to-Common-File-Types)
   - [Peek FileId](#Peek-FileId)
   - [Available Decode Options](#Available-Decode-Options)
2. [Encoding](#Encoding)
   - [Encode RAW Protocol Messages](#Encode-RAW-Protocol-Messages)
   - [Encode Common File Types](#Encode-Common-File-Types)
   - [Available Encode Options](#Available-Encode-Options)

## Decoding

### Decode RAW Protocol Messages

Decode as RAW Protocol Messages allows us to interact with FIT files through their original protocol message structures
without the needs to do conversions.

```go
package main

import (
    "bufio"
    "context"
    "fmt"
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/untyped/fieldnum"
)

func main() {
    f, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    dec := decoder.New(bufio.NewReader(f))

    fit, err := dec.Decode()
    if err != nil {
        panic(err)
    }

    fmt.Printf("FileHeader DataSize: %d\n", fit.FileHeader.DataSize)
    fmt.Printf("Messages count: %d\n", len(fit.Messages))
    field, _ := fit.Messages[0].FieldByNum(fieldnum.FileIdType) // FileId is always the first message.
    fmt.Printf("File Type: %v\n", field.Value) // 4 = activity

    // Output:
    // FileHeader DataSize: 94080
    // Messages count: 3611
    // File Type: 4
}
```

#### Decode Chained FIT FIle

If you are uncertain if it's a chained fit file. Create a loop and use dec.Next() to check whether next sequence of bytes are still a valid FIT sequence.

```go
    ...

    dec := decoder.New(bufio.NewReader(f))
    for dec.Next() {
        fit, err := dec.Decode()
        if err != nil {
            return err
        }
        // do something with fit variable
    }

    ...
```

### Decode to Common File Types

Decode to Common File Types enables us to interact with FIT files through common file types such as Activity Files, Course Files, Workout Files, and more, which group protocol messages based on specific purposes.

_Note: Currently only 3 common file types are defined: Activity, Course & Workout, but you can create your own file types using our building block and register it in the listener as an option using WithFileSets()_

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/filedef"
)

func main() {
    f, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // The listener will receive every decoded message from the decoder as soon as it is decoded
    // and transform it into an filedef's File.
    al := filedef.NewListener()
    defer al.Close() // release channel used by listener

    dec := decoder.New(bufio.NewReader(f),
        decoder.WithMesgListener(al), // Add activity listener to the decoder
        decoder.WithBroadcastOnly(),  // Direct the decoder to only broadcast the messages without retaining them.
    )

    for dec.Next() {
        _, err = dec.Decode()
        if err != nil {
            panic(err)
        }

        // The resulting File can be retrieved after decoding process completed.
        switch file := al.File().(type) {
        case *filedef.Course:
            // do something if it's a course file
        case *filedef.Workout:
            // do something if it's a workout file
        case *filedef.Activity:
            fmt.Printf("File Type: %s\n", file.FileId.Type)
            fmt.Printf("Sessions count: %d\n", len(file.Sessions))
            fmt.Printf("Laps count: %d\n", len(file.Laps))
            fmt.Printf("Records count: %d\n", len(file.Records))

            i := 100
            fmt.Printf("\nSample value of record[%d]:\n", i)
            fmt.Printf("  Lat: %v semicircles\n", file.Records[i].PositionLat)
            fmt.Printf("  Long: %v semicircles\n", file.Records[i].PositionLong)
            fmt.Printf("  Speed: %g m/s\n", float64(file.Records[i].Speed)/1000)
            fmt.Printf("  HeartRate: %v bpm\n", file.Records[i].HeartRate)
            fmt.Printf("  Cadence: %v rpm\n", file.Records[i].Cadence)

            // Output:
            // File Type: activity
            // Sessions count: 1
            // Laps count: 1
            // Records count: 3601
            //
            // Sample value of record[100]:
            //   Lat: 0 semicircles
            //   Long: 10717 semicircles
            //   Speed: 1 m/s
            //   HeartRate: 126 bpm
            //   Cadence: 100 rpm
        }
    }
}
```

### Peek FileId

We don't need to decode the entire Fit file to verify its type. Instead, we can use the 'PeekFileId' method to check the corresponding type. After invoking this method, we can decide whether to proceed with decoding the file or to stop. If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/filedef"
    "github.com/muktihari/fit/profile/typedef"
)

func main() {
    f, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    al := filedef.NewListener()
    defer al.Close() // release channel used by listener

    dec := decoder.New(bufio.NewReader(f),
        decoder.WithMesgListener(al),
        decoder.WithBroadcastOnly(),
    )

    fileId, err := dec.PeekFileId()
    if err != nil {
        panic(err)
    }

    fmt.Printf("File Type: %s\n", fileId.Type)

    // Output:
    // File Type: activity

    if fileId.Type != typedef.FileActivity {
        return // Let's stop.
    }

    // It's an Activity File, let's Decode it.
    _, err = dec.Decode()
    if err != nil {
        panic(err)
    }

    activity := al.File().(*filedef.Activity)

    fmt.Printf("Sessions count: %d\n", len(activity.Sessions))
    fmt.Printf("Laps count: %d\n", len(activity.Laps))
    fmt.Printf("Records count: %d\n", len(activity.Records))
    // ...
}
```

### Available Decode Options

1. **WithFactory**: allow us to use custom Factory, for example if we are working with multiple manufacturer specific messages at the same time.

   Example

   ```go
   fac := factory.New()
   fac.RegisterMesg(proto.Message{
       Num:  65281,
       Name: "Manufacturer specific message",
       Fields: []proto.Field{
           {
               FieldBase: &proto.FieldBase{
                   Num:    253,
                   Name:   "Timestamp",
                   Type:   profile.Uint32,
                   Size:   4,
                   Scale:  1,
                   Offset: 0,
                   Units:  "s",
               },
           },
       },
   })

   dec := decoder.New(f, decoder.WithFactory(fac))
   ```

1. **WithMesgListener**: adds message listener to the listener pool so that we can receive the messages as soon as it is decoded.

   Example:

   ```go
   al := filedef.NewListener()
   dec := decoder.New(f,
       decoder.WithMesgListener(al),
   )
   ```

1. **WithMesgDefListener**: adds message definition listener to the listener pool so that we can receive the message definitions as soon as it is decoded.

   Example:

   ```go
   conv := fitcsv.NewConverter(bw)
   defer conv.Wait()

   dec := decoder.New(f,
       decoder.WithMesgDefListener(conv),
       decoder.WithMesgListener(conv),
   )
   ```

1. **WithBroadcastOnly**: directs the decoder to only broadcast the messages without retaining them.

   Example:

   ```go
   al := filedef.NewListener()
   dec := decoder.New(f,
       decoder.WithMesgListener(al),
       decoder.WithBroadcastOnly(),
   )
   ```

1. **WithIgnoreChecksum**: directs the decoder to ignore the checksum, which is useful when we want to retrieve the data without considering its integrity.

   Example:
   Example:

   ```go
   dec := decoder.New(f, decoder.WithIgnoreChecksum())
   ```

## Encoding

Note: By default, Encoder use protocol version 1.0 (proto.V1), if you want to use protocol version 2.0 (proto.V2), please specify it using Encode Option: WithProtocolVersion. See [Available Encode Options](#Available-Encode-Options)

### Encode RAW Protocol Messages

Example of encoding fit by self declaring the protocol messages, this is to show how we can compose the message using this SDK.

```go

package main

import (
	"context"
	"os"
	"time"

	"github.com/muktihari/fit/encoder"
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/kit/bufferedwriter"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

func main() {
    f, err := os.OpenFile("CoffeeRide.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    now := time.Now()

    fit := &proto.Fit{
        Messages: []proto.Message{
            factory.CreateMesg(mesgnum.FileId).WithFieldValues(map[byte]any{
                fieldnum.FileIdTimeCreated:  datetime.ToUint32(now),
                fieldnum.FileIdManufacturer: typedef.ManufacturerBryton,
                fieldnum.FileIdProductName:  "Bryton Active App",
            }),
            factory.CreateMesg(mesgnum.Activity).WithFieldValues(map[byte]any{
                fieldnum.ActivityType:        typedef.ActivityManual,
                fieldnum.ActivityTimestamp:   datetime.ToUint32(now),
                fieldnum.ActivityNumSessions: uint16(1),
			}),
            factory.CreateMesg(mesgnum.Session).WithFieldValues(map[byte]any{
                fieldnum.SessionAvgSpeed:     uint16(1000),
                fieldnum.SessionAvgCadence:   uint8(78),
                fieldnum.SessionAvgHeartRate: uint8(100),
            }),
            // We can use WithFields as well, see `proto` for details.
            // When using WithFields, it's recommended to use CreateMesgOnly to create a mesg with empty fields
            // to reduce unecessary alloc since the fields will be replaced anyway, see `factory` for details.
            factory.CreateMesgOnly(mesgnum.Record).WithFields(
                factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
                factory.CreateField(mesgnum.Record, fieldnum.RecordCadence).WithValue(uint8(78)),
                factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(100)),
            ),
        },
    }

    bw := bufferedwriter.NewSize(f, 1000<<10) // write to file every 1MB
    defer bw.Flush() // write any buffered data to the underlying io.Writer.

    enc := encoder.New(bw)
    if err := enc.Encode(fit); err != nil {
        panic(err)
    }
}
```

### Encode Common File Types

Example of encoding fit by self declaring the protocol messages but using common File Types building block.

```go
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
    f, err := os.OpenFile("NewActivity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
    if err != nil {
        panic(err)
    }

    now := time.Now()

    activity := filedef.NewActivity()

    activity.FileId = *mesgdef.NewFileId(nil).
        SetType(typedef.FileActivity).
        SetTimeCreated(now).
        SetManufacturer(typedef.ManufacturerSuunto).
        SetProduct(56). // Suunto 5 Peak
        SetProductName("Suunto 5 Peak")

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

    bw := bufferedwriter.New(f)
    defer bw.Flush()

    enc := encoder.New(bw)
    if err := enc.Encode(&fit); err != nil {
        panic(err)
    }
}


```

Example decoding FIT file into common file `Activity File`, edit the manufacturer and product, and then encode it again.

```go
    ...
    /* Assume you retrive the activity file from decoding a Fit file. */

    fin, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }

    al := filedef.NewListener()
    defer al.Close()

    dec := decoder.New(bufio.NewReader(fin),
        decoder.WithMesgListener(al),
        decoder.WithBroadcastOnly(),
    )

    _, err = dec.Decode()
    if err != nil {
        panic(err)
    }

    activity := al.File().(*filedef.Activity)

    /* Do something with the Activity File, for example changing manufacturer and product like this */
    activity.FileId.Manufacturer = typedef.ManufacturerGarmin
    activity.FileId.Product = uint16(typedef.GarminProductEdge530)

    // Convert back to RAW Protocol Messages
    fit := activity.ToFit(factory.StandardFactory())

    fout, err := os.OpenFile("NewActivity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
    if err != nil {
        panic(err)
    }

    bw := bufferedwriter.New(fout)
    defer bw.Flush()

    // Encode FIT to file
    enc := encoder.New(bw)
    if err := enc.Encode(&fit); err != nil {
        panic(err)
    }

    ...
```

### Available Encode Options

1. **WithProtocolVersion**: directs the Encoder to use specific Protocol Version (default: proto.V1).

   Example:

   ```go
   enc := encoder.New(f, encoder.WithProtocolVersion(proto.V2))
   ```

1. **WithMessageValidator**: directs the Encoder to use this message validator instead of the default one.

   Example:

   ```go
   enc := encoder.New(f, encoder.WithMessageValidator(
       encoder.NewMessageValidator(encoder.ValidatorWithPreserveInvalidValues())),
   )
   ```

1. **WithBigEndian**: directs the Encoder to encode values in Big-Endian bytes order (default: Little-Endian).

   Example:

   ```go
   enc := encoder.New(f, encoder.WithBigEndian())
   ```

1. **WithCompressedTimestampHeader**: directs the Encoder to compress timestamp in header to reduce file size.  
   Saves 7 bytes per message: 3 bytes for field definition and 4 bytes for the uint32 timestamp value.

   Example:

   ```go
   enc := encoder.New(f, encoder.WithCompressedTimestampHeader())
   ```

1. **WithNormalHeader**: directs the Encoder to use NormalHeader for encoding the message using
   specified multiple local message typedef. By default, the Encoder uses local message type 0.
   This option allows users to specify values between 0-15 (while entering zero is equivalent to using
   the default option, nothing is changed). Using multiple local message types optimizes file size by
   avoiding the need to interleave different message typedef.

   Note: To minimize the required RAM for decoding, it's recommended to use a minimal number of
   local message types in a file. For instance, embedded devices may only support decoding data
   from local message type 0. Additionally, multiple local message types should be avoided in
   file types like settings, where messages of the same type can be grouped together.

   Example:

   ```go
   enc := encoder.New(f, encoder.WithNormalHeader(4)) // 0-15
   ```

   Note: we can only use either WithCompressedTimestampHeader or WithNormalHeader, can't use both at the same time.
