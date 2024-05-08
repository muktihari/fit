# FIT SDK for Go

![GitHub Workflow Status](https://github.com/muktihari/fit/workflows/CI/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/muktihari/fit.svg)](https://pkg.go.dev/github.com/muktihari/fit)
[![CodeCov](https://codecov.io/gh/muktihari/fit/branch/master/graph/badge.svg)](https://codecov.io/gh/muktihari/fit)
[![Go Report Card](https://goreportcard.com/badge/github.com/muktihari/fit)](https://goreportcard.com/report/github.com/muktihari/fit)
[![Profile Version](https://img.shields.io/badge/profile-v21.133-lightblue.svg?style=flat)](https://developer.garmin.com/fit/download)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8731/badge)](https://www.bestpractices.dev/projects/8731)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/muktihari/fit/badge)](https://securityscorecards.dev/viewer/?uri=github.com/muktihari/fit)

This project hosts the Go implementation for [The Flexible and Interoperable Data Transfer (FIT) protocol](developer.garmin.com/fit), which is a protocol developed by Garmin for storing and sharing data originating from sports, fitness, and health devices. Activities recorded using devices such as smartwatch and cycling computer are now mostly in a FIT file format (\*.fit).

## Motivation

The FIT protocol, known for its compact size as a binary file format, is the preferred choice for manufacturers to use in their embedded devices. However, despite its widespread adoption, Garmin has not yet released an official SDK for Go, and existing third-party libraries for decoding and encoding the FIT protocol lack the semantics of the official SDK.

One of the key semantics they are missing is the ability for users to retrieve raw protocol messages; instead, they decode directly into predefined message structs grouped by [common file types](https://developer.garmin.com/fit/file-types). Furthermore, existing third-party libraries do not seem to fully support FIT Protocol V2, and their ability to produce variant options of the FIT protocol is limited. For instance, creating FIT files with compressed timestamps or FIT files with multiple local message types, which significantly reduces the resulting FIT files' size, is missing.

Without diminishing respect for the existing libraries created nearly a decade ago, at a time when the capabilities of Go were limited, we believe a new approach is necessary. This is where this SDK comes in, bridging the gap and enabling Go developers to seamlessly interact with the FIT protocol.

## Usage

Please see [Usage](/docs/usage.md).

## Protocol Version 2.0 is supported

Version 2.0 introduced [**Developer Fields**](https://developer.garmin.com/fit/cookbook/developer-data) as a way to add custom data fields to existing messages. We strives to support **Developer Fields** and carefully thought about how to implement it since the inception of the SDK. While this may still need to be battle-tested to ensure correctness, this is generally work and usable.

Here is the sample of what **Developer Fields** would look like in a **.fit** that have been converted to **.csv** by [fitconv](/cmd/fitconv/README.md). The **device_info** message has some **Developer Fields** defined in **field_description** (the ones that are being bold):

<table class="table table-bordered table-hover table-condensed">
<thead>
<tr>
    <th>Type</th>
    <th>Local Number</th>
    <th>Message</th>
    <th>Field 1</th>
    <th>Value 1</th>
    <th>Units 1</th>
    <th>Field 2</th>
    <th>Value 2</th>
    <th>Units 2</th>
    <th>Field 3</th>
    <th>Value 3</th>
    <th>Units 3</th>
    <th>Field 4</th>
    <th>Value 4</th>
    <th>Units 4</th>
    <th>Field 5</th>
    <th>Value 5</th>
    <th>Units 5</th>
    <th>Field 6</th>
    <th>Value 6</th>
    <th>Units 6</th>
</tr>
</thead>
<tbody>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>developer_data_id</td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>application_id</td>
    <td>16</td>
    <td> </td>
    <td>application_version</td>
    <td>1</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>developer_data_id</td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>application_id</td>
    <td>&lt;omitted&gt;</td>
    <td> </td>
    <td>application_version</td>
    <td>40113</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>1</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>1</td>
    <td></td>
    <td>field_name</td>
    <td>13</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>5</td>
    <td></td>
    <td>field_name</td>
    <td><strong>device_model</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>1</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>1</td>
    <td></td>
    <td>field_name</td>
    <td>20</td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>4</td>
    <td></td>
    <td>field_name</td>
    <td><strong>device_manufacturer</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>6</td>
    <td></td>
    <td>field_name</td>
    <td><strong>device_os_version</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>field_description</td>
    <td>fit_base_type_id</td>
    <td>7</td>
    <td> </td>
    <td>developer_data_index</td>
    <td>1</td>
    <td> </td>
    <td>field_definition_number</td>
    <td>7</td>
    <td></td>
    <td>field_name</td>
    <td><strong>mobile_app_version</strong></td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
    <td> </td>
</tr>
<tr>
    <td>Definition</td>
    <td>0</td>
    <td>device_info</td>
    <td>manufacturer</td>
    <td>1</td>
    <td> </td>
    <td>product</td>
    <td>1</td>
    <td> </td>
    <td>device_model</td>
    <td>11</td>
    <td></td>
    <td>device_manufacturer</td>
    <td>6</td>
    <td> </td>
    <td>device_os_version</td>
    <td>5</td>
    <td></td>
    <td>mobile_app_version</td>
    <td>8</td>
    <td> </td>
</tr>
<tr>
    <td>Data</td>
    <td>0</td>
    <td>device_info</td>
    <td>manufacturer</td>
    <td>265</td>
    <td> </td>
    <td>product</td>
    <td>101</td>
    <td> </td>
    <td><strong>device_model</strong></td>
    <td><strong>iPhone14,4</strong></td>
    <td> </td>
    <td><strong>device_manufacturer</strong></td>
    <td><strong>apple</strong></td>
    <td> </td>
    <td><strong>device_os_version</strong></td>
    <td><strong>16</strong>.6</td>
    <td> </td>
    <td><strong>mobile_app_version</strong></td>
    <td><strong>332.0.0</strong></td>
    <td> </td>
</tr>
</tbody>
</table>

## CLIs

We provide some CLI programs to interact with FIT files that can be found in [cmd](/cmd/doc.go) folder.

1. **fitactivity**: Combines multiple FIT activity files into one continuous FIT activity (and conceal the start and end GPS positions for privacy). [README.md](/cmd/fitactivity/README.md)
2. **fitconv**: Converts FIT files to CSV format, enabling us to read the FIT data in a human-readable format. Conversely, it also converts CSV files back to FIT format, enabling us to create or edit FIT files in CSV form. The programs is designed to work seamlessly with CSVs produced by the Official FIT SDK's _FitCSVTool.jar_. [README.md](/cmd/fitconv/README.md)
3. **fitprint**: Generates comprehensive human-readable **\*.txt** file containing details extracted from FIT files. [README.md](/cmd/fitprint/README.md)
4. **fitdump**: Dumps the FIT file(s) into segmented bytes in a **\*.txt** file format. [README.md](/cmd/fitdump/README.md)

Some programs are automatically built during release; for Linux, Windows, and macOS platforms. They are available for download in [Release's Assets](https://github.com/muktihari/fit/releases).

## Custom FIT SDK

A FIT file may contain manufacturer specific messages (`Product Profile`) that are not defined in `Global Profile (Profile.xlsx)` since it's specific to a manufacturer (other manufacturers may have different meaning for that messages)

To be able to decode or create the manufacturer specific messages, we provide options to pick based on your need:

1. Register Manufacturer Specific Messages at Runtime

   For those who prefer using this SDK as it is without need to generate their own custom SDK, we provide `factory` package as an abstraction to hold the profile messages. See [Register at Runtime](/docs/runtime_registration.md).

2. Generate Custom FIT SDK

   Please see [Generate Custom FIT SDK](/docs/generating_code.md#Generate-Custom-FIT-SDK)

## Benchmark

We do not aim to compete with anyone; rather, we have created this FIT SDK with the intention of providing an alternative. However, having a benchmark can show us how relevant we are to the world.

Here is a benchmark for decoding and encoding [big_activity.fit](./testdata/big_activity.fit) using this FIT SDK in comparison to [github.com/tormoder/fit](https://github.com/tormoder/fit), the long-standing Go library for decoding and encoding FIT files. See [internal/cmd/benchfit/benchfit_test.go](./internal/cmd/benchfit/benchfit_test.go)

```sh
cd internal/cmd/benchfit
go test -bench=. -benchmem
```

```js
goos: darwin
goarch: amd64
pkg: benchfit
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
BenchmarkDecode/muktihari/fit_raw-4  12    95655408 ns/op   77092784 B/op    100047 allocs/op
BenchmarkDecode/muktihari/fit-4      13    87700988 ns/op   52699547 B/op    101064 allocs/op
BenchmarkDecode/tormoder/fit-4       10   106331934 ns/op   84108932 B/op    700051 allocs/op
BenchmarkEncode/muktihari/fit_raw-4  15    75523944 ns/op     131522 B/op        14 allocs/op
BenchmarkEncode/muktihari/fit-4       8   147434340 ns/op   44139516 B/op    100018 allocs/op
BenchmarkEncode/tormoder/fit-4        1  1301732705 ns/op  101992544 B/op  12100312 allocs/op
PASS
ok      benchfit        10.958s
```

NOTE: The `1st` on the list, "raw", means we decode the file into the original FIT protocol message structure (similar to the Official FIT SDK implementation in other languages). While the `2nd` decodes messages to **Activity File** struct, which should be equivalent to what the `3rd` does.

We decode slightly faster and encode significantly faster. We allocate far fewer objects on the heap and have a smaller memory footprint for both decoding and encoding.

## Contributing

Please see [CONTRIBUTING.md](/CONTRIBUTING.md).
Thank you, [contributors](https://github.com/muktihari/fit/graphs/contributors)!
