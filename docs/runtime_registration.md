# Register Manufacturer Specific Types & Messages at Runtime

The `typedef` and `factory` packages contains predefined types and messages declared in [Profile.xlsx] provided in the Official FIT SDK. These packages can be updated through [code generation](./generating_code.md) using `fitgen` by providing your custom `Profile.xlsx`.

However, code generation is not for everyone, when using code generation they need to maintain their custom SDK always align with this SDK release. We make a simple approach to allowing register manufacturer specific massage at runtime. It has limitations but it should be sufficient for simple case.

Note: The data presented here is only for demonstration, we don't known any specific messages since we haven't work with the manufacturers we mention below.

## 1. Typedef

Note: If you don't need stringer for your constants or you have your own custom stringer, you can skip this.

To track your own **MesgNum** constants and its string representative, you can add your types in typedef. However, unlike `factory` that you can have different Factory instances for different manufacturers, `typedef` is shared globally, so you can only work with one manufacturer for this (this is one of our current limitation that we mention earlier). Other than **MesgNum**, you can register **File** type as well.

```go

package main

import "github.com/muktihari/fit/profile/typedef"

func main() {
    // Register specific MesgNum
    typedef.MesgNumRegister(68, "Internal Message")
    typedef.MesgNumRegister(65282, "Product Information")

    // If your company have specific File Type, you can register it as well.
    typedef.FileRegister(247, "Internal File")
    ...
}

```

## 2. Factory

Let's say you will work with FIT files from two different manufacturers, or maybe you are the manufacturer itself that want to use this SDK to build a service in Go. You can define messages that only your company use in the FIT you are created.

The available range is between `0xFF00 - 0xFFFE (65280 - 65534)`, but we found some company are using lower number such as `68`, so we accommodate that as long as the number is not already defined in original `Profile.xlsx`.

```go
package main

import (
    "app/bryton"
    "app/garmin"
    "os"

    "github.com/muktihari/fit/decoder"
    "github.com/muktihari/fit/encoder"
    "github.com/muktihari/fit/factory"
    "github.com/muktihari/fit/profile"
    "github.com/muktihari/fit/profile/filedef"
    "github.com/muktihari/fit/proto"
)

func main() {
    brytonFactory := factory.New()

    brytonFactory.RegisterMesg(proto.Message{
        Num: 68, // I found this mesg num used by Bryton & Garmin in their FIT file.
        Fields: []proto.Field{
            {
                FieldBase: &proto.FieldBase{
                    Num:    0,
                    Name:   "Software Version",
                    Type:   profile.Uint16,
                    Scale:  1,
                    Offset: 0,
                },
            },
        },
    })

    brytonFactory.RegisterMesg(proto.Message{
        Num: 65290,
        Fields: []proto.Field{
            {
                FieldBase: &proto.FieldBase{
                    Num:    0,
                    Name:   "Max Heart Rate",
                    Type:   profile.Uint8,
                    Scale:  1,
                    Offset: 0,
                },
            },
        },
    })

    garminFactory := factory.New()

    garminFactory.RegisterMesg(proto.Message{
        Num: 65282, // I found this mesg num used by Garmin in their FIT file.
        Fields: []proto.Field{
            {
                FieldBase: &proto.FieldBase{
                    Num:    0,
                    Name:   "Region",
                    Type:   profile.String,
                    Scale:  1,
                    Offset: 0,
                },
            },
            {
                FieldBase: &proto.FieldBase{
                    Num:    1,
                    Name:   "Product Year",
                    Type:   profile.Uint8,
                    Scale:  1,
                    Offset: 0,
                },
            },
        },
    })

    // Add the factory to your service
    brytonService := bryton.NewService(brytonFactory)
    garminService := garmin.NewService(garminFactory)

    ...

    // Or if you just want to decode FIT files right away, add it to decoder
    brytonDec := decoder.New(f, decoder.WithFactory(brytonFactory))
    garminDec := decoder.New(f, decoder.WithFactory(garminFactory))

    ...
}

```

Note: When using this approach, by default, you can only work with RAW messages since `mesgdef` is only available through code generation.
