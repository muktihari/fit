# Fit SDK for Go

![GitHub Workflow Status](https://github.com/muktihari/fit/workflows/CI/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/muktihari/fit.svg)](https://pkg.go.dev/github.com/muktihari/fit)
[![CodeCov](https://codecov.io/gh/muktihari/fit/branch/master/graph/badge.svg)](https://codecov.io/gh/muktihari/fit)
[![Go Report Card](https://goreportcard.com/badge/github.com/muktihari/fit)](https://goreportcard.com/report/github.com/muktihari/fit)
[![SDK Version](https://img.shields.io/badge/sdkversion-21.126-lightblue.svg?style=flat)](https://developer.garmin.com/fit/download)

The Flexible and Interoperable Data Transfer (FIT) protocol is a protocol developed by Garmin for storing and sharing data originating from sports, fitness, and health devices.
When recording an activity using devices such as cycling computer, smartwatch, and similiar devices, chances are the resulting file is often in FIT file format (\*.fit).
The FIT file is a binary file format known for its compact size, making it the preferred choice for manufacturers to use in their embedded devices.
However, despite having gained widespread adoption, Garmin has not yet released an official SDK for Go, this is where this SDK comes in to bridge the gap, enabling Go developers to be able to interact with FIT file format.

More about FIT: [developer.garmin.com/fit](https://developer.garmin.com/fit)

This SDK is designed with efficiency in mind, but it places a higher priority on clarity, simplicity and extensibility.

## Usage

Please see [Usage](/docs/usage.md).

## CLIs

We provides some CLI programs to interact with FIT files that can be found in `cmd` folder.

1. `fitactivity`: Combines multiple FIT activity files into one continuous FIT activity (and conceal the start and end GPS positions for privacy). [README.md](/cmd/fitactivity/README.md)
2. `fitconv`: Converts FIT file to a CSV file, allowing us to read the FIT file in a human-readable format. [README.md](/cmd/fitconv/README.md)
3. `fitprint`: Print FIT file into stdout in human-readable format. [README.md](/cmd/fitprint/README.md)

Now, the `fitactivity` and `fitconv` programs are automatically built during release; for Linux, Windows, and macOS platforms. They are available for download in [Release's Assets](https://github.com/muktihari/fit/releases).

## Generate Custom FIT SDK

Please see [Generate Custom FIT SDK](/docs/generating_code.md#Generate-Custom-FIT-SDK)

## Contributing

Please see [CONTRIBUTING.md](/CONTRIBUTING.md).
Thank you, [contributors](https://github.com/muktihari/fit/graphs/contributors)!
