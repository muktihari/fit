# Usage

Table of Contents:

1. [Dicoding](./usage.md#Decoding)
   - [Decode RAW Protocol Messages](#Decode-RAW-Protocol-Messages)
   - [Decode to Common File Types](#Decode-to-Common-File-Types)
   - [Peek FileId](#Peek-FileId)
   - [Discard FIT File Sequences](#Discard-FIT-File-Sequences)
   - [Check Integrity](#Check-Integrity)
   - [Available Decode Options](#Available-Decode-Options)
   - [RawDecoder (Low-Level Abstraction)](#RawDecoder-Low-Level-Abstraction)
2. [Encoding](#Encoding)
   - [Encode RAW Protocol Messages](#Encode-RAW-Protocol-Messages)
   - [Encode Common File Types](#Encode-Common-File-Types)
   - [Available Encode Options](#Available-Encode-Options)
   - [Stream Encoder](#Stream-Encoder)

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
    fmt.Printf("File Type: %v\n", fit.Messages[0].FieldValueByNum(fieldnum.FileIdType).Any()) // FileId is always the first message; 4 = activity

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

Decode to Common File Types enables us to interact with FIT files through common file types such as Activity Files, Course Files, Workout Files, and [more](../profile/filedef/doc.go), which group protocol messages based on specific purposes.

1. To get started, the simpliest way to create an common file type is to decode the FIT file in its raw protocol messages then pass the messages to create the desired common file type.

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

    dec := decoder.New(bufio.NewReader(f))

    for dec.Next() {
        fit, err := dec.Decode()
        if err != nil {
            panic(err)
        }

        activity := filedef.NewActivity(fit.Messages...)

        fmt.Printf("File Type: %s\n", activity.FileId.Type)
        fmt.Printf("Sessions count: %d\n", len(activity.Sessions))
        fmt.Printf("Laps count: %d\n", len(activity.Laps))
        fmt.Printf("Records count: %d\n", len(activity.Records))

        i := 100
        fmt.Printf("\nSample value of record[%d]:\n", i)
        fmt.Printf("  Lat: %v semicircles\n", activity.Records[i].PositionLat)
        fmt.Printf("  Long: %v semicircles\n", activity.Records[i].PositionLong)
        fmt.Printf("  Speed: %g m/s\n", float64(activity.Records[i].Speed)/1000)
        fmt.Printf("  HeartRate: %v bpm\n", activity.Records[i].HeartRate)
        fmt.Printf("  Cadence: %v rpm\n", activity.Records[i].Cadence)

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
```

2. While the previous example is work for most use cases and probably can be your goto choice to use for small scale, it's slighly inefficient as we only utilize one goroutine (the main goroutine) and also we need to allocate the `fit.Messages` before creating the `activity` file itself. For bigger scale, or in scale that require a streaming process, we can define `filedef's Listener` to create the `activity` file. This not only reduce the need to allocate `fit.Messages` but also we can receive the message as it is decode in other goroutine. As the decoder decode the message, we can create the message in another process concurrently.

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

**The ability to broadcast every message as it is decoded is one of biggest advantage of using this SDK, we can define custom listener to process the message as we like and in a streaming fashion, as long as it satisfies the [Listener](https://github.com/muktihari/fit/blob/master/decoder/listener.go) interface.**

### Peek FileId

We don't need to decode the entire FIT file to verify its type. Instead, we can use the 'PeekFileId' method to check the corresponding type. After invoking this method, we can decide whether to proceed with decoding the file or to stop. If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.

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

### Discard FIT File Sequences

When handling a chained fit file, sometimes we only want to decode a certain file type, let's say a Course File, while discarding other file types. Instead of unecessarily decode all the file types but we don't use all of them, we can just discard it. Discard directs the Decoder to efficiently just discard the bytes without doing unnecessary work.

```go
    ...

    for dec.Next() {
        fileId, err := dec.PeekFileId()
        if err != nil {
            return err
        }
        if fileId.Type != typedef.FileCourse {
            if err := dec.Discard(); err != nil { // not a Course File, discard this sequence!
                return err
            }
            continue
        }
        fit, err := dec.Decode()
        if err != nil {
            return err
        }
        // do something with fit variable
    }
    ...
```

### Check Integrity

Check integrity checks whether FIT File is not corrupted or contains missing data that can invalidates the entire file. Example of when we need to check the integrity is when dealing with **Course File** that contains turn-by-turn navigation, we don't want to guide a person halfway to their destination or guide them to unintended route, do we? The same applies for other file types where it is critical that the contents of the file should be valid, such as Workout, User Profile, Device Settings, etc. For Activity File, most cases, we don't need to check the integrity.

This returns the number of sequences completed and any error encountered. The number of sequences completed can help recovering valid FIT sequences in a chained FIT that contains invalid or corrupted data.

More about this: [https://developer.garmin.com/fit/cookbook/isfit-checkintegrity-read/](https://developer.garmin.com/fit/cookbook/isfit-checkintegrity-read/)

```go
package main

import (
    "bufio"
    "io"
    "os"

    "github.com/muktihari/fit/decoder"
)

func main() {
    f, err := os.Open("Course.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    dec := decoder.New(bufio.NewReader(f))

    // Fyi, we can invoke PeekFileId() first to check the type of the first FIT File sequence before checking the integrity.
    // For most cases, we wouldn't want to check integrity of Activity File.
    // However, you can only Peek FileId of the first FIT sequence if it's a chained FIT files.
    if _, err := dec.CheckIntegrity(); err != nil {
        panic(err)
    }

    // After invoking CheckIntegrity and users want to reuse `dec` to Decode the FIT file,
    // `f` should be reset since `f` has been fully read. The following method will do:
    _, err = f.Seek(0, io.SeekStart)
    if err != nil {
        panic(err)
    }

    for dec.Next() {
        fit, err := dec.Decode()
        if err != nil {
            panic(err)
        }
        _ = fit // Do something with fit
    }
}
```

### Available Decode Options

1.  **WithFactory**: allow us to use custom Factory, for example if we are working with multiple manufacturer specific messages at the same time.

    Example

    ```go
    fac := factory.New()
    fac.RegisterMesg(proto.Message{
        Num:  65281,
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

1.  **WithMesgListener**: adds message listener to the listener pool so that we can receive the messages as soon as it is decoded.

    Example:

    ```go
    al := filedef.NewListener()
    dec := decoder.New(f,
        decoder.WithMesgListener(al),
    )
    ```

1.  **WithMesgDefListener**: adds message definition listener to the listener pool so that we can receive the message definitions as soon as it is decoded.

    Example:

    ```go
    conv := fitcsv.NewConverter(bw)
    defer conv.Wait()

    dec := decoder.New(f,
        decoder.WithMesgDefListener(conv),
        decoder.WithMesgListener(conv),
    )
    ```

1.  **WithBroadcastOnly**: directs the decoder to only broadcast the messages without retaining them.

    Example:

    ```go
    al := filedef.NewListener()
    dec := decoder.New(f,
        decoder.WithMesgListener(al),
        decoder.WithBroadcastOnly(),
    )
    ```

1.  **WithIgnoreChecksum**: directs the decoder to ignore the checksum, which is useful when we want to retrieve the data without considering its integrity.

    Example:

    ```go
    dec := decoder.New(f, decoder.WithIgnoreChecksum())
    ```

1.  **WithNoComponentExpansion**: directs the Decoder to not expand the components.

    Example:

    ```go
    dec := decoder.New(f, decoder.WithNoComponentExpansion())
    ```

1.  **WithLogWriter**: specifies where the log messages will be written to. By default, the Decoder writes log message to io.Discard.
    The Decoder will only write log messages when it encountered a bad encoded FIT file such as:

    - Field Definition's Size is less than its basetype's size. e.g. Size 1 byte but having basetype uint32 (4 bytes).
    - Encounter a Developer Field without prior Field Description Message.

    ```go
    dec := decoder.New(f, decoder.WithLogWriter(os.Stdout))
    ```

### RawDecoder (Low-Level Abstraction)

Raw Decoder provides a way to split the bytes based on its scope (File Header, Message Definition, Message Data and CRC) as the building block to work with the FIT data in its scoped bytes.

The idea is to allow us to use a minimal viable decoder for performance and memory-critical situations, where every computation or memory usage is constrained. RawDecoder itself is using constant memory < 131 KB and the Decode method has zero heap alloc (except errors) while it may use additional small stack memory. The implementation of the callback function is also expected to have minimal overhead. Theoretically, from the memory usage alone, this can run on an embedded device, for instance, using [tinygo](https://tinygo.org) or other compilers, but no attempt has been made.

Here is the simple example to check integrity using this building block:

```go
package main

import (
    "bufio"
    "encoding/binary"
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/kit/hash/crc16"
)

func main() {
    f, err := os.Open("./testdata/from_garmin_forums/triathlon_summary_first.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    dec := decoder.NewRaw()
    hash16 := crc16.New(nil)

    _, err = dec.Decode(bufio.NewReader(f), func(flag decoder.RawFlag, b []byte) error {
        switch flag {
        case decoder.RawFlagFileHeader:
            if err := proto.Validate(b[1]); err != nil {
                return err
            }
            if binary.LittleEndian.Uint32(b[4:8]) == 0 {
                return decoder.ErrDataSizeZero
            }
            if b[0] == 14 {
                hash16.Write(b[:12])
                if binary.LittleEndian.Uint16(b[12:14]) != hash16.Sum16() {
                    return decoder.ErrCRCChecksumMismatch
                }
                hash16.Reset()
            }
        case decoder.RawFlagMesgDef, decoder.RawFlagMesgData:
            hash16.Write(b)
        case decoder.RawFlagCRC:
            if binary.LittleEndian.Uint16(b[:2]) != hash16.Sum16() {
                return decoder.ErrCRCChecksumMismatch
            }
            hash16.Reset()
        }
        return nil
    })

    if err != nil {
        panic(err)
    }
}

```

Example of using RawDecoder to count how many messages in a FIT File (in case it matters):

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "time"

    "github.com/muktihari/fit/decoder"
)

func main() {
    defer func(begin time.Time) {
        fmt.Printf("took: %s\n", time.Since(begin))
    }(time.Now())

    f, err := os.Open("./testdata/from_garmin_forums/triathlon_summary_first.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    dec := decoder.NewRaw()

    var lenMesgsPerSequence []int
    var lenMesgs int

    _, err = dec.Decode(bufio.NewReader(f), func(flag decoder.RawFlag, b []byte) error {
        switch flag {
        case decoder.RawFlagFileHeader:
            lenMesgs = 0
        case decoder.RawFlagMesgData:
            lenMesgs++
        case decoder.RawFlagCRC:
            lenMesgsPerSequence = append(lenMesgsPerSequence, lenMesgs)
        }
        return nil
    })
    if err != nil {
        panic(err)
    }

    for i := range lenMesgsPerSequence {
        fmt.Printf("seq[%d] has %d messages\n", i, lenMesgsPerSequence[i])
    }

    // Output:
    // seq[0] has 4177 messages
    // took: 312.651µs
}
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

    fit := &proto.FIT{
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
            // to reduce unnecessary alloc since the fields will be replaced anyway, see `factory` for details.
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
    defer f.Close()

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
    fit := activity.ToFIT(nil)

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
    /* Assume you retrive the activity file from decoding a FIT file. */

    fin, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer fin.Close()

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
    activity.FileId.Product = typedef.GarminProductEdge530.Uint16()

    // Convert back to RAW Protocol Messages
    fit := activity.ToFIT(nil)

    fout, err := os.OpenFile("NewActivity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
    if err != nil {
        panic(err)
    }
    defer fout.Close()

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

   - Encode Invalid Values (default: Invalid values are omitted)

     Example:

     ```go
     enc := encoder.New(f, encoder.WithMessageValidator(
         encoder.NewMessageValidator(
             encoder.ValidatorWithPreserveInvalidValues(),
         )),
     )
     ```

   - Validate Developer Field with custom Factory (default: factory.StandardFactory())

     If Developer Field contains a valid NativeMesgNum and NativeFieldNum, the value should be treated as native value (scale, offset, etc shall apply). To be able to check Nativeness, we need to look up the message's field in the factory and only then we can validate it such as validating scale & offset of the value. By applying this option, user can now use their own Factory.

     Example:

     ```go
     fac := factory.New()
     /* fill manufacturer specific messages in fac */
     enc := encoder.New(f, encoder.WithMessageValidator(
         encoder.NewMessageValidator(
             encoder.ValidatorWithFactory(fac),
         )),
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
   specified multiple local message type. By default, the Encoder uses local message type 0.
   This option allows users to specify values between 0-15 (while entering zero is equivalent to using
   the default option, nothing is changed). Using multiple local message types optimizes file size by
   avoiding the need to interleave different message definition.

   Note: To minimize the required RAM for decoding, it's recommended to use a minimal number of
   local message types in a file. For instance, embedded devices may only support decoding data
   from local message type 0. Additionally, multiple local message types should be avoided in
   file types like settings, where messages of the same type can be grouped together.

   Example:

   ```go
   enc := encoder.New(f, encoder.WithNormalHeader(4)) // 0-15
   ```

   Note: we can only use either WithCompressedTimestampHeader or WithNormalHeader, can't use both at the same time.

### Stream Encoder

This is a new feature to enable encode per message basis or in streaming fashion rather than bulk per `proto.FIT`. To enable this, the Encoder's Writer should either implement io.WriterAt or io.WriteSeeker, since we need to be able to update FileHeader (the first 14 bytes) for every sequence completed. This is another building block that we can use.

```go
package main

import (
    "os"
    "time"

    "github.com/muktihari/fit/encoder"
    "github.com/muktihari/fit/factory"
    "github.com/muktihari/fit/kit/bufferedwriter"
    "github.com/muktihari/fit/kit/datetime"
    "github.com/muktihari/fit/profile/untyped/fieldnum"
    "github.com/muktihari/fit/profile/untyped/mesgnum"
)

func main() {
    f, err := os.OpenFile("Activity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o777)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // f (*os.File) is implementing both io.WriterAt and io.WriteSeeker so it can be used directly.
    // But to reduce syscall, let's wrap it using buffered writer instead.
    bw := bufferedwriter.New(f)
    defer bw.Flush()

    streamEnc, err := encoder.New(bw).StreamEncoder()
    if err != nil {
        panic(err)
    }

    // Simplified example, writing only this mesg.
    mesg := factory.CreateMesgOnly(mesgnum.FileId).WithFields(
        factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(time.Now())),
    )

    // Write per message, we can use this to write message as it arrives.
    // For example, message retrieved from decoder's Listener can be write right away
    // without waiting all messages to be received.
    if err := streamEnc.WriteMessage(&mesg); err != nil {
        panic(err)
    }

    /* Write more messages */

    // This should be invoked for every sequence of FIT File (not every message) to finalize.
    if err := streamEnc.SequenceCompleted(); err != nil {
        panic(err)
    }
}

```
