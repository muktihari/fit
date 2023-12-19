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

## Protocol Version 2.0 is supported

Version 2.0 introduced [**Developer Fields**](https://developer.garmin.com/fit/cookbook/developer-data) as a way to add custom data fields to existing messages. We strives to support **Developer Fields** and carefully thought about how to implement it since the inception of the SDK. While this may still need to be battle-tested to ensure correctness, this is generally work and usable.

Here is the sample of what **Developer Fields** would look like in a **.fit** that have been converted to **.csv** by `fitconv`. The **device_info** message has some **Developer Fields** defined in **field_description**:

| Type       | Local Number | Message           | Field 1              | Value 1 | Units 1 | Field 2              | Value 2                                                                                                                             | Units 2 | Field 3                 | Value 3    | Units 3 | Field 4             | Value 4             | Units 4 | Field 5           | Value 5 | Units 5 | Field 6            | Value 6 | Units 6 |
| ---------- | ------------ | ----------------- | -------------------- | ------- | ------- | -------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ------- | ----------------------- | ---------- | ------- | ------------------- | ------------------- | ------- | ----------------- | ------- | ------- | ------------------ | ------- | ------- |
| Definition | 0            | developer_data_id | developer_data_index | 1       |         | application_id       | 16                                                                                                                                  |         | application_version     | 1          |         |                     |                     |         |                   |         |         |                    |         |         |
| Data       | 0            | developer_data_id | developer_data_index | 1       |         | application_id       | 32&#124;99&#124;111&#124;109&#124;46&#124;115&#124;116&#124;114&#124;97&#124;118&#124;97&#124;46&#124;105&#124;111&#124;115&#124;32 |         | application_version     | 40113      |         |                     |                     |         |                   |         |         |                    |         |         |
| Definition | 0            | field_description | fit_base_type_id     | 1       |         | developer_data_index | 1                                                                                                                                   |         | field_definition_number | 1          |         | field_name          | 13                  |         |                   |         |         |                    |         |         |
| Data       | 0            | field_description | fit_base_type_id     | 7       |         | developer_data_index | 1                                                                                                                                   |         | field_definition_number | 5          |         | field_name          | device_model        |         |                   |         |         |                    |         |         |
| Definition | 0            | field_description | fit_base_type_id     | 1       |         | developer_data_index | 1                                                                                                                                   |         | field_definition_number | 1          |         | field_name          | 20                  |         |                   |         |         |                    |         |         |
| Data       | 0            | field_description | fit_base_type_id     | 7       |         | developer_data_index | 1                                                                                                                                   |         | field_definition_number | 4          |         | field_name          | device_manufacturer |         |                   |         |         |                    |         |         |
| Data       | 0            | field_description | fit_base_type_id     | 7       |         | developer_data_index | 1                                                                                                                                   |         | field_definition_number | 6          |         | field_name          | device_os_version   |         |                   |         |         |                    |         |         |
| Data       | 0            | field_description | fit_base_type_id     | 7       |         | developer_data_index | 1                                                                                                                                   |         | field_definition_number | 7          |         | field_name          | mobile_app_version  |         |                   |         |         |                    |         |         |
| Definition | 0            | device_info       | manufacturer         | 1       |         | product              | 1                                                                                                                                   |         | device_model            | 11         |         | device_manufacturer | 6                   |         | device_os_version | 5       |         | mobile_app_version | 8       |         |
| Data       | 0            | device_info       | manufacturer         | 265     |         | product              | 101                                                                                                                                 |         | device_model            | iPhone14,4 |         | device_manufacturer | apple               |         | device_os_version | 16.6    |         | mobile_app_version | 332.0.0 |

## CLIs

We provides some CLI programs to interact with FIT files that can be found in `cmd` folder.

1. `fitactivity`: Combines multiple FIT activity files into one continuous FIT activity (and conceal the start and end GPS positions for privacy). [README.md](/cmd/fitactivity/README.md)
2. `fitconv`: Converts FIT file to a CSV file, allowing us to read the FIT file in a human-readable format. [README.md](/cmd/fitconv/README.md)

The programs are automatically built during release; for Linux, Windows, and macOS platforms. They are available for download in [Release's Assets](https://github.com/muktihari/fit/releases).

## Custom FIT SDK

A FIT file may contain manufacturer specific messages that are not defined in `Global Profile (Profile.xlsx)` since it's specific to a manufacturer (other manufacturers may have different meaning for that messages)

To be able to decode or create the manufacturer specific messages, we provide options to pick based on your need:

1. Register manufacturer specific messages at runtime

   For those who prefer using this SDK as it is without need to generate their own custom SDK, we provide `factory` package as an abstraction to hold the profile messages. For example, please see [usage.md/#Available-Decode-Options (WithFactory)](/docs/usage.md#Available-Decode-Options).

2. Generate custom FIT SDK

   Please see [Generate Custom FIT SDK](/docs/generating_code.md#Generate-Custom-FIT-SDK)

## Contributing

Please see [CONTRIBUTING.md](/CONTRIBUTING.md).
Thank you, [contributors](https://github.com/muktihari/fit/graphs/contributors)!
