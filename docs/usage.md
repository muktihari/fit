# Usage

Table of Contents:

1. [Decoding](#Decoding)
   - [Decode Protocol Messages](#Decode-Protocol-Messages)
   - [Decode with Message Listener](#Decode-with-Message-Listener)
     - [Creating your own Message Listener](#Creating-your-own-Message-Listener)
   - [Decode into Common File Types](#Decode-into-Common-File-Types)
     - [Using your own custom File Types](#Using-your-own-custom-File-Types)
   - [Decode Chained FIT Files](#Decode-Chained-FIT-Files)
   - [Peek FileHeader](#Peek-FileHeader)
   - [Peek FileId](#Peek-FileId)
   - [Discard FIT File Sequences](#Discard-FIT-File-Sequences)
   - [Check Integrity](#Check-Integrity)
   - [Available Decode Options](#Available-Decode-Options)
   - [Advance Decoder Usage](#Advance-Decoder-Usage)
   - [RawDecoder (Low-Level Abstraction)](#RawDecoder-Low-Level-Abstraction)
2. [Encoding](#Encoding)
   - [Encode Protocol Messages](#Encode-Protocol-Messages)
   - [Encode Common File Types](#Encode-Common-File-Types)
   - [Available Encode Options](#Available-Encode-Options)
   - [Stream Encoder](#Stream-Encoder)
   - [Advance Encoder Usage](#Advance-Encoder-Usage)

## Decoding

NOTE: Decoder already implements efficient io.Reader buffering, so there's no need to wrap io.Reader (such as *os.File) using *bufio.Reader. Doing so will only reduce performance.

### Decode Protocol Messages

Decode protocol messages allows us to interact with FIT files directly through their original protocol messages' structure.

```go
package main

import (
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

    dec := decoder.New(f)

    fit, err := dec.Decode()
    if err != nil {
        panic(err)
    }

    fmt.Printf("FileHeader DataSize: %d\n", fit.FileHeader.DataSize)
    fmt.Printf("Messages count: %d\n", len(fit.Messages))
    // FileId is always the first message; 4 = activity
    fmt.Printf("File Type: %v\n",
        fit.Messages[0].FieldValueByNum(fieldnum.FileIdType).Any())

    // Output:
    // FileHeader DataSize: 94080
    // Messages count: 3611
    // File Type: 4
}
```

### Decode with Message Listener

The decoder has the ability to broadcast messages to all registered message listeners, enable us to consume the message as soon as it is decoded. We can register one or more listeners, whether it's a [filedef.Listener](https://github.com/muktihari/fit/blob/master/profile/filedef/listener.go) or your own custom listener (as long as it satisfies the [decoder.MesgListener](https://github.com/muktihari/fit/blob/master/decoder/listener.go) interface) using **WithMesgListener** option.

Please note that by default, the message passed to **OnMesg** function is short-lived object, you need to copy the **mesg.Fields** and **mesg.DeveloperFields** if you want to retain them, or direct the decoder to copy them for us using **WithBroadcastMesgCopy**.

#### Creating your own Message Listener

Here is the simple example of creating message listener. A RecordCounter to count how many Record in a FIT file:

```go
package main

import (
    "fmt"
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/untyped/mesgnum"
    "github.com/muktihari/fit/proto"
)

type RecordCounter struct{ Count int }

var _ decoder.MesgListener = (*RecordCounter)(nil)

func (c *RecordCounter) OnMesg(mesg proto.Message) {
    if mesg.Num == mesgnum.Record {
        c.Count++
    }
}

func main() {
    f, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    rc := new(RecordCounter)

    dec := decoder.New(f,
        // Add activity listener to the decoder:
        decoder.WithMesgListener(rc),
        // Direct the decoder to only broadcast
        // the messages without retaining them:
        decoder.WithBroadcastOnly(),
    )

    _, err = dec.Decode()
    if err != nil {
        panic(err)
    }

    fmt.Printf("We have %d records\n", rc.Count)
}

```

Another example is when we just want to retrieve the Record Messages from a FIT file:

```go
package main

import (
    "fmt"
    "os"
    "slices"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/untyped/mesgnum"
    "github.com/muktihari/fit/proto"
)

type RecordRetriever struct{ Records []proto.Message }

var _ decoder.MesgListener = (*RecordRetriever)(nil)

func (c *RecordRetriever) OnMesg(mesg proto.Message) {
    if mesg.Num == mesgnum.Record {
        // Unless WithBroadcastMesgCopy is used, we must copy
        // the slices since mesg is a short-lived object.
        mesg.Fields = slices.Clone(mesg.Fields)
        mesg.DeveloperFields = slices.Clone(mesg.DeveloperFields)

        c.Records = append(c.Records, mesg)
    }
}

func main() {
    f, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    rr := new(RecordRetriever)

    dec := decoder.New(f,
        // Add activity listener to the decoder:
        decoder.WithMesgListener(rr),
        // Direct the decoder to only broadcast
        // the messages without retaining them:
        decoder.WithBroadcastOnly(),
    )

    _, err = dec.Decode()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Records: %d\n", len(rr.Records))
}
```

### Decode into Common File Types

Decode into Common File Types enables us to interact with FIT files through common file types such as Activity Files, Course Files, Workout Files, and [more](../profile/filedef/doc.go), which group protocol messages based on specific purposes.

1. To get started, the simpliest (but least efficient) way to create an common file type is to decode the FIT file in its protocol messages then pass the messages to create the desired common file type.

```go
package main

import (
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

    dec := decoder.New(f)

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
    fmt.Printf("  Distance: %g m\n", activity.Records[i].DistanceScaled())
    fmt.Printf("  Lat: %d semicircles\n", activity.Records[i].PositionLat)
    fmt.Printf("  Long: %d semicircles\n", activity.Records[i].PositionLong)
    fmt.Printf("  Speed: %g m/s\n", activity.Records[i].SpeedScaled())
    fmt.Printf("  HeartRate: %d bpm\n", activity.Records[i].HeartRate)
    fmt.Printf("  Cadence: %d rpm\n", activity.Records[i].Cadence)

    // Output:
    // File Type: activity
    // Sessions count: 1
    // Laps count: 1
    // Records count: 3601
    //
    // Sample value of record[100]:
    //   Distance: 100 m
    //   Lat: 0 semicircles
    //   Long: 10717 semicircles
    //   Speed: 1 m/s
    //   HeartRate: 126 bpm
    //   Cadence: 100 rpm
}
```

2. The better way to decode common file types using our building block is by using `filedef.Listener` then register it to the `Decoder`, the listener will process every message that being broadcasted by the `Decoder` concurrently.

```go
package main

import (
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

    // The listener will receive every message from the decoder
    // as soon as it is decoded and transform it into an filedef.File.
    lis := filedef.NewListener()
    defer lis.Close() // release channel used by listener

    dec := decoder.New(f,
        // Add activity listener to the decoder:
        decoder.WithMesgListener(lis),
        // Direct the decoder to only broadcast
        // the messages without retaining them:
        decoder.WithBroadcastOnly(),
    )

    _, err = dec.Decode()
    if err != nil {
        panic(err)
    }

    // The resulting File can be retrieved after decoding process completed.
    // filedef.File is just an interface, we can do type assertion like this.
    file, ok := lis.File().(*filedef.Activity)
    if !ok {
        fmt.Printf("%T is not an Activity File\n", lis.File())
        return
    }

    fmt.Printf("File Type: %s\n", file.FileId.Type)
    fmt.Printf("Sessions count: %d\n", len(file.Sessions))
    fmt.Printf("Laps count: %d\n", len(file.Laps))
    fmt.Printf("Records count: %d\n", len(file.Records))
    i := 100
    fmt.Printf("\nSample value of record[%d]:\n", i)
    fmt.Printf("  Distance: %g m\n", file.Records[i].DistanceScaled())
    fmt.Printf("  Lat: %g degrees\n", file.Records[i].PositionLatDegrees())
    fmt.Printf("  Long: %g degrees\n", file.Records[i].PositionLongDegrees())
    fmt.Printf("  Speed: %g m/s\n", file.Records[i].SpeedScaled())
    fmt.Printf("  HeartRate: %d bpm\n", file.Records[i].HeartRate)
    fmt.Printf("  Cadence: %d rpm\n", file.Records[i].Cadence)
    // Output:
    // File Type: activity
    // Sessions count: 1
    // Laps count: 1
    // Records count: 3601
    //
    // Sample value of record[100]:
    //   Distance: 100 m
    //   Lat: 0 degrees
    //   Long: 0.0008982885628938675 degrees
    //   Speed: 1 m/s
    //   HeartRate: 126 bpm
    //   Cadence: 100 rpm
}
```

#### Using your own custom File Types

Not only does [filedef.Listener](https://github.com/muktihari/fit/blob/master/profile/filedef/listener.go#L16C1-L23C2) implements [decoder.MesgListener](https://github.com/muktihari/fit/blob/master/decoder/listener.go) interface, but it also allows you to create your own custom File as long as it satisfies [filedef.File](https://github.com/muktihari/fit/blob/master/profile/filedef/filedef.go#L17C1-L22C2) interface. For example, if you only want to retrieve some messages, you can build your own custom Activity File like this instead of using our default predefined [Activity File](https://github.com/muktihari/fit/blob/master/profile/filedef/activity.go#L19C1-L46C2):

```go
package main

import (
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/filedef"
    "github.com/muktihari/fit/profile/mesgdef"
    "github.com/muktihari/fit/profile/typedef"
    "github.com/muktihari/fit/profile/untyped/mesgnum"
    "github.com/muktihari/fit/proto"
)

var _ filedef.File = (*Activity)(nil)

type Activity struct {
    FileId   mesgdef.FileId
    Activity *mesgdef.Activity
    Sessions []*mesgdef.Session
    Laps     []*mesgdef.Lap
    Records  []*mesgdef.Record
}

func (f *Activity) Add(mesg proto.Message) {
    switch mesg.Num {
    case mesgnum.FileId:
        f.FileId = *mesgdef.NewFileId(&mesg)
    case mesgnum.Activity:
        f.Activity = mesgdef.NewActivity(&mesg)
    case mesgnum.Session:
        f.Sessions = append(f.Sessions, mesgdef.NewSession(&mesg))
    case mesgnum.Lap:
        f.Laps = append(f.Laps, mesgdef.NewLap(&mesg))
    case mesgnum.Record:
        f.Records = append(f.Records, mesgdef.NewRecord(&mesg))
    }
}

func (f *Activity) ToFIT(options *mesgdef.Options) proto.FIT {
    size := 2 + len(f.Sessions) + len(f.Laps) + len(f.Records)
    fit := proto.FIT{Messages: make([]proto.Message, 0, size)}
    fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))
    if f.Activity != nil {
        fit.Messages = append(fit.Messages, f.Activity.ToMesg(options))
    }
    for i := range f.Sessions {
        fit.Messages = append(fit.Messages, f.Sessions[i].ToMesg(options))
    }
    for i := range f.Laps {
        fit.Messages = append(fit.Messages, f.Laps[i].ToMesg(options))
    }
    for i := range f.Records {
        fit.Messages = append(fit.Messages, f.Records[i].ToMesg(options))
    }
    filedef.SortMessagesByTimestamp(fit.Messages)
    return fit
}

func main() {
    f, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    lis := filedef.NewListener(
        // Replace default filedef.Activity with our custom Activity.
        filedef.WithFileFunc(typedef.FileActivity,
            func() filedef.File { return new(Activity) }),
    )
    defer lis.Close()

    dec := decoder.New(f,
        decoder.WithMesgListener(lis),
        decoder.WithBroadcastOnly(),
    )

    _, err = dec.Decode()
    if err != nil {
        panic(err)
    }

    // The resulting File will be *Activity instead of *filedef.Activity.
    file := lis.File().(*Activity)
    fmt.Printf("Distance: %.2f km\n",
        file.Sessions[0].TotalDistanceScaled()/1000)
}
```

### Decode Chained FIT Files

A single invocation of `Decode()` will process exactly one FIT sequence. To decode a chained FIT file containing multiple FIT data, invoke Decode() or DecodeWithContext() method multiple times. For convenience, we can wrap it with the Next() method as follows (optional):

```go
    ...
    dec := decoder.New(f)
    for dec.Next() {
        fit, err := dec.Decode()
        if err != nil {
            return err
        }
        // do something with fit variable
    }
    ...
```

```go
    ...
    lis := filedef.NewListener()
    defer lis.Close()

    dec := decoder.New(f,
        decoder.WithMesgListener(lis),
        decoder.WithBroadcastOnly(),
    )

    for dec.Next() {
        _, err := dec.Decode()
        if err != nil {
            return err
        }

        file, ok := lis.File().(*filedef.Activity)
        if !ok {
            return
        }
        // do something with file variable
    }
    ...
```

You can also use [PeekFileHeader()](#-Peek-FileHeader), [PeekFileId()](#Peek-FileId) and [Discard()](#Discard-FIT-File-Sequences) methods below inside the loop.

### Peek FileHeader

We can verify whether the given file is a FIT file by checking the File Header (first 12-14 bytes). PeekFileHeader decodes only up to FileHeader (first 12-14 bytes) without decoding the whole reader. If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.

NOTE: The FileHeader retrieved from this method is only valid before Decode method is completed, you will need to copy it if you want to use it later by dereferencing the pointer before calling Decode method.

```go
package main

import (
    "fmt"
    "os"

    "github.com/muktihari/fit/decoder"
)

func main() {
    f, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    dec := decoder.New(f)

    fileHeader, err := dec.PeekFileHeader()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%v\n", fileHeader)

    // Output:
    // &{14 32 2147 94080 .FIT 17310}
}
```

### Peek FileId

We don't need to decode the entire FIT file to verify its type. Instead, we can use the 'PeekFileId' method to check the corresponding type. After invoking this method, we can decide whether to proceed with decoding the file or to stop. If we choose to continue, Decode picks up where this left then continue decoding next messages instead of starting from zero.

```go
package main

import (
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

    lis := filedef.NewListener()
    defer lis.Close() // release channel used by listener

    dec := decoder.New(f,
        decoder.WithMesgListener(lis),
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

    activity := lis.File().(*filedef.Activity)

    fmt.Printf("Sessions count: %d\n", len(activity.Sessions))
    fmt.Printf("Laps count: %d\n", len(activity.Laps))
    fmt.Printf("Records count: %d\n", len(activity.Records))
    // ...
}
```

### Discard FIT File Sequences

When handling a chained FIT file, sometimes we only want to decode a certain file type, let's say a Course File, while discarding other file types. Instead of unecessarily decode all the file types but we don't use all of them, we can just discard it. Discard directs the Decoder to efficiently just discard the bytes without doing unnecessary work.

```go
    ...

    for dec.Next() {
        fileId, err := dec.PeekFileId()
        if err != nil {
            return err
        }
        if fileId.Type != typedef.FileCourse {
            // not a Course File, discard this sequence!
            if err := dec.Discard(); err != nil {
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

This method ensures the entire FIT File is valid; any trailing bytes will invalidate its integrity, an error will be returned. This also ensures that the FIT File we're working on doesn't contain any malicious byte code.

This returns the number of sequences completed and any error encountered. The number of sequences completed can help recovering valid FIT sequences in a chained FIT that contains invalid or corrupted data.

More about this: [https://developer.garmin.com/fit/cookbook/isfit-checkintegrity-read/](https://developer.garmin.com/fit/cookbook/isfit-checkintegrity-read/)

```go
package main

import (
    "io"
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/typedef"
)

func main() {
    f, err := os.Open("Course.fit")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    dec := decoder.New(f)

    // If needed, we can invoke PeekFileId() first to check the type of FIT File.
    // Checking the integrity of an Activity File is generally not mandatory in most cases.
    // However, we can only peek the first FIT sequence if it's a chained FIT files.
    fileId, err := dec.PeekFileId()
    if err != nil {
        panic(err)
    }
    if fileId.Type == typedef.FileActivity {
        return // Skip
    }

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
        Num: 65281,
        Fields: []proto.Field{
            {
                FieldBase: &proto.FieldBase{
                    Num:        253,
                    Name:       "Timestamp",
                    Type:       profile.Uint32,
                    BaseType:   basetype.Uint32,
                    Array:      false,
                    Accumulate: false,
                    Scale:      1,
                    Offset:     0,
                    Units:      "s",
                },
            },
        },
    })

    dec := decoder.New(f, decoder.WithFactory(fac))
    ```

1.  **WithMesgListener**: adds message listener to the listener pool so that we can receive the messages as soon as it is decoded.

    Example:

    ```go
    lis := filedef.NewListener()
    defer lis.Close()

    dec := decoder.New(f,
        decoder.WithMesgListener(lis),
    )
    ```

1.  **WithMesgDefListener**: adds message definition listener to the listener pool so that we can receive the message definitions as soon as it is decoded.

    Example:

    ```go
    conv := fitcsv.NewFITToCSVConv(bw)
    defer conv.Wait()

    dec := decoder.New(f,
        decoder.WithMesgDefListener(conv),
        decoder.WithMesgListener(conv),
    )
    ```

1.  **WithBroadcastOnly**: directs the decoder to only broadcast the messages without retaining them.

    Example:

    ```go
    lis := filedef.NewListener()
    defer lis.Close()

    dec := decoder.New(f,
        decoder.WithMesgListener(lis),
        decoder.WithBroadcastOnly(),
    )
    ```

1.  **WithBroadcastMesgCopy**: directs the Decoder to copy the mesg before passing it to listeners. It was the default behavior on <= v0.14.0.

    ```go
    lis := NewLegacyNonBlockingListener() /* any legacy non-blocking listener created on version <= v0.14.0) */
    dec := decoder.New(f,
        decoder.WithMesgListener(lis),
        decoder.WithBroadcastOnly(),
        decoder.WithBroadcastMesgCopy(),
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

1.  **WithLogWriter**: specifies where the log messages will be written to. By default, the Decoder do not write any log if log writer is not specified. The Decoder will only write log messages when it encountered a bad encoded FIT file such as:

    - Field Definition's Size (or Developer Field Definition's Size) is zero.
    - Field Definition's Size (or Developer Field Definition's Size) is less than its basetype's size.
      e.g. Size 1 byte but having basetype uint32 (4 bytes).
    - Encounter a Developer Field without prior Field Description Message.

    ```go
    dec := decoder.New(f, decoder.WithLogWriter(os.Stdout))
    ```

### Advance Decoder Usage

Decoder is programmed to be a reusable object to reduce memory allocation, similar to bytes.Buffer that has Reset method, Decoder has one as well. We can use it with sync.Pool. This might be useful for a system where decoding happens frequently but some of the processes can occur one after another (such as a server).

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/profile/mesgdef"
)

var (
    pool    = sync.Pool{New: func() any { return decoder.New(nil) }}
    lispool = sync.Pool{New: func() any { return filedef.NewListener() }}
)

func main() {
    srv := http.NewServeMux()
    srv.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        dec := pool.Get().(*decoder.Decoder)
        defer pool.Put(dec)  // put decoder back to the pool

        lis := lispool.Get().(*filedef.Listener)
        defer lispool.Put(lis) // put listener back to the pool
        defer lis.Close()      // release channels used by listener

        // Assign reader and options just like
        // when using decoder.New().
        dec.Reset(r.Body,
            decoder.WithMesgListener(lis),
            decoder.WithBroadcastOnly(),
        )
        defer r.Body.Close()

        for dec.Next() {
            _, err := dec.Decode()
            if err != nil {
                writeError(w, err)
                return
            }
            activity, ok := lis.File().(*filedef.Activity)
            if !ok {
                writeError(w, fmt.Errorf("not an activity file"))
            }
            fmt.Fprintf(w, "%+v\n", activity.FileId)
        }
    })
    log.Fatal(http.ListenAndServe(":8080", srv))
}

func writeError(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("{\"err\":%q}", err)))
}
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
    hash16 := crc16.New()

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

Note:

- By default, Encoder will use protocol version in FileHeader for each FIT file, if it's not specified, it will use protocol version 1.0 (proto.V1). If you want to use specific protocol version for the entire encoding regardless the value in FileHeader, please use this Encode Option: WithProtocolVersion. See [Available Encode Options](#Available-Encode-Options)
- Encoder already implements efficient io.Writer buffering, DO NOT wrap io.Writer (such as \*os.File) with buffer such as using \*bufio.Writer; Doing so will greatly reduce performance.

### Encode Protocol Messages

Example of encoding FIT by self declaring the protocol messages, this is to show how we can compose the message using this SDK.

```go
package main

import (
    "os"
    "time"

    "github.com/muktihari/fit/encoder"
    "github.com/muktihari/fit/factory"
    "github.com/muktihari/fit/kit/datetime"
    "github.com/muktihari/fit/profile/typedef"
    "github.com/muktihari/fit/profile/untyped/fieldnum"
    "github.com/muktihari/fit/profile/untyped/mesgnum"
    "github.com/muktihari/fit/proto"
)

func main() {
    now := time.Now()
    fit := proto.FIT{
        Messages: []proto.Message{
            {Num: mesgnum.FileId, Fields: []proto.Field{
                factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity.Byte()),
                factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(now)),
                factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerBryton.Uint16()),
                factory.CreateField(mesgnum.FileId, fieldnum.FileIdProduct).WithValue(uint16(1901)), // Bryton Rider 420
                factory.CreateField(mesgnum.FileId, fieldnum.FileIdProductName).WithValue("Bryton Rider 420"),
            }},
            {Num: mesgnum.Record, Fields: []proto.Field{
                factory.CreateField(mesgnum.Record, fieldnum.RecordTimestamp).WithValue(datetime.ToUint32(now)),
                factory.CreateField(mesgnum.Record, fieldnum.RecordDistance).WithValue(uint32(100)),
                factory.CreateField(mesgnum.Record, fieldnum.RecordSpeed).WithValue(uint16(1000)),
                factory.CreateField(mesgnum.Record, fieldnum.RecordCadence).WithValue(uint8(78)),
                factory.CreateField(mesgnum.Record, fieldnum.RecordHeartRate).WithValue(uint8(100)),
            }},
            {Num: mesgnum.Session, Fields: []proto.Field{
                factory.CreateField(mesgnum.Session, fieldnum.SessionTimestamp).WithValue(datetime.ToUint32(now)),
                factory.CreateField(mesgnum.Session, fieldnum.SessionTotalDistance).WithValue(uint32(100)),
                factory.CreateField(mesgnum.Session, fieldnum.SessionAvgSpeed).WithValue(uint16(1000)),
                factory.CreateField(mesgnum.Session, fieldnum.SessionAvgCadence).WithValue(uint8(78)),
                factory.CreateField(mesgnum.Session, fieldnum.SessionAvgHeartRate).WithValue(uint8(100)),
            }},
            {Num: mesgnum.Activity, Fields: []proto.Field{
                factory.CreateField(mesgnum.Activity, fieldnum.ActivityTimestamp).WithValue(datetime.ToUint32(now)),
                factory.CreateField(mesgnum.Activity, fieldnum.ActivityType).WithValue(typedef.ActivityManual.Byte()),
                factory.CreateField(mesgnum.Activity, fieldnum.ActivityLocalTimestamp).WithValue(datetime.ToUint32(now.Add(7 * time.Hour))), // GMT+7
                factory.CreateField(mesgnum.Activity, fieldnum.ActivityNumSessions).WithValue(uint16(1)),
            }},
        },
    }

    f, err := os.OpenFile("CoffeeRide.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    enc := encoder.New(f)
    if err := enc.Encode(&fit); err != nil {
        panic(err)
    }
}
```

### Encode Common File Types

Example of encoding FIT created using Common File Types building block.

```go
package main

import (
    "os"
    "time"

    "github.com/muktihari/fit/encoder"
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

    f, err := os.OpenFile("NewActivity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    enc := encoder.New(f)
    if err := enc.Encode(&fit); err != nil {
        panic(err)
    }
}
```

Example decoding FIT file into common file `Activity File`, edit the manufacturer and product, and then encode it again.

```go
package main

import (
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/encoder"
    "github.com/muktihari/fit/profile/filedef"
    "github.com/muktihari/fit/profile/typedef"
)

func main() {
    fin, err := os.Open("Activity.fit")
    if err != nil {
        panic(err)
    }
    defer fin.Close()

    lis := filedef.NewListener()
    defer lis.Close()

    dec := decoder.New(fin,
        decoder.WithMesgListener(lis),
        decoder.WithBroadcastOnly(),
    )

    _, err = dec.Decode()
    if err != nil {
        panic(err)
    }

    activity := lis.File().(*filedef.Activity)

    // Do something with the Activity File, for example
    // changing manufacturer and product like this
    activity.FileId.Manufacturer = typedef.ManufacturerGarmin
    activity.FileId.Product = typedef.GarminProductEdge530.Uint16()

    // Convert back to Protocol Messages
    fit := activity.ToFIT(nil)

    fout, err := os.OpenFile("NewActivity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
    if err != nil {
        panic(err)
    }
    defer fout.Close()

    enc := encoder.New(fout)
    if err := enc.Encode(&fit); err != nil {
        panic(err)
    }
}
```

### Available Encode Options

1. **WithProtocolVersion**: directs the Encoder to use specific ProtocolVersion for the entire encoding. By default, Encoder will use ProtocolVersion in FileHeader for each FIT file, if it's not specified, it will use proto.V1. This option overrides the FileHeader's ProtocolVersion and forces all FIT files to use this ProtocolVersion during encoding.

   Example:

   ```go
   enc := encoder.New(f, encoder.WithProtocolVersion(proto.V2))
   ```

1. **WithWriteBufferSize**: directs the Encoder to use this buffer size for writing to io.Writer instead of default 4096. When size <= 0, the Encoder will write directly to io.Writer without buffering.

   Example:

   ```go
   enc := encoder.New(f, encoder.WithWriteBufferSize(0))    // No buffer

   enc := encoder.New(f, encoder.WithWriteBufferSize(8192)) // 8192 buffer
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

   - If you love to live dangerously, you can always bypass the message validator. It might speed up the encoding process, but be warned—it's prone to errors. So, only go for it if you know what you're doing!

   ```go
     // Define your own message validator, for example
     // this validator will always return nil error.
     type messageValidatorBypass struct{}

     func (messageValidatorBypass) Validate(mesg *proto.Message) error { return nil }
     func (messageValidatorBypass) Reset()

     ...

     enc := encoder.New(f, encoder.WithMessageValidator(&messageValidatorBypass{}))
   ```

1. **WithBigEndian**: directs the Encoder to encode values in Big-Endian bytes order (default: Little-Endian).

   Example:

   ```go
   enc := encoder.New(f, encoder.WithBigEndian())
   ```

1. **WithHeaderOption**: directs the Encoder to use this option instead of default HeaderOptionNormal and local message type zero.

   - If HeaderOptionNormal is selected, valid local message type value is 0-15; invalid values will be treated as 15.
   - If HeaderOptionCompressedTimestamp is selected, valid local message type value is 0-3; invalid values will be treated as 3.

     Saves 7 bytes per message when its timestamp is compressed: 3 bytes for field definition and 4 bytes for the uint32 timestamp value.

   - Otherwise, no change will be made and the Encoder will use default values.

   NOTE: To minimize the required RAM for decoding, it's recommended to use a minimal number of local message type.
   For instance, embedded devices may only support decoding data from local message type 0. Additionally,
   multiple local message types should be avoided in file types like settings, where messages of the same type
   can be grouped together.

   Example:

   ```go
   enc := encoder.New(f, encoder.WithHeaderOption(encoder.HeaderOptionNormal, 15)) // valid 0-15

   enc := encoder.New(f, encoder.WithHeaderOption(encoder.HeaderOptionCompressedTimestamp, 3)) // valid 0-3
   ```

### Stream Encoder

This feature enables us to encode per message basis or in streaming fashion rather than bulk per `proto.FIT`. To enable this, the Encoder's Writer should either implement io.WriterAt or io.WriteSeeker, since we need to be able to update FileHeader (the first 14 bytes) for every sequence completed. This is another building block that we can use.

```go
package main

import (
    "os"
    "time"

    "github.com/muktihari/fit/encoder"
    "github.com/muktihari/fit/factory"
    "github.com/muktihari/fit/kit/datetime"
    "github.com/muktihari/fit/profile/typedef"
    "github.com/muktihari/fit/profile/untyped/fieldnum"
    "github.com/muktihari/fit/profile/untyped/mesgnum"
    "github.com/muktihari/fit/proto"
)

func main() {
    f, err := os.OpenFile("Activity.fit", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o777)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    streamEnc, err := encoder.New(f).StreamEncoder()
    if err != nil {
        panic(err)
    }

    // Simplified example, writing only this mesg.
    mesg := proto.Message{Num: mesgnum.FileId, Fields: []proto.Field{
        factory.CreateField(mesgnum.FileId, fieldnum.FileIdType).WithValue(typedef.FileActivity.Byte()),
        factory.CreateField(mesgnum.FileId, fieldnum.FileIdManufacturer).WithValue(typedef.ManufacturerDevelopment.Uint16()),
        factory.CreateField(mesgnum.FileId, fieldnum.FileIdProduct).WithValue(uint16(0)),
        factory.CreateField(mesgnum.FileId, fieldnum.FileIdTimeCreated).WithValue(datetime.ToUint32(time.Now())),
    }}

    // Write per message, we can use this to write message as it arrives.
    // For example, message retrieved from decoder's Listener can be
    // write right away without waiting all messages to be received.
    if err := streamEnc.WriteMessage(&mesg); err != nil {
        panic(err)
    }

    /* Write more messages */

    // After all messages have been written, invoke this to finalize.
    if err := streamEnc.SequenceCompleted(); err != nil {
        panic(err)
    }
}
```

### Advance Encoder Usage

Encoder and StreamEncoder are programmed to be a reusable object to reduce memory allocation, similar to bytes.Buffer that has Reset method, Encoder and StreamEncoder have one as well. We can use it with sync.Pool. This might be useful for a system where encoding happens frequently but some of the processes can occur one after another (such as a server).

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "sync"

    "github.com/muktihari/fit/encoder"
    "github.com/muktihari/fit/profile/filedef"
)



var (
    pool    = sync.Pool{New: func() any { return encoder.New(nil) }}
    bufpool = sync.Pool{New: func() any { return &bufferAt{new(bytes.Buffer)} }}
)

func main() {
    srv := http.NewServeMux()
    srv.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        enc := pool.Get().(*encoder.Encoder)
        defer pool.Put(enc)  // put encoder back to the pool

        // Inmem writer for faster encoding since w does not
        // implement io.WriterAt or io.WriteSeeker.
        buf := bufpool.Get().(*bufferAt)
        defer bufpool.Put(buf)
        buf.Reset()

        // Assign writer and options just like
        // when using encoder.New().
        enc.Reset(buf)

        // Assume the data is in filedef.Activity JSON format encoding.
        var activity filedef.Activity
        err := json.NewDecoder(r.Body).Decode(&activity)
        r.Body.Close()
        if err != nil {
            writeError(w, err)
            return
        }

        fit := activity.ToFIT(nil)
        if err := enc.Encode(&fit); err != nil {
            writeError(w, err)
            return
        }

        if _, err := buf.WriteTo(w); err != nil {
            writeError(w, err)
            return
        }
    })
    log.Fatal(http.ListenAndServe(":8080", srv))
}

// bufferAt is a wrapper to enable io.WriterAt functionality.
type bufferAt struct{ *bytes.Buffer }

var _ io.WriterAt = (*bufferAt)(nil)

func (b *bufferAt) WriteAt(p []byte, off int64) (int, error) {
    return copy(b.Bytes()[off:], p), nil
}

func writeError(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(fmt.Sprintf("{\"err\":%q}", err)))
}
```
